# CaptableTransferred

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message is the human sentence the cap table wrote, e.g. \&quot;Share transferred\&quot;. | [optional] 
**NewShareId** | Pointer to **string** | NewShareID names the certificate a PARTIAL transfer created. It is null on a full transfer, which reassigns the existing certificate instead of splitting it — so null here means \&quot;no new certificate\&quot;, never \&quot;the transfer failed\&quot;. | [optional] 
**Success** | Pointer to **bool** | Success is true when the transfer was applied. | [optional] 
**Transferred** | Pointer to **int64** | Transferred is how many shares moved. | [optional] 

## Methods

### NewCaptableTransferred

`func NewCaptableTransferred() *CaptableTransferred`

NewCaptableTransferred instantiates a new CaptableTransferred object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableTransferredWithDefaults

`func NewCaptableTransferredWithDefaults() *CaptableTransferred`

NewCaptableTransferredWithDefaults instantiates a new CaptableTransferred object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CaptableTransferred) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CaptableTransferred) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CaptableTransferred) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CaptableTransferred) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewShareId

`func (o *CaptableTransferred) GetNewShareId() string`

GetNewShareId returns the NewShareId field if non-nil, zero value otherwise.

### GetNewShareIdOk

`func (o *CaptableTransferred) GetNewShareIdOk() (*string, bool)`

GetNewShareIdOk returns a tuple with the NewShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewShareId

`func (o *CaptableTransferred) SetNewShareId(v string)`

SetNewShareId sets NewShareId field to given value.

### HasNewShareId

`func (o *CaptableTransferred) HasNewShareId() bool`

HasNewShareId returns a boolean if a field has been set.

### GetSuccess

`func (o *CaptableTransferred) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CaptableTransferred) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CaptableTransferred) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CaptableTransferred) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTransferred

`func (o *CaptableTransferred) GetTransferred() int64`

GetTransferred returns the Transferred field if non-nil, zero value otherwise.

### GetTransferredOk

`func (o *CaptableTransferred) GetTransferredOk() (*int64, bool)`

GetTransferredOk returns a tuple with the Transferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferred

`func (o *CaptableTransferred) SetTransferred(v int64)`

SetTransferred sets Transferred field to given value.

### HasTransferred

`func (o *CaptableTransferred) HasTransferred() bool`

HasTransferred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


