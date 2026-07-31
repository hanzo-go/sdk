# CloudCatalogPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudEntry**](CloudEntry.md) | Data is the page of matching entries, most recently updated first. | [optional] 
**Facets** | Pointer to **map[string]map[string]int32** | Facets counts the whole matching set along every browse axis, so a rail a client renders is a rail that has results behind it. Keyed axis → value → count. | [optional] 
**Total** | Pointer to **int32** | Total is how many entries matched BEFORE paging — what a pager sizes itself on. | [optional] 

## Methods

### NewCloudCatalogPage

`func NewCloudCatalogPage() *CloudCatalogPage`

NewCloudCatalogPage instantiates a new CloudCatalogPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCatalogPageWithDefaults

`func NewCloudCatalogPageWithDefaults() *CloudCatalogPage`

NewCloudCatalogPageWithDefaults instantiates a new CloudCatalogPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudCatalogPage) GetData() []CloudEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudCatalogPage) GetDataOk() (*[]CloudEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudCatalogPage) SetData(v []CloudEntry)`

SetData sets Data field to given value.

### HasData

`func (o *CloudCatalogPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFacets

`func (o *CloudCatalogPage) GetFacets() map[string]map[string]int32`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *CloudCatalogPage) GetFacetsOk() (*map[string]map[string]int32, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *CloudCatalogPage) SetFacets(v map[string]map[string]int32)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *CloudCatalogPage) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### GetTotal

`func (o *CloudCatalogPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudCatalogPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudCatalogPage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudCatalogPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


