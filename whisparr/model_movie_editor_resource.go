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

// MovieEditorResource struct for MovieEditorResource
type MovieEditorResource struct {
	MovieIds []*int32 `json:"movieIds,omitempty"`
	Monitored NullableBool `json:"monitored,omitempty"`
	QualityProfileId NullableInt32 `json:"qualityProfileId,omitempty"`
	MinimumAvailability *MovieStatusType `json:"minimumAvailability,omitempty"`
	RootFolderPath NullableString `json:"rootFolderPath,omitempty"`
	Tags []*int32 `json:"tags,omitempty"`
	ApplyTags *ApplyTags `json:"applyTags,omitempty"`
	MoveFiles *bool `json:"moveFiles,omitempty"`
	DeleteFiles *bool `json:"deleteFiles,omitempty"`
	AddImportExclusion *bool `json:"addImportExclusion,omitempty"`
}

// NewMovieEditorResource instantiates a new MovieEditorResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieEditorResource() *MovieEditorResource {
	this := MovieEditorResource{}
	return &this
}

// NewMovieEditorResourceWithDefaults instantiates a new MovieEditorResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieEditorResourceWithDefaults() *MovieEditorResource {
	this := MovieEditorResource{}
	return &this
}

// GetMovieIds returns the MovieIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieEditorResource) GetMovieIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.MovieIds
}

// GetMovieIdsOk returns a tuple with the MovieIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieEditorResource) GetMovieIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.MovieIds) {
    return nil, false
	}
	return o.MovieIds, true
}

// HasMovieIds returns a boolean if a field has been set.
func (o *MovieEditorResource) HasMovieIds() bool {
	if o != nil && isNil(o.MovieIds) {
		return true
	}

	return false
}

// SetMovieIds gets a reference to the given []int32 and assigns it to the MovieIds field.
func (o *MovieEditorResource) SetMovieIds(v []*int32) {
	o.MovieIds = v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieEditorResource) GetMonitored() bool {
	if o == nil || isNil(o.Monitored.Get()) {
		var ret bool
		return ret
	}
	return *o.Monitored.Get()
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieEditorResource) GetMonitoredOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.Monitored.Get(), o.Monitored.IsSet()
}

// HasMonitored returns a boolean if a field has been set.
func (o *MovieEditorResource) HasMonitored() bool {
	if o != nil && o.Monitored.IsSet() {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given NullableBool and assigns it to the Monitored field.
func (o *MovieEditorResource) SetMonitored(v bool) {
	o.Monitored.Set(&v)
}
// SetMonitoredNil sets the value for Monitored to be an explicit nil
func (o *MovieEditorResource) SetMonitoredNil() {
	o.Monitored.Set(nil)
}

// UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
func (o *MovieEditorResource) UnsetMonitored() {
	o.Monitored.Unset()
}

// GetQualityProfileId returns the QualityProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieEditorResource) GetQualityProfileId() int32 {
	if o == nil || isNil(o.QualityProfileId.Get()) {
		var ret int32
		return ret
	}
	return *o.QualityProfileId.Get()
}

// GetQualityProfileIdOk returns a tuple with the QualityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieEditorResource) GetQualityProfileIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.QualityProfileId.Get(), o.QualityProfileId.IsSet()
}

// HasQualityProfileId returns a boolean if a field has been set.
func (o *MovieEditorResource) HasQualityProfileId() bool {
	if o != nil && o.QualityProfileId.IsSet() {
		return true
	}

	return false
}

// SetQualityProfileId gets a reference to the given NullableInt32 and assigns it to the QualityProfileId field.
func (o *MovieEditorResource) SetQualityProfileId(v int32) {
	o.QualityProfileId.Set(&v)
}
// SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil
func (o *MovieEditorResource) SetQualityProfileIdNil() {
	o.QualityProfileId.Set(nil)
}

// UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil
func (o *MovieEditorResource) UnsetQualityProfileId() {
	o.QualityProfileId.Unset()
}

