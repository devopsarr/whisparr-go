# \ParseApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParse**](ParseApi.md#GetParse) | **Get** /api/v3/parse | 



## GetParse

> ParseResource GetParse(ctx).Title(title).Path(path).Execute()



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
    title := "title_example" // string |  (optional)
    path := "path_example" // string |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParseApi.GetParse(context.Background()).Title(title).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParseApi.GetParse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetParse`: ParseResource
    fmt.Fprintf(os.Stdout, "Response from `ParseApi.GetParse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetParseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** |  | 
 **path** | **string** |  | 

### Return type

[**ParseResource**](ParseResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

