# EvaluateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | Pointer to **string** | DistinctID is the identity the flags are evaluated for. Required. | [optional] 
**Groups** | Pointer to **interface{}** |  | [optional] 
**PersonProperties** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewEvaluateIn

`func NewEvaluateIn() *EvaluateIn`

NewEvaluateIn instantiates a new EvaluateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluateInWithDefaults

`func NewEvaluateInWithDefaults() *EvaluateIn`

NewEvaluateInWithDefaults instantiates a new EvaluateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *EvaluateIn) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *EvaluateIn) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *EvaluateIn) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *EvaluateIn) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetGroups

`func (o *EvaluateIn) GetGroups() interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EvaluateIn) GetGroupsOk() (*interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EvaluateIn) SetGroups(v interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EvaluateIn) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *EvaluateIn) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *EvaluateIn) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetPersonProperties

`func (o *EvaluateIn) GetPersonProperties() interface{}`

GetPersonProperties returns the PersonProperties field if non-nil, zero value otherwise.

### GetPersonPropertiesOk

`func (o *EvaluateIn) GetPersonPropertiesOk() (*interface{}, bool)`

GetPersonPropertiesOk returns a tuple with the PersonProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonProperties

`func (o *EvaluateIn) SetPersonProperties(v interface{})`

SetPersonProperties sets PersonProperties field to given value.

### HasPersonProperties

`func (o *EvaluateIn) HasPersonProperties() bool`

HasPersonProperties returns a boolean if a field has been set.

### SetPersonPropertiesNil

`func (o *EvaluateIn) SetPersonPropertiesNil(b bool)`

 SetPersonPropertiesNil sets the value for PersonProperties to be an explicit nil

### UnsetPersonProperties
`func (o *EvaluateIn) UnsetPersonProperties()`

UnsetPersonProperties ensures that no value is present for PersonProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