// GetMinimumAvailability returns the MinimumAvailability field value if set, zero value otherwise.
func (o *MovieEditorResource) GetMinimumAvailability() MovieStatusType {
	if o == nil || isNil(o.MinimumAvailability) {
		var ret MovieStatusType
		return ret
	}
	return *o.MinimumAvailability
}

// GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieEditorResource) GetMinimumAvailabilityOk() (*MovieStatusType, bool) {
	if o == nil || isNil(o.MinimumAvailability) {
    return nil, false
	}
	return o.MinimumAvailability, true
}

// HasMinimumAvailability returns a boolean if a field has been set.
func (o *MovieEditorResource) HasMinimumAvailability() bool {
	if o != nil && !isNil(o.MinimumAvailability) {
		return true
	}

	return false
}

// SetMinimumAvailability gets a reference to the given MovieStatusType and assigns it to the MinimumAvailability field.
func (o *MovieEditorResource) SetMinimumAvailability(v MovieStatusType) {
	o.MinimumAvailability = &v
}

// GetRootFolderPath returns the RootFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieEditorResource) GetRootFolderPath() string {
	if o == nil || isNil(o.RootFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderPath.Get()
}

// GetRootFolderPathOk returns a tuple with the RootFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieEditorResource) GetRootFolderPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RootFolderPath.Get(), o.RootFolderPath.IsSet()
}

// HasRootFolderPath returns a boolean if a field has been set.
func (o *MovieEditorResource) HasRootFolderPath() bool {
	if o != nil && o.RootFolderPath.IsSet() {
		return true
	}

	return false
}

// SetRootFolderPath gets a reference to the given NullableString and assigns it to the RootFolderPath field.
func (o *MovieEditorResource) SetRootFolderPath(v string) {
	o.RootFolderPath.Set(&v)
}
// SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil
func (o *MovieEditorResource) SetRootFolderPathNil() {
	o.RootFolderPath.Set(nil)
}

// UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
func (o *MovieEditorResource) UnsetRootFolderPath() {
	o.RootFolderPath.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieEditorResource) GetTags() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieEditorResource) GetTagsOk() ([]*int32, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MovieEditorResource) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *MovieEditorResource) SetTags(v []*int32) {
	o.Tags = v
}

// GetApplyTags returns the ApplyTags field value if set, zero value otherwise.
func (o *MovieEditorResource) GetApplyTags() ApplyTags {
	if o == nil || isNil(o.ApplyTags) {
		var ret ApplyTags
		return ret
	}
	return *o.ApplyTags
}

// GetApplyTagsOk returns a tuple with the ApplyTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieEditorResource) GetApplyTagsOk() (*ApplyTags, bool) {
	if o == nil || isNil(o.ApplyTags) {
    return nil, false
	}
	return o.ApplyTags, true
}

// HasApplyTags returns a boolean if a field has been set.
func (o *MovieEditorResource) HasApplyTags() bool {
	if o != nil && !isNil(o.ApplyTags) {
		return true
	}

	return false
}

// SetApplyTags gets a reference to the given ApplyTags and assigns it to the ApplyTags field.
func (o *MovieEditorResource) SetApplyTags(v ApplyTags) {
	o.ApplyTags = &v
}

// GetMoveFiles returns the MoveFiles field value if set, zero value otherwise.
func (o *MovieEditorResource) GetMoveFiles() bool {
	if o == nil || isNil(o.MoveFiles) {
		var ret bool
		return ret
	}
	return *o.MoveFiles
}

// GetMoveFilesOk returns a tuple with the MoveFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieEditorResource) GetMoveFilesOk() (*bool, bool) {
	if o == nil || isNil(o.MoveFiles) {
    return nil, false
	}
	return o.MoveFiles, true
}

// HasMoveFiles returns a boolean if a field has been set.
func (o *MovieEditorResource) HasMoveFiles() bool {
	if o != nil && !isNil(o.MoveFiles) {
		return true
	}

	return false
}

