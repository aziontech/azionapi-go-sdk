# Go API client for edgefunctions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import edgefunctions "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `edgefunctions.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), edgefunctions.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `edgefunctions.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), edgefunctions.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `edgefunctions.ContextOperationServerIndices` and `edgefunctions.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), edgefunctions.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), edgefunctions.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.azionapi.net*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EdgeFunctionsAPI* | [**EdgeFunctionsGet**](docs/EdgeFunctionsAPI.md#edgefunctionsget) | **Get** /edge_functions | edge_functions
*EdgeFunctionsAPI* | [**EdgeFunctionsIdDelete**](docs/EdgeFunctionsAPI.md#edgefunctionsiddelete) | **Delete** /edge_functions/{id} | edge_functions
*EdgeFunctionsAPI* | [**EdgeFunctionsIdGet**](docs/EdgeFunctionsAPI.md#edgefunctionsidget) | **Get** /edge_functions/{id} | edge_functions
*EdgeFunctionsAPI* | [**EdgeFunctionsIdPatch**](docs/EdgeFunctionsAPI.md#edgefunctionsidpatch) | **Patch** /edge_functions/{id} | edge_functions
*EdgeFunctionsAPI* | [**EdgeFunctionsIdPut**](docs/EdgeFunctionsAPI.md#edgefunctionsidput) | **Put** /edge_functions/{id} | edge_functions
*EdgeFunctionsAPI* | [**EdgeFunctionsPost**](docs/EdgeFunctionsAPI.md#edgefunctionspost) | **Post** /edge_functions | edge_functions


## Documentation For Models

 - [BadRequestResponse](docs/BadRequestResponse.md)
 - [CreateEdgeFunctionRequest](docs/CreateEdgeFunctionRequest.md)
 - [EdgeFunctionResponse](docs/EdgeFunctionResponse.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [Links](docs/Links.md)
 - [ListEdgeFunctionResponse](docs/ListEdgeFunctionResponse.md)
 - [PatchEdgeFunctionRequest](docs/PatchEdgeFunctionRequest.md)
 - [PutEdgeFunctionRequest](docs/PutEdgeFunctionRequest.md)
 - [Results](docs/Results.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: tokenAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		edgefunctions.ContextAPIKeys,
		map[string]edgefunctions.APIKey{
			"tokenAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



