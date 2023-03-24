# \SeriesImportApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSeriesImport**](SeriesImportApi.md#CreateSeriesImport) | **Post** /api/v3/series/import | 



## CreateSeriesImport

> CreateSeriesImport(ctx).SeriesResource(seriesResource).Execute()



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
    seriesResource := []whisparrClient.SeriesResource{*whisparrClient.NewSeriesResource()} // []SeriesResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesImportApi.CreateSeriesImport(context.Background()).SeriesResource(seriesResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesImportApi.CreateSeriesImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeriesImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesResource** | [**[]SeriesResource**](SeriesResource.md) |  | 

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

