# \LanguageProfileApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLanguageProfile**](LanguageProfileApi.md#CreateLanguageProfile) | **Post** /api/v3/languageprofile | 
[**DeleteLanguageProfile**](LanguageProfileApi.md#DeleteLanguageProfile) | **Delete** /api/v3/languageprofile/{id} | 
[**GetLanguageProfileById**](LanguageProfileApi.md#GetLanguageProfileById) | **Get** /api/v3/languageprofile/{id} | 
[**ListLanguageProfile**](LanguageProfileApi.md#ListLanguageProfile) | **Get** /api/v3/languageprofile | 
[**UpdateLanguageProfile**](LanguageProfileApi.md#UpdateLanguageProfile) | **Put** /api/v3/languageprofile/{id} | 



## CreateLanguageProfile

> LanguageProfileResource CreateLanguageProfile(ctx).LanguageProfileResource(languageProfileResource).Execute()



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
    languageProfileResource := *whisparrClient.NewLanguageProfileResource() // LanguageProfileResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageProfileApi.CreateLanguageProfile(context.Background()).LanguageProfileResource(languageProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.CreateLanguageProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanguageProfile`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.CreateLanguageProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanguageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **languageProfileResource** | [**LanguageProfileResource**](LanguageProfileResource.md) |  | 

### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanguageProfile

> DeleteLanguageProfile(ctx, id).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.DeleteLanguageProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.DeleteLanguageProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLanguageProfileRequest struct via the builder pattern


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


## GetLanguageProfileById

> LanguageProfileResource GetLanguageProfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.GetLanguageProfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.GetLanguageProfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanguageProfileById`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.GetLanguageProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguageProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLanguageProfile

> []LanguageProfileResource ListLanguageProfile(ctx).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.ListLanguageProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.ListLanguageProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLanguageProfile`: []LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.ListLanguageProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLanguageProfileRequest struct via the builder pattern


### Return type

[**[]LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLanguageProfile

> LanguageProfileResource UpdateLanguageProfile(ctx, id).LanguageProfileResource(languageProfileResource).Execute()



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
    languageProfileResource := *whisparrClient.NewLanguageProfileResource() // LanguageProfileResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageProfileApi.UpdateLanguageProfile(context.Background(), id).LanguageProfileResource(languageProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.UpdateLanguageProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLanguageProfile`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.UpdateLanguageProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLanguageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **languageProfileResource** | [**LanguageProfileResource**](LanguageProfileResource.md) |  | 

### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

