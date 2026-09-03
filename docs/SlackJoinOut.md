# SlackJoinOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Already** | Pointer to **int64** | Already counts the channels it was a member of before, kept apart from Joined because only one of the two is a change. | [optional] 
**Failed** | Pointer to [**[]JoinFailure**](JoinFailure.md) | Failed is per-channel, so one refusal does not hide the rest of the walk. | [optional] 
**Joined** | Pointer to **[]string** | Joined names the channels this run walked into — the change it made. | [optional] 
**Listed** | Pointer to **int64** | Listed is every public, unarchived channel the workspace has. | [optional] 

## Methods

### NewSlackJoinOut

`func NewSlackJoinOut() *SlackJoinOut`

NewSlackJoinOut instantiates a new SlackJoinOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackJoinOutWithDefaults

`func NewSlackJoinOutWithDefaults() *SlackJoinOut`

NewSlackJoinOutWithDefaults instantiates a new SlackJoinOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlready

`func (o *SlackJoinOut) GetAlready() int64`

GetAlready returns the Already field if non-nil, zero value otherwise.

### GetAlreadyOk

`func (o *SlackJoinOut) GetAlreadyOk() (*int64, bool)`

GetAlreadyOk returns a tuple with the Already field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlready

`func (o *SlackJoinOut) SetAlready(v int64)`

SetAlready sets Already field to given value.

### HasAlready

`func (o *SlackJoinOut) HasAlready() bool`

HasAlready returns a boolean if a field has been set.

### GetFailed

`func (o *SlackJoinOut) GetFailed() []JoinFailure`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *SlackJoinOut) GetFailedOk() (*[]JoinFailure, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *SlackJoinOut) SetFailed(v []JoinFailure)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *SlackJoinOut) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetJoined

`func (o *SlackJoinOut) GetJoined() []string`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *SlackJoinOut) GetJoinedOk() (*[]string, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *SlackJoinOut) SetJoined(v []string)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *SlackJoinOut) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### GetListed

`func (o *SlackJoinOut) GetListed() int64`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *SlackJoinOut) GetListedOk() (*int64, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *SlackJoinOut) SetListed(v int64)`

SetListed sets Listed field to given value.

### HasListed

`func (o *SlackJoinOut) HasListed() bool`

HasListed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


