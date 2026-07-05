# MqPullMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to **int32** | Number of messages to pull. | [optional] [default to 1]
**Expires** | Pointer to **string** | Maximum time to wait for messages (e.g., \&quot;5s\&quot;, \&quot;30s\&quot;). Defaults to 30s.  | [optional] 
**NoWait** | Pointer to **bool** | Return immediately if no messages are available instead of waiting.  | [optional] [default to false]

## Methods

### NewMqPullMessagesRequest

`func NewMqPullMessagesRequest() *MqPullMessagesRequest`

NewMqPullMessagesRequest instantiates a new MqPullMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqPullMessagesRequestWithDefaults

`func NewMqPullMessagesRequestWithDefaults() *MqPullMessagesRequest`

NewMqPullMessagesRequestWithDefaults instantiates a new MqPullMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *MqPullMessagesRequest) GetBatch() int32`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *MqPullMessagesRequest) GetBatchOk() (*int32, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *MqPullMessagesRequest) SetBatch(v int32)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *MqPullMessagesRequest) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetExpires

`func (o *MqPullMessagesRequest) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *MqPullMessagesRequest) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *MqPullMessagesRequest) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *MqPullMessagesRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetNoWait

`func (o *MqPullMessagesRequest) GetNoWait() bool`

GetNoWait returns the NoWait field if non-nil, zero value otherwise.

### GetNoWaitOk

`func (o *MqPullMessagesRequest) GetNoWaitOk() (*bool, bool)`

GetNoWaitOk returns a tuple with the NoWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWait

`func (o *MqPullMessagesRequest) SetNoWait(v bool)`

SetNoWait sets NoWait field to given value.

### HasNoWait

`func (o *MqPullMessagesRequest) HasNoWait() bool`

HasNoWait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


