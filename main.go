package main

import (
	"os"

	"github.com/smartcontractkit/chainlink/core/logger"
	"go.uber.org/zap/zapcore"
)

func init() {
	logger.SetLogger(logger.CreateProductionLogger("", false, zapcore.DebugLevel, false))
}

func main() {
	logger.Info("Starting the IRITA adapter")

	chainID := os.Getenv("ILA_CHAIN_ID")
	endpointRPC := os.Getenv("ILA_ENDPOINT_RPC")
	endpointGRPC := os.Getenv("ILA_ENDPOINT_GRPC")
	txFee := os.Getenv("ILA_TX_FEE")
	mnemonic := os.Getenv("ILA_KEY_MNEMONIC")
	listenAddr := os.Getenv("ILA_LISTEN_ADDR")

	endpoint := Endpoint{
		ChainID: chainID,
		RPC:     endpointRPC,
		GRPC:    endpointGRPC,
		Fee:     txFee,
	}

	keyParams := KeyParams{
		Mnemonic: mnemonic,
		Name:     DefaultKeyName,
		Password: DefaultKeyPass,
	}

	adapter, err := NewIritaAdapter(endpoint, keyParams)
	if err != nil {
		logger.Errorf("Failed to create the IRITA adapter: %s", err)
		return
	}

	RunWebServer(listenAddr, adapter.handle)
}
