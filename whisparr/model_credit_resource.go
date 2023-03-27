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

// CreditResource struct for CreditResource
type CreditResource struct {
	Id *int32 `json:"id,omitempty"`
	PersonName NullableString `json:"personName,omitempty"`
	CreditTmdbId NullableString `json:"creditTmdbId,omitempty"`
	PersonTmdbId *int32 `json:"personTmdbId,omitempty"`
	MovieMetadataId *int32 `json:"movieMetadataId,omitempty"`
	Images []*MediaCover `json:"images,omitempty"`
	Department NullableString `json:"department,omitempty"`
	Job NullableString `json:"job,omitempty"`
	Character NullableString `json:"character,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Type *CreditType `json:"type,omitempty"`
}

// NewCreditResource instantiates a new CreditResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditResource() *CreditResource {
	this := CreditResource{}
	return &this
}

// NewCreditResourceWithDefaults instantiates a new CreditResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditResourceWithDefaults() *CreditResource {
	this := CreditResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreditResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreditResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CreditResource) SetId(v int32) {
	o.Id = &v
}

// GetPersonName returns the PersonName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditResource) GetPersonName() string {
	if o == nil || isNil(o.PersonName.Get()) {
		var ret string
		return ret
	}
	return *o.PersonName.Get()
}

// GetPersonNameOk returns a tuple with the PersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditResource) GetPersonNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PersonName.Get(), o.PersonName.IsSet()
}

// HasPersonName returns a boolean if a field has been set.
func (o *CreditResource) HasPersonName() bool {
	if o != nil && o.PersonName.IsSet() {
		return true
	}

	return false
}

// SetPersonName gets a reference to the given NullableString and assigns it to the PersonName field.
func (o *CreditResource) SetPersonName(v string) {
	o.PersonName.Set(&v)
}
// SetPersonNameNil sets the value for PersonName to be an explicit nil
func (o *CreditResource) SetPersonNameNil() {
	o.PersonName.Set(nil)
}

// UnsetPersonName ensures that no value is present for PersonName, not even an explicit nil
func (o *CreditResource) UnsetPersonName() {
	o.PersonName.Unset()
}

// GetCreditTmdbId returns the CreditTmdbId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditResource) GetCreditTmdbId() string {
	if o == nil || isNil(o.CreditTmdbId.Get()) {
		var ret string
		return ret
	}
	return *o.CreditTmdbId.Get()
}

// GetCreditTmdbIdOk returns a tuple with the CreditTmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditResource) GetCreditTmdbIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CreditTmdbId.Get(), o.CreditTmdbId.IsSet()
}

// HasCreditTmdbId returns a boolean if a field has been set.
func (o *CreditResource) HasCreditTmdbId() bool {
	if o != nil && o.CreditTmdbId.IsSet() {
		return true
	}

	return false
}

// SetCreditTmdbId gets a reference to the given NullableString and assigns it to the CreditTmdbId field.
func (o *CreditResource) SetCreditTmdbId(v string) {
	o.CreditTmdbId.Set(&v)
}
// SetCreditTmdbIdNil sets the value for CreditTmdbId to be an explicit nil
func (o *CreditResource) SetCreditTmdbIdNil() {
	o.CreditTmdbId.Set(nil)
}

// UnsetCreditTmdbId ensures that no value is present for CreditTmdbId, not even an explicit nil
func (o *CreditResource) UnsetCreditTmdbId() {
	o.CreditTmdbId.Unset()
}

// GetPersonTmdbId returns the PersonTmdbId field value if set, zero value otherwise.
func (o *CreditResource) GetPersonTmdbId() int32 {
	if o == nil || isNil(o.PersonTmdbId) {
		var ret int32
		return ret
	}
	return *o.PersonTmdbId
}

// GetPersonTmdbIdOk returns a tuple with the PersonTmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditResource) GetPersonTmdbIdOk() (*int32, bool) {
	if o == nil || isNil(o.PersonTmdbId) {
    return nil, false
	}
	return o.PersonTmdbId, true
}

// HasPersonTmdbId returns a boolean if a field has been set.
func (o *CreditResource) HasPersonTmdbId() bool {
	if o != nil && !isNil(o.PersonTmdbId) {
		return true
	}

	return false
}

// SetPersonTmdbId gets a reference to the given int32 and assigns it to the PersonTmdbId field.
func (o *CreditResource) SetPersonTmdbId(v int32) {
	o.PersonTmdbId = &v
}

// GetMovieMetadataId returns the MovieMetadataId field value if set, zero value otherwise.
func (o *CreditResource) GetMovieMetadataId() int32 {
	if o == nil || isNil(o.MovieMetadataId) {
		var ret int32
		return ret
	}
	return *o.MovieMetadataId
}

// GetMovieMetadataIdOk returns a tuple with the MovieMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditResource) GetMovieMetadataIdOk() (*int32, bool) {
	if o == nil || isNil(o.MovieMetadataId) {
    return nil, false
	}
	return o.MovieMetadataId, true
}

// HasMovieMetadataId returns a boolean if a field has been set.
func (o *CreditResource) HasMovieMetadataId() bool {
	if o != nil && !isNil(o.MovieMetadataId) {
		return true
	}

	return false
}

// SetMovieMetadataId gets a reference to the given int32 and assigns it to the MovieMetadataId field.
func (o *CreditResource) SetMovieMetadataId(v int32) {
	o.MovieMetadataId = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditResource) GetImages() []*MediaCover {
	if o == nil {
		var ret []*MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditResource) GetImagesOk() ([]*MediaCover, bool) {
	if o == nil || isNil(o.Images) {
    return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *CreditResource) HasImages() bool {
	if o != nil && isNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *CreditResource) SetImages(v []*MediaCover) {
	o.Images = v
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditResource) GetDepartment() string {
	if o == nil || isNil(o.Department.Get()) {
		var ret string
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditResource) GetDepartmentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *CreditResource) HasDepartment() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableString and assigns it to the Department field.
func (o *CreditResource) SetDepartment(v string) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *CreditResource) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *CreditResource) UnsetDepartment() {
	o.Department.Unset()
}

// GetJob returns the Job field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditResource) GetJob() string {
	if o == nil || isNil(o.Job.Get()) {
		var ret string
		return ret
	}
	return *o.Job.Get()
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditResource) GetJobOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Job.Get(), o.Job.IsSet()
}

// HasJob returns a boolean if a field has been set.
func (o *CreditResource) HasJob() bool {
	if o != nil && o.Job.IsSet() {
		return true
	}

	return false
}

// SetJob gets a reference to the given NullableString and assigns it to the Job field.
func (o *CreditResource) SetJob(v string) {
	o.Job.Set(&v)
}
// SetJobNil sets the value for Job to be an explicit nil
func (o *CreditResource) SetJobNil() {
	o.Job.Set(nil)
}

// UnsetJob ensures that no value is present for Job, not even an explicit nil
func (o *CreditResource) UnsetJob() {
	o.Job.Unset()
}

// GetCharacter returns the Character field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditResource) GetCharacter() string {
	if o == nil || isNil(o.Character.Get()) {
		var ret string
		return ret
	}
	return *o.Character.Get()
}

// GetCharacterOk returns a tuple with the Character field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditResource) GetCharacterOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Character.Get(), o.Character.IsSet()
}

// HasCharacter returns a boolean if a field has been set.
func (o *CreditResource) HasCharacter() bool {
	if o != nil && o.Character.IsSet() {
		return true
	}

	return false
}

// SetCharacter gets a reference to the given NullableString and assigns it to the Character field.
func (o *CreditResource) SetCharacter(v string) {
	o.Character.Set(&v)
}
// SetCharacterNil sets the value for Character to be an explicit nil
func (o *CreditResource) SetCharacterNil() {
	o.Character.Set(nil)
}

// UnsetCharacter ensures that no value is present for Character, not even an explicit nil
func (o *CreditResource) UnsetCharacter() {
	o.Character.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *CreditResource) GetOrder() int32 {
	if o == nil || isNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditResource) GetOrderOk() (*int32, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *CreditResource) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *CreditResource) SetOrder(v int32) {
	o.Order = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreditResource) GetType() CreditType {
	if o == nil || isNil(o.Type) {
		var ret CreditType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditResource) GetTypeOk() (*CreditType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreditResource) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CreditType and assigns it to the Type field.
func (o *CreditResource) SetType(v CreditType) {
	o.Type = &v
}

func (o CreditResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.PersonName.IsSet() {
		toSerialize["personName"] = o.PersonName.Get()
	}
	if o.CreditTmdbId.IsSet() {
		toSerialize["creditTmdbId"] = o.CreditTmdbId.Get()
	}
	if !isNil(o.PersonTmdbId) {
		toSerialize["personTmdbId"] = o.PersonTmdbId
	}
	if !isNil(o.MovieMetadataId) {
		toSerialize["movieMetadataId"] = o.MovieMetadataId
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Department.IsSet() {
		toSerialize["department"] = o.Department.Get()
	}
	if o.Job.IsSet() {
		toSerialize["job"] = o.Job.Get()
	}
	if o.Character.IsSet() {
		toSerialize["character"] = o.Character.Get()
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCreditResource struct {
	value *CreditResource
	isSet bool
}

func (v NullableCreditResource) Get() *CreditResource {
	return v.value
}

func (v *NullableCreditResource) Set(val *CreditResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditResource(val *CreditResource) *NullableCreditResource {
	return &NullableCreditResource{value: val, isSet: true}
}

func (v NullableCreditResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

