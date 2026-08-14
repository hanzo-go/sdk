# O11yO11yTestNotificationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertCount** | Pointer to **int32** | AlertCount is how many series would alert for the tested rule. | [optional] 
**Message** | Pointer to **string** | Message is a human-readable status, e.g. \&quot;notification sent\&quot;. | [optional] 

## Methods

### NewO11yO11yTestNotificationResult

`func NewO11yO11yTestNotificationResult() *O11yO11yTestNotificationResult`

NewO11yO11yTestNotificationResult instantiates a new O11yO11yTestNotificationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTestNotificationResultWithDefaults

`func NewO11yO11yTestNotificationResultWithDefaults() *O11yO11yTestNotificationResult`

NewO11yO11yTestNotificationResultWithDefaults instantiates a new O11yO11yTestNotificationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertCount

`func (o *O11yO11yTestNotificationResult) GetAlertCount() int32`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *O11yO11yTestNotificationResult) GetAlertCountOk() (*int32, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *O11yO11yTestNotificationResult) SetAlertCount(v int32)`

SetAlertCount sets AlertCount field to given value.

### HasAlertCount

`func (o *O11yO11yTestNotificationResult) HasAlertCount() bool`

HasAlertCount returns a boolean if a field has been set.

### GetMessage

`func (o *O11yO11yTestNotificationResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yO11yTestNotificationResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yO11yTestNotificationResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yO11yTestNotificationResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


