# CaptableCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the created row&#39;s id. | [optional] 
**Message** | Pointer to **string** | Message is the human sentence the cap table wrote, e.g. \&quot;Share issued\&quot;. | [optional] 
**Success** | Pointer to **bool** | Success is true when the row was written. | [optional] 

## Methods

### NewCaptableCreated

`func NewCaptableCreated() *CaptableCreated`

NewCaptableCreated instantiates a new CaptableCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCreatedWithDefaults

`func NewCaptableCreatedWithDefaults() *CaptableCreated`

NewCaptableCreatedWithDefaults instantiates a new CaptableCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaptableCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableCreated) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableCreated) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *CaptableCreated) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CaptableCreated) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CaptableCreated) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CaptableCreated) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *CaptableCreated) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CaptableCreated) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CaptableCreated) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CaptableCreated) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


