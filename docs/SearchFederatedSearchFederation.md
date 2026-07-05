# SearchFederatedSearchFederation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** |  | [optional] [default to 0]
**Limit** | Pointer to **int32** |  | [optional] [default to 20]
**FacetsByIndex** | Pointer to **map[string][]string** |  | [optional] 
**MergeFacets** | Pointer to [**SearchFederatedSearchFederationMergeFacets**](SearchFederatedSearchFederationMergeFacets.md) |  | [optional] 

## Methods

### NewSearchFederatedSearchFederation

`func NewSearchFederatedSearchFederation() *SearchFederatedSearchFederation`

NewSearchFederatedSearchFederation instantiates a new SearchFederatedSearchFederation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFederatedSearchFederationWithDefaults

`func NewSearchFederatedSearchFederationWithDefaults() *SearchFederatedSearchFederation`

NewSearchFederatedSearchFederationWithDefaults instantiates a new SearchFederatedSearchFederation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *SearchFederatedSearchFederation) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchFederatedSearchFederation) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchFederatedSearchFederation) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchFederatedSearchFederation) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchFederatedSearchFederation) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchFederatedSearchFederation) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchFederatedSearchFederation) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchFederatedSearchFederation) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFacetsByIndex

`func (o *SearchFederatedSearchFederation) GetFacetsByIndex() map[string][]string`

GetFacetsByIndex returns the FacetsByIndex field if non-nil, zero value otherwise.

### GetFacetsByIndexOk

`func (o *SearchFederatedSearchFederation) GetFacetsByIndexOk() (*map[string][]string, bool)`

GetFacetsByIndexOk returns a tuple with the FacetsByIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetsByIndex

`func (o *SearchFederatedSearchFederation) SetFacetsByIndex(v map[string][]string)`

SetFacetsByIndex sets FacetsByIndex field to given value.

### HasFacetsByIndex

`func (o *SearchFederatedSearchFederation) HasFacetsByIndex() bool`

HasFacetsByIndex returns a boolean if a field has been set.

### GetMergeFacets

`func (o *SearchFederatedSearchFederation) GetMergeFacets() SearchFederatedSearchFederationMergeFacets`

GetMergeFacets returns the MergeFacets field if non-nil, zero value otherwise.

### GetMergeFacetsOk

`func (o *SearchFederatedSearchFederation) GetMergeFacetsOk() (*SearchFederatedSearchFederationMergeFacets, bool)`

GetMergeFacetsOk returns a tuple with the MergeFacets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeFacets

`func (o *SearchFederatedSearchFederation) SetMergeFacets(v SearchFederatedSearchFederationMergeFacets)`

SetMergeFacets sets MergeFacets field to given value.

### HasMergeFacets

`func (o *SearchFederatedSearchFederation) HasMergeFacets() bool`

HasMergeFacets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


