# \ImportListConfigApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigImportlist**](ImportListConfigApi.md#GetConfigImportlist) | **Get** /api/v3/config/importlist | 
[**GetConfigImportlistById**](ImportListConfigApi.md#GetConfigImportlistById) | **Get** /api/v3/config/importlist/{id} | 
[**UpdateConfigImportlist**](ImportListConfigApi.md#UpdateConfigImportlist) | **Put** /api/v3/config/importlist/{id} | 



## GetConfigImportlist

> ImportListConfigResource GetConfigImportlist(ctx).Execute()



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
    resp, r, err := apiClient.ImportListConfigApi.GetConfigImportlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListConfigApi.GetConfigImportlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigImportlist`: ImportListConfigResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListConfigApi.GetConfigImportlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigImportlistRequest struct via the builder pattern


### Return type

[**ImportListConfigResource**](ImportListConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigImportlistById

> ImportListConfigResource GetConfigImportlistById(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListConfigApi.GetConfigImportlistById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListConfigApi.GetConfigImportlistById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigImportlistById`: ImportListConfigResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListConfigApi.GetConfigImportlistById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigImportlistByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportListConfigResource**](ImportListConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigImportlist

> ImportListConfigResource UpdateConfigImportlist(ctx, id).ImportListConfigResource(importListConfigResource).Execute()



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
    importListConfigResource := *whisparrClient.NewImportListConfigResource() // ImportListConfigResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListConfigApi.UpdateConfigImportlist(context.Background(), id).ImportListConfigResource(importListConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListConfigApi.UpdateConfigImportlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigImportlist`: ImportListConfigResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListConfigApi.UpdateConfigImportlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigImportlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importListConfigResource** | [**ImportListConfigResource**](ImportListConfigResource.md) |  | 

### Return type

[**ImportListConfigResource**](ImportListConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

