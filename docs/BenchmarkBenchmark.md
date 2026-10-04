# BenchmarkBenchmark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axis** | Pointer to **string** | what capability it measures | [optional] 
**Id** | Pointer to **string** | the id every other op on this surface takes | [optional] 
**Items** | Pointer to **int64** | how many items it holds, when the set is fixed | [optional] 
**Native** | Pointer to **bool** | whether the standardized harness runs it today | [optional] 
**Source** | Pointer to **string** | where the items come from | [optional] 
**Title** | Pointer to **string** | the benchmark&#39;s published name | [optional] 

## Methods

### NewBenchmarkBenchmark

`func NewBenchmarkBenchmark() *BenchmarkBenchmark`

NewBenchmarkBenchmark instantiates a new BenchmarkBenchmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkBenchmarkWithDefaults

`func NewBenchmarkBenchmarkWithDefaults() *BenchmarkBenchmark`

NewBenchmarkBenchmarkWithDefaults instantiates a new BenchmarkBenchmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxis

`func (o *BenchmarkBenchmark) GetAxis() string`

GetAxis returns the Axis field if non-nil, zero value otherwise.

### GetAxisOk

`func (o *BenchmarkBenchmark) GetAxisOk() (*string, bool)`

GetAxisOk returns a tuple with the Axis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis

`func (o *BenchmarkBenchmark) SetAxis(v string)`

SetAxis sets Axis field to given value.

### HasAxis

`func (o *BenchmarkBenchmark) HasAxis() bool`

HasAxis returns a boolean if a field has been set.

### GetId

`func (o *BenchmarkBenchmark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BenchmarkBenchmark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BenchmarkBenchmark) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BenchmarkBenchmark) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *BenchmarkBenchmark) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BenchmarkBenchmark) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BenchmarkBenchmark) SetItems(v int64)`

SetItems sets Items field to given value.

### HasItems

`func (o *BenchmarkBenchmark) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNative

`func (o *BenchmarkBenchmark) GetNative() bool`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *BenchmarkBenchmark) GetNativeOk() (*bool, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *BenchmarkBenchmark) SetNative(v bool)`

SetNative sets Native field to given value.

### HasNative

`func (o *BenchmarkBenchmark) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetSource

`func (o *BenchmarkBenchmark) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BenchmarkBenchmark) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BenchmarkBenchmark) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BenchmarkBenchmark) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *BenchmarkBenchmark) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BenchmarkBenchmark) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BenchmarkBenchmark) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BenchmarkBenchmark) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


