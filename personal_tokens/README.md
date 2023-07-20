# Go API client for personal_tokens

The Personal Tokens API allows you to manage your existing personal tokens.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import personal_tokens "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), personal_tokens.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), personal_tokens.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), personal_tokens.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), personal_tokens.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.azionapi.net*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PersonalTokenApi* | [**CreatePersonalToken**](docs/PersonalTokenApi.md#createpersonaltoken) | **Post** /iam/personal_tokens | Create a new personal token
*PersonalTokenApi* | [**DeletePersonalToken**](docs/PersonalTokenApi.md#deletepersonaltoken) | **Delete** /iam/personal_tokens/{personalTokenId} | Delete a personal token by id
*PersonalTokenApi* | [**GetPersonalToken**](docs/PersonalTokenApi.md#getpersonaltoken) | **Get** /iam/personal_tokens/{personalTokenId} | Get a personal token info
*PersonalTokenApi* | [**ListPersonalToken**](docs/PersonalTokenApi.md#listpersonaltoken) | **Get** /iam/personal_tokens | List all existing personal token


## Documentation For Models

 - [CreatePersonalTokenRequest](docs/CreatePersonalTokenRequest.md)
 - [CreatePersonalTokenResponse](docs/CreatePersonalTokenResponse.md)
 - [PersonalTokenResponseGet](docs/PersonalTokenResponseGet.md)
 - [PersonalTokenResponseWithResults](docs/PersonalTokenResponseWithResults.md)
 - [PersonalTokenResponseWithResultsLinks](docs/PersonalTokenResponseWithResultsLinks.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


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


