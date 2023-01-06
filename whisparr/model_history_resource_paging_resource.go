/*
Whisparr

Whisparr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whisparr

import (
	"encoding/json"
)

// HistoryResourcePagingResource struct for HistoryResourcePagingResource
type HistoryResourcePagingResource struct {
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	SortKey NullableString `json:"sortKey,omitempty"`
	SortDirection *SortDirection `json:"sortDirection,omitempty"`
	Filters []*PagingResourceFilter `json:"filters,omitempty"`
	TotalRecords *int32 `json:"totalRecords,omitempty"`
	Records []*HistoryResource `json:"records,omitempty"`
}

// NewHistoryResourcePagingResource instantiates a new HistoryResourcePagingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryResourcePagingResource() *HistoryResourcePagingResource {
	this := HistoryResourcePagingResource{}
	return &this
}

// NewHistoryResourcePagingResourceWithDefaults instantiates a new HistoryResourcePagingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryResourcePagingResourceWithDefaults() *HistoryResourcePagingResource {
	this := HistoryResourcePagingResource{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *HistoryResourcePagingResource) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResourcePagingResource) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *HistoryResourcePagingResource) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *HistoryResourcePagingResource) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *HistoryResourcePagingResource) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResourcePagingResource) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *HistoryResourcePagingResource) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *HistoryResourcePagingResource) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResourcePagingResource) GetSortKey() string {
	if o == nil || isNil(o.SortKey.Get()) {
		var ret string
		return ret
	}
	return *o.SortKey.Get()
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResourcePagingResource) GetSortKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SortKey.Get(), o.SortKey.IsSet()
}

// HasSortKey returns a boolean if a field has been set.
func (o *HistoryResourcePagingResource) HasSortKey() bool {
	if o != nil && o.SortKey.IsSet() {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given NullableString and assigns it to the SortKey field.
func (o *HistoryResourcePagingResource) SetSortKey(v string) {
	o.SortKey.Set(&v)
}
// SetSortKeyNil sets the value for SortKey to be an explicit nil
func (o *HistoryResourcePagingResource) SetSortKeyNil() {
	o.SortKey.Set(nil)
}

// UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
func (o *HistoryResourcePagingResource) UnsetSortKey() {
	o.SortKey.Unset()
}

// GetSortDirection returns the SortDirection field value if set, zero value otherwise.
func (o *HistoryResourcePagingResource) GetSortDirection() SortDirection {
	if o == nil || isNil(o.SortDirection) {
		var ret SortDirection
		return ret
	}
	return *o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool) {
	if o == nil || isNil(o.SortDirection) {
    return nil, false
	}
	return o.SortDirection, true
}

// HasSortDirection returns a boolean if a field has been set.
func (o *HistoryResourcePagingResource) HasSortDirection() bool {
	if o != nil && !isNil(o.SortDirection) {
		return true
	}

	return false
}

// SetSortDirection gets a reference to the given SortDirection and assigns it to the SortDirection field.
func (o *HistoryResourcePagingResource) SetSortDirection(v SortDirection) {
	o.SortDirection = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResourcePagingResource) GetFilters() []*PagingResourceFilter {
	if o == nil {
		var ret []*PagingResourceFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResourcePagingResource) GetFiltersOk() ([]*PagingResourceFilter, bool) {
	if o == nil || isNil(o.Filters) {
    return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *HistoryResourcePagingResource) HasFilters() bool {
	if o != nil && isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []PagingResourceFilter and assigns it to the Filters field.
func (o *HistoryResourcePagingResource) SetFilters(v []*PagingResourceFilter) {
	o.Filters = v
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *HistoryResourcePagingResource) GetTotalRecords() int32 {
	if o == nil || isNil(o.TotalRecords) {
		var ret int32
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResourcePagingResource) GetTotalRecordsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalRecords) {
    return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *HistoryResourcePagingResource) HasTotalRecords() bool {
	if o != nil && !isNil(o.TotalRecords) {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given int32 and assigns it to the TotalRecords field.
func (o *HistoryResourcePagingResource) SetTotalRecords(v int32) {
	o.TotalRecords = &v
}

// GetRecords returns the Records field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResourcePagingResource) GetRecords() []*HistoryResource {
	if o == nil {
		var ret []*HistoryResource
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResourcePagingResource) GetRecordsOk() ([]*HistoryResource, bool) {
	if o == nil || isNil(o.Records) {
    return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *HistoryResourcePagingResource) HasRecords() bool {
	if o != nil && isNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []HistoryResource and assigns it to the Records field.
func (o *HistoryResourcePagingResource) SetRecords(v []*HistoryResource) {
	o.Records = v
}

func (o HistoryResourcePagingResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.SortKey.IsSet() {
		toSerialize["sortKey"] = o.SortKey.Get()
	}
	if !isNil(o.SortDirection) {
		toSerialize["sortDirection"] = o.SortDirection
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if !isNil(o.TotalRecords) {
		toSerialize["totalRecords"] = o.TotalRecords
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableHistoryResourcePagingResource struct {
	value *HistoryResourcePagingResource
	isSet bool
}

func (v NullableHistoryResourcePagingResource) Get() *HistoryResourcePagingResource {
	return v.value
}

func (v *NullableHistoryResourcePagingResource) Set(val *HistoryResourcePagingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryResourcePagingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryResourcePagingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryResourcePagingResource(val *HistoryResourcePagingResource) *NullableHistoryResourcePagingResource {
	return &NullableHistoryResourcePagingResource{value: val, isSet: true}
}

func (v NullableHistoryResourcePagingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryResourcePagingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


