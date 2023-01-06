# QueueResourcePagingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**SortKey** | Pointer to **NullableString** |  | [optional] 
**SortDirection** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] 
**Filters** | Pointer to [**[]PagingResourceFilter**](PagingResourceFilter.md) |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**Records** | Pointer to [**[]QueueResource**](QueueResource.md) |  | [optional] 

## Methods

### NewQueueResourcePagingResource

`func NewQueueResourcePagingResource() *QueueResourcePagingResource`

NewQueueResourcePagingResource instantiates a new QueueResourcePagingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueResourcePagingResourceWithDefaults

`func NewQueueResourcePagingResourceWithDefaults() *QueueResourcePagingResource`

NewQueueResourcePagingResourceWithDefaults instantiates a new QueueResourcePagingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *QueueResourcePagingResource) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QueueResourcePagingResource) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QueueResourcePagingResource) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *QueueResourcePagingResource) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *QueueResourcePagingResource) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *QueueResourcePagingResource) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *QueueResourcePagingResource) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *QueueResourcePagingResource) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *QueueResourcePagingResource) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *QueueResourcePagingResource) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *QueueResourcePagingResource) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *QueueResourcePagingResource) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### SetSortKeyNil

`func (o *QueueResourcePagingResource) SetSortKeyNil(b bool)`

 SetSortKeyNil sets the value for SortKey to be an explicit nil

### UnsetSortKey
`func (o *QueueResourcePagingResource) UnsetSortKey()`

UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
### GetSortDirection

`func (o *QueueResourcePagingResource) GetSortDirection() SortDirection`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *QueueResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *QueueResourcePagingResource) SetSortDirection(v SortDirection)`

SetSortDirection sets SortDirection field to given value.

### HasSortDirection

`func (o *QueueResourcePagingResource) HasSortDirection() bool`

HasSortDirection returns a boolean if a field has been set.

### GetFilters

`func (o *QueueResourcePagingResource) GetFilters() []PagingResourceFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueueResourcePagingResource) GetFiltersOk() (*[]PagingResourceFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueueResourcePagingResource) SetFilters(v []PagingResourceFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueueResourcePagingResource) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *QueueResourcePagingResource) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *QueueResourcePagingResource) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetTotalRecords

`func (o *QueueResourcePagingResource) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *QueueResourcePagingResource) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *QueueResourcePagingResource) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *QueueResourcePagingResource) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetRecords

`func (o *QueueResourcePagingResource) GetRecords() []QueueResource`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *QueueResourcePagingResource) GetRecordsOk() (*[]QueueResource, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *QueueResourcePagingResource) SetRecords(v []QueueResource)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *QueueResourcePagingResource) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### SetRecordsNil

`func (o *QueueResourcePagingResource) SetRecordsNil(b bool)`

 SetRecordsNil sets the value for Records to be an explicit nil

### UnsetRecords
`func (o *QueueResourcePagingResource) UnsetRecords()`

UnsetRecords ensures that no value is present for Records, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


