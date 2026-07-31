# AgentsMetricsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**[]AgentsSeriesLine**](AgentsSeriesLine.md) |  | [optional] 
**Resource** | Pointer to [**AgentsResourceUsage**](AgentsResourceUsage.md) |  | [optional] 

## Methods

### NewAgentsMetricsView

`func NewAgentsMetricsView() *AgentsMetricsView`

NewAgentsMetricsView instantiates a new AgentsMetricsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsMetricsViewWithDefaults

`func NewAgentsMetricsViewWithDefaults() *AgentsMetricsView`

NewAgentsMetricsViewWithDefaults instantiates a new AgentsMetricsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *AgentsMetricsView) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *AgentsMetricsView) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *AgentsMetricsView) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *AgentsMetricsView) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *AgentsMetricsView) GetSeries() []AgentsSeriesLine`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *AgentsMetricsView) GetSeriesOk() (*[]AgentsSeriesLine, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *AgentsMetricsView) SetSeries(v []AgentsSeriesLine)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *AgentsMetricsView) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetResource

`func (o *AgentsMetricsView) GetResource() AgentsResourceUsage`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AgentsMetricsView) GetResourceOk() (*AgentsResourceUsage, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AgentsMetricsView) SetResource(v AgentsResourceUsage)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AgentsMetricsView) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


