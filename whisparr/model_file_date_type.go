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

// FileDateType the model 'FileDateType'
type FileDateType string

// List of FileDateType
const (
	FILEDATETYPE_NONE FileDateType = "none"
	FILEDATETYPE_CINEMAS FileDateType = "cinemas"
	FILEDATETYPE_RELEASE FileDateType = "release"
)

// All allowed values of FileDateType enum
var AllowedFileDateTypeEnumValues = []FileDateType{
	"none",
	"cinemas",
	"release",
}

func (v *FileDateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileDateType(value)
	for _, existing := range AllowedFileDateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileDateType", value)
}

// NewFileDateTypeFromValue returns a pointer to a valid FileDateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileDateTypeFromValue(v string) (*FileDateType, error) {
	ev := FileDateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileDateType: valid values are %v", v, AllowedFileDateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileDateType) IsValid() bool {
	for _, existing := range AllowedFileDateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileDateType value
func (v FileDateType) Ptr() *FileDateType {
	return &v
}

type NullableFileDateType struct {
	value *FileDateType
	isSet bool
}

func (v NullableFileDateType) Get() *FileDateType {
	return v.value
}

func (v *NullableFileDateType) Set(val *FileDateType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDateType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDateType(val *FileDateType) *NullableFileDateType {
	return &NullableFileDateType{value: val, isSet: true}
}

func (v NullableFileDateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

