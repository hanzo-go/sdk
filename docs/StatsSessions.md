# StatsSessions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSessions** | Pointer to [**map[string][]StatsUser**](array.md) | ActiveSessions maps a workspace uuid to its connected sessions. It carries only the token&#39;s OWN workspace, and is empty for a token that names none. | [optional] 

## Methods

### NewStatsSessions

`func NewStatsSessions() *StatsSessions`

NewStatsSessions instantiates a new StatsSessions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsSessionsWithDefaults

`func NewStatsSessionsWithDefaults() *StatsSessions`

NewStatsSessionsWithDefaults instantiates a new StatsSessions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSessions

`func (o *StatsSessions) GetActiveSessions() map[string][]StatsUser`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *StatsSessions) GetActiveSessionsOk() (*map[string][]StatsUser, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *StatsSessions) SetActiveSessions(v map[string][]StatsUser)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *StatsSessions) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


