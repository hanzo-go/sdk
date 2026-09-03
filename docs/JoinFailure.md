# JoinFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is Slack&#39;s id for the room that refused. | [optional] 
**Error** | Pointer to **string** | Error is Slack&#39;s own code, carried through unchanged. | [optional] 
**Name** | Pointer to **string** | Name is that room&#39;s human name, so the operator does not have to look the id up. | [optional] 

## Methods

### NewJoinFailure

`func NewJoinFailure() *JoinFailure`

NewJoinFailure instantiates a new JoinFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinFailureWithDefaults

`func NewJoinFailureWithDefaults() *JoinFailure`

NewJoinFailureWithDefaults instantiates a new JoinFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *JoinFailure) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *JoinFailure) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *JoinFailure) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *JoinFailure) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetError

`func (o *JoinFailure) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *JoinFailure) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *JoinFailure) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *JoinFailure) HasError() bool`

HasError returns a boolean if a field has been set.

### GetName

`func (o *JoinFailure) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JoinFailure) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JoinFailure) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JoinFailure) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


