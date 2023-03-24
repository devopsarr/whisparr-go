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

// HealthResource struct for HealthResource
type HealthResource struct {
	Id *int32 `json:"id,omitempty"`
	Source NullableString `json:"source,omitempty"`
	Type *HealthCheckResult `json:"type,omitempty"`
	Message NullableString `json:"message,omitempty"`
	WikiUrl *string `json:"wikiUrl,omitempty"`
}

// NewHealthResource instantiates a new HealthResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthResource() *HealthResource {
	this := HealthResource{}
	return &this
}

// NewHealthResourceWithDefaults instantiates a new HealthResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthResourceWithDefaults() *HealthResource {
	this := HealthResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HealthResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HealthResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *HealthResource) SetId(v int32) {
	o.Id = &v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthResource) GetSource() string {
	if o == nil || isNil(o.Source.Get()) {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthResource) GetSourceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *HealthResource) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *HealthResource) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *HealthResource) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *HealthResource) UnsetSource() {
	o.Source.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HealthResource) GetType() HealthCheckResult {
	if o == nil || isNil(o.Type) {
		var ret HealthCheckResult
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthResource) GetTypeOk() (*HealthCheckResult, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HealthResource) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given HealthCheckResult and assigns it to the Type field.
func (o *HealthResource) SetType(v HealthCheckResult) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthResource) GetMessage() string {
	if o == nil || isNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthResource) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *HealthResource) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *HealthResource) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *HealthResource) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *HealthResource) UnsetMessage() {
	o.Message.Unset()
}

// GetWikiUrl returns the WikiUrl field value if set, zero value otherwise.
func (o *HealthResource) GetWikiUrl() string {
	if o == nil || isNil(o.WikiUrl) {
		var ret string
		return ret
	}
	return *o.WikiUrl
}

// GetWikiUrlOk returns a tuple with the WikiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthResource) GetWikiUrlOk() (*string, bool) {
	if o == nil || isNil(o.WikiUrl) {
    return nil, false
	}
	return o.WikiUrl, true
}

// HasWikiUrl returns a boolean if a field has been set.
func (o *HealthResource) HasWikiUrl() bool {
	if o != nil && !isNil(o.WikiUrl) {
		return true
	}

	return false
}

// SetWikiUrl gets a reference to the given string and assigns it to the WikiUrl field.
func (o *HealthResource) SetWikiUrl(v string) {
	o.WikiUrl = &v
}

func (o HealthResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !isNil(o.WikiUrl) {
		toSerialize["wikiUrl"] = o.WikiUrl
	}
	return json.Marshal(toSerialize)
}

type NullableHealthResource struct {
	value *HealthResource
	isSet bool
}

func (v NullableHealthResource) Get() *HealthResource {
	return v.value
}

func (v *NullableHealthResource) Set(val *HealthResource) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthResource) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthResource(val *HealthResource) *NullableHealthResource {
	return &NullableHealthResource{value: val, isSet: true}
}

func (v NullableHealthResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


