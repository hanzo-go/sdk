# SessionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | Pointer to [**[]SessionView**](SessionView.md) | Sessions is the matching sessions, each with its event and child counts and a one-line preview of its latest event. | [optional] 

## Methods

### NewSessionList

`func NewSessionList() *SessionList`

NewSessionList instantiates a new SessionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionListWithDefaults

`func NewSessionListWithDefaults() *SessionList`

NewSessionListWithDefaults instantiates a new SessionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *SessionList) GetSessions() []SessionView`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *SessionList) GetSessionsOk() (*[]SessionView, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *SessionList) SetSessions(v []SessionView)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *SessionList) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


