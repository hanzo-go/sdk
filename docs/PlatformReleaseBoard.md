# PlatformReleaseBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Releases** | Pointer to [**[]PlatformReleaseRow**](PlatformReleaseRow.md) | Releases are the deployments that genuinely reached the cluster. | [optional] 

## Methods

### NewPlatformReleaseBoard

`func NewPlatformReleaseBoard() *PlatformReleaseBoard`

NewPlatformReleaseBoard instantiates a new PlatformReleaseBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformReleaseBoardWithDefaults

`func NewPlatformReleaseBoardWithDefaults() *PlatformReleaseBoard`

NewPlatformReleaseBoardWithDefaults instantiates a new PlatformReleaseBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleases

`func (o *PlatformReleaseBoard) GetReleases() []PlatformReleaseRow`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *PlatformReleaseBoard) GetReleasesOk() (*[]PlatformReleaseRow, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *PlatformReleaseBoard) SetReleases(v []PlatformReleaseRow)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *PlatformReleaseBoard) HasReleases() bool`

HasReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


