# TrainMetricsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TrainSeries**](TrainSeries.md) | Data is one series per metric name. | [optional] 

## Methods

### NewTrainMetricsOut

`func NewTrainMetricsOut() *TrainMetricsOut`

NewTrainMetricsOut instantiates a new TrainMetricsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainMetricsOutWithDefaults

`func NewTrainMetricsOutWithDefaults() *TrainMetricsOut`

NewTrainMetricsOutWithDefaults instantiates a new TrainMetricsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TrainMetricsOut) GetData() []TrainSeries`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrainMetricsOut) GetDataOk() (*[]TrainSeries, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrainMetricsOut) SetData(v []TrainSeries)`

SetData sets Data field to given value.

### HasData

`func (o *TrainMetricsOut) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


