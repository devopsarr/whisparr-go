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

// Source the model 'Source'
type Source string

// List of Source
const (
	SOURCE_UNKNOWN Source = "unknown"
	SOURCE_CAM Source = "cam"
	SOURCE_TELESYNC Source = "telesync"
	SOURCE_TELECINE Source = "telecine"
	SOURCE_WORKPRINT Source = "workprint"
	SOURCE_DVD Source = "dvd"
	SOURCE_TV Source = "tv"
	SOURCE_WEBDL Source = "webdl"
	SOURCE_WEBRIP Source = "webrip"
	SOURCE_BLURAY Source = "bluray"
)

// All allowed values of Source enum
var AllowedSourceEnumValues = []Source{
	"unknown",
	"cam",
	"telesync",
	"telecine",
	"workprint",
	"dvd",
	"tv",
	"webdl",
	"webrip",
	"bluray",
}

func (v *Source) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Source(value)
	for _, existing := range AllowedSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Source", value)
}

// NewSourceFromValue returns a pointer to a valid Source
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSourceFromValue(v string) (*Source, error) {
	ev := Source(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Source: valid values are %v", v, AllowedSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Source) IsValid() bool {
	for _, existing := range AllowedSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Source value
func (v Source) Ptr() *Source {
	return &v
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource(val *Source) *NullableSource {
	return &NullableSource{value: val, isSet: true}
}

func (v NullableSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

