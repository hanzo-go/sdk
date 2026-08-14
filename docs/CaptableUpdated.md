# CaptableUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message is the human sentence the cap table wrote, e.g. \&quot;Company updated\&quot;. | [optional] 
**Success** | Pointer to **bool** | Success is true when the update was applied. | [optional] 

## Methods

### NewCaptableUpdated

`func NewCaptableUpdated() *CaptableUpdated`

NewCaptableUpdated instantiates a new CaptableUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableUpdatedWithDefaults

`func NewCaptableUpdatedWithDefaults() *CaptableUpdated`

NewCaptableUpdatedWithDefaults instantiates a new CaptableUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CaptableUpdated) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CaptableUpdated) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CaptableUpdated) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CaptableUpdated) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *CaptableUpdated) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CaptableUpdated) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CaptableUpdated) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CaptableUpdated) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


