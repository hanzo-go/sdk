# CaptableCaptableInvested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the investment record&#39;s id. | [optional] 
**Message** | Pointer to **string** | Message is the human sentence the cap table wrote. | [optional] 
**NewShareId** | Pointer to **string** | NewShareID names the certificate the investment issued, when the round carries a price per share. Null when the round prices later, which is a recorded investment and not a failure. | [optional] 
**Success** | Pointer to **bool** | Success is true when the investment was recorded. | [optional] 

## Methods

### NewCaptableCaptableInvested

`func NewCaptableCaptableInvested() *CaptableCaptableInvested`

NewCaptableCaptableInvested instantiates a new CaptableCaptableInvested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableInvestedWithDefaults

`func NewCaptableCaptableInvestedWithDefaults() *CaptableCaptableInvested`

NewCaptableCaptableInvestedWithDefaults instantiates a new CaptableCaptableInvested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaptableCaptableInvested) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableCaptableInvested) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableCaptableInvested) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableCaptableInvested) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *CaptableCaptableInvested) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CaptableCaptableInvested) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CaptableCaptableInvested) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CaptableCaptableInvested) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewShareId

`func (o *CaptableCaptableInvested) GetNewShareId() string`

GetNewShareId returns the NewShareId field if non-nil, zero value otherwise.

### GetNewShareIdOk

`func (o *CaptableCaptableInvested) GetNewShareIdOk() (*string, bool)`

GetNewShareIdOk returns a tuple with the NewShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewShareId

`func (o *CaptableCaptableInvested) SetNewShareId(v string)`

SetNewShareId sets NewShareId field to given value.

### HasNewShareId

`func (o *CaptableCaptableInvested) HasNewShareId() bool`

HasNewShareId returns a boolean if a field has been set.

### GetSuccess

`func (o *CaptableCaptableInvested) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CaptableCaptableInvested) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CaptableCaptableInvested) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CaptableCaptableInvested) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


