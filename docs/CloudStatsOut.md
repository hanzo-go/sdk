# CloudStatsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Admin is the upstream service&#39;s server-panel flag, always false here. | [optional] 
**Metrics** | Pointer to **map[string]interface{}** | Metrics is the upstream transactor&#39;s metrics block. This server does not populate it, so it is always the empty object — the front reads the key, not its contents. | [optional] 
**Statistics** | Pointer to [**CloudStatsSessions**](CloudStatsSessions.md) | Statistics carries the live sessions. | [optional] 

## Methods

### NewCloudStatsOut

`func NewCloudStatsOut() *CloudStatsOut`

NewCloudStatsOut instantiates a new CloudStatsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStatsOutWithDefaults

`func NewCloudStatsOutWithDefaults() *CloudStatsOut`

NewCloudStatsOutWithDefaults instantiates a new CloudStatsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *CloudStatsOut) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *CloudStatsOut) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *CloudStatsOut) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *CloudStatsOut) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetMetrics

`func (o *CloudStatsOut) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CloudStatsOut) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CloudStatsOut) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CloudStatsOut) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetStatistics

`func (o *CloudStatsOut) GetStatistics() CloudStatsSessions`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *CloudStatsOut) GetStatisticsOk() (*CloudStatsSessions, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *CloudStatsOut) SetStatistics(v CloudStatsSessions)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *CloudStatsOut) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


