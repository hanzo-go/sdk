# ArgoConnectionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptedAt** | Pointer to **string** | AttemptedAt is when the connection was last probed. Always absent: nothing is probed, and a fabricated timestamp would claim a check that never ran. | [optional] 
**Message** | Pointer to **string** | Message is why a connection failed. Always absent, since none does. | [optional] 
**Status** | Pointer to **string** | Status is ArgoCD&#39;s ConnectionStatus — Successful, Failed or Unknown. Always Successful here: the destination is the cluster this process is already running in, so it is reachable by construction and there is no credential to probe. | [optional] 

## Methods

### NewArgoConnectionState

`func NewArgoConnectionState() *ArgoConnectionState`

NewArgoConnectionState instantiates a new ArgoConnectionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoConnectionStateWithDefaults

`func NewArgoConnectionStateWithDefaults() *ArgoConnectionState`

NewArgoConnectionStateWithDefaults instantiates a new ArgoConnectionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptedAt

`func (o *ArgoConnectionState) GetAttemptedAt() string`

GetAttemptedAt returns the AttemptedAt field if non-nil, zero value otherwise.

### GetAttemptedAtOk

`func (o *ArgoConnectionState) GetAttemptedAtOk() (*string, bool)`

GetAttemptedAtOk returns a tuple with the AttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedAt

`func (o *ArgoConnectionState) SetAttemptedAt(v string)`

SetAttemptedAt sets AttemptedAt field to given value.

### HasAttemptedAt

`func (o *ArgoConnectionState) HasAttemptedAt() bool`

HasAttemptedAt returns a boolean if a field has been set.

### GetMessage

`func (o *ArgoConnectionState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArgoConnectionState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArgoConnectionState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArgoConnectionState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ArgoConnectionState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArgoConnectionState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArgoConnectionState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArgoConnectionState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


