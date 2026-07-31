# VectorMatchCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Match** | Pointer to [**VectorMatch**](VectorMatch.md) |  | [optional] 

## Methods

### NewVectorMatchCondition

`func NewVectorMatchCondition() *VectorMatchCondition`

NewVectorMatchCondition instantiates a new VectorMatchCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorMatchConditionWithDefaults

`func NewVectorMatchConditionWithDefaults() *VectorMatchCondition`

NewVectorMatchConditionWithDefaults instantiates a new VectorMatchCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VectorMatchCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VectorMatchCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VectorMatchCondition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VectorMatchCondition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMatch

`func (o *VectorMatchCondition) GetMatch() VectorMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *VectorMatchCondition) GetMatchOk() (*VectorMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *VectorMatchCondition) SetMatch(v VectorMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *VectorMatchCondition) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


