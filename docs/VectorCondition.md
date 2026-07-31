# VectorCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Match** | Pointer to [**VectorMatch**](VectorMatch.md) |  | [optional] 
**Range** | Pointer to [**VectorRange**](VectorRange.md) |  | [optional] 

## Methods

### NewVectorCondition

`func NewVectorCondition() *VectorCondition`

NewVectorCondition instantiates a new VectorCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorConditionWithDefaults

`func NewVectorConditionWithDefaults() *VectorCondition`

NewVectorConditionWithDefaults instantiates a new VectorCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VectorCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VectorCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VectorCondition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VectorCondition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMatch

`func (o *VectorCondition) GetMatch() VectorMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *VectorCondition) GetMatchOk() (*VectorMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *VectorCondition) SetMatch(v VectorMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *VectorCondition) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetRange

`func (o *VectorCondition) GetRange() VectorRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *VectorCondition) GetRangeOk() (*VectorRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *VectorCondition) SetRange(v VectorRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *VectorCondition) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


