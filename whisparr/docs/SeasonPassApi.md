# \SeasonPassApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSeasonPass**](SeasonPassApi.md#CreateSeasonPass) | **Post** /api/v3/seasonpass | 



## CreateSeasonPass

> CreateSeasonPass(ctx).SeasonPassResource(seasonPassResource).Execute()



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
    seasonPassResource := *whisparrClient.NewSeasonPassResource() // SeasonPassResource |  (optional)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeasonPassApi.CreateSeasonPass(context.Background()).SeasonPassResource(seasonPassResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeasonPassApi.CreateSeasonPass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeasonPassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seasonPassResource** | [**SeasonPassResource**](SeasonPassResource.md) |  | 

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

