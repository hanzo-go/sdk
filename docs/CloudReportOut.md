# CloudReportOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivered** | Pointer to **bool** | Delivered is true when a waiting durable owner received this result. False means there was none to deliver to — an unknown or already-finished run — which is a clean no-op, not an error. | [optional] 

## Methods

### NewCloudReportOut

`func NewCloudReportOut() *CloudReportOut`

NewCloudReportOut instantiates a new CloudReportOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudReportOutWithDefaults

`func NewCloudReportOutWithDefaults() *CloudReportOut`

NewCloudReportOutWithDefaults instantiates a new CloudReportOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivered

`func (o *CloudReportOut) GetDelivered() bool`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *CloudReportOut) GetDeliveredOk() (*bool, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *CloudReportOut) SetDelivered(v bool)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *CloudReportOut) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


