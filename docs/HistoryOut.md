# HistoryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **string** | Benchmark is the catalog id these histories are about. | [optional] 
**Data** | Pointer to [**[]ModelHistory**](ModelHistory.md) | Data is one entry per model, ordered by model name. | [optional] 
**Total** | Pointer to **int32** | Total is how many models Data holds. | [optional] 

## Methods

### NewHistoryOut

`func NewHistoryOut() *HistoryOut`

NewHistoryOut instantiates a new HistoryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryOutWithDefaults

`func NewHistoryOutWithDefaults() *HistoryOut`

NewHistoryOutWithDefaults instantiates a new HistoryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *HistoryOut) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *HistoryOut) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *HistoryOut) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *HistoryOut) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetData

`func (o *HistoryOut) GetData() []ModelHistory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryOut) GetDataOk() (*[]ModelHistory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoryOut) SetData(v []ModelHistory)`

SetData sets Data field to given value.

### HasData

`func (o *HistoryOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *HistoryOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HistoryOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HistoryOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HistoryOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


