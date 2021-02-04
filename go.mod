module github.com/bianjieai/irita-link-adapter

go 1.13

require (
	github.com/gin-gonic/gin v1.6.0
	github.com/irisnet/service-sdk-go v0.0.0-20201030091855-7f57f83f8c6c
	github.com/smartcontractkit/chainlink v0.9.4
	go.uber.org/zap v1.16.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-200930
)
