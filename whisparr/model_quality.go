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

// Quality struct for Quality
type Quality struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Source *QualitySource `json:"source,omitempty"`
	Resolution *int32 `json:"resolution,omitempty"`
	Modifier *Modifier `json:"modifier,omitempty"`
}

// NewQuality instantiates a new Quality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuality() *Quality {
	this := Quality{}
	return &this
}

// NewQualityWithDefaults instantiates a new Quality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityWithDefaults() *Quality {
	this := Quality{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Quality) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quality) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Quality) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Quality) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Quality) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Quality) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Quality) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Quality) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Quality) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Quality) UnsetName() {
	o.Name.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Quality) GetSource() QualitySource {
	if o == nil || isNil(o.Source) {
		var ret QualitySource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quality) GetSourceOk() (*QualitySource, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Quality) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given QualitySource and assigns it to the Source field.
func (o *Quality) SetSource(v QualitySource) {
	o.Source = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *Quality) GetResolution() int32 {
	if o == nil || isNil(o.Resolution) {
		var ret int32
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quality) GetResolutionOk() (*int32, bool) {
	if o == nil || isNil(o.Resolution) {
    return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *Quality) HasResolution() bool {
	if o != nil && !isNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given int32 and assigns it to the Resolution field.
func (o *Quality) SetResolution(v int32) {
	o.Resolution = &v
}

// GetModifier returns the Modifier field value if set, zero value otherwise.
func (o *Quality) GetModifier() Modifier {
	if o == nil || isNil(o.Modifier) {
		var ret Modifier
		return ret
	}
	return *o.Modifier
}

// GetModifierOk returns a tuple with the Modifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quality) GetModifierOk() (*Modifier, bool) {
	if o == nil || isNil(o.Modifier) {
    return nil, false
	}
	return o.Modifier, true
}

// HasModifier returns a boolean if a field has been set.
func (o *Quality) HasModifier() bool {
	if o != nil && !isNil(o.Modifier) {
		return true
	}

	return false
}

// SetModifier gets a reference to the given Modifier and assigns it to the Modifier field.
func (o *Quality) SetModifier(v Modifier) {
	o.Modifier = &v
}

func (o Quality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !isNil(o.Modifier) {
		toSerialize["modifier"] = o.Modifier
	}
	return json.Marshal(toSerialize)
}

type NullableQuality struct {
	value *Quality
	isSet bool
}

func (v NullableQuality) Get() *Quality {
	return v.value
}

func (v *NullableQuality) Set(val *Quality) {
	v.value = val
	v.isSet = true
}

func (v NullableQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuality(val *Quality) *NullableQuality {
	return &NullableQuality{value: val, isSet: true}
}

func (v NullableQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


