# O11yO11yMetricAlerts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]O11yO11yMetricAlert**](O11yO11yMetricAlert.md) | Alerts are the alert rules referencing the metric. | [optional] 

## Methods

### NewO11yO11yMetricAlerts

`func NewO11yO11yMetricAlerts() *O11yO11yMetricAlerts`

NewO11yO11yMetricAlerts instantiates a new O11yO11yMetricAlerts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricAlertsWithDefaults

`func NewO11yO11yMetricAlertsWithDefaults() *O11yO11yMetricAlerts`

NewO11yO11yMetricAlertsWithDefaults instantiates a new O11yO11yMetricAlerts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *O11yO11yMetricAlerts) GetAlerts() []O11yO11yMetricAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *O11yO11yMetricAlerts) GetAlertsOk() (*[]O11yO11yMetricAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *O11yO11yMetricAlerts) SetAlerts(v []O11yO11yMetricAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *O11yO11yMetricAlerts) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


