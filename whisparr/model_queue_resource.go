/*
Radarr

Radarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whisparr

import (
	"encoding/json"
	"time"
)

// QueueResource struct for QueueResource
type QueueResource struct {
	Id *int32 `json:"id,omitempty"`
	MovieId NullableInt32 `json:"movieId,omitempty"`
	Movie *MovieResource `json:"movie,omitempty"`
	Languages []*Language `json:"languages,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	CustomFormats []*CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Sizeleft *float64 `json:"sizeleft,omitempty"`
	Timeleft *string `json:"timeleft,omitempty"`
	EstimatedCompletionTime NullableTime `json:"estimatedCompletionTime,omitempty"`
	Added NullableTime `json:"added,omitempty"`
	Status NullableString `json:"status,omitempty"`
	TrackedDownloadStatus *TrackedDownloadStatus `json:"trackedDownloadStatus,omitempty"`
	TrackedDownloadState *TrackedDownloadState `json:"trackedDownloadState,omitempty"`
	StatusMessages []*TrackedDownloadStatusMessage `json:"statusMessages,omitempty"`
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	DownloadId NullableString `json:"downloadId,omitempty"`
	Protocol *DownloadProtocol `json:"protocol,omitempty"`
	DownloadClient NullableString `json:"downloadClient,omitempty"`
	DownloadClientHasPostImportCategory *bool `json:"downloadClientHasPostImportCategory,omitempty"`
	Indexer NullableString `json:"indexer,omitempty"`
	OutputPath NullableString `json:"outputPath,omitempty"`
}

// NewQueueResource instantiates a new QueueResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueResource() *QueueResource {
	this := QueueResource{}
	return &this
}

// NewQueueResourceWithDefaults instantiates a new QueueResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueResourceWithDefaults() *QueueResource {
	this := QueueResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QueueResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QueueResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *QueueResource) SetId(v int32) {
	o.Id = &v
}

// GetMovieId returns the MovieId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetMovieId() int32 {
	if o == nil || isNil(o.MovieId.Get()) {
		var ret int32
		return ret
	}
	return *o.MovieId.Get()
}

// GetMovieIdOk returns a tuple with the MovieId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetMovieIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MovieId.Get(), o.MovieId.IsSet()
}

// HasMovieId returns a boolean if a field has been set.
func (o *QueueResource) HasMovieId() bool {
	if o != nil && o.MovieId.IsSet() {
		return true
	}

	return false
}

// SetMovieId gets a reference to the given NullableInt32 and assigns it to the MovieId field.
func (o *QueueResource) SetMovieId(v int32) {
	o.MovieId.Set(&v)
}
// SetMovieIdNil sets the value for MovieId to be an explicit nil
func (o *QueueResource) SetMovieIdNil() {
	o.MovieId.Set(nil)
}

// UnsetMovieId ensures that no value is present for MovieId, not even an explicit nil
func (o *QueueResource) UnsetMovieId() {
	o.MovieId.Unset()
}

// GetMovie returns the Movie field value if set, zero value otherwise.
func (o *QueueResource) GetMovie() MovieResource {
	if o == nil || isNil(o.Movie) {
		var ret MovieResource
		return ret
	}
	return *o.Movie
}

// GetMovieOk returns a tuple with the Movie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetMovieOk() (*MovieResource, bool) {
	if o == nil || isNil(o.Movie) {
    return nil, false
	}
	return o.Movie, true
}

// HasMovie returns a boolean if a field has been set.
func (o *QueueResource) HasMovie() bool {
	if o != nil && !isNil(o.Movie) {
		return true
	}

	return false
}

// SetMovie gets a reference to the given MovieResource and assigns it to the Movie field.
func (o *QueueResource) SetMovie(v MovieResource) {
	o.Movie = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetLanguages() []*Language {
	if o == nil {
		var ret []*Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetLanguagesOk() ([]*Language, bool) {
	if o == nil || isNil(o.Languages) {
    return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *QueueResource) HasLanguages() bool {
	if o != nil && isNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *QueueResource) SetLanguages(v []*Language) {
	o.Languages = v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *QueueResource) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *QueueResource) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *QueueResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetCustomFormats() []*CustomFormatResource {
	if o == nil {
		var ret []*CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetCustomFormatsOk() ([]*CustomFormatResource, bool) {
	if o == nil || isNil(o.CustomFormats) {
    return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *QueueResource) HasCustomFormats() bool {
	if o != nil && isNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *QueueResource) SetCustomFormats(v []*CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *QueueResource) GetCustomFormatScore() int32 {
	if o == nil || isNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || isNil(o.CustomFormatScore) {
    return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *QueueResource) HasCustomFormatScore() bool {
	if o != nil && !isNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *QueueResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *QueueResource) GetSize() float64 {
	if o == nil || isNil(o.Size) {
		var ret float64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetSizeOk() (*float64, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *QueueResource) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float64 and assigns it to the Size field.
func (o *QueueResource) SetSize(v float64) {
	o.Size = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *QueueResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *QueueResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *QueueResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *QueueResource) UnsetTitle() {
	o.Title.Unset()
}

// GetSizeleft returns the Sizeleft field value if set, zero value otherwise.
func (o *QueueResource) GetSizeleft() float64 {
	if o == nil || isNil(o.Sizeleft) {
		var ret float64
		return ret
	}
	return *o.Sizeleft
}

// GetSizeleftOk returns a tuple with the Sizeleft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetSizeleftOk() (*float64, bool) {
	if o == nil || isNil(o.Sizeleft) {
    return nil, false
	}
	return o.Sizeleft, true
}

// HasSizeleft returns a boolean if a field has been set.
func (o *QueueResource) HasSizeleft() bool {
	if o != nil && !isNil(o.Sizeleft) {
		return true
	}

	return false
}

// SetSizeleft gets a reference to the given float64 and assigns it to the Sizeleft field.
func (o *QueueResource) SetSizeleft(v float64) {
	o.Sizeleft = &v
}

// GetTimeleft returns the Timeleft field value if set, zero value otherwise.
func (o *QueueResource) GetTimeleft() string {
	if o == nil || isNil(o.Timeleft) {
		var ret string
		return ret
	}
	return *o.Timeleft
}

// GetTimeleftOk returns a tuple with the Timeleft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetTimeleftOk() (*string, bool) {
	if o == nil || isNil(o.Timeleft) {
    return nil, false
	}
	return o.Timeleft, true
}

// HasTimeleft returns a boolean if a field has been set.
func (o *QueueResource) HasTimeleft() bool {
	if o != nil && !isNil(o.Timeleft) {
		return true
	}

	return false
}

// SetTimeleft gets a reference to the given string and assigns it to the Timeleft field.
func (o *QueueResource) SetTimeleft(v string) {
	o.Timeleft = &v
}

// GetEstimatedCompletionTime returns the EstimatedCompletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetEstimatedCompletionTime() time.Time {
	if o == nil || isNil(o.EstimatedCompletionTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedCompletionTime.Get()
}

// GetEstimatedCompletionTimeOk returns a tuple with the EstimatedCompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetEstimatedCompletionTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.EstimatedCompletionTime.Get(), o.EstimatedCompletionTime.IsSet()
}

// HasEstimatedCompletionTime returns a boolean if a field has been set.
func (o *QueueResource) HasEstimatedCompletionTime() bool {
	if o != nil && o.EstimatedCompletionTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedCompletionTime gets a reference to the given NullableTime and assigns it to the EstimatedCompletionTime field.
func (o *QueueResource) SetEstimatedCompletionTime(v time.Time) {
	o.EstimatedCompletionTime.Set(&v)
}
// SetEstimatedCompletionTimeNil sets the value for EstimatedCompletionTime to be an explicit nil
func (o *QueueResource) SetEstimatedCompletionTimeNil() {
	o.EstimatedCompletionTime.Set(nil)
}

// UnsetEstimatedCompletionTime ensures that no value is present for EstimatedCompletionTime, not even an explicit nil
func (o *QueueResource) UnsetEstimatedCompletionTime() {
	o.EstimatedCompletionTime.Unset()
}

// GetAdded returns the Added field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetAdded() time.Time {
	if o == nil || isNil(o.Added.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Added.Get()
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetAddedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.Added.Get(), o.Added.IsSet()
}

// HasAdded returns a boolean if a field has been set.
func (o *QueueResource) HasAdded() bool {
	if o != nil && o.Added.IsSet() {
		return true
	}

	return false
}

// SetAdded gets a reference to the given NullableTime and assigns it to the Added field.
func (o *QueueResource) SetAdded(v time.Time) {
	o.Added.Set(&v)
}
// SetAddedNil sets the value for Added to be an explicit nil
func (o *QueueResource) SetAddedNil() {
	o.Added.Set(nil)
}

// UnsetAdded ensures that no value is present for Added, not even an explicit nil
func (o *QueueResource) UnsetAdded() {
	o.Added.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetStatus() string {
	if o == nil || isNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *QueueResource) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *QueueResource) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *QueueResource) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *QueueResource) UnsetStatus() {
	o.Status.Unset()
}

// GetTrackedDownloadStatus returns the TrackedDownloadStatus field value if set, zero value otherwise.
func (o *QueueResource) GetTrackedDownloadStatus() TrackedDownloadStatus {
	if o == nil || isNil(o.TrackedDownloadStatus) {
		var ret TrackedDownloadStatus
		return ret
	}
	return *o.TrackedDownloadStatus
}

// GetTrackedDownloadStatusOk returns a tuple with the TrackedDownloadStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetTrackedDownloadStatusOk() (*TrackedDownloadStatus, bool) {
	if o == nil || isNil(o.TrackedDownloadStatus) {
    return nil, false
	}
	return o.TrackedDownloadStatus, true
}

// HasTrackedDownloadStatus returns a boolean if a field has been set.
func (o *QueueResource) HasTrackedDownloadStatus() bool {
	if o != nil && !isNil(o.TrackedDownloadStatus) {
		return true
	}

	return false
}

// SetTrackedDownloadStatus gets a reference to the given TrackedDownloadStatus and assigns it to the TrackedDownloadStatus field.
func (o *QueueResource) SetTrackedDownloadStatus(v TrackedDownloadStatus) {
	o.TrackedDownloadStatus = &v
}

// GetTrackedDownloadState returns the TrackedDownloadState field value if set, zero value otherwise.
func (o *QueueResource) GetTrackedDownloadState() TrackedDownloadState {
	if o == nil || isNil(o.TrackedDownloadState) {
		var ret TrackedDownloadState
		return ret
	}
	return *o.TrackedDownloadState
}

// GetTrackedDownloadStateOk returns a tuple with the TrackedDownloadState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetTrackedDownloadStateOk() (*TrackedDownloadState, bool) {
	if o == nil || isNil(o.TrackedDownloadState) {
    return nil, false
	}
	return o.TrackedDownloadState, true
}

// HasTrackedDownloadState returns a boolean if a field has been set.
func (o *QueueResource) HasTrackedDownloadState() bool {
	if o != nil && !isNil(o.TrackedDownloadState) {
		return true
	}

	return false
}

// SetTrackedDownloadState gets a reference to the given TrackedDownloadState and assigns it to the TrackedDownloadState field.
func (o *QueueResource) SetTrackedDownloadState(v TrackedDownloadState) {
	o.TrackedDownloadState = &v
}

// GetStatusMessages returns the StatusMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetStatusMessages() []*TrackedDownloadStatusMessage {
	if o == nil {
		var ret []*TrackedDownloadStatusMessage
		return ret
	}
	return o.StatusMessages
}

// GetStatusMessagesOk returns a tuple with the StatusMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetStatusMessagesOk() ([]*TrackedDownloadStatusMessage, bool) {
	if o == nil || isNil(o.StatusMessages) {
    return nil, false
	}
	return o.StatusMessages, true
}

// HasStatusMessages returns a boolean if a field has been set.
func (o *QueueResource) HasStatusMessages() bool {
	if o != nil && isNil(o.StatusMessages) {
		return true
	}

	return false
}

// SetStatusMessages gets a reference to the given []TrackedDownloadStatusMessage and assigns it to the StatusMessages field.
func (o *QueueResource) SetStatusMessages(v []*TrackedDownloadStatusMessage) {
	o.StatusMessages = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetErrorMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *QueueResource) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *QueueResource) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *QueueResource) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *QueueResource) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetDownloadId() string {
	if o == nil || isNil(o.DownloadId.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadId.Get()
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetDownloadIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DownloadId.Get(), o.DownloadId.IsSet()
}

// HasDownloadId returns a boolean if a field has been set.
func (o *QueueResource) HasDownloadId() bool {
	if o != nil && o.DownloadId.IsSet() {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given NullableString and assigns it to the DownloadId field.
func (o *QueueResource) SetDownloadId(v string) {
	o.DownloadId.Set(&v)
}
// SetDownloadIdNil sets the value for DownloadId to be an explicit nil
func (o *QueueResource) SetDownloadIdNil() {
	o.DownloadId.Set(nil)
}

// UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
func (o *QueueResource) UnsetDownloadId() {
	o.DownloadId.Unset()
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *QueueResource) GetProtocol() DownloadProtocol {
	if o == nil || isNil(o.Protocol) {
		var ret DownloadProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetProtocolOk() (*DownloadProtocol, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *QueueResource) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given DownloadProtocol and assigns it to the Protocol field.
func (o *QueueResource) SetProtocol(v DownloadProtocol) {
	o.Protocol = &v
}

// GetDownloadClient returns the DownloadClient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetDownloadClient() string {
	if o == nil || isNil(o.DownloadClient.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadClient.Get()
}

// GetDownloadClientOk returns a tuple with the DownloadClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetDownloadClientOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DownloadClient.Get(), o.DownloadClient.IsSet()
}

// HasDownloadClient returns a boolean if a field has been set.
func (o *QueueResource) HasDownloadClient() bool {
	if o != nil && o.DownloadClient.IsSet() {
		return true
	}

	return false
}

// SetDownloadClient gets a reference to the given NullableString and assigns it to the DownloadClient field.
func (o *QueueResource) SetDownloadClient(v string) {
	o.DownloadClient.Set(&v)
}
// SetDownloadClientNil sets the value for DownloadClient to be an explicit nil
func (o *QueueResource) SetDownloadClientNil() {
	o.DownloadClient.Set(nil)
}

// UnsetDownloadClient ensures that no value is present for DownloadClient, not even an explicit nil
func (o *QueueResource) UnsetDownloadClient() {
	o.DownloadClient.Unset()
}

// GetDownloadClientHasPostImportCategory returns the DownloadClientHasPostImportCategory field value if set, zero value otherwise.
func (o *QueueResource) GetDownloadClientHasPostImportCategory() bool {
	if o == nil || isNil(o.DownloadClientHasPostImportCategory) {
		var ret bool
		return ret
	}
	return *o.DownloadClientHasPostImportCategory
}

// GetDownloadClientHasPostImportCategoryOk returns a tuple with the DownloadClientHasPostImportCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResource) GetDownloadClientHasPostImportCategoryOk() (*bool, bool) {
	if o == nil || isNil(o.DownloadClientHasPostImportCategory) {
    return nil, false
	}
	return o.DownloadClientHasPostImportCategory, true
}

// HasDownloadClientHasPostImportCategory returns a boolean if a field has been set.
func (o *QueueResource) HasDownloadClientHasPostImportCategory() bool {
	if o != nil && !isNil(o.DownloadClientHasPostImportCategory) {
		return true
	}

	return false
}

// SetDownloadClientHasPostImportCategory gets a reference to the given bool and assigns it to the DownloadClientHasPostImportCategory field.
func (o *QueueResource) SetDownloadClientHasPostImportCategory(v bool) {
	o.DownloadClientHasPostImportCategory = &v
}

// GetIndexer returns the Indexer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetIndexer() string {
	if o == nil || isNil(o.Indexer.Get()) {
		var ret string
		return ret
	}
	return *o.Indexer.Get()
}

// GetIndexerOk returns a tuple with the Indexer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetIndexerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Indexer.Get(), o.Indexer.IsSet()
}

// HasIndexer returns a boolean if a field has been set.
func (o *QueueResource) HasIndexer() bool {
	if o != nil && o.Indexer.IsSet() {
		return true
	}

	return false
}

// SetIndexer gets a reference to the given NullableString and assigns it to the Indexer field.
func (o *QueueResource) SetIndexer(v string) {
	o.Indexer.Set(&v)
}
// SetIndexerNil sets the value for Indexer to be an explicit nil
func (o *QueueResource) SetIndexerNil() {
	o.Indexer.Set(nil)
}

// UnsetIndexer ensures that no value is present for Indexer, not even an explicit nil
func (o *QueueResource) UnsetIndexer() {
	o.Indexer.Unset()
}

// GetOutputPath returns the OutputPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResource) GetOutputPath() string {
	if o == nil || isNil(o.OutputPath.Get()) {
		var ret string
		return ret
	}
	return *o.OutputPath.Get()
}

// GetOutputPathOk returns a tuple with the OutputPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResource) GetOutputPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.OutputPath.Get(), o.OutputPath.IsSet()
}

// HasOutputPath returns a boolean if a field has been set.
func (o *QueueResource) HasOutputPath() bool {
	if o != nil && o.OutputPath.IsSet() {
		return true
	}

	return false
}

// SetOutputPath gets a reference to the given NullableString and assigns it to the OutputPath field.
func (o *QueueResource) SetOutputPath(v string) {
	o.OutputPath.Set(&v)
}
// SetOutputPathNil sets the value for OutputPath to be an explicit nil
func (o *QueueResource) SetOutputPathNil() {
	o.OutputPath.Set(nil)
}

// UnsetOutputPath ensures that no value is present for OutputPath, not even an explicit nil
func (o *QueueResource) UnsetOutputPath() {
	o.OutputPath.Unset()
}

func (o QueueResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.MovieId.IsSet() {
		toSerialize["movieId"] = o.MovieId.Get()
	}
	if !isNil(o.Movie) {
		toSerialize["movie"] = o.Movie
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
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
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !isNil(o.Sizeleft) {
		toSerialize["sizeleft"] = o.Sizeleft
	}
	if !isNil(o.Timeleft) {
		toSerialize["timeleft"] = o.Timeleft
	}
	if o.EstimatedCompletionTime.IsSet() {
		toSerialize["estimatedCompletionTime"] = o.EstimatedCompletionTime.Get()
	}
	if o.Added.IsSet() {
		toSerialize["added"] = o.Added.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !isNil(o.TrackedDownloadStatus) {
		toSerialize["trackedDownloadStatus"] = o.TrackedDownloadStatus
	}
	if !isNil(o.TrackedDownloadState) {
		toSerialize["trackedDownloadState"] = o.TrackedDownloadState
	}
	if o.StatusMessages != nil {
		toSerialize["statusMessages"] = o.StatusMessages
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.DownloadId.IsSet() {
		toSerialize["downloadId"] = o.DownloadId.Get()
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if o.DownloadClient.IsSet() {
		toSerialize["downloadClient"] = o.DownloadClient.Get()
	}
	if !isNil(o.DownloadClientHasPostImportCategory) {
		toSerialize["downloadClientHasPostImportCategory"] = o.DownloadClientHasPostImportCategory
	}
	if o.Indexer.IsSet() {
		toSerialize["indexer"] = o.Indexer.Get()
	}
	if o.OutputPath.IsSet() {
		toSerialize["outputPath"] = o.OutputPath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableQueueResource struct {
	value *QueueResource
	isSet bool
}

func (v NullableQueueResource) Get() *QueueResource {
	return v.value
}

func (v *NullableQueueResource) Set(val *QueueResource) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueResource) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueResource(val *QueueResource) *NullableQueueResource {
	return &NullableQueueResource{value: val, isSet: true}
}

func (v NullableQueueResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


