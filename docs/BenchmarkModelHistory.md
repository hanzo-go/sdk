# BenchmarkModelHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Model is the system these runs measured. | [optional] 
**Points** | Pointer to [**[]BenchmarkRunPoint**](BenchmarkRunPoint.md) | Points is every run, oldest first. | [optional] 
**Trend** | Pointer to **float64** | Trend is the change from the first run to the last, absent when there has only been one. It answers the question a list of points makes you compute. | [optional] 

## Methods

### NewBenchmarkModelHistory

`func NewBenchmarkModelHistory() *BenchmarkModelHistory`

NewBenchmarkModelHistory instantiates a new BenchmarkModelHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkModelHistoryWithDefaults

`func NewBenchmarkModelHistoryWithDefaults() *BenchmarkModelHistory`

NewBenchmarkModelHistoryWithDefaults instantiates a new BenchmarkModelHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *BenchmarkModelHistory) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BenchmarkModelHistory) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BenchmarkModelHistory) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BenchmarkModelHistory) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPoints

`func (o *BenchmarkModelHistory) GetPoints() []BenchmarkRunPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *BenchmarkModelHistory) GetPointsOk() (*[]BenchmarkRunPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *BenchmarkModelHistory) SetPoints(v []BenchmarkRunPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *BenchmarkModelHistory) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetTrend

`func (o *BenchmarkModelHistory) GetTrend() float64`

GetTrend returns the Trend field if non-nil, zero value otherwise.

### GetTrendOk

`func (o *BenchmarkModelHistory) GetTrendOk() (*float64, bool)`

GetTrendOk returns a tuple with the Trend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrend

`func (o *BenchmarkModelHistory) SetTrend(v float64)`

SetTrend sets Trend field to given value.

### HasTrend

`func (o *BenchmarkModelHistory) HasTrend() bool`

HasTrend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


