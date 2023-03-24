# \ImportListMoviesApi

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImportlistMovie**](ImportListMoviesApi.md#CreateImportlistMovie) | **Post** /api/v3/importlist/movie | 
[**GetImportlistMovie**](ImportListMoviesApi.md#GetImportlistMovie) | **Get** /api/v3/importlist/movie | 



## CreateImportlistMovie

> CreateImportlistMovie(ctx).MovieResource(movieResource).Execute()



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
    movieResource := []whisparrClient.MovieResource{*whisparrClient.NewMovieResource()} // []MovieResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListMoviesApi.CreateImportlistMovie(context.Background()).MovieResource(movieResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListMoviesApi.CreateImportlistMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImportlistMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieResource** | [**[]MovieResource**](MovieResource.md) |  | 

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


## GetImportlistMovie

> GetImportlistMovie(ctx).IncludeRecommendations(includeRecommendations).Execute()



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
    includeRecommendations := true // bool |  (optional) (default to false)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListMoviesApi.GetImportlistMovie(context.Background()).IncludeRecommendations(includeRecommendations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListMoviesApi.GetImportlistMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImportlistMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeRecommendations** | **bool** |  | [default to false]

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

