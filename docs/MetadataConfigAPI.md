# \MetadataConfigAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetadataConfig**](MetadataConfigAPI.md#GetMetadataConfig) | **Get** /api/v3/config/metadata | 
[**GetMetadataConfigById**](MetadataConfigAPI.md#GetMetadataConfigById) | **Get** /api/v3/config/metadata/{id} | 
[**UpdateMetadataConfig**](MetadataConfigAPI.md#UpdateMetadataConfig) | **Put** /api/v3/config/metadata/{id} | 



## GetMetadataConfig

> MetadataConfigResource GetMetadataConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	whisparrClient "github.com/devopsarr/whisparr-go/whisparr"
)

func main() {

	configuration := whisparrClient.NewConfiguration()
	apiClient := whisparrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataConfigAPI.GetMetadataConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataConfigAPI.GetMetadataConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataConfig`: MetadataConfigResource
	fmt.Fprintf(os.Stdout, "Response from `MetadataConfigAPI.GetMetadataConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataConfigRequest struct via the builder pattern


### Return type

[**MetadataConfigResource**](MetadataConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataConfigById

> MetadataConfigResource GetMetadataConfigById(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	whisparrClient "github.com/devopsarr/whisparr-go/whisparr"
)

func main() {
	id := int32(56) // int32 | 

	configuration := whisparrClient.NewConfiguration()
	apiClient := whisparrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataConfigAPI.GetMetadataConfigById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataConfigAPI.GetMetadataConfigById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataConfigById`: MetadataConfigResource
	fmt.Fprintf(os.Stdout, "Response from `MetadataConfigAPI.GetMetadataConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataConfigResource**](MetadataConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetadataConfig

> MetadataConfigResource UpdateMetadataConfig(ctx, id).MetadataConfigResource(metadataConfigResource).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	whisparrClient "github.com/devopsarr/whisparr-go/whisparr"
)

func main() {
	id := "id_example" // string | 
	metadataConfigResource := *whisparrClient.NewMetadataConfigResource() // MetadataConfigResource |  (optional)

	configuration := whisparrClient.NewConfiguration()
	apiClient := whisparrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataConfigAPI.UpdateMetadataConfig(context.Background(), id).MetadataConfigResource(metadataConfigResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataConfigAPI.UpdateMetadataConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMetadataConfig`: MetadataConfigResource
	fmt.Fprintf(os.Stdout, "Response from `MetadataConfigAPI.UpdateMetadataConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetadataConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataConfigResource** | [**MetadataConfigResource**](MetadataConfigResource.md) |  | 

### Return type

[**MetadataConfigResource**](MetadataConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

