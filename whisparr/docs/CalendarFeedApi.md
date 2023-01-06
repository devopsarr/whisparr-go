# \CalendarFeedApi

All URIs are relative to *http://localhost:6969*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedV3CalendarWhisparrIcs**](CalendarFeedApi.md#GetFeedV3CalendarWhisparrIcs) | **Get** /feed/v3/calendar/whisparr.ics | 



## GetFeedV3CalendarWhisparrIcs

> GetFeedV3CalendarWhisparrIcs(ctx).PastDays(pastDays).FutureDays(futureDays).TagList(tagList).Unmonitored(unmonitored).Execute()



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
    pastDays := int32(56) // int32 |  (optional) (default to 7)
    futureDays := int32(56) // int32 |  (optional) (default to 28)
    tagList := "tagList_example" // string |  (optional) (default to "")
    unmonitored := true // bool |  (optional) (default to false)

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalendarFeedApi.GetFeedV3CalendarWhisparrIcs(context.Background()).PastDays(pastDays).FutureDays(futureDays).TagList(tagList).Unmonitored(unmonitored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarFeedApi.GetFeedV3CalendarWhisparrIcs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedV3CalendarWhisparrIcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pastDays** | **int32** |  | [default to 7]
 **futureDays** | **int32** |  | [default to 28]
 **tagList** | **string** |  | [default to &quot;&quot;]
 **unmonitored** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

