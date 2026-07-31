# CloudCaptableUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message is the human sentence the cap table wrote, e.g. \&quot;Company updated\&quot;. | [optional] 
**Success** | Pointer to **bool** | Success is true when the update was applied. | [optional] 

## Methods

### NewCloudCaptableUpdated

`func NewCloudCaptableUpdated() *CloudCaptableUpdated`

NewCloudCaptableUpdated instantiates a new CloudCaptableUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableUpdatedWithDefaults

`func NewCloudCaptableUpdatedWithDefaults() *CloudCaptableUpdated`

NewCloudCaptableUpdatedWithDefaults instantiates a new CloudCaptableUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CloudCaptableUpdated) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudCaptableUpdated) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudCaptableUpdated) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudCaptableUpdated) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *CloudCaptableUpdated) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CloudCaptableUpdated) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CloudCaptableUpdated) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CloudCaptableUpdated) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


