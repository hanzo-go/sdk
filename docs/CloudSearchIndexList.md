# CloudSearchIndexList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexes** | Pointer to [**[]CloudSearchIndex**](CloudSearchIndex.md) | Indexes is one row per Meilisearch index, sorted by name. Empty — never absent — when the search service cannot be reached. | [optional] 

## Methods

### NewCloudSearchIndexList

`func NewCloudSearchIndexList() *CloudSearchIndexList`

NewCloudSearchIndexList instantiates a new CloudSearchIndexList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSearchIndexListWithDefaults

`func NewCloudSearchIndexListWithDefaults() *CloudSearchIndexList`

NewCloudSearchIndexListWithDefaults instantiates a new CloudSearchIndexList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexes

`func (o *CloudSearchIndexList) GetIndexes() []CloudSearchIndex`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *CloudSearchIndexList) GetIndexesOk() (*[]CloudSearchIndex, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *CloudSearchIndexList) SetIndexes(v []CloudSearchIndex)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *CloudSearchIndexList) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


