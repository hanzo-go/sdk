# CloudSessionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | Pointer to [**[]CloudSessionView**](CloudSessionView.md) | Sessions is the matching sessions, each with its event and child counts and a one-line preview of its latest event. | [optional] 

## Methods

### NewCloudSessionList

`func NewCloudSessionList() *CloudSessionList`

NewCloudSessionList instantiates a new CloudSessionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSessionListWithDefaults

`func NewCloudSessionListWithDefaults() *CloudSessionList`

NewCloudSessionListWithDefaults instantiates a new CloudSessionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *CloudSessionList) GetSessions() []CloudSessionView`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *CloudSessionList) GetSessionsOk() (*[]CloudSessionView, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *CloudSessionList) SetSessions(v []CloudSessionView)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *CloudSessionList) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


