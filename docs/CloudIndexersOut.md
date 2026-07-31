# CloudIndexersOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexers** | Pointer to [**[]CloudIndexerView**](CloudIndexerView.md) | Indexers is one row per reachable chain indexer, or an empty list when the indexer is unreachable — never a fabricated row. | [optional] 

## Methods

### NewCloudIndexersOut

`func NewCloudIndexersOut() *CloudIndexersOut`

NewCloudIndexersOut instantiates a new CloudIndexersOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIndexersOutWithDefaults

`func NewCloudIndexersOutWithDefaults() *CloudIndexersOut`

NewCloudIndexersOutWithDefaults instantiates a new CloudIndexersOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexers

`func (o *CloudIndexersOut) GetIndexers() []CloudIndexerView`

GetIndexers returns the Indexers field if non-nil, zero value otherwise.

### GetIndexersOk

`func (o *CloudIndexersOut) GetIndexersOk() (*[]CloudIndexerView, bool)`

GetIndexersOk returns a tuple with the Indexers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexers

`func (o *CloudIndexersOut) SetIndexers(v []CloudIndexerView)`

SetIndexers sets Indexers field to given value.

### HasIndexers

`func (o *CloudIndexersOut) HasIndexers() bool`

HasIndexers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


