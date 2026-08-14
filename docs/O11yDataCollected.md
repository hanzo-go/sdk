# O11yDataCollected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]O11yCollectedLogAttribute**](O11yCollectedLogAttribute.md) |  | [optional] 
**Metrics** | Pointer to [**[]O11yCollectedMetric**](O11yCollectedMetric.md) |  | [optional] 

## Methods

### NewO11yDataCollected

`func NewO11yDataCollected() *O11yDataCollected`

NewO11yDataCollected instantiates a new O11yDataCollected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDataCollectedWithDefaults

`func NewO11yDataCollectedWithDefaults() *O11yDataCollected`

NewO11yDataCollectedWithDefaults instantiates a new O11yDataCollected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yDataCollected) GetLogs() []O11yCollectedLogAttribute`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yDataCollected) GetLogsOk() (*[]O11yCollectedLogAttribute, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yDataCollected) SetLogs(v []O11yCollectedLogAttribute)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yDataCollected) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yDataCollected) GetMetrics() []O11yCollectedMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yDataCollected) GetMetricsOk() (*[]O11yCollectedMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yDataCollected) SetMetrics(v []O11yCollectedMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yDataCollected) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


