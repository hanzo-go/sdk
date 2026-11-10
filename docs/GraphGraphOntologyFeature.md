# GraphGraphOntologyFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Group is the feature block: identity (0..19), communication (20..59) or photos (60..143). | [optional] 
**Index** | Pointer to **int64** | Index is the feature&#39;s 0-based position (0..143) in metadata.json. | [optional] 
**Name** | Pointer to **string** | Name is Apple&#39;s feature identifier, such as sameLastName or photosPersonPosteriorProbabilityGivenSaturdayNight. | [optional] 

## Methods

### NewGraphGraphOntologyFeature

`func NewGraphGraphOntologyFeature() *GraphGraphOntologyFeature`

NewGraphGraphOntologyFeature instantiates a new GraphGraphOntologyFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphOntologyFeatureWithDefaults

`func NewGraphGraphOntologyFeatureWithDefaults() *GraphGraphOntologyFeature`

NewGraphGraphOntologyFeatureWithDefaults instantiates a new GraphGraphOntologyFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GraphGraphOntologyFeature) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GraphGraphOntologyFeature) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GraphGraphOntologyFeature) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GraphGraphOntologyFeature) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIndex

`func (o *GraphGraphOntologyFeature) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GraphGraphOntologyFeature) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GraphGraphOntologyFeature) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *GraphGraphOntologyFeature) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *GraphGraphOntologyFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GraphGraphOntologyFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GraphGraphOntologyFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GraphGraphOntologyFeature) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


