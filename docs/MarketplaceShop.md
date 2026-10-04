# MarketplaceShop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facets** | Pointer to [**MarketplaceFacets**](MarketplaceFacets.md) | Facets counts the matches by kind, category, price and rating, each counted as if its own filter were not applied. | [optional] 
**Listings** | Pointer to [**[]MarketplaceShopListing**](MarketplaceShopListing.md) | Listings is this page, newest first. | [optional] 
**Total** | Pointer to **int64** | Total is how many listings the search matched, past this page. | [optional] 

## Methods

### NewMarketplaceShop

`func NewMarketplaceShop() *MarketplaceShop`

NewMarketplaceShop instantiates a new MarketplaceShop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceShopWithDefaults

`func NewMarketplaceShopWithDefaults() *MarketplaceShop`

NewMarketplaceShopWithDefaults instantiates a new MarketplaceShop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacets

`func (o *MarketplaceShop) GetFacets() MarketplaceFacets`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *MarketplaceShop) GetFacetsOk() (*MarketplaceFacets, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *MarketplaceShop) SetFacets(v MarketplaceFacets)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *MarketplaceShop) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### GetListings

`func (o *MarketplaceShop) GetListings() []MarketplaceShopListing`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *MarketplaceShop) GetListingsOk() (*[]MarketplaceShopListing, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *MarketplaceShop) SetListings(v []MarketplaceShopListing)`

SetListings sets Listings field to given value.

### HasListings

`func (o *MarketplaceShop) HasListings() bool`

HasListings returns a boolean if a field has been set.

### GetTotal

`func (o *MarketplaceShop) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MarketplaceShop) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MarketplaceShop) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MarketplaceShop) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


