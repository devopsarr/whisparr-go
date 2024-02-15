# \BlocklistAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBlocklist**](BlocklistAPI.md#DeleteBlocklist) | **Delete** /api/v3/blocklist/{id} | 
[**DeleteBlocklistBulk**](BlocklistAPI.md#DeleteBlocklistBulk) | **Delete** /api/v3/blocklist/bulk | 
[**GetBlocklist**](BlocklistAPI.md#GetBlocklist) | **Get** /api/v3/blocklist | 
[**ListBlocklistMovie**](BlocklistAPI.md#ListBlocklistMovie) | **Get** /api/v3/blocklist/movie | 



## DeleteBlocklist

> DeleteBlocklist(ctx, id).Execute()



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
	r, err := apiClient.BlocklistAPI.DeleteBlocklist(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.DeleteBlocklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlocklistRequest struct via the builder pattern


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


## DeleteBlocklistBulk

> DeleteBlocklistBulk(ctx).BlocklistBulkResource(blocklistBulkResource).Execute()



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
	blocklistBulkResource := *whisparrClient.NewBlocklistBulkResource() // BlocklistBulkResource |  (optional)

	configuration := whisparrClient.NewConfiguration()
	apiClient := whisparrClient.NewAPIClient(configuration)
	r, err := apiClient.BlocklistAPI.DeleteBlocklistBulk(context.Background()).BlocklistBulkResource(blocklistBulkResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.DeleteBlocklistBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlocklistBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocklistBulkResource** | [**BlocklistBulkResource**](BlocklistBulkResource.md) |  | 

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


## GetBlocklist

> BlocklistResourcePagingResource GetBlocklist(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).Execute()



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	sortKey := "sortKey_example" // string |  (optional)
	sortDirection := whisparrClient.SortDirection("default") // SortDirection |  (optional)

	configuration := whisparrClient.NewConfiguration()
	apiClient := whisparrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocklistAPI.GetBlocklist(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.GetBlocklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocklist`: BlocklistResourcePagingResource
	fmt.Fprintf(os.Stdout, "Response from `BlocklistAPI.GetBlocklist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 

### Return type

[**BlocklistResourcePagingResource**](BlocklistResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlocklistMovie

> []BlocklistResource ListBlocklistMovie(ctx).MovieId(movieId).Execute()



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
	movieId := int32(56) // int32 |  (optional)

	configuration := whisparrClient.NewConfiguration()
	apiClient := whisparrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocklistAPI.ListBlocklistMovie(context.Background()).MovieId(movieId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.ListBlocklistMovie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlocklistMovie`: []BlocklistResource
	fmt.Fprintf(os.Stdout, "Response from `BlocklistAPI.ListBlocklistMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlocklistMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 

### Return type

[**[]BlocklistResource**](BlocklistResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

