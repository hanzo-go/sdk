# Benchmark

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

### NewBenchmark

`func NewBenchmark() *Benchmark`

NewBenchmark instantiates a new Benchmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkWithDefaults

`func NewBenchmarkWithDefaults() *Benchmark`

NewBenchmarkWithDefaults instantiates a new Benchmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxis

`func (o *Benchmark) GetAxis() string`

GetAxis returns the Axis field if non-nil, zero value otherwise.

### GetAxisOk

`func (o *Benchmark) GetAxisOk() (*string, bool)`

GetAxisOk returns a tuple with the Axis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis

`func (o *Benchmark) SetAxis(v string)`

SetAxis sets Axis field to given value.

### HasAxis

`func (o *Benchmark) HasAxis() bool`

HasAxis returns a boolean if a field has been set.

### GetId

`func (o *Benchmark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Benchmark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Benchmark) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Benchmark) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *Benchmark) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Benchmark) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Benchmark) SetItems(v int64)`

SetItems sets Items field to given value.

### HasItems

`func (o *Benchmark) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNative

`func (o *Benchmark) GetNative() bool`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *Benchmark) GetNativeOk() (*bool, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *Benchmark) SetNative(v bool)`

SetNative sets Native field to given value.

### HasNative

`func (o *Benchmark) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetSource

`func (o *Benchmark) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Benchmark) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Benchmark) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Benchmark) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *Benchmark) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Benchmark) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Benchmark) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Benchmark) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


