# VectorFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Must** | Pointer to [**[]VectorCondition**](VectorCondition.md) |  | [optional] 
**Should** | Pointer to [**[]VectorCondition**](VectorCondition.md) |  | [optional] 
**MustNot** | Pointer to [**[]VectorCondition**](VectorCondition.md) |  | [optional] 

## Methods

### NewVectorFilter

`func NewVectorFilter() *VectorFilter`

NewVectorFilter instantiates a new VectorFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorFilterWithDefaults

`func NewVectorFilterWithDefaults() *VectorFilter`

NewVectorFilterWithDefaults instantiates a new VectorFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMust

`func (o *VectorFilter) GetMust() []VectorCondition`

GetMust returns the Must field if non-nil, zero value otherwise.

### GetMustOk

`func (o *VectorFilter) GetMustOk() (*[]VectorCondition, bool)`

GetMustOk returns a tuple with the Must field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMust

`func (o *VectorFilter) SetMust(v []VectorCondition)`

SetMust sets Must field to given value.

### HasMust

`func (o *VectorFilter) HasMust() bool`

HasMust returns a boolean if a field has been set.

### GetShould

`func (o *VectorFilter) GetShould() []VectorCondition`

GetShould returns the Should field if non-nil, zero value otherwise.

### GetShouldOk

`func (o *VectorFilter) GetShouldOk() (*[]VectorCondition, bool)`

GetShouldOk returns a tuple with the Should field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShould

`func (o *VectorFilter) SetShould(v []VectorCondition)`

SetShould sets Should field to given value.

### HasShould

`func (o *VectorFilter) HasShould() bool`

HasShould returns a boolean if a field has been set.

### GetMustNot

`func (o *VectorFilter) GetMustNot() []VectorCondition`

GetMustNot returns the MustNot field if non-nil, zero value otherwise.

### GetMustNotOk

`func (o *VectorFilter) GetMustNotOk() (*[]VectorCondition, bool)`

GetMustNotOk returns a tuple with the MustNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustNot

`func (o *VectorFilter) SetMustNot(v []VectorCondition)`

SetMustNot sets MustNot field to given value.

### HasMustNot

`func (o *VectorFilter) HasMustNot() bool`

HasMustNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


