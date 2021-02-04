#!/usr/bin/make -f

export GO111MODULE = on
export GOPROXY = https://goproxy.io

install:
	@echo "installing irita-link-adapter..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/irita-link-adapter
