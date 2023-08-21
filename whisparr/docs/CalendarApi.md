# \CalendarApi

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCalendarById**](CalendarApi.md#GetCalendarById) | **Get** /api/v3/calendar/{id} | 
[**ListCalendar**](CalendarApi.md#ListCalendar) | **Get** /api/v3/calendar | 



## GetCalendarById

> MovieResource GetCalendarById(ctx, id).Execute()



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
    resp, r, err := apiClient.CalendarApi.GetCalendarById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.GetCalendarById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCalendarById`: MovieResource
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.GetCalendarById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCalendarByIdRequest struct via the builder pattern


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


## ListCalendar

> []MovieResource ListCalendar(ctx).Start(start).End(end).Unmonitored(unmonitored).Tags(tags).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    whisparrClient "./openapi"
)

func main() {
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    unmonitored := true // bool |  (optional) (default to false)
    tags := "tags_example" // string |  (optional) (default to "")

    configuration := whisparrClient.NewConfiguration()
    apiClient := whisparrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalendarApi.ListCalendar(context.Background()).Start(start).End(end).Unmonitored(unmonitored).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.ListCalendar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCalendar`: []MovieResource
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.ListCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 
 **unmonitored** | **bool** |  | [default to false]
 **tags** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]MovieResource**](MovieResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

