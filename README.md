# IRITA Chainlink Adapter

The project is intended for the IRITA adapter for Chainlink.

## Get Started

### Install

```bash
make install
```

### Configuration

#### Environment variables

| Key | Description |
|-----|-------------|
| `ILA_KEY_MNEMONIC` | IRITA key mnemonic |
| `ILA_CHAIN_ID` | Chain ID of the IRITA network |
| `ILA_ENDPOINT_RPC` | Endpoint RPC address for the IRITA node to connect to |
| `ILA_ENDPOINT_GRPC` | Endpoint gRPC address for the IRITA node to connect to |
| `ILA_LISTEN_ADDR` | Address on which the IRITA adapter will be listening, default to `0.0.0.0:8080` |

### Start

```bash
irita-link-adapter
```
