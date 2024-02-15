# \MediaCoverAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMediaCoverByFilename**](MediaCoverAPI.md#GetMediaCoverByFilename) | **Get** /api/v3/mediacover/{movieId}/{filename} | 



## GetMediaCoverByFilename

> GetMediaCoverByFilename(ctx, movieId, filename).Execute()



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
	movieId := int32(56) // int32 | 
	filename := "filename_example" // string | 

	configuration := whisparrClient.NewConfiguration()
	apiClient := whisparrClient.NewAPIClient(configuration)
	r, err := apiClient.MediaCoverAPI.GetMediaCoverByFilename(context.Background(), movieId, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverAPI.GetMediaCoverByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **int32** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaCoverByFilenameRequest struct via the builder pattern


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

