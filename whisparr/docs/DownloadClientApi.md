# \DownloadClientApi

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDownloadClient**](DownloadClientApi.md#CreateDownloadClient) | **Post** /api/v3/downloadclient | 
[**CreateDownloadClientActionByName**](DownloadClientApi.md#CreateDownloadClientActionByName) | **Post** /api/v3/downloadclient/action/{name} | 
[**DeleteDownloadClient**](DownloadClientApi.md#DeleteDownloadClient) | **Delete** /api/v3/downloadclient/{id} | 
[**DeleteDownloadClientBulk**](DownloadClientApi.md#DeleteDownloadClientBulk) | **Delete** /api/v3/downloadclient/bulk | 
[**GetDownloadClientById**](DownloadClientApi.md#GetDownloadClientById) | **Get** /api/v3/downloadclient/{id} | 
[**ListDownloadClient**](DownloadClientApi.md#ListDownloadClient) | **Get** /api/v3/downloadclient | 
[**ListDownloadClientSchema**](DownloadClientApi.md#ListDownloadClientSchema) | **Get** /api/v3/downloadclient/schema | 
[**PutDownloadClientBulk**](DownloadClientApi.md#PutDownloadClientBulk) | **Put** /api/v3/downloadclient/bulk | 
[**TestDownloadClient**](DownloadClientApi.md#TestDownloadClient) | **Post** /api/v3/downloadclient/test | 
[**TestallDownloadClient**](DownloadClientApi.md#TestallDownloadClient) | **Post** /api/v3/downloadclient/testall | 
[**UpdateDownloadClient**](DownloadClientApi.md#UpdateDownloadClient) | **Put** /api/v3/downloadclient/{id} | 



## CreateDownloadClient

> DownloadClientResource CreateDownloadClient(ctx).ForceSave(forceSave).DownloadClientResource(downloadClientResource).Execute()



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
    forceSave := true // bool |  (optional) (default to false)
    downloadClientResource := *whisparrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.CreateDownloadClient(context.Background()).ForceSave(forceSave).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.CreateDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDownloadClient`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.CreateDownloadClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceSave** | **bool** |  | [default to false]
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDownloadClientActionByName

> CreateDownloadClientActionByName(ctx, name).DownloadClientResource(downloadClientResource).Execute()



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
    name := "name_example" // string | 
    downloadClientResource := *whisparrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.CreateDownloadClientActionByName(context.Background(), name).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.CreateDownloadClientActionByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownloadClientActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDownloadClient

> DeleteDownloadClient(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.DeleteDownloadClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.DeleteDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDownloadClientBulk

> DeleteDownloadClientBulk(ctx).DownloadClientBulkResource(downloadClientBulkResource).Execute()



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
    downloadClientBulkResource := *whisparrClient.NewDownloadClientBulkResource() // DownloadClientBulkResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.DeleteDownloadClientBulk(context.Background()).DownloadClientBulkResource(downloadClientBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.DeleteDownloadClientBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDownloadClientBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientBulkResource** | [**DownloadClientBulkResource**](DownloadClientBulkResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadClientById

> DownloadClientResource GetDownloadClientById(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.GetDownloadClientById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.GetDownloadClientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadClientById`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.GetDownloadClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDownloadClient

> []DownloadClientResource ListDownloadClient(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ListDownloadClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ListDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDownloadClient`: []DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ListDownloadClient`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDownloadClientRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDownloadClientSchema

> []DownloadClientResource ListDownloadClientSchema(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ListDownloadClientSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ListDownloadClientSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDownloadClientSchema`: []DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ListDownloadClientSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDownloadClientSchemaRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDownloadClientBulk

> DownloadClientResource PutDownloadClientBulk(ctx).DownloadClientBulkResource(downloadClientBulkResource).Execute()



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
    downloadClientBulkResource := *whisparrClient.NewDownloadClientBulkResource() // DownloadClientBulkResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.PutDownloadClientBulk(context.Background()).DownloadClientBulkResource(downloadClientBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.PutDownloadClientBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDownloadClientBulk`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.PutDownloadClientBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDownloadClientBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientBulkResource** | [**DownloadClientBulkResource**](DownloadClientBulkResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestDownloadClient

> TestDownloadClient(ctx).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *whisparrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.TestDownloadClient(context.Background()).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.TestDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestallDownloadClient

> TestallDownloadClient(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.TestallDownloadClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.TestallDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallDownloadClientRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDownloadClient

> DownloadClientResource UpdateDownloadClient(ctx, id).ForceSave(forceSave).DownloadClientResource(downloadClientResource).Execute()



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
    forceSave := true // bool |  (optional) (default to false)
    downloadClientResource := *whisparrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.UpdateDownloadClient(context.Background(), id).ForceSave(forceSave).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.UpdateDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDownloadClient`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.UpdateDownloadClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceSave** | **bool** |  | [default to false]
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

