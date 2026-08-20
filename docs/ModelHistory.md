# ModelHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Model is the system these runs measured. | [optional] 
**Points** | Pointer to [**[]RunPoint**](RunPoint.md) | Points is every run, oldest first. | [optional] 
**Trend** | Pointer to **float32** | Trend is the change from the first run to the last, absent when there has only been one. It answers the question a list of points makes you compute. | [optional] 

## Methods

### NewModelHistory

`func NewModelHistory() *ModelHistory`

NewModelHistory instantiates a new ModelHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelHistoryWithDefaults

`func NewModelHistoryWithDefaults() *ModelHistory`

NewModelHistoryWithDefaults instantiates a new ModelHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ModelHistory) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelHistory) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelHistory) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModelHistory) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPoints

`func (o *ModelHistory) GetPoints() []RunPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *ModelHistory) GetPointsOk() (*[]RunPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *ModelHistory) SetPoints(v []RunPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *ModelHistory) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetTrend

`func (o *ModelHistory) GetTrend() float32`

GetTrend returns the Trend field if non-nil, zero value otherwise.

### GetTrendOk

`func (o *ModelHistory) GetTrendOk() (*float32, bool)`

GetTrendOk returns a tuple with the Trend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrend

`func (o *ModelHistory) SetTrend(v float32)`

SetTrend sets Trend field to given value.

### HasTrend

`func (o *ModelHistory) HasTrend() bool`

HasTrend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


