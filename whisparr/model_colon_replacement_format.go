/*
Whisparr

Whisparr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whisparr

import (
	"encoding/json"
	"fmt"
)

// ColonReplacementFormat the model 'ColonReplacementFormat'
type ColonReplacementFormat string

// List of ColonReplacementFormat
const (
	COLONREPLACEMENTFORMAT_DELETE ColonReplacementFormat = "delete"
	COLONREPLACEMENTFORMAT_DASH ColonReplacementFormat = "dash"
	COLONREPLACEMENTFORMAT_SPACE_DASH ColonReplacementFormat = "spaceDash"
	COLONREPLACEMENTFORMAT_SPACE_DASH_SPACE ColonReplacementFormat = "spaceDashSpace"
)

// All allowed values of ColonReplacementFormat enum
var AllowedColonReplacementFormatEnumValues = []ColonReplacementFormat{
	"delete",
	"dash",
	"spaceDash",
	"spaceDashSpace",
}

func (v *ColonReplacementFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ColonReplacementFormat(value)
	for _, existing := range AllowedColonReplacementFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ColonReplacementFormat", value)
}

// NewColonReplacementFormatFromValue returns a pointer to a valid ColonReplacementFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewColonReplacementFormatFromValue(v string) (*ColonReplacementFormat, error) {
	ev := ColonReplacementFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ColonReplacementFormat: valid values are %v", v, AllowedColonReplacementFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ColonReplacementFormat) IsValid() bool {
	for _, existing := range AllowedColonReplacementFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ColonReplacementFormat value
func (v ColonReplacementFormat) Ptr() *ColonReplacementFormat {
	return &v
}

type NullableColonReplacementFormat struct {
	value *ColonReplacementFormat
	isSet bool
}

func (v NullableColonReplacementFormat) Get() *ColonReplacementFormat {
	return v.value
}

func (v *NullableColonReplacementFormat) Set(val *ColonReplacementFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableColonReplacementFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableColonReplacementFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColonReplacementFormat(val *ColonReplacementFormat) *NullableColonReplacementFormat {
	return &NullableColonReplacementFormat{value: val, isSet: true}
}

func (v NullableColonReplacementFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColonReplacementFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

