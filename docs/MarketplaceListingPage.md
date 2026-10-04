# MarketplaceListingPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listings** | Pointer to [**[]MarketplaceListing**](MarketplaceListing.md) | Listings is every listing this org has published, private ones included (Public says which are discoverable by others). | [optional] 

## Methods

### NewMarketplaceListingPage

`func NewMarketplaceListingPage() *MarketplaceListingPage`

NewMarketplaceListingPage instantiates a new MarketplaceListingPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceListingPageWithDefaults

`func NewMarketplaceListingPageWithDefaults() *MarketplaceListingPage`

NewMarketplaceListingPageWithDefaults instantiates a new MarketplaceListingPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListings

`func (o *MarketplaceListingPage) GetListings() []MarketplaceListing`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *MarketplaceListingPage) GetListingsOk() (*[]MarketplaceListing, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *MarketplaceListingPage) SetListings(v []MarketplaceListing)`

SetListings sets Listings field to given value.

### HasListings

`func (o *MarketplaceListingPage) HasListings() bool`

HasListings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


