# HistoryResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MovieId** | Pointer to **int32** |  | [optional] 
**SourceTitle** | Pointer to **NullableString** |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**QualityCutoffNotMet** | Pointer to **bool** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**EventType** | Pointer to [**MovieHistoryEventType**](MovieHistoryEventType.md) |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 
**Movie** | Pointer to [**MovieResource**](MovieResource.md) |  | [optional] 

## Methods

### NewHistoryResource

`func NewHistoryResource() *HistoryResource`

NewHistoryResource instantiates a new HistoryResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryResourceWithDefaults

`func NewHistoryResourceWithDefaults() *HistoryResource`

NewHistoryResourceWithDefaults instantiates a new HistoryResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HistoryResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMovieId

`func (o *HistoryResource) GetMovieId() int32`

GetMovieId returns the MovieId field if non-nil, zero value otherwise.

### GetMovieIdOk

`func (o *HistoryResource) GetMovieIdOk() (*int32, bool)`

GetMovieIdOk returns a tuple with the MovieId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieId

`func (o *HistoryResource) SetMovieId(v int32)`

SetMovieId sets MovieId field to given value.

### HasMovieId

`func (o *HistoryResource) HasMovieId() bool`

HasMovieId returns a boolean if a field has been set.

### GetSourceTitle

`func (o *HistoryResource) GetSourceTitle() string`

GetSourceTitle returns the SourceTitle field if non-nil, zero value otherwise.

### GetSourceTitleOk

`func (o *HistoryResource) GetSourceTitleOk() (*string, bool)`

GetSourceTitleOk returns a tuple with the SourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTitle

`func (o *HistoryResource) SetSourceTitle(v string)`

SetSourceTitle sets SourceTitle field to given value.

### HasSourceTitle

`func (o *HistoryResource) HasSourceTitle() bool`

HasSourceTitle returns a boolean if a field has been set.

### SetSourceTitleNil

`func (o *HistoryResource) SetSourceTitleNil(b bool)`

 SetSourceTitleNil sets the value for SourceTitle to be an explicit nil

### UnsetSourceTitle
`func (o *HistoryResource) UnsetSourceTitle()`

UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
### GetLanguages

`func (o *HistoryResource) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *HistoryResource) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *HistoryResource) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *HistoryResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *HistoryResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *HistoryResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetQuality

`func (o *HistoryResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *HistoryResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *HistoryResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *HistoryResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetCustomFormats

`func (o *HistoryResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *HistoryResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *HistoryResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *HistoryResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *HistoryResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *HistoryResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetQualityCutoffNotMet

`func (o *HistoryResource) GetQualityCutoffNotMet() bool`

GetQualityCutoffNotMet returns the QualityCutoffNotMet field if non-nil, zero value otherwise.

### GetQualityCutoffNotMetOk

`func (o *HistoryResource) GetQualityCutoffNotMetOk() (*bool, bool)`

GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityCutoffNotMet

`func (o *HistoryResource) SetQualityCutoffNotMet(v bool)`

SetQualityCutoffNotMet sets QualityCutoffNotMet field to given value.

### HasQualityCutoffNotMet

`func (o *HistoryResource) HasQualityCutoffNotMet() bool`

HasQualityCutoffNotMet returns a boolean if a field has been set.

### GetDate

`func (o *HistoryResource) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HistoryResource) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HistoryResource) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *HistoryResource) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDownloadId

`func (o *HistoryResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *HistoryResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *HistoryResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *HistoryResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *HistoryResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *HistoryResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetEventType

`func (o *HistoryResource) GetEventType() MovieHistoryEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *HistoryResource) GetEventTypeOk() (*MovieHistoryEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *HistoryResource) SetEventType(v MovieHistoryEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *HistoryResource) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetData

`func (o *HistoryResource) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryResource) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoryResource) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *HistoryResource) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *HistoryResource) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *HistoryResource) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMovie

`func (o *HistoryResource) GetMovie() MovieResource`

GetMovie returns the Movie field if non-nil, zero value otherwise.

### GetMovieOk

`func (o *HistoryResource) GetMovieOk() (*MovieResource, bool)`

GetMovieOk returns a tuple with the Movie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovie

`func (o *HistoryResource) SetMovie(v MovieResource)`

SetMovie sets Movie field to given value.

### HasMovie

`func (o *HistoryResource) HasMovie() bool`

HasMovie returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


