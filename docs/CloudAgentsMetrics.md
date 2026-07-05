# CloudAgentsMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**[]CloudAgentsMetricsSeriesInner**](CloudAgentsMetricsSeriesInner.md) |  | [optional] 
**Resource** | Pointer to [**CloudAgentsMetricsResource**](CloudAgentsMetricsResource.md) |  | [optional] 

## Methods

### NewCloudAgentsMetrics

`func NewCloudAgentsMetrics() *CloudAgentsMetrics`

NewCloudAgentsMetrics instantiates a new CloudAgentsMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsMetricsWithDefaults

`func NewCloudAgentsMetricsWithDefaults() *CloudAgentsMetrics`

NewCloudAgentsMetricsWithDefaults instantiates a new CloudAgentsMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *CloudAgentsMetrics) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudAgentsMetrics) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudAgentsMetrics) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudAgentsMetrics) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *CloudAgentsMetrics) GetSeries() []CloudAgentsMetricsSeriesInner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CloudAgentsMetrics) GetSeriesOk() (*[]CloudAgentsMetricsSeriesInner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CloudAgentsMetrics) SetSeries(v []CloudAgentsMetricsSeriesInner)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *CloudAgentsMetrics) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetResource

`func (o *CloudAgentsMetrics) GetResource() CloudAgentsMetricsResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CloudAgentsMetrics) GetResourceOk() (*CloudAgentsMetricsResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CloudAgentsMetrics) SetResource(v CloudAgentsMetricsResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CloudAgentsMetrics) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


