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

// AddMovieOptions struct for AddMovieOptions
type AddMovieOptions struct {
	IgnoreEpisodesWithFiles *bool `json:"ignoreEpisodesWithFiles,omitempty"`
	IgnoreEpisodesWithoutFiles *bool `json:"ignoreEpisodesWithoutFiles,omitempty"`
	Monitor *MonitorTypes `json:"monitor,omitempty"`
	SearchForMovie *bool `json:"searchForMovie,omitempty"`
	AddMethod *AddMovieMethod `json:"addMethod,omitempty"`
}

// NewAddMovieOptions instantiates a new AddMovieOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMovieOptions() *AddMovieOptions {
	this := AddMovieOptions{}
	return &this
}

// NewAddMovieOptionsWithDefaults instantiates a new AddMovieOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMovieOptionsWithDefaults() *AddMovieOptions {
	this := AddMovieOptions{}
	return &this
}

// GetIgnoreEpisodesWithFiles returns the IgnoreEpisodesWithFiles field value if set, zero value otherwise.
func (o *AddMovieOptions) GetIgnoreEpisodesWithFiles() bool {
	if o == nil || isNil(o.IgnoreEpisodesWithFiles) {
		var ret bool
		return ret
	}
	return *o.IgnoreEpisodesWithFiles
}

// GetIgnoreEpisodesWithFilesOk returns a tuple with the IgnoreEpisodesWithFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMovieOptions) GetIgnoreEpisodesWithFilesOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreEpisodesWithFiles) {
    return nil, false
	}
	return o.IgnoreEpisodesWithFiles, true
}

// HasIgnoreEpisodesWithFiles returns a boolean if a field has been set.
func (o *AddMovieOptions) HasIgnoreEpisodesWithFiles() bool {
	if o != nil && !isNil(o.IgnoreEpisodesWithFiles) {
		return true
	}

	return false
}

// SetIgnoreEpisodesWithFiles gets a reference to the given bool and assigns it to the IgnoreEpisodesWithFiles field.
func (o *AddMovieOptions) SetIgnoreEpisodesWithFiles(v bool) {
	o.IgnoreEpisodesWithFiles = &v
}

// GetIgnoreEpisodesWithoutFiles returns the IgnoreEpisodesWithoutFiles field value if set, zero value otherwise.
func (o *AddMovieOptions) GetIgnoreEpisodesWithoutFiles() bool {
	if o == nil || isNil(o.IgnoreEpisodesWithoutFiles) {
		var ret bool
		return ret
	}
	return *o.IgnoreEpisodesWithoutFiles
}

// GetIgnoreEpisodesWithoutFilesOk returns a tuple with the IgnoreEpisodesWithoutFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMovieOptions) GetIgnoreEpisodesWithoutFilesOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreEpisodesWithoutFiles) {
    return nil, false
	}
	return o.IgnoreEpisodesWithoutFiles, true
}

// HasIgnoreEpisodesWithoutFiles returns a boolean if a field has been set.
func (o *AddMovieOptions) HasIgnoreEpisodesWithoutFiles() bool {
	if o != nil && !isNil(o.IgnoreEpisodesWithoutFiles) {
		return true
	}

	return false
}

// SetIgnoreEpisodesWithoutFiles gets a reference to the given bool and assigns it to the IgnoreEpisodesWithoutFiles field.
func (o *AddMovieOptions) SetIgnoreEpisodesWithoutFiles(v bool) {
	o.IgnoreEpisodesWithoutFiles = &v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *AddMovieOptions) GetMonitor() MonitorTypes {
	if o == nil || isNil(o.Monitor) {
		var ret MonitorTypes
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMovieOptions) GetMonitorOk() (*MonitorTypes, bool) {
	if o == nil || isNil(o.Monitor) {
    return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *AddMovieOptions) HasMonitor() bool {
	if o != nil && !isNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given MonitorTypes and assigns it to the Monitor field.
func (o *AddMovieOptions) SetMonitor(v MonitorTypes) {
	o.Monitor = &v
}

// GetSearchForMovie returns the SearchForMovie field value if set, zero value otherwise.
func (o *AddMovieOptions) GetSearchForMovie() bool {
	if o == nil || isNil(o.SearchForMovie) {
		var ret bool
		return ret
	}
	return *o.SearchForMovie
}

// GetSearchForMovieOk returns a tuple with the SearchForMovie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMovieOptions) GetSearchForMovieOk() (*bool, bool) {
	if o == nil || isNil(o.SearchForMovie) {
    return nil, false
	}
	return o.SearchForMovie, true
}

// HasSearchForMovie returns a boolean if a field has been set.
func (o *AddMovieOptions) HasSearchForMovie() bool {
	if o != nil && !isNil(o.SearchForMovie) {
		return true
	}

	return false
}

// SetSearchForMovie gets a reference to the given bool and assigns it to the SearchForMovie field.
func (o *AddMovieOptions) SetSearchForMovie(v bool) {
	o.SearchForMovie = &v
}

// GetAddMethod returns the AddMethod field value if set, zero value otherwise.
func (o *AddMovieOptions) GetAddMethod() AddMovieMethod {
	if o == nil || isNil(o.AddMethod) {
		var ret AddMovieMethod
		return ret
	}
	return *o.AddMethod
}

// GetAddMethodOk returns a tuple with the AddMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMovieOptions) GetAddMethodOk() (*AddMovieMethod, bool) {
	if o == nil || isNil(o.AddMethod) {
    return nil, false
	}
	return o.AddMethod, true
}

// HasAddMethod returns a boolean if a field has been set.
func (o *AddMovieOptions) HasAddMethod() bool {
	if o != nil && !isNil(o.AddMethod) {
		return true
	}

	return false
}

// SetAddMethod gets a reference to the given AddMovieMethod and assigns it to the AddMethod field.
func (o *AddMovieOptions) SetAddMethod(v AddMovieMethod) {
	o.AddMethod = &v
}

func (o AddMovieOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IgnoreEpisodesWithFiles) {
		toSerialize["ignoreEpisodesWithFiles"] = o.IgnoreEpisodesWithFiles
	}
	if !isNil(o.IgnoreEpisodesWithoutFiles) {
		toSerialize["ignoreEpisodesWithoutFiles"] = o.IgnoreEpisodesWithoutFiles
	}
	if !isNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	if !isNil(o.SearchForMovie) {
		toSerialize["searchForMovie"] = o.SearchForMovie
	}
	if !isNil(o.AddMethod) {
		toSerialize["addMethod"] = o.AddMethod
	}
	return json.Marshal(toSerialize)
}

type NullableAddMovieOptions struct {
	value *AddMovieOptions
	isSet bool
}

func (v NullableAddMovieOptions) Get() *AddMovieOptions {
	return v.value
}

func (v *NullableAddMovieOptions) Set(val *AddMovieOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMovieOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMovieOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMovieOptions(val *AddMovieOptions) *NullableAddMovieOptions {
	return &NullableAddMovieOptions{value: val, isSet: true}
}

func (v NullableAddMovieOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMovieOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


