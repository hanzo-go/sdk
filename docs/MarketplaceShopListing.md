# MarketplaceShopListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category groups it; empty is ungrouped. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when it was published, unix seconds. | [optional] 
**Currency** | Pointer to **string** | Currency labels Price. | [optional] 
**Description** | Pointer to **string** | Description is the long copy. | [optional] 
**Docs** | Pointer to **string** | Docs is where its documentation lives, when the seller said. | [optional] 
**Id** | Pointer to **string** | ID is the listing. | [optional] 
**Kind** | Pointer to **string** | Kind is what it sells: agent, persona, app, skill, mcp or tool. | [optional] 
**Links** | Pointer to [**MarketplaceLinks**](MarketplaceLinks.md) | Links are the ways to act on it. | [optional] 
**Price** | Pointer to **interface{}** |  | [optional] 
**Public** | Pointer to **bool** | Public is always true here. | [optional] 
**PublisherOrg** | Pointer to **string** | PublisherOrg is the org that sells it. | [optional] 
**Reputation** | Pointer to [**MarketplaceReputation**](MarketplaceReputation.md) | Reputation is what buyers made of this listing. | [optional] 
**Seller** | Pointer to [**MarketplaceSeller**](MarketplaceSeller.md) | Seller is what anyone may know about the seller. | [optional] 
**Title** | Pointer to **string** | Title is the shop-window name. | [optional] 
**Tool** | Pointer to **string** | Tool names the thing sold, as the seller named it. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when it was last edited, unix seconds. | [optional] 

## Methods

### NewMarketplaceShopListing

`func NewMarketplaceShopListing() *MarketplaceShopListing`

NewMarketplaceShopListing instantiates a new MarketplaceShopListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceShopListingWithDefaults

`func NewMarketplaceShopListingWithDefaults() *MarketplaceShopListing`

NewMarketplaceShopListingWithDefaults instantiates a new MarketplaceShopListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *MarketplaceShopListing) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplaceShopListing) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplaceShopListing) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplaceShopListing) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketplaceShopListing) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketplaceShopListing) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketplaceShopListing) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketplaceShopListing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *MarketplaceShopListing) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplaceShopListing) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplaceShopListing) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplaceShopListing) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *MarketplaceShopListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceShopListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceShopListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceShopListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocs

`func (o *MarketplaceShopListing) GetDocs() string`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *MarketplaceShopListing) GetDocsOk() (*string, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *MarketplaceShopListing) SetDocs(v string)`

SetDocs sets Docs field to given value.

### HasDocs

`func (o *MarketplaceShopListing) HasDocs() bool`

HasDocs returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceShopListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceShopListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceShopListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceShopListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *MarketplaceShopListing) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketplaceShopListing) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketplaceShopListing) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MarketplaceShopListing) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLinks

`func (o *MarketplaceShopListing) GetLinks() MarketplaceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MarketplaceShopListing) GetLinksOk() (*MarketplaceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MarketplaceShopListing) SetLinks(v MarketplaceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MarketplaceShopListing) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPrice

`func (o *MarketplaceShopListing) GetPrice() interface{}`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketplaceShopListing) GetPriceOk() (*interface{}, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketplaceShopListing) SetPrice(v interface{})`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarketplaceShopListing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *MarketplaceShopListing) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *MarketplaceShopListing) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPublic

`func (o *MarketplaceShopListing) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *MarketplaceShopListing) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *MarketplaceShopListing) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *MarketplaceShopListing) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPublisherOrg

`func (o *MarketplaceShopListing) GetPublisherOrg() string`

GetPublisherOrg returns the PublisherOrg field if non-nil, zero value otherwise.

### GetPublisherOrgOk

`func (o *MarketplaceShopListing) GetPublisherOrgOk() (*string, bool)`

GetPublisherOrgOk returns a tuple with the PublisherOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherOrg

`func (o *MarketplaceShopListing) SetPublisherOrg(v string)`

SetPublisherOrg sets PublisherOrg field to given value.

### HasPublisherOrg

`func (o *MarketplaceShopListing) HasPublisherOrg() bool`

HasPublisherOrg returns a boolean if a field has been set.

### GetReputation

`func (o *MarketplaceShopListing) GetReputation() MarketplaceReputation`

GetReputation returns the Reputation field if non-nil, zero value otherwise.

### GetReputationOk

`func (o *MarketplaceShopListing) GetReputationOk() (*MarketplaceReputation, bool)`

GetReputationOk returns a tuple with the Reputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputation

`func (o *MarketplaceShopListing) SetReputation(v MarketplaceReputation)`

SetReputation sets Reputation field to given value.

### HasReputation

`func (o *MarketplaceShopListing) HasReputation() bool`

HasReputation returns a boolean if a field has been set.

### GetSeller

`func (o *MarketplaceShopListing) GetSeller() MarketplaceSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *MarketplaceShopListing) GetSellerOk() (*MarketplaceSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *MarketplaceShopListing) SetSeller(v MarketplaceSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *MarketplaceShopListing) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetTitle

`func (o *MarketplaceShopListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketplaceShopListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketplaceShopListing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketplaceShopListing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *MarketplaceShopListing) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *MarketplaceShopListing) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *MarketplaceShopListing) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *MarketplaceShopListing) HasTool() bool`

HasTool returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketplaceShopListing) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketplaceShopListing) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketplaceShopListing) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketplaceShopListing) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


