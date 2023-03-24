/*
Radarr

Radarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whisparr

import (
	"encoding/json"
)

// IndexerFlagResource struct for IndexerFlagResource
type IndexerFlagResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	NameLower NullableString `json:"nameLower,omitempty"`
}

// NewIndexerFlagResource instantiates a new IndexerFlagResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerFlagResource() *IndexerFlagResource {
	this := IndexerFlagResource{}
	return &this
}

// NewIndexerFlagResourceWithDefaults instantiates a new IndexerFlagResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerFlagResourceWithDefaults() *IndexerFlagResource {
	this := IndexerFlagResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndexerFlagResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerFlagResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndexerFlagResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IndexerFlagResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerFlagResource) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerFlagResource) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IndexerFlagResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IndexerFlagResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IndexerFlagResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IndexerFlagResource) UnsetName() {
	o.Name.Unset()
}

// GetNameLower returns the NameLower field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerFlagResource) GetNameLower() string {
	if o == nil || isNil(o.NameLower.Get()) {
		var ret string
		return ret
	}
	return *o.NameLower.Get()
}

// GetNameLowerOk returns a tuple with the NameLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerFlagResource) GetNameLowerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NameLower.Get(), o.NameLower.IsSet()
}

// HasNameLower returns a boolean if a field has been set.
func (o *IndexerFlagResource) HasNameLower() bool {
	if o != nil && o.NameLower.IsSet() {
		return true
	}

	return false
}

// SetNameLower gets a reference to the given NullableString and assigns it to the NameLower field.
func (o *IndexerFlagResource) SetNameLower(v string) {
	o.NameLower.Set(&v)
}
// SetNameLowerNil sets the value for NameLower to be an explicit nil
func (o *IndexerFlagResource) SetNameLowerNil() {
	o.NameLower.Set(nil)
}

// UnsetNameLower ensures that no value is present for NameLower, not even an explicit nil
func (o *IndexerFlagResource) UnsetNameLower() {
	o.NameLower.Unset()
}

func (o IndexerFlagResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.NameLower.IsSet() {
		toSerialize["nameLower"] = o.NameLower.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIndexerFlagResource struct {
	value *IndexerFlagResource
	isSet bool
}

func (v NullableIndexerFlagResource) Get() *IndexerFlagResource {
	return v.value
}

func (v *NullableIndexerFlagResource) Set(val *IndexerFlagResource) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerFlagResource) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerFlagResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerFlagResource(val *IndexerFlagResource) *NullableIndexerFlagResource {
	return &NullableIndexerFlagResource{value: val, isSet: true}
}

func (v NullableIndexerFlagResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerFlagResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


