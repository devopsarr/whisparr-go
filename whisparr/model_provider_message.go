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

// ProviderMessage struct for ProviderMessage
type ProviderMessage struct {
	Message NullableString `json:"message,omitempty"`
	Type *ProviderMessageType `json:"type,omitempty"`
}

// NewProviderMessage instantiates a new ProviderMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderMessage() *ProviderMessage {
	this := ProviderMessage{}
	return &this
}

// NewProviderMessageWithDefaults instantiates a new ProviderMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderMessageWithDefaults() *ProviderMessage {
	this := ProviderMessage{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderMessage) GetMessage() string {
	if o == nil || isNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderMessage) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ProviderMessage) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ProviderMessage) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ProviderMessage) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ProviderMessage) UnsetMessage() {
	o.Message.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProviderMessage) GetType() ProviderMessageType {
	if o == nil || isNil(o.Type) {
		var ret ProviderMessageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderMessage) GetTypeOk() (*ProviderMessageType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProviderMessage) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProviderMessageType and assigns it to the Type field.
func (o *ProviderMessage) SetType(v ProviderMessageType) {
	o.Type = &v
}

func (o ProviderMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableProviderMessage struct {
	value *ProviderMessage
	isSet bool
}

func (v NullableProviderMessage) Get() *ProviderMessage {
	return v.value
}

func (v *NullableProviderMessage) Set(val *ProviderMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderMessage(val *ProviderMessage) *NullableProviderMessage {
	return &NullableProviderMessage{value: val, isSet: true}
}

func (v NullableProviderMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


