# \MediaManagementConfigApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMediaManagementConfig**](MediaManagementConfigApi.md#GetMediaManagementConfig) | **Get** /api/v3/config/mediamanagement | 
[**GetMediaManagementConfigById**](MediaManagementConfigApi.md#GetMediaManagementConfigById) | **Get** /api/v3/config/mediamanagement/{id} | 
[**UpdateMediaManagementConfig**](MediaManagementConfigApi.md#UpdateMediaManagementConfig) | **Put** /api/v3/config/mediamanagement/{id} | 



## GetMediaManagementConfig

> MediaManagementConfigResource GetMediaManagementConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    whisparrClient "./openapi"
)

func main() {

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaManagementConfigApi.GetMediaManagementConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.GetMediaManagementConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMediaManagementConfig`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.GetMediaManagementConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaManagementConfigRequest struct via the builder pattern


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaManagementConfigById

> MediaManagementConfigResource GetMediaManagementConfigById(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    whisparrClient "./openapi"
)

func main() {
    id := int32(56) // int32 | 

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaManagementConfigApi.GetMediaManagementConfigById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.GetMediaManagementConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMediaManagementConfigById`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.GetMediaManagementConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaManagementConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMediaManagementConfig

> MediaManagementConfigResource UpdateMediaManagementConfig(ctx, id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    whisparrClient "./openapi"
)

func main() {
    id := "id_example" // string | 
    mediaManagementConfigResource := *whisparrClient.NewMediaManagementConfigResource() // MediaManagementConfigResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaManagementConfigApi.UpdateMediaManagementConfig(context.Background(), id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.UpdateMediaManagementConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMediaManagementConfig`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.UpdateMediaManagementConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMediaManagementConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaManagementConfigResource** | [**MediaManagementConfigResource**](MediaManagementConfigResource.md) |  | 

### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

