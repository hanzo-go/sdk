# BenchmarkBenchmarkCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BenchmarkBenchmark**](BenchmarkBenchmark.md) | Data is one row per benchmark, in the catalog&#39;s own order. | [optional] 
**Total** | Pointer to **int64** | Total is how many rows Data holds. | [optional] 

## Methods

### NewBenchmarkBenchmarkCatalog

`func NewBenchmarkBenchmarkCatalog() *BenchmarkBenchmarkCatalog`

NewBenchmarkBenchmarkCatalog instantiates a new BenchmarkBenchmarkCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkBenchmarkCatalogWithDefaults

`func NewBenchmarkBenchmarkCatalogWithDefaults() *BenchmarkBenchmarkCatalog`

NewBenchmarkBenchmarkCatalogWithDefaults instantiates a new BenchmarkBenchmarkCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BenchmarkBenchmarkCatalog) GetData() []BenchmarkBenchmark`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BenchmarkBenchmarkCatalog) GetDataOk() (*[]BenchmarkBenchmark, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BenchmarkBenchmarkCatalog) SetData(v []BenchmarkBenchmark)`

SetData sets Data field to given value.

### HasData

`func (o *BenchmarkBenchmarkCatalog) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *BenchmarkBenchmarkCatalog) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BenchmarkBenchmarkCatalog) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BenchmarkBenchmarkCatalog) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BenchmarkBenchmarkCatalog) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


