# ReleaseBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Releases** | Pointer to [**[]ReleaseRow**](ReleaseRow.md) | Releases are the deployments that genuinely reached the cluster. | [optional] 

## Methods

### NewReleaseBoard

`func NewReleaseBoard() *ReleaseBoard`

NewReleaseBoard instantiates a new ReleaseBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseBoardWithDefaults

`func NewReleaseBoardWithDefaults() *ReleaseBoard`

NewReleaseBoardWithDefaults instantiates a new ReleaseBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleases

`func (o *ReleaseBoard) GetReleases() []ReleaseRow`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *ReleaseBoard) GetReleasesOk() (*[]ReleaseRow, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *ReleaseBoard) SetReleases(v []ReleaseRow)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *ReleaseBoard) HasReleases() bool`

HasReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


