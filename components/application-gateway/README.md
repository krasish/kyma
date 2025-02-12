# Application Gateway

## Overview

This is the repository for the Kyma Application Gateway.

## Prerequisites

The Application Gateway requires Go 1.8 or higher.

## Installation

To install the Application Gateway, follow these steps:

1. `git clone git@github.com:kyma-project/kyma.git`
2. `cd kyma/components/application-gateway`
3. `CGO_ENABLED=0 go build ./cmd/applicationgateway`

## Usage

This section explains how to use the Application Gateway.

### Start the Application Gateway

To start the Application Gateway, run this command:

```
./applicationgateway
```

The Application Gateway has the following parameters:
- **proxyPort** is the port that acts as a proxy for the calls from services and Functions to an external solution. The default port is `8080`.
- **externalAPIPort** is the port that exposes the API allowing to check component status. The default port is `8081`.
- **application** is the Application name used to write and read information about services. The default Application is `default-ec`.
- **namespace** is the Namespace in which the Application Gateway is deployed. The default Namespace is `kyma-system`.
- **requestTimeout** is the timeout for requests sent through the Application Gateway, expressed in seconds. The default value is `1`.
- **skipVerify** is the flag for skipping the verification of certificates for the proxy targets. The default value is `false`.
- **requestLogging** is the flag for logging incoming requests. The default value is `false`.
- **proxyTimeout** is the timeout for requests sent through the proxy, expressed in seconds. The default value is `10`.
- **proxyCacheTTL** is the time to live of the remote API information stored in the proxy cache, expressed in seconds. The default value is `120`.

## Development

This section explains the development process.

### Generate mocks

Prerequisites:

 - [Mockery](https://github.com/vektra/mockery) 2.0 or higher

To generate mocks, run:

```sh
go generate ./...
```

When adding a new interface to be mocked or when a mock of an existing interface is not being generated, add the following line directly above the interface declaration:

```
//go:generate mockery --name {INTERFACE_NAME}
```

### Tests

This section outlines the testing details.

#### Unit tests

To run the unit tests, use the following command:

```
go test `go list ./internal/... ./cmd/...`
```
### Generate Kubernetes clients for custom resources

1. Create a directory structure for each client, similar to the one in `pkg/apis`. For example, when generating a client for EgressRule in Istio, the directory structure looks like this: `pkg/apis/istio/v1alpha2`.
2. After creating the directories, define the following files:
    - `doc.go`
    - `register.go`
    - `types.go` - define the custom structs that reflect the fields of the custom resource.

See an example in `pkg/apis/istio/v1alpha2`.

3. Go to the project root directory and run `./hack/update-codegen.sh`. The script generates a new client in `pkg/apis/client/clientset`.

### Contribution

To learn how you can contribute to this project, see the [Contributing](/CONTRIBUTING.md) document.
