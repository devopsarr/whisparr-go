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

// HttpUri struct for HttpUri
type HttpUri struct {
	FullUri NullableString `json:"fullUri,omitempty"`
	Scheme NullableString `json:"scheme,omitempty"`
	Host NullableString `json:"host,omitempty"`
	Port NullableInt32 `json:"port,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Query NullableString `json:"query,omitempty"`
	Fragment NullableString `json:"fragment,omitempty"`
}

// NewHttpUri instantiates a new HttpUri object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpUri() *HttpUri {
	this := HttpUri{}
	return &this
}

// NewHttpUriWithDefaults instantiates a new HttpUri object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpUriWithDefaults() *HttpUri {
	this := HttpUri{}
	return &this
}

// GetFullUri returns the FullUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpUri) GetFullUri() string {
	if o == nil || isNil(o.FullUri.Get()) {
		var ret string
		return ret
	}
	return *o.FullUri.Get()
}

// GetFullUriOk returns a tuple with the FullUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpUri) GetFullUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FullUri.Get(), o.FullUri.IsSet()
}

// HasFullUri returns a boolean if a field has been set.
func (o *HttpUri) HasFullUri() bool {
	if o != nil && o.FullUri.IsSet() {
		return true
	}

	return false
}

// SetFullUri gets a reference to the given NullableString and assigns it to the FullUri field.
func (o *HttpUri) SetFullUri(v string) {
	o.FullUri.Set(&v)
}
// SetFullUriNil sets the value for FullUri to be an explicit nil
func (o *HttpUri) SetFullUriNil() {
	o.FullUri.Set(nil)
}

// UnsetFullUri ensures that no value is present for FullUri, not even an explicit nil
func (o *HttpUri) UnsetFullUri() {
	o.FullUri.Unset()
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpUri) GetScheme() string {
	if o == nil || isNil(o.Scheme.Get()) {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpUri) GetSchemeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *HttpUri) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *HttpUri) SetScheme(v string) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *HttpUri) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *HttpUri) UnsetScheme() {
	o.Scheme.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpUri) GetHost() string {
	if o == nil || isNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpUri) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *HttpUri) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *HttpUri) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *HttpUri) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *HttpUri) UnsetHost() {
	o.Host.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpUri) GetPort() int32 {
	if o == nil || isNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpUri) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *HttpUri) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *HttpUri) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *HttpUri) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *HttpUri) UnsetPort() {
	o.Port.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpUri) GetPath() string {
	if o == nil || isNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpUri) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *HttpUri) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *HttpUri) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *HttpUri) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *HttpUri) UnsetPath() {
	o.Path.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpUri) GetQuery() string {
	if o == nil || isNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpUri) GetQueryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *HttpUri) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *HttpUri) SetQuery(v string) {
	o.Query.Set(&v)
}
// SetQueryNil sets the value for Query to be an explicit nil
func (o *HttpUri) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *HttpUri) UnsetQuery() {
	o.Query.Unset()
}

// GetFragment returns the Fragment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpUri) GetFragment() string {
	if o == nil || isNil(o.Fragment.Get()) {
		var ret string
		return ret
	}
	return *o.Fragment.Get()
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpUri) GetFragmentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Fragment.Get(), o.Fragment.IsSet()
}

// HasFragment returns a boolean if a field has been set.
func (o *HttpUri) HasFragment() bool {
	if o != nil && o.Fragment.IsSet() {
		return true
	}

	return false
}

// SetFragment gets a reference to the given NullableString and assigns it to the Fragment field.
func (o *HttpUri) SetFragment(v string) {
	o.Fragment.Set(&v)
}
// SetFragmentNil sets the value for Fragment to be an explicit nil
func (o *HttpUri) SetFragmentNil() {
	o.Fragment.Set(nil)
}

// UnsetFragment ensures that no value is present for Fragment, not even an explicit nil
func (o *HttpUri) UnsetFragment() {
	o.Fragment.Unset()
}

func (o HttpUri) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FullUri.IsSet() {
		toSerialize["fullUri"] = o.FullUri.Get()
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.Fragment.IsSet() {
		toSerialize["fragment"] = o.Fragment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHttpUri struct {
	value *HttpUri
	isSet bool
}

func (v NullableHttpUri) Get() *HttpUri {
	return v.value
}

func (v *NullableHttpUri) Set(val *HttpUri) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpUri) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpUri) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpUri(val *HttpUri) *NullableHttpUri {
	return &NullableHttpUri{value: val, isSet: true}
}

func (v NullableHttpUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpUri) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


