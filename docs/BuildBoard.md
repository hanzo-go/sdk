# BuildBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Builds** | Pointer to [**[]BuildRow**](BuildRow.md) | Builds are the org&#39;s real BuildKit build records, newest first. | [optional] 

## Methods

### NewBuildBoard

`func NewBuildBoard() *BuildBoard`

NewBuildBoard instantiates a new BuildBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildBoardWithDefaults

`func NewBuildBoardWithDefaults() *BuildBoard`

NewBuildBoardWithDefaults instantiates a new BuildBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilds

`func (o *BuildBoard) GetBuilds() []BuildRow`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *BuildBoard) GetBuildsOk() (*[]BuildRow, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *BuildBoard) SetBuilds(v []BuildRow)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *BuildBoard) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


