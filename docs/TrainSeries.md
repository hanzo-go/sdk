# TrainSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the metric. | [optional] 
**Points** | Pointer to **[][]float64** | Points are [step, value, at] in step order. | [optional] 

## Methods

### NewTrainSeries

`func NewTrainSeries() *TrainSeries`

NewTrainSeries instantiates a new TrainSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainSeriesWithDefaults

`func NewTrainSeriesWithDefaults() *TrainSeries`

NewTrainSeriesWithDefaults instantiates a new TrainSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrainSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainSeries) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainSeries) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoints

`func (o *TrainSeries) GetPoints() [][]float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *TrainSeries) GetPointsOk() (*[][]float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *TrainSeries) SetPoints(v [][]float64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *TrainSeries) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


