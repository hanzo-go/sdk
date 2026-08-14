# GpuAlertList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to **[]map[string]interface{}** | Alerts is always empty, and typed as a raw list because Visor exposes no alert inventory for this surface to shape: there is nothing to describe until there is something to return. | [optional] 

## Methods

### NewGpuAlertList

`func NewGpuAlertList() *GpuAlertList`

NewGpuAlertList instantiates a new GpuAlertList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuAlertListWithDefaults

`func NewGpuAlertListWithDefaults() *GpuAlertList`

NewGpuAlertListWithDefaults instantiates a new GpuAlertList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *GpuAlertList) GetAlerts() []map[string]interface{}`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *GpuAlertList) GetAlertsOk() (*[]map[string]interface{}, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *GpuAlertList) SetAlerts(v []map[string]interface{})`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *GpuAlertList) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


