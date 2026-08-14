# VectorCollectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]VectorCollection**](VectorCollection.md) | Collections is one row per Qdrant collection, sorted by name. Empty — never absent — when the vector service cannot be reached. | [optional] 

## Methods

### NewVectorCollectionList

`func NewVectorCollectionList() *VectorCollectionList`

NewVectorCollectionList instantiates a new VectorCollectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorCollectionListWithDefaults

`func NewVectorCollectionListWithDefaults() *VectorCollectionList`

NewVectorCollectionListWithDefaults instantiates a new VectorCollectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *VectorCollectionList) GetCollections() []VectorCollection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *VectorCollectionList) GetCollectionsOk() (*[]VectorCollection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *VectorCollectionList) SetCollections(v []VectorCollection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *VectorCollectionList) HasCollections() bool`

HasCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


