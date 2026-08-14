# O11yO11yTestNotificationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11yTestNotificationResult**](O11yO11yTestNotificationResult.md) | Data holds the fired-series count and a status message. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yTestNotificationOut

`func NewO11yO11yTestNotificationOut() *O11yO11yTestNotificationOut`

NewO11yO11yTestNotificationOut instantiates a new O11yO11yTestNotificationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTestNotificationOutWithDefaults

`func NewO11yO11yTestNotificationOutWithDefaults() *O11yO11yTestNotificationOut`

NewO11yO11yTestNotificationOutWithDefaults instantiates a new O11yO11yTestNotificationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yTestNotificationOut) GetData() O11yO11yTestNotificationResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yTestNotificationOut) GetDataOk() (*O11yO11yTestNotificationResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yTestNotificationOut) SetData(v O11yO11yTestNotificationResult)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yTestNotificationOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yTestNotificationOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yTestNotificationOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yTestNotificationOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yTestNotificationOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


