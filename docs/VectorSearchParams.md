# VectorSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HnswEf** | Pointer to **int32** |  | [optional] 
**Exact** | Pointer to **bool** |  | [optional] [default to false]
**Quantization** | Pointer to [**VectorSearchParamsQuantization**](VectorSearchParamsQuantization.md) |  | [optional] 

## Methods

### NewVectorSearchParams

`func NewVectorSearchParams() *VectorSearchParams`

NewVectorSearchParams instantiates a new VectorSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchParamsWithDefaults

`func NewVectorSearchParamsWithDefaults() *VectorSearchParams`

NewVectorSearchParamsWithDefaults instantiates a new VectorSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHnswEf

`func (o *VectorSearchParams) GetHnswEf() int32`

GetHnswEf returns the HnswEf field if non-nil, zero value otherwise.

### GetHnswEfOk

`func (o *VectorSearchParams) GetHnswEfOk() (*int32, bool)`

GetHnswEfOk returns a tuple with the HnswEf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHnswEf

`func (o *VectorSearchParams) SetHnswEf(v int32)`

SetHnswEf sets HnswEf field to given value.

### HasHnswEf

`func (o *VectorSearchParams) HasHnswEf() bool`

HasHnswEf returns a boolean if a field has been set.

### GetExact

`func (o *VectorSearchParams) GetExact() bool`

GetExact returns the Exact field if non-nil, zero value otherwise.

### GetExactOk

`func (o *VectorSearchParams) GetExactOk() (*bool, bool)`

GetExactOk returns a tuple with the Exact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExact

`func (o *VectorSearchParams) SetExact(v bool)`

SetExact sets Exact field to given value.

### HasExact

`func (o *VectorSearchParams) HasExact() bool`

HasExact returns a boolean if a field has been set.

### GetQuantization

`func (o *VectorSearchParams) GetQuantization() VectorSearchParamsQuantization`

GetQuantization returns the Quantization field if non-nil, zero value otherwise.

### GetQuantizationOk

`func (o *VectorSearchParams) GetQuantizationOk() (*VectorSearchParamsQuantization, bool)`

GetQuantizationOk returns a tuple with the Quantization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantization

`func (o *VectorSearchParams) SetQuantization(v VectorSearchParamsQuantization)`

SetQuantization sets Quantization field to given value.

### HasQuantization

`func (o *VectorSearchParams) HasQuantization() bool`

HasQuantization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


