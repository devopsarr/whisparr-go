# \QualityProfileSchemaApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQualityprofileSchema**](QualityProfileSchemaApi.md#GetQualityprofileSchema) | **Get** /api/v3/qualityprofile/schema | 



## GetQualityprofileSchema

> QualityProfileResource GetQualityprofileSchema(ctx).Execute()



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
    resp, r, err := apiClient.QualityProfileSchemaApi.GetQualityprofileSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileSchemaApi.GetQualityprofileSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQualityprofileSchema`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileSchemaApi.GetQualityprofileSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQualityprofileSchemaRequest struct via the builder pattern


### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

