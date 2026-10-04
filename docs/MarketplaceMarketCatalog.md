# MarketplaceMarketCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]MarketplaceMarketItem**](MarketplaceMarketItem.md) | Items is every capability the caller can see in their own (org, project), each carrying any public listing&#39;s shop metadata and whether it is installed. | [optional] 

## Methods

### NewMarketplaceMarketCatalog

`func NewMarketplaceMarketCatalog() *MarketplaceMarketCatalog`

NewMarketplaceMarketCatalog instantiates a new MarketplaceMarketCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceMarketCatalogWithDefaults

`func NewMarketplaceMarketCatalogWithDefaults() *MarketplaceMarketCatalog`

NewMarketplaceMarketCatalogWithDefaults instantiates a new MarketplaceMarketCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MarketplaceMarketCatalog) GetItems() []MarketplaceMarketItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MarketplaceMarketCatalog) GetItemsOk() (*[]MarketplaceMarketItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MarketplaceMarketCatalog) SetItems(v []MarketplaceMarketItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *MarketplaceMarketCatalog) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


