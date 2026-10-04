# TeamStatsSessions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSessions** | Pointer to [**map[string][]TeamStatsUser**](array.md) | ActiveSessions maps a space uuid to its connected sessions. It carries only the token&#39;s OWN space, and is empty for a token that names none. | [optional] 

## Methods

### NewTeamStatsSessions

`func NewTeamStatsSessions() *TeamStatsSessions`

NewTeamStatsSessions instantiates a new TeamStatsSessions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamStatsSessionsWithDefaults

`func NewTeamStatsSessionsWithDefaults() *TeamStatsSessions`

NewTeamStatsSessionsWithDefaults instantiates a new TeamStatsSessions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSessions

`func (o *TeamStatsSessions) GetActiveSessions() map[string][]TeamStatsUser`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *TeamStatsSessions) GetActiveSessionsOk() (*map[string][]TeamStatsUser, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *TeamStatsSessions) SetActiveSessions(v map[string][]TeamStatsUser)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *TeamStatsSessions) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


