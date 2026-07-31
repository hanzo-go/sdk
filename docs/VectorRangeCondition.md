# VectorRangeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Range** | Pointer to [**VectorRange**](VectorRange.md) |  | [optional] 

## Methods

### NewVectorRangeCondition

`func NewVectorRangeCondition() *VectorRangeCondition`

NewVectorRangeCondition instantiates a new VectorRangeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorRangeConditionWithDefaults

`func NewVectorRangeConditionWithDefaults() *VectorRangeCondition`

NewVectorRangeConditionWithDefaults instantiates a new VectorRangeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VectorRangeCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VectorRangeCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VectorRangeCondition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VectorRangeCondition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRange

`func (o *VectorRangeCondition) GetRange() VectorRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *VectorRangeCondition) GetRangeOk() (*VectorRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *VectorRangeCondition) SetRange(v VectorRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *VectorRangeCondition) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


