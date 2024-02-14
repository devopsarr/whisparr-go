/*
Radarr

Radarr API docs

API version: b08981dee068e1ed23e4f45a0d8fe70ef7bf7703
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whisparr

import (
	"encoding/json"
)

// AlternativeTitleResource struct for AlternativeTitleResource
type AlternativeTitleResource struct {
	Id *int32 `json:"id,omitempty"`
	SourceType *SourceType `json:"sourceType,omitempty"`
	MovieMetadataId *int32 `json:"movieMetadataId,omitempty"`
	Title NullableString `json:"title,omitempty"`
	CleanTitle NullableString `json:"cleanTitle,omitempty"`
}

// NewAlternativeTitleResource instantiates a new AlternativeTitleResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeTitleResource() *AlternativeTitleResource {
	this := AlternativeTitleResource{}
	return &this
}

// NewAlternativeTitleResourceWithDefaults instantiates a new AlternativeTitleResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeTitleResourceWithDefaults() *AlternativeTitleResource {
	this := AlternativeTitleResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlternativeTitleResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitleResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlternativeTitleResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlternativeTitleResource) SetId(v int32) {
	o.Id = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *AlternativeTitleResource) GetSourceType() SourceType {
	if o == nil || isNil(o.SourceType) {
		var ret SourceType
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitleResource) GetSourceTypeOk() (*SourceType, bool) {
	if o == nil || isNil(o.SourceType) {
    return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *AlternativeTitleResource) HasSourceType() bool {
	if o != nil && !isNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given SourceType and assigns it to the SourceType field.
func (o *AlternativeTitleResource) SetSourceType(v SourceType) {
	o.SourceType = &v
}

// GetMovieMetadataId returns the MovieMetadataId field value if set, zero value otherwise.
func (o *AlternativeTitleResource) GetMovieMetadataId() int32 {
	if o == nil || isNil(o.MovieMetadataId) {
		var ret int32
		return ret
	}
	return *o.MovieMetadataId
}

// GetMovieMetadataIdOk returns a tuple with the MovieMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitleResource) GetMovieMetadataIdOk() (*int32, bool) {
	if o == nil || isNil(o.MovieMetadataId) {
    return nil, false
	}
	return o.MovieMetadataId, true
}

// HasMovieMetadataId returns a boolean if a field has been set.
func (o *AlternativeTitleResource) HasMovieMetadataId() bool {
	if o != nil && !isNil(o.MovieMetadataId) {
		return true
	}

	return false
}

// SetMovieMetadataId gets a reference to the given int32 and assigns it to the MovieMetadataId field.
func (o *AlternativeTitleResource) SetMovieMetadataId(v int32) {
	o.MovieMetadataId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlternativeTitleResource) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlternativeTitleResource) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AlternativeTitleResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AlternativeTitleResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AlternativeTitleResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AlternativeTitleResource) UnsetTitle() {
	o.Title.Unset()
}

// GetCleanTitle returns the CleanTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlternativeTitleResource) GetCleanTitle() string {
	if o == nil || isNil(o.CleanTitle.Get()) {
		var ret string
		return ret
	}
	return *o.CleanTitle.Get()
}

// GetCleanTitleOk returns a tuple with the CleanTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlternativeTitleResource) GetCleanTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CleanTitle.Get(), o.CleanTitle.IsSet()
}

// HasCleanTitle returns a boolean if a field has been set.
func (o *AlternativeTitleResource) HasCleanTitle() bool {
	if o != nil && o.CleanTitle.IsSet() {
		return true
	}

	return false
}

// SetCleanTitle gets a reference to the given NullableString and assigns it to the CleanTitle field.
func (o *AlternativeTitleResource) SetCleanTitle(v string) {
	o.CleanTitle.Set(&v)
}
// SetCleanTitleNil sets the value for CleanTitle to be an explicit nil
func (o *AlternativeTitleResource) SetCleanTitleNil() {
	o.CleanTitle.Set(nil)
}

// UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
func (o *AlternativeTitleResource) UnsetCleanTitle() {
	o.CleanTitle.Unset()
}

func (o AlternativeTitleResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !isNil(o.MovieMetadataId) {
		toSerialize["movieMetadataId"] = o.MovieMetadataId
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.CleanTitle.IsSet() {
		toSerialize["cleanTitle"] = o.CleanTitle.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAlternativeTitleResource struct {
	value *AlternativeTitleResource
	isSet bool
}

func (v NullableAlternativeTitleResource) Get() *AlternativeTitleResource {
	return v.value
}

func (v *NullableAlternativeTitleResource) Set(val *AlternativeTitleResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeTitleResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeTitleResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeTitleResource(val *AlternativeTitleResource) *NullableAlternativeTitleResource {
	return &NullableAlternativeTitleResource{value: val, isSet: true}
}

func (v NullableAlternativeTitleResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeTitleResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


