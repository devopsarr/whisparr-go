# MovieCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 

## Methods

### NewMovieCollection

`func NewMovieCollection() *MovieCollection`

NewMovieCollection instantiates a new MovieCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieCollectionWithDefaults

`func NewMovieCollectionWithDefaults() *MovieCollection`

NewMovieCollectionWithDefaults instantiates a new MovieCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MovieCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MovieCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MovieCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MovieCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MovieCollection) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MovieCollection) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTmdbId

`func (o *MovieCollection) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *MovieCollection) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *MovieCollection) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *MovieCollection) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetImages

`func (o *MovieCollection) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *MovieCollection) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *MovieCollection) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *MovieCollection) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *MovieCollection) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *MovieCollection) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


