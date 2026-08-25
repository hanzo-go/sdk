# CaptableInvested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the investment record&#39;s id. | [optional] 
**Message** | Pointer to **string** | Message is the human sentence the cap table wrote. | [optional] 
**NewShareId** | Pointer to **string** | NewShareID names the certificate the investment issued, when the round carries a price per share. Null when the round prices later, which is a recorded investment and not a failure. | [optional] 
**Success** | Pointer to **bool** | Success is true when the investment was recorded. | [optional] 

## Methods

### NewCaptableInvested

`func NewCaptableInvested() *CaptableInvested`

NewCaptableInvested instantiates a new CaptableInvested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableInvestedWithDefaults

`func NewCaptableInvestedWithDefaults() *CaptableInvested`

NewCaptableInvestedWithDefaults instantiates a new CaptableInvested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaptableInvested) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableInvested) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableInvested) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableInvested) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *CaptableInvested) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CaptableInvested) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CaptableInvested) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CaptableInvested) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewShareId

`func (o *CaptableInvested) GetNewShareId() string`

GetNewShareId returns the NewShareId field if non-nil, zero value otherwise.

### GetNewShareIdOk

`func (o *CaptableInvested) GetNewShareIdOk() (*string, bool)`

GetNewShareIdOk returns a tuple with the NewShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewShareId

`func (o *CaptableInvested) SetNewShareId(v string)`

SetNewShareId sets NewShareId field to given value.

### HasNewShareId

`func (o *CaptableInvested) HasNewShareId() bool`

HasNewShareId returns a boolean if a field has been set.

### GetSuccess

`func (o *CaptableInvested) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CaptableInvested) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CaptableInvested) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CaptableInvested) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


