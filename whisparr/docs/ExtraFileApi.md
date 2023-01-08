# \ExtraFileApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExtraFile**](ExtraFileApi.md#ListExtraFile) | **Get** /api/v3/extrafile | 



## ListExtraFile

> []ExtraFileResource ListExtraFile(ctx).MovieId(movieId).Execute()



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
    movieId := int32(56) // int32 |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtraFileApi.ListExtraFile(context.Background()).MovieId(movieId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtraFileApi.ListExtraFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtraFile`: []ExtraFileResource
    fmt.Fprintf(os.Stdout, "Response from `ExtraFileApi.ListExtraFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtraFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 

### Return type

[**[]ExtraFileResource**](ExtraFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

