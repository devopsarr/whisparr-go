# AlternativeTitleResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SourceType** | Pointer to [**SourceType**](SourceType.md) |  | [optional] 
**MovieId** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**SourceId** | Pointer to **int32** |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**VoteCount** | Pointer to **int32** |  | [optional] 
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 

## Methods

### NewAlternativeTitleResource

`func NewAlternativeTitleResource() *AlternativeTitleResource`

NewAlternativeTitleResource instantiates a new AlternativeTitleResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeTitleResourceWithDefaults

`func NewAlternativeTitleResourceWithDefaults() *AlternativeTitleResource`

NewAlternativeTitleResourceWithDefaults instantiates a new AlternativeTitleResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlternativeTitleResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlternativeTitleResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlternativeTitleResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlternativeTitleResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceType

`func (o *AlternativeTitleResource) GetSourceType() SourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AlternativeTitleResource) GetSourceTypeOk() (*SourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AlternativeTitleResource) SetSourceType(v SourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AlternativeTitleResource) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetMovieId

`func (o *AlternativeTitleResource) GetMovieId() int32`

GetMovieId returns the MovieId field if non-nil, zero value otherwise.

### GetMovieIdOk

`func (o *AlternativeTitleResource) GetMovieIdOk() (*int32, bool)`

GetMovieIdOk returns a tuple with the MovieId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieId

`func (o *AlternativeTitleResource) SetMovieId(v int32)`

SetMovieId sets MovieId field to given value.

### HasMovieId

`func (o *AlternativeTitleResource) HasMovieId() bool`

HasMovieId returns a boolean if a field has been set.

### GetTitle

`func (o *AlternativeTitleResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlternativeTitleResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlternativeTitleResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlternativeTitleResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AlternativeTitleResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AlternativeTitleResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCleanTitle

`func (o *AlternativeTitleResource) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *AlternativeTitleResource) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *AlternativeTitleResource) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *AlternativeTitleResource) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *AlternativeTitleResource) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *AlternativeTitleResource) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetSourceId

`func (o *AlternativeTitleResource) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AlternativeTitleResource) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AlternativeTitleResource) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AlternativeTitleResource) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetVotes

`func (o *AlternativeTitleResource) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *AlternativeTitleResource) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *AlternativeTitleResource) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *AlternativeTitleResource) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVoteCount

`func (o *AlternativeTitleResource) GetVoteCount() int32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *AlternativeTitleResource) GetVoteCountOk() (*int32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *AlternativeTitleResource) SetVoteCount(v int32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *AlternativeTitleResource) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### GetLanguage

`func (o *AlternativeTitleResource) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AlternativeTitleResource) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AlternativeTitleResource) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AlternativeTitleResource) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


