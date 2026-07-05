# SearchFederatedSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]SearchSearchQueryWithIndex**](SearchSearchQueryWithIndex.md) |  | 
**Federation** | Pointer to [**SearchFederatedSearchFederation**](SearchFederatedSearchFederation.md) |  | [optional] 

## Methods

### NewSearchFederatedSearch

`func NewSearchFederatedSearch(queries []SearchSearchQueryWithIndex, ) *SearchFederatedSearch`

NewSearchFederatedSearch instantiates a new SearchFederatedSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFederatedSearchWithDefaults

`func NewSearchFederatedSearchWithDefaults() *SearchFederatedSearch`

NewSearchFederatedSearchWithDefaults instantiates a new SearchFederatedSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *SearchFederatedSearch) GetQueries() []SearchSearchQueryWithIndex`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SearchFederatedSearch) GetQueriesOk() (*[]SearchSearchQueryWithIndex, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SearchFederatedSearch) SetQueries(v []SearchSearchQueryWithIndex)`

SetQueries sets Queries field to given value.


### GetFederation

`func (o *SearchFederatedSearch) GetFederation() SearchFederatedSearchFederation`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *SearchFederatedSearch) GetFederationOk() (*SearchFederatedSearchFederation, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *SearchFederatedSearch) SetFederation(v SearchFederatedSearchFederation)`

SetFederation sets Federation field to given value.

### HasFederation

`func (o *SearchFederatedSearch) HasFederation() bool`

HasFederation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


