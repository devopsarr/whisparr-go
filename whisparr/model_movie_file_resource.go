/*
Radarr

Radarr API docs

API version: b08981dee068e1ed23e4f45a0d8fe70ef7bf7703
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whisparr

import (
	"encoding/json"
	"time"
)

// MovieFileResource struct for MovieFileResource
type MovieFileResource struct {
	Id *int32 `json:"id,omitempty"`
	MovieId *int32 `json:"movieId,omitempty"`
	RelativePath NullableString `json:"relativePath,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Size *int64 `json:"size,omitempty"`
	DateAdded *time.Time `json:"dateAdded,omitempty"`
	SceneName NullableString `json:"sceneName,omitempty"`
	IndexerFlags *int32 `json:"indexerFlags,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	CustomFormats []*CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
	MediaInfo *MediaInfoResource `json:"mediaInfo,omitempty"`
	OriginalFilePath NullableString `json:"originalFilePath,omitempty"`
	QualityCutoffNotMet *bool `json:"qualityCutoffNotMet,omitempty"`
	Languages []*Language `json:"languages,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	Edition NullableString `json:"edition,omitempty"`
}

// NewMovieFileResource instantiates a new MovieFileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieFileResource() *MovieFileResource {
	this := MovieFileResource{}
	return &this
}

// NewMovieFileResourceWithDefaults instantiates a new MovieFileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieFileResourceWithDefaults() *MovieFileResource {
	this := MovieFileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MovieFileResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MovieFileResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MovieFileResource) SetId(v int32) {
	o.Id = &v
}

// GetMovieId returns the MovieId field value if set, zero value otherwise.
func (o *MovieFileResource) GetMovieId() int32 {
	if o == nil || isNil(o.MovieId) {
		var ret int32
		return ret
	}
	return *o.MovieId
}

// GetMovieIdOk returns a tuple with the MovieId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetMovieIdOk() (*int32, bool) {
	if o == nil || isNil(o.MovieId) {
    return nil, false
	}
	return o.MovieId, true
}

// HasMovieId returns a boolean if a field has been set.
func (o *MovieFileResource) HasMovieId() bool {
	if o != nil && !isNil(o.MovieId) {
		return true
	}

	return false
}

// SetMovieId gets a reference to the given int32 and assigns it to the MovieId field.
func (o *MovieFileResource) SetMovieId(v int32) {
	o.MovieId = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileResource) GetRelativePath() string {
	if o == nil || isNil(o.RelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.RelativePath.Get()
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileResource) GetRelativePathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RelativePath.Get(), o.RelativePath.IsSet()
}

// HasRelativePath returns a boolean if a field has been set.
func (o *MovieFileResource) HasRelativePath() bool {
	if o != nil && o.RelativePath.IsSet() {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given NullableString and assigns it to the RelativePath field.
func (o *MovieFileResource) SetRelativePath(v string) {
	o.RelativePath.Set(&v)
}
// SetRelativePathNil sets the value for RelativePath to be an explicit nil
func (o *MovieFileResource) SetRelativePathNil() {
	o.RelativePath.Set(nil)
}

// UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
func (o *MovieFileResource) UnsetRelativePath() {
	o.RelativePath.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileResource) GetPath() string {
	if o == nil || isNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileResource) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *MovieFileResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *MovieFileResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *MovieFileResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *MovieFileResource) UnsetPath() {
	o.Path.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MovieFileResource) GetSize() int64 {
	if o == nil || isNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetSizeOk() (*int64, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MovieFileResource) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *MovieFileResource) SetSize(v int64) {
	o.Size = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *MovieFileResource) GetDateAdded() time.Time {
	if o == nil || isNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateAdded) {
    return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *MovieFileResource) HasDateAdded() bool {
	if o != nil && !isNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *MovieFileResource) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileResource) GetSceneName() string {
	if o == nil || isNil(o.SceneName.Get()) {
		var ret string
		return ret
	}
	return *o.SceneName.Get()
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileResource) GetSceneNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SceneName.Get(), o.SceneName.IsSet()
}

// HasSceneName returns a boolean if a field has been set.
func (o *MovieFileResource) HasSceneName() bool {
	if o != nil && o.SceneName.IsSet() {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given NullableString and assigns it to the SceneName field.
func (o *MovieFileResource) SetSceneName(v string) {
	o.SceneName.Set(&v)
}
// SetSceneNameNil sets the value for SceneName to be an explicit nil
func (o *MovieFileResource) SetSceneNameNil() {
	o.SceneName.Set(nil)
}

// UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
func (o *MovieFileResource) UnsetSceneName() {
	o.SceneName.Unset()
}

// GetIndexerFlags returns the IndexerFlags field value if set, zero value otherwise.
func (o *MovieFileResource) GetIndexerFlags() int32 {
	if o == nil || isNil(o.IndexerFlags) {
		var ret int32
		return ret
	}
	return *o.IndexerFlags
}

// GetIndexerFlagsOk returns a tuple with the IndexerFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetIndexerFlagsOk() (*int32, bool) {
	if o == nil || isNil(o.IndexerFlags) {
    return nil, false
	}
	return o.IndexerFlags, true
}

// HasIndexerFlags returns a boolean if a field has been set.
func (o *MovieFileResource) HasIndexerFlags() bool {
	if o != nil && !isNil(o.IndexerFlags) {
		return true
	}

	return false
}

// SetIndexerFlags gets a reference to the given int32 and assigns it to the IndexerFlags field.
func (o *MovieFileResource) SetIndexerFlags(v int32) {
	o.IndexerFlags = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *MovieFileResource) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *MovieFileResource) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *MovieFileResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileResource) GetCustomFormats() []*CustomFormatResource {
	if o == nil {
		var ret []*CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileResource) GetCustomFormatsOk() ([]*CustomFormatResource, bool) {
	if o == nil || isNil(o.CustomFormats) {
    return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *MovieFileResource) HasCustomFormats() bool {
	if o != nil && isNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *MovieFileResource) SetCustomFormats(v []*CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *MovieFileResource) GetCustomFormatScore() int32 {
	if o == nil || isNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || isNil(o.CustomFormatScore) {
    return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *MovieFileResource) HasCustomFormatScore() bool {
	if o != nil && !isNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *MovieFileResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *MovieFileResource) GetMediaInfo() MediaInfoResource {
	if o == nil || isNil(o.MediaInfo) {
		var ret MediaInfoResource
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetMediaInfoOk() (*MediaInfoResource, bool) {
	if o == nil || isNil(o.MediaInfo) {
    return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *MovieFileResource) HasMediaInfo() bool {
	if o != nil && !isNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfoResource and assigns it to the MediaInfo field.
func (o *MovieFileResource) SetMediaInfo(v MediaInfoResource) {
	o.MediaInfo = &v
}

// GetOriginalFilePath returns the OriginalFilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileResource) GetOriginalFilePath() string {
	if o == nil || isNil(o.OriginalFilePath.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalFilePath.Get()
}

// GetOriginalFilePathOk returns a tuple with the OriginalFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileResource) GetOriginalFilePathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.OriginalFilePath.Get(), o.OriginalFilePath.IsSet()
}

// HasOriginalFilePath returns a boolean if a field has been set.
func (o *MovieFileResource) HasOriginalFilePath() bool {
	if o != nil && o.OriginalFilePath.IsSet() {
		return true
	}

	return false
}

// SetOriginalFilePath gets a reference to the given NullableString and assigns it to the OriginalFilePath field.
func (o *MovieFileResource) SetOriginalFilePath(v string) {
	o.OriginalFilePath.Set(&v)
}
// SetOriginalFilePathNil sets the value for OriginalFilePath to be an explicit nil
func (o *MovieFileResource) SetOriginalFilePathNil() {
	o.OriginalFilePath.Set(nil)
}

// UnsetOriginalFilePath ensures that no value is present for OriginalFilePath, not even an explicit nil
func (o *MovieFileResource) UnsetOriginalFilePath() {
	o.OriginalFilePath.Unset()
}

// GetQualityCutoffNotMet returns the QualityCutoffNotMet field value if set, zero value otherwise.
func (o *MovieFileResource) GetQualityCutoffNotMet() bool {
	if o == nil || isNil(o.QualityCutoffNotMet) {
		var ret bool
		return ret
	}
	return *o.QualityCutoffNotMet
}

// GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileResource) GetQualityCutoffNotMetOk() (*bool, bool) {
	if o == nil || isNil(o.QualityCutoffNotMet) {
    return nil, false
	}
	return o.QualityCutoffNotMet, true
}

// HasQualityCutoffNotMet returns a boolean if a field has been set.
func (o *MovieFileResource) HasQualityCutoffNotMet() bool {
	if o != nil && !isNil(o.QualityCutoffNotMet) {
		return true
	}

	return false
}

// SetQualityCutoffNotMet gets a reference to the given bool and assigns it to the QualityCutoffNotMet field.
func (o *MovieFileResource) SetQualityCutoffNotMet(v bool) {
	o.QualityCutoffNotMet = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileResource) GetLanguages() []*Language {
	if o == nil {
		var ret []*Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileResource) GetLanguagesOk() ([]*Language, bool) {
	if o == nil || isNil(o.Languages) {
    return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *MovieFileResource) HasLanguages() bool {
	if o != nil && isNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *MovieFileResource) SetLanguages(v []*Language) {
	o.Languages = v
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileResource) GetReleaseGroup() string {
	if o == nil || isNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileResource) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *MovieFileResource) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *MovieFileResource) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *MovieFileResource) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *MovieFileResource) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetEdition returns the Edition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileResource) GetEdition() string {
	if o == nil || isNil(o.Edition.Get()) {
		var ret string
		return ret
	}
	return *o.Edition.Get()
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileResource) GetEditionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Edition.Get(), o.Edition.IsSet()
}

// HasEdition returns a boolean if a field has been set.
func (o *MovieFileResource) HasEdition() bool {
	if o != nil && o.Edition.IsSet() {
		return true
	}

	return false
}

// SetEdition gets a reference to the given NullableString and assigns it to the Edition field.
func (o *MovieFileResource) SetEdition(v string) {
	o.Edition.Set(&v)
}
// SetEditionNil sets the value for Edition to be an explicit nil
func (o *MovieFileResource) SetEditionNil() {
	o.Edition.Set(nil)
}

// UnsetEdition ensures that no value is present for Edition, not even an explicit nil
func (o *MovieFileResource) UnsetEdition() {
	o.Edition.Unset()
}

func (o MovieFileResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.MovieId) {
		toSerialize["movieId"] = o.MovieId
	}
	if o.RelativePath.IsSet() {
		toSerialize["relativePath"] = o.RelativePath.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.DateAdded) {
		toSerialize["dateAdded"] = o.DateAdded
	}
	if o.SceneName.IsSet() {
		toSerialize["sceneName"] = o.SceneName.Get()
	}
	if !isNil(o.IndexerFlags) {
		toSerialize["indexerFlags"] = o.IndexerFlags
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !isNil(o.CustomFormatScore) {
		toSerialize["customFormatScore"] = o.CustomFormatScore
	}
	if !isNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}
	if o.OriginalFilePath.IsSet() {
		toSerialize["originalFilePath"] = o.OriginalFilePath.Get()
	}
	if !isNil(o.QualityCutoffNotMet) {
		toSerialize["qualityCutoffNotMet"] = o.QualityCutoffNotMet
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if o.Edition.IsSet() {
		toSerialize["edition"] = o.Edition.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMovieFileResource struct {
	value *MovieFileResource
	isSet bool
}

func (v NullableMovieFileResource) Get() *MovieFileResource {
	return v.value
}

func (v *NullableMovieFileResource) Set(val *MovieFileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieFileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieFileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieFileResource(val *MovieFileResource) *NullableMovieFileResource {
	return &NullableMovieFileResource{value: val, isSet: true}
}

func (v NullableMovieFileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieFileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


