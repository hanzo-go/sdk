# O11yO11yQueueCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** | Attribute is the span attribute or telemetry the check looked for. | [optional] 
**ErrorMessage** | Pointer to **string** | Message says what is missing when the check fails; empty on a pass. Its wire key is error_message. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;1\&quot; when the telemetry is present, \&quot;0\&quot; when it is not. | [optional] 

## Methods

### NewO11yO11yQueueCheck

`func NewO11yO11yQueueCheck() *O11yO11yQueueCheck`

NewO11yO11yQueueCheck instantiates a new O11yO11yQueueCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueueCheckWithDefaults

`func NewO11yO11yQueueCheckWithDefaults() *O11yO11yQueueCheck`

NewO11yO11yQueueCheckWithDefaults instantiates a new O11yO11yQueueCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *O11yO11yQueueCheck) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *O11yO11yQueueCheck) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *O11yO11yQueueCheck) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *O11yO11yQueueCheck) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetErrorMessage

`func (o *O11yO11yQueueCheck) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *O11yO11yQueueCheck) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *O11yO11yQueueCheck) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *O11yO11yQueueCheck) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yQueueCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yQueueCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yQueueCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yQueueCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


