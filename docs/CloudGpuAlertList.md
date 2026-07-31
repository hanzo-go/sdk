# CloudGpuAlertList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to **[]map[string]interface{}** | Alerts is always empty, and typed as a raw list because Visor exposes no alert inventory for this surface to shape: there is nothing to describe until there is something to return. | [optional] 

## Methods

### NewCloudGpuAlertList

`func NewCloudGpuAlertList() *CloudGpuAlertList`

NewCloudGpuAlertList instantiates a new CloudGpuAlertList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGpuAlertListWithDefaults

`func NewCloudGpuAlertListWithDefaults() *CloudGpuAlertList`

NewCloudGpuAlertListWithDefaults instantiates a new CloudGpuAlertList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *CloudGpuAlertList) GetAlerts() []map[string]interface{}`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *CloudGpuAlertList) GetAlertsOk() (*[]map[string]interface{}, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *CloudGpuAlertList) SetAlerts(v []map[string]interface{})`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *CloudGpuAlertList) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


