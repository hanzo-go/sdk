# EnvironmentBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**[]EnvironmentRow**](EnvironmentRow.md) | Environments are the org&#39;s deploy targets, in first-seen order. | [optional] 

## Methods

### NewEnvironmentBoard

`func NewEnvironmentBoard() *EnvironmentBoard`

NewEnvironmentBoard instantiates a new EnvironmentBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentBoardWithDefaults

`func NewEnvironmentBoardWithDefaults() *EnvironmentBoard`

NewEnvironmentBoardWithDefaults instantiates a new EnvironmentBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *EnvironmentBoard) GetEnvironments() []EnvironmentRow`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *EnvironmentBoard) GetEnvironmentsOk() (*[]EnvironmentRow, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *EnvironmentBoard) SetEnvironments(v []EnvironmentRow)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *EnvironmentBoard) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


