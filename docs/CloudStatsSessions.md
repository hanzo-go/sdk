# CloudStatsSessions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSessions** | Pointer to [**map[string][]CloudStatsUser**](array.md) | ActiveSessions maps a workspace uuid to its connected sessions. It carries only the token&#39;s OWN workspace, and is empty for a token that names none. | [optional] 

## Methods

### NewCloudStatsSessions

`func NewCloudStatsSessions() *CloudStatsSessions`

NewCloudStatsSessions instantiates a new CloudStatsSessions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStatsSessionsWithDefaults

`func NewCloudStatsSessionsWithDefaults() *CloudStatsSessions`

NewCloudStatsSessionsWithDefaults instantiates a new CloudStatsSessions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSessions

`func (o *CloudStatsSessions) GetActiveSessions() map[string][]CloudStatsUser`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *CloudStatsSessions) GetActiveSessionsOk() (*map[string][]CloudStatsUser, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *CloudStatsSessions) SetActiveSessions(v map[string][]CloudStatsUser)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *CloudStatsSessions) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


