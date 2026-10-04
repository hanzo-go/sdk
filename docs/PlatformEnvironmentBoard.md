# PlatformEnvironmentBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**[]PlatformEnvironmentRow**](PlatformEnvironmentRow.md) | Environments are the org&#39;s deploy targets, in first-seen order. | [optional] 

## Methods

### NewPlatformEnvironmentBoard

`func NewPlatformEnvironmentBoard() *PlatformEnvironmentBoard`

NewPlatformEnvironmentBoard instantiates a new PlatformEnvironmentBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformEnvironmentBoardWithDefaults

`func NewPlatformEnvironmentBoardWithDefaults() *PlatformEnvironmentBoard`

NewPlatformEnvironmentBoardWithDefaults instantiates a new PlatformEnvironmentBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *PlatformEnvironmentBoard) GetEnvironments() []PlatformEnvironmentRow`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PlatformEnvironmentBoard) GetEnvironmentsOk() (*[]PlatformEnvironmentRow, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PlatformEnvironmentBoard) SetEnvironments(v []PlatformEnvironmentRow)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PlatformEnvironmentBoard) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


