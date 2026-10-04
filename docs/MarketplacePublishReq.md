# MarketplacePublishReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category groups the listing in the shop window. | [optional] 
**Currency** | Pointer to **string** | Currency denominates Price. | [optional] 
**Description** | Pointer to **string** | Description is the long copy, at most 4096 characters. | [optional] 
**Docs** | Pointer to **string** | Docs is where the listing&#39;s documentation lives: an https URL or a path on docs.hanzo.ai. | [optional] 
**Kind** | Pointer to **string** | Kind is what is sold: agent, persona, app, skill, mcp or tool (the default). A tool is a platform tool, sold per call by the platform alone. | [optional] 
**Price** | Pointer to **string** | Price is the price as a decimal USD string, exact to 18 places — per call for a tool, what a job is offered at otherwise. \&quot;0.0025\&quot; is a quarter of a cent and stays one. Empty or \&quot;0\&quot; publishes it free; a positive price requires Recipient. | [optional] 
**Public** | Pointer to **bool** | Public makes the listing discoverable by other orgs. Private otherwise. | [optional] 
**Recipient** | Pointer to **string** | Recipient is the seller&#39;s payout wallet id, in the publishing org — the wallet x402 pays. Required for a monetized listing. | [optional] 
**Title** | Pointer to **string** | Title is the shop-window name, 1-200 characters. Required. | [optional] 
**Tool** | Pointer to **string** | Tool names the thing sold, in the seller&#39;s own org: an agent&#39;s name or id, an app&#39;s project slug or id, a skill&#39;s name or id, an MCP server&#39;s name or id registered by hand, or a platform tool&#39;s registry name. It must be the seller&#39;s — there are no phantom listings. | [optional] 

## Methods

### NewMarketplacePublishReq

`func NewMarketplacePublishReq() *MarketplacePublishReq`

NewMarketplacePublishReq instantiates a new MarketplacePublishReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacePublishReqWithDefaults

`func NewMarketplacePublishReqWithDefaults() *MarketplacePublishReq`

NewMarketplacePublishReqWithDefaults instantiates a new MarketplacePublishReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *MarketplacePublishReq) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplacePublishReq) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplacePublishReq) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplacePublishReq) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCurrency

`func (o *MarketplacePublishReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplacePublishReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplacePublishReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplacePublishReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *MarketplacePublishReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplacePublishReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplacePublishReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplacePublishReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocs

`func (o *MarketplacePublishReq) GetDocs() string`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *MarketplacePublishReq) GetDocsOk() (*string, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *MarketplacePublishReq) SetDocs(v string)`

SetDocs sets Docs field to given value.

### HasDocs

`func (o *MarketplacePublishReq) HasDocs() bool`

HasDocs returns a boolean if a field has been set.

### GetKind

`func (o *MarketplacePublishReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketplacePublishReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketplacePublishReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MarketplacePublishReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPrice

`func (o *MarketplacePublishReq) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketplacePublishReq) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketplacePublishReq) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarketplacePublishReq) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPublic

`func (o *MarketplacePublishReq) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *MarketplacePublishReq) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *MarketplacePublishReq) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *MarketplacePublishReq) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRecipient

`func (o *MarketplacePublishReq) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *MarketplacePublishReq) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *MarketplacePublishReq) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *MarketplacePublishReq) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetTitle

`func (o *MarketplacePublishReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketplacePublishReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketplacePublishReq) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketplacePublishReq) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *MarketplacePublishReq) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *MarketplacePublishReq) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *MarketplacePublishReq) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *MarketplacePublishReq) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


