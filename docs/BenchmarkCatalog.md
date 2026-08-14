# BenchmarkCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Benchmark**](Benchmark.md) | Data is one row per benchmark, in the catalog&#39;s own order. | [optional] 
**Total** | Pointer to **int32** | Total is how many rows Data holds. | [optional] 

## Methods

### NewBenchmarkCatalog

`func NewBenchmarkCatalog() *BenchmarkCatalog`

NewBenchmarkCatalog instantiates a new BenchmarkCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkCatalogWithDefaults

`func NewBenchmarkCatalogWithDefaults() *BenchmarkCatalog`

NewBenchmarkCatalogWithDefaults instantiates a new BenchmarkCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BenchmarkCatalog) GetData() []Benchmark`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BenchmarkCatalog) GetDataOk() (*[]Benchmark, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BenchmarkCatalog) SetData(v []Benchmark)`

SetData sets Data field to given value.

### HasData

`func (o *BenchmarkCatalog) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *BenchmarkCatalog) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BenchmarkCatalog) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BenchmarkCatalog) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BenchmarkCatalog) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


