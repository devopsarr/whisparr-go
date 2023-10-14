# \NamingConfigAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNamingConfig**](NamingConfigAPI.md#GetNamingConfig) | **Get** /api/v3/config/naming | 
[**GetNamingConfigById**](NamingConfigAPI.md#GetNamingConfigById) | **Get** /api/v3/config/naming/{id} | 
[**GetNamingConfigExamples**](NamingConfigAPI.md#GetNamingConfigExamples) | **Get** /api/v3/config/naming/examples | 
[**UpdateNamingConfig**](NamingConfigAPI.md#UpdateNamingConfig) | **Put** /api/v3/config/naming/{id} | 



## GetNamingConfig

> NamingConfigResource GetNamingConfig(ctx).Execute()



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
    resp, r, err := apiClient.NamingConfigAPI.GetNamingConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigAPI.GetNamingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamingConfig`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigAPI.GetNamingConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingConfigRequest struct via the builder pattern


### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingConfigById

> NamingConfigResource GetNamingConfigById(ctx, id).Execute()



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
    resp, r, err := apiClient.NamingConfigAPI.GetNamingConfigById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigAPI.GetNamingConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamingConfigById`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigAPI.GetNamingConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingConfigExamples

> GetNamingConfigExamples(ctx).RenameMovies(renameMovies).ReplaceIllegalCharacters(replaceIllegalCharacters).ColonReplacementFormat(colonReplacementFormat).StandardMovieFormat(standardMovieFormat).MovieFolderFormat(movieFolderFormat).IncludeQuality(includeQuality).ReplaceSpaces(replaceSpaces).Separator(separator).NumberStyle(numberStyle).Id(id).ResourceName(resourceName).Execute()



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
    renameMovies := true // bool |  (optional)
    replaceIllegalCharacters := true // bool |  (optional)
    colonReplacementFormat := whisparrClient.ColonReplacementFormat("delete") // ColonReplacementFormat |  (optional)
    standardMovieFormat := "standardMovieFormat_example" // string |  (optional)
    movieFolderFormat := "movieFolderFormat_example" // string |  (optional)
    includeQuality := true // bool |  (optional)
    replaceSpaces := true // bool |  (optional)
    separator := "separator_example" // string |  (optional)
    numberStyle := "numberStyle_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    resourceName := "resourceName_example" // string |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamingConfigAPI.GetNamingConfigExamples(context.Background()).RenameMovies(renameMovies).ReplaceIllegalCharacters(replaceIllegalCharacters).ColonReplacementFormat(colonReplacementFormat).StandardMovieFormat(standardMovieFormat).MovieFolderFormat(movieFolderFormat).IncludeQuality(includeQuality).ReplaceSpaces(replaceSpaces).Separator(separator).NumberStyle(numberStyle).Id(id).ResourceName(resourceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigAPI.GetNamingConfigExamples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingConfigExamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameMovies** | **bool** |  | 
 **replaceIllegalCharacters** | **bool** |  | 
 **colonReplacementFormat** | [**ColonReplacementFormat**](ColonReplacementFormat.md) |  | 
 **standardMovieFormat** | **string** |  | 
 **movieFolderFormat** | **string** |  | 
 **includeQuality** | **bool** |  | 
 **replaceSpaces** | **bool** |  | 
 **separator** | **string** |  | 
 **numberStyle** | **string** |  | 
 **id** | **int32** |  | 
 **resourceName** | **string** |  | 

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


## UpdateNamingConfig

> NamingConfigResource UpdateNamingConfig(ctx, id).NamingConfigResource(namingConfigResource).Execute()



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
    namingConfigResource := *whisparrClient.NewNamingConfigResource() // NamingConfigResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamingConfigAPI.UpdateNamingConfig(context.Background(), id).NamingConfigResource(namingConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigAPI.UpdateNamingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNamingConfig`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigAPI.UpdateNamingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namingConfigResource** | [**NamingConfigResource**](NamingConfigResource.md) |  | 

### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

