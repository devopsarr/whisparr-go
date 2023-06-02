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

// NotificationBulkResource struct for NotificationBulkResource
type NotificationBulkResource struct {
	Ids []*int32 `json:"ids,omitempty"`
	Tags []*int32 `json:"tags,omitempty"`
	ApplyTags *ApplyTags `json:"applyTags,omitempty"`
}

// NewNotificationBulkResource instantiates a new NotificationBulkResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationBulkResource() *NotificationBulkResource {
	this := NotificationBulkResource{}
	return &this
}

// NewNotificationBulkResourceWithDefaults instantiates a new NotificationBulkResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationBulkResourceWithDefaults() *NotificationBulkResource {
	this := NotificationBulkResource{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationBulkResource) GetIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationBulkResource) GetIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *NotificationBulkResource) HasIds() bool {
	if o != nil && isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *NotificationBulkResource) SetIds(v []*int32) {
	o.Ids = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationBulkResource) GetTags() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationBulkResource) GetTagsOk() ([]*int32, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NotificationBulkResource) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *NotificationBulkResource) SetTags(v []*int32) {
	o.Tags = v
}

// GetApplyTags returns the ApplyTags field value if set, zero value otherwise.
func (o *NotificationBulkResource) GetApplyTags() ApplyTags {
	if o == nil || isNil(o.ApplyTags) {
		var ret ApplyTags
		return ret
	}
	return *o.ApplyTags
}

// GetApplyTagsOk returns a tuple with the ApplyTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBulkResource) GetApplyTagsOk() (*ApplyTags, bool) {
	if o == nil || isNil(o.ApplyTags) {
    return nil, false
	}
	return o.ApplyTags, true
}

// HasApplyTags returns a boolean if a field has been set.
func (o *NotificationBulkResource) HasApplyTags() bool {
	if o != nil && !isNil(o.ApplyTags) {
		return true
	}

	return false
}

// SetApplyTags gets a reference to the given ApplyTags and assigns it to the ApplyTags field.
func (o *NotificationBulkResource) SetApplyTags(v ApplyTags) {
	o.ApplyTags = &v
}

func (o NotificationBulkResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.ApplyTags) {
		toSerialize["applyTags"] = o.ApplyTags
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationBulkResource struct {
	value *NotificationBulkResource
	isSet bool
}

func (v NullableNotificationBulkResource) Get() *NotificationBulkResource {
	return v.value
}

func (v *NullableNotificationBulkResource) Set(val *NotificationBulkResource) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationBulkResource) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationBulkResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationBulkResource(val *NotificationBulkResource) *NullableNotificationBulkResource {
	return &NullableNotificationBulkResource{value: val, isSet: true}
}

func (v NullableNotificationBulkResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationBulkResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


