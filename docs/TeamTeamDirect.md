# TeamTeamDirect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**TeamTeamChannel**](TeamTeamChannel.md) | Channel is the direct message, as the channel listing answers it. | [optional] 
**Created** | Pointer to **bool** | Created is true when this call opened the conversation (201) and false when it already existed (200). | [optional] 

## Methods

### NewTeamTeamDirect

`func NewTeamTeamDirect() *TeamTeamDirect`

NewTeamTeamDirect instantiates a new TeamTeamDirect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamDirectWithDefaults

`func NewTeamTeamDirectWithDefaults() *TeamTeamDirect`

NewTeamTeamDirectWithDefaults instantiates a new TeamTeamDirect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *TeamTeamDirect) GetChannel() TeamTeamChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TeamTeamDirect) GetChannelOk() (*TeamTeamChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TeamTeamDirect) SetChannel(v TeamTeamChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TeamTeamDirect) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreated

`func (o *TeamTeamDirect) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TeamTeamDirect) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TeamTeamDirect) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TeamTeamDirect) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


