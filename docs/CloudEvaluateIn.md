# CloudEvaluateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | Pointer to **string** | DistinctID is the identity the flags are evaluated for. Required. | [optional] 
**Groups** | Pointer to **interface{}** |  | [optional] 
**PersonProperties** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCloudEvaluateIn

`func NewCloudEvaluateIn() *CloudEvaluateIn`

NewCloudEvaluateIn instantiates a new CloudEvaluateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEvaluateInWithDefaults

`func NewCloudEvaluateInWithDefaults() *CloudEvaluateIn`

NewCloudEvaluateInWithDefaults instantiates a new CloudEvaluateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *CloudEvaluateIn) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *CloudEvaluateIn) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *CloudEvaluateIn) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *CloudEvaluateIn) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetGroups

`func (o *CloudEvaluateIn) GetGroups() interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CloudEvaluateIn) GetGroupsOk() (*interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CloudEvaluateIn) SetGroups(v interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CloudEvaluateIn) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *CloudEvaluateIn) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *CloudEvaluateIn) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetPersonProperties

`func (o *CloudEvaluateIn) GetPersonProperties() interface{}`

GetPersonProperties returns the PersonProperties field if non-nil, zero value otherwise.

### GetPersonPropertiesOk

`func (o *CloudEvaluateIn) GetPersonPropertiesOk() (*interface{}, bool)`

GetPersonPropertiesOk returns a tuple with the PersonProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonProperties

`func (o *CloudEvaluateIn) SetPersonProperties(v interface{})`

SetPersonProperties sets PersonProperties field to given value.

### HasPersonProperties

`func (o *CloudEvaluateIn) HasPersonProperties() bool`

HasPersonProperties returns a boolean if a field has been set.

### SetPersonPropertiesNil

`func (o *CloudEvaluateIn) SetPersonPropertiesNil(b bool)`

 SetPersonPropertiesNil sets the value for PersonProperties to be an explicit nil

### UnsetPersonProperties
`func (o *CloudEvaluateIn) UnsetPersonProperties()`

UnsetPersonProperties ensures that no value is present for PersonProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


