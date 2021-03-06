package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	servicesdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	"github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

type Request struct {
	RequestID string `json:"request_id"`
	Result    string `json:"result"`
}

type ServiceResponseResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ServiceResponseOutput struct {
	Header models.JSON `json:"header"`
	Body   models.JSON `json:"body"`
}

type KeyParams struct {
	Mnemonic string
	Name     string
	Password string
	Address  types.AccAddress
}

type Endpoint struct {
	ChainID string
	RPC     string
	GRPC    string
}

type IritaAdapter struct {
	Client    servicesdk.ServiceClient
	KeyParams KeyParams
}

func NewIritaAdapter(endpoint Endpoint, keyParams KeyParams) (*IritaAdapter, error) {
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.ModeOption(types.Commit),
	}

	cfg, err := types.NewClientConfig(
		endpoint.RPC,
		endpoint.GRPC,
		endpoint.ChainID,
		options...,
	)
	if err != nil {
		return nil, err
	}

	client := servicesdk.NewServiceClient(cfg)

	address, err := client.Recover(keyParams.Name, keyParams.Password, keyParams.Mnemonic)
	if err != nil {
		return nil, err
	}

	keyParams.Address, _ = types.AccAddressFromBech32(address)

	return &IritaAdapter{
		Client:    client,
		KeyParams: keyParams,
	}, nil
}

func (adapter IritaAdapter) handle(req Request) (interface{}, error) {
	logger.Infof("Request received: %+v", req)

	requestIDBz, err := hex.DecodeString(req.RequestID)
	if err != nil {
		return nil, err
	}

	result, output := adapter.buildServiceResponse(req.Result)

	msg := service.MsgRespondService{
		RequestId: requestIDBz,
		Result:    result,
		Output:    output,
		Provider:  adapter.KeyParams.Address,
	}

	baseTx := types.BaseTx{
		From:     adapter.KeyParams.Name,
		Password: adapter.KeyParams.Password,
	}

	res, err := adapter.Client.BuildAndSend([]types.Msg{&msg}, baseTx)
	if err != nil {
		return nil, fmt.Errorf("Failed to send transaction: %s", err.Error())
	}

	logger.Infof("Transaction sent successfully: %s", res.Hash)

	return res, nil
}

func (adapter IritaAdapter) buildServiceResponse(payload string) (result, output string) {
	code := 200
	message := ""

	_, err := models.ParseJSON([]byte(payload))
	if err != nil {
		code = 500
		message = "failed to process request"
	}

	if code == 200 {
		output, _ = buildServiceResponseOutput("{}", payload)
	}

	result, _ = buildServiceResponseResult(code, message)

	return result, output
}

func buildServiceResponseResult(code int, message string) (string, error) {
	result := ServiceResponseResult{
		Code:    code,
		Message: message,
	}

	bz, err := json.Marshal(result)
	return string(bz), err
}

func buildServiceResponseOutput(header, body string) (string, error) {
	headerJS, err := models.ParseJSON([]byte(header))
	if err != nil {
		return "", err
	}

	bodyJS, err := models.ParseJSON([]byte(body))
	if err != nil {
		return "", err
	}

	output := ServiceResponseOutput{
		Header: headerJS,
		Body:   bodyJS,
	}

	bz, err := json.Marshal(output)
	return string(bz), err
}
