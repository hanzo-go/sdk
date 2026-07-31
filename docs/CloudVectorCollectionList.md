# CloudVectorCollectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]CloudVectorCollection**](CloudVectorCollection.md) | Collections is one row per Qdrant collection, sorted by name. Empty — never absent — when the vector service cannot be reached. | [optional] 

## Methods

### NewCloudVectorCollectionList

`func NewCloudVectorCollectionList() *CloudVectorCollectionList`

NewCloudVectorCollectionList instantiates a new CloudVectorCollectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVectorCollectionListWithDefaults

`func NewCloudVectorCollectionListWithDefaults() *CloudVectorCollectionList`

NewCloudVectorCollectionListWithDefaults instantiates a new CloudVectorCollectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *CloudVectorCollectionList) GetCollections() []CloudVectorCollection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *CloudVectorCollectionList) GetCollectionsOk() (*[]CloudVectorCollection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *CloudVectorCollectionList) SetCollections(v []CloudVectorCollection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *CloudVectorCollectionList) HasCollections() bool`

HasCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


