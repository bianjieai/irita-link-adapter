module github.com/bianjieai/irita-link-adapter

go 1.16

require (
	github.com/bianjieai/irita-sdk-go v1.1.1-0.20210707070124-79ed0124b3de
	github.com/gin-gonic/gin v1.6.0
	github.com/smartcontractkit/chainlink v0.9.4
	go.uber.org/zap v1.16.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413
)
