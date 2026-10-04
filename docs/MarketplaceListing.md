# MarketplaceListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category groups the listing in the shop window. Free text — no vocabulary, nothing validates it — and unlike Description it is silently cut to 4096 bytes rather than refused. Empty means ungrouped. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the listing was published, in Unix SECONDS, minted at insert. Every listing read orders by it descending, so it is the shop&#39;s ordering key as well as its age. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code Price is quoted in; Create defaults it to \&quot;USD\&quot; when the publisher names none. It is a LABEL that travels to the shop window: publish parses Price with money.ParseUSD and the x402 terms carry no currency, so another code here changes what is displayed, not what is charged. | [optional] 
**Description** | Pointer to **string** | Description is the long copy. Publish REFUSES one past 4096 bytes rather than truncating it, so what is stored is what was sent; empty is allowed. | [optional] 
**Docs** | Pointer to **string** | Docs is where the listing&#39;s documentation lives: an https URL, or a path on docs.hanzo.ai. Empty when the seller named none. | [optional] 
**Id** | Pointer to **string** | ID is the listing&#39;s id, minted here as \&quot;lst_\&quot; + 16 hex characters. A publisher cannot choose it: Create overwrites whatever arrives. It is unique within PublisherOrg (the primary key is the pair), and it is the path segment DELETE /v1/marketplace/listings/:id takes. | [optional] 
**Kind** | Pointer to **string** | Kind is what the listing sells: agent, persona, app, skill, mcp or tool. Listings published before kinds existed sold tools, and read as tool. | [optional] 
**Price** | Pointer to **interface{}** |  | [optional] 
**Public** | Pointer to **bool** | Public is whether other orgs can discover the listing. It also decides ENFORCEMENT: only public rows reach the price table, so a private listing with a price charges nobody. False leaves the row visible to its publisher alone. | [optional] 
**PublisherOrg** | Pointer to **string** | PublisherOrg is the org that published the listing, taken from the validated principal and never off the wire. It is also the PAYEE org — Recipient is resolved inside it — and the isolation key: a publisher reads and deletes only rows carrying its own org. | [optional] 
**Recipient** | Pointer to **string** | Recipient is the seller&#39;s payout WALLET id in PublisherOrg: the wallet x402 pays for a call or a job bought through this listing. | [optional] 
**Ref** | Pointer to **string** | Ref is the id the thing&#39;s owning app knows it by, resolved when the listing was published — the agent id for an agent named by its name. A tool&#39;s is the x402 resource a dispatch of it is priced by: the source it resolves from and its name, so a price follows the tool that runs and not a name another org&#39;s own row may share. | [optional] 
**Title** | Pointer to **string** | Title is the shop-window name, required and refused past 200 bytes. It is what discovery paints over the tool&#39;s registry name. | [optional] 
**Tool** | Pointer to **string** | Tool names the thing sold, as the seller named it: the agent, app, skill or MCP server&#39;s name or id, or the tool&#39;s registry name. It was proved the seller&#39;s own when the listing was published. For a tool it is what install switches on. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the listing was last edited, unix seconds; its creation when it never was. | [optional] 

## Methods

### NewMarketplaceListing

`func NewMarketplaceListing() *MarketplaceListing`

NewMarketplaceListing instantiates a new MarketplaceListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceListingWithDefaults

`func NewMarketplaceListingWithDefaults() *MarketplaceListing`

NewMarketplaceListingWithDefaults instantiates a new MarketplaceListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *MarketplaceListing) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplaceListing) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplaceListing) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplaceListing) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketplaceListing) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketplaceListing) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketplaceListing) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketplaceListing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *MarketplaceListing) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplaceListing) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplaceListing) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplaceListing) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *MarketplaceListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocs

`func (o *MarketplaceListing) GetDocs() string`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *MarketplaceListing) GetDocsOk() (*string, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *MarketplaceListing) SetDocs(v string)`

SetDocs sets Docs field to given value.

### HasDocs

`func (o *MarketplaceListing) HasDocs() bool`

HasDocs returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *MarketplaceListing) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketplaceListing) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketplaceListing) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MarketplaceListing) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPrice

`func (o *MarketplaceListing) GetPrice() interface{}`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketplaceListing) GetPriceOk() (*interface{}, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketplaceListing) SetPrice(v interface{})`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarketplaceListing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *MarketplaceListing) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *MarketplaceListing) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPublic

`func (o *MarketplaceListing) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *MarketplaceListing) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *MarketplaceListing) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *MarketplaceListing) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPublisherOrg

`func (o *MarketplaceListing) GetPublisherOrg() string`

GetPublisherOrg returns the PublisherOrg field if non-nil, zero value otherwise.

### GetPublisherOrgOk

`func (o *MarketplaceListing) GetPublisherOrgOk() (*string, bool)`

GetPublisherOrgOk returns a tuple with the PublisherOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherOrg

`func (o *MarketplaceListing) SetPublisherOrg(v string)`

SetPublisherOrg sets PublisherOrg field to given value.

### HasPublisherOrg

`func (o *MarketplaceListing) HasPublisherOrg() bool`

HasPublisherOrg returns a boolean if a field has been set.

### GetRecipient

`func (o *MarketplaceListing) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *MarketplaceListing) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *MarketplaceListing) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *MarketplaceListing) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetRef

`func (o *MarketplaceListing) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MarketplaceListing) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MarketplaceListing) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MarketplaceListing) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetTitle

`func (o *MarketplaceListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketplaceListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketplaceListing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketplaceListing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *MarketplaceListing) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *MarketplaceListing) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *MarketplaceListing) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *MarketplaceListing) HasTool() bool`

HasTool returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketplaceListing) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketplaceListing) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketplaceListing) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketplaceListing) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


