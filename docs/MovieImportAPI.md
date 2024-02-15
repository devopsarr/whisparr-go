# \MovieImportAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMovieImport**](MovieImportAPI.md#CreateMovieImport) | **Post** /api/v3/movie/import | 
[**GetMovieImportById**](MovieImportAPI.md#GetMovieImportById) | **Get** /api/v3/movie/import/{id} | 



## CreateMovieImport

> CreateMovieImport(ctx).MovieResource(movieResource).Execute()



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
	movieResource := []whisparrClient.MovieResource{*whisparrClient.NewMovieResource()} // []MovieResource |  (optional)

	configuration := whisparrClient.NewConfiguration()
	apiClient := whisparrClient.NewAPIClient(configuration)
	r, err := apiClient.MovieImportAPI.CreateMovieImport(context.Background()).MovieResource(movieResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieImportAPI.CreateMovieImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMovieImportRequest struct via the builder pattern


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


## GetMovieImportById

> MovieResource GetMovieImportById(ctx, id).Execute()



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
	resp, r, err := apiClient.MovieImportAPI.GetMovieImportById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieImportAPI.GetMovieImportById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieImportById`: MovieResource
	fmt.Fprintf(os.Stdout, "Response from `MovieImportAPI.GetMovieImportById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieImportByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MovieResource**](MovieResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

