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

// Rejection struct for Rejection
type Rejection struct {
	Reason NullableString `json:"reason,omitempty"`
	Type *RejectionType `json:"type,omitempty"`
}

// NewRejection instantiates a new Rejection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejection() *Rejection {
	this := Rejection{}
	return &this
}

// NewRejectionWithDefaults instantiates a new Rejection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectionWithDefaults() *Rejection {
	this := Rejection{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rejection) GetReason() string {
	if o == nil || isNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rejection) GetReasonOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *Rejection) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *Rejection) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *Rejection) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *Rejection) UnsetReason() {
	o.Reason.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Rejection) GetType() RejectionType {
	if o == nil || isNil(o.Type) {
		var ret RejectionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rejection) GetTypeOk() (*RejectionType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Rejection) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RejectionType and assigns it to the Type field.
func (o *Rejection) SetType(v RejectionType) {
	o.Type = &v
}

func (o Rejection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRejection struct {
	value *Rejection
	isSet bool
}

func (v NullableRejection) Get() *Rejection {
	return v.value
}

func (v *NullableRejection) Set(val *Rejection) {
	v.value = val
	v.isSet = true
}

func (v NullableRejection) IsSet() bool {
	return v.isSet
}

func (v *NullableRejection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejection(val *Rejection) *NullableRejection {
	return &NullableRejection{value: val, isSet: true}
}

func (v NullableRejection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