// SetMoveFiles gets a reference to the given bool and assigns it to the MoveFiles field.
func (o *MovieEditorResource) SetMoveFiles(v bool) {
	o.MoveFiles = &v
}

// GetDeleteFiles returns the DeleteFiles field value if set, zero value otherwise.
func (o *MovieEditorResource) GetDeleteFiles() bool {
	if o == nil || isNil(o.DeleteFiles) {
		var ret bool
		return ret
	}
	return *o.DeleteFiles
}

// GetDeleteFilesOk returns a tuple with the DeleteFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieEditorResource) GetDeleteFilesOk() (*bool, bool) {
	if o == nil || isNil(o.DeleteFiles) {
    return nil, false
	}
	return o.DeleteFiles, true
}

// HasDeleteFiles returns a boolean if a field has been set.
func (o *MovieEditorResource) HasDeleteFiles() bool {
	if o != nil && !isNil(o.DeleteFiles) {
		return true
	}

	return false
}

// SetDeleteFiles gets a reference to the given bool and assigns it to the DeleteFiles field.
func (o *MovieEditorResource) SetDeleteFiles(v bool) {
	o.DeleteFiles = &v
}

// GetAddImportExclusion returns the AddImportExclusion field value if set, zero value otherwise.
func (o *MovieEditorResource) GetAddImportExclusion() bool {
	if o == nil || isNil(o.AddImportExclusion) {
		var ret bool
		return ret
	}
	return *o.AddImportExclusion
}

// GetAddImportExclusionOk returns a tuple with the AddImportExclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieEditorResource) GetAddImportExclusionOk() (*bool, bool) {
	if o == nil || isNil(o.AddImportExclusion) {
    return nil, false
	}
	return o.AddImportExclusion, true
}

// HasAddImportExclusion returns a boolean if a field has been set.
func (o *MovieEditorResource) HasAddImportExclusion() bool {
	if o != nil && !isNil(o.AddImportExclusion) {
		return true
	}

	return false
}

// SetAddImportExclusion gets a reference to the given bool and assigns it to the AddImportExclusion field.
func (o *MovieEditorResource) SetAddImportExclusion(v bool) {
	o.AddImportExclusion = &v
}

func (o MovieEditorResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MovieIds != nil {
		toSerialize["movieIds"] = o.MovieIds
	}
	if o.Monitored.IsSet() {
		toSerialize["monitored"] = o.Monitored.Get()
	}
	if o.QualityProfileId.IsSet() {
		toSerialize["qualityProfileId"] = o.QualityProfileId.Get()
	}
	if !isNil(o.MinimumAvailability) {
		toSerialize["minimumAvailability"] = o.MinimumAvailability
	}
	if o.RootFolderPath.IsSet() {
		toSerialize["rootFolderPath"] = o.RootFolderPath.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.ApplyTags) {
		toSerialize["applyTags"] = o.ApplyTags
	}
	if !isNil(o.MoveFiles) {
		toSerialize["moveFiles"] = o.MoveFiles
	}
	if !isNil(o.DeleteFiles) {
		toSerialize["deleteFiles"] = o.DeleteFiles
	}
	if !isNil(o.AddImportExclusion) {
		toSerialize["addImportExclusion"] = o.AddImportExclusion
	}
	return json.Marshal(toSerialize)
}

type NullableMovieEditorResource struct {
	value *MovieEditorResource
	isSet bool
}

func (v NullableMovieEditorResource) Get() *MovieEditorResource {
	return v.value
}

func (v *NullableMovieEditorResource) Set(val *MovieEditorResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieEditorResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieEditorResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieEditorResource(val *MovieEditorResource) *NullableMovieEditorResource {
	return &NullableMovieEditorResource{value: val, isSet: true}
}

func (v NullableMovieEditorResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieEditorResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

