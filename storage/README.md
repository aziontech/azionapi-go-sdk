# Go API client for storage

REST API OpenAPI documentation for the Object Storage

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0 (v1)
- Package version: 1.0.0
- Generator version: 7.4.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import storage "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `storage.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), storage.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `storage.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), storage.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `storage.ContextOperationServerIndices` and `storage.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), storage.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), storage.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.azion.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*StorageAPI* | [**StorageApiBucketsCreate**](docs/StorageAPI.md#storageapibucketscreate) | **Post** /v4/storage/buckets | Create a new bucket
*StorageAPI* | [**StorageApiBucketsDestroy**](docs/StorageAPI.md#storageapibucketsdestroy) | **Delete** /v4/storage/buckets/{name} | Delete a bucket
*StorageAPI* | [**StorageApiBucketsList**](docs/StorageAPI.md#storageapibucketslist) | **Get** /v4/storage/buckets | List buckets
*StorageAPI* | [**StorageApiBucketsObjectsCreate**](docs/StorageAPI.md#storageapibucketsobjectscreate) | **Post** /v4/storage/buckets/{bucket_name}/objects/{object_key} | Create new object key
*StorageAPI* | [**StorageApiBucketsObjectsDestroy**](docs/StorageAPI.md#storageapibucketsobjectsdestroy) | **Delete** /v4/storage/buckets/{bucket_name}/objects/{object_key} | Delete object key
*StorageAPI* | [**StorageApiBucketsObjectsList**](docs/StorageAPI.md#storageapibucketsobjectslist) | **Get** /v4/storage/buckets/{bucket_name}/objects | List buckets objects
*StorageAPI* | [**StorageApiBucketsObjectsRetrieve**](docs/StorageAPI.md#storageapibucketsobjectsretrieve) | **Get** /v4/storage/buckets/{bucket_name}/objects/{object_key} | Download object
*StorageAPI* | [**StorageApiBucketsObjectsUpdate**](docs/StorageAPI.md#storageapibucketsobjectsupdate) | **Put** /v4/storage/buckets/{bucket_name}/objects/{object_key} | Update the object key
*StorageAPI* | [**StorageApiBucketsPartialUpdate**](docs/StorageAPI.md#storageapibucketspartialupdate) | **Patch** /v4/storage/buckets/{name} | Update bucket info


## Documentation For Models

 - [Bucket](docs/Bucket.md)
 - [BucketCreate](docs/BucketCreate.md)
 - [BucketObject](docs/BucketObject.md)
 - [BucketUpdate](docs/BucketUpdate.md)
 - [EdgeAccessEnum](docs/EdgeAccessEnum.md)
 - [ObjectResponseData](docs/ObjectResponseData.md)
 - [PaginatedBucketList](docs/PaginatedBucketList.md)
 - [PaginatedBucketObjectList](docs/PaginatedBucketObjectList.md)
 - [ResponseBucket](docs/ResponseBucket.md)
 - [StateEnum](docs/StateEnum.md)
 - [SuccessBucketOperation](docs/SuccessBucketOperation.md)
 - [SuccessObjectOperation](docs/SuccessObjectOperation.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		storage.ContextAPIKeys,
		map[string]storage.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
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


