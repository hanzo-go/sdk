# O11yIntegrationAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to **[]interface{}** |  | [optional] 
**Dashboards** | Pointer to **[]map[string]map[string]interface{}** |  | [optional] 
**Logs** | Pointer to [**O11yLogsAssets**](O11yLogsAssets.md) |  | [optional] 

## Methods

### NewO11yIntegrationAssets

`func NewO11yIntegrationAssets() *O11yIntegrationAssets`

NewO11yIntegrationAssets instantiates a new O11yIntegrationAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yIntegrationAssetsWithDefaults

`func NewO11yIntegrationAssetsWithDefaults() *O11yIntegrationAssets`

NewO11yIntegrationAssetsWithDefaults instantiates a new O11yIntegrationAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *O11yIntegrationAssets) GetAlerts() []interface{}`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *O11yIntegrationAssets) GetAlertsOk() (*[]interface{}, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *O11yIntegrationAssets) SetAlerts(v []interface{})`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *O11yIntegrationAssets) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDashboards

`func (o *O11yIntegrationAssets) GetDashboards() []map[string]map[string]interface{}`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *O11yIntegrationAssets) GetDashboardsOk() (*[]map[string]map[string]interface{}, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *O11yIntegrationAssets) SetDashboards(v []map[string]map[string]interface{})`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *O11yIntegrationAssets) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### GetLogs

`func (o *O11yIntegrationAssets) GetLogs() O11yLogsAssets`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yIntegrationAssets) GetLogsOk() (*O11yLogsAssets, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yIntegrationAssets) SetLogs(v O11yLogsAssets)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yIntegrationAssets) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


