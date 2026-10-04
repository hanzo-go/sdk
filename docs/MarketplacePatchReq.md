# MarketplacePatchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category groups the listing in the shop window. | [optional] 
**Currency** | Pointer to **string** | Currency denominates Price. | [optional] 
**Description** | Pointer to **string** | Description is the long copy, at most 4096 characters. | [optional] 
**Docs** | Pointer to **string** | Docs is where the listing&#39;s documentation lives; empty clears it. | [optional] 
**Id** | Pointer to **string** | ID is the listing to edit, from the path. | [optional] 
**Price** | Pointer to **string** | Price is the decimal USD price, exact to 18 places; \&quot;0\&quot; makes it free. | [optional] 
**Public** | Pointer to **bool** | Public makes the listing discoverable by other orgs, or withdraws it. | [optional] 
**Recipient** | Pointer to **string** | Recipient is the payout wallet id, in the publishing org. | [optional] 
**Title** | Pointer to **string** | Title is the shop-window name, 1-200 characters. | [optional] 

## Methods

### NewMarketplacePatchReq

`func NewMarketplacePatchReq() *MarketplacePatchReq`

NewMarketplacePatchReq instantiates a new MarketplacePatchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacePatchReqWithDefaults

`func NewMarketplacePatchReqWithDefaults() *MarketplacePatchReq`

NewMarketplacePatchReqWithDefaults instantiates a new MarketplacePatchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *MarketplacePatchReq) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplacePatchReq) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplacePatchReq) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplacePatchReq) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCurrency

`func (o *MarketplacePatchReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplacePatchReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplacePatchReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplacePatchReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *MarketplacePatchReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplacePatchReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplacePatchReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplacePatchReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocs

`func (o *MarketplacePatchReq) GetDocs() string`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *MarketplacePatchReq) GetDocsOk() (*string, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *MarketplacePatchReq) SetDocs(v string)`

SetDocs sets Docs field to given value.

### HasDocs

`func (o *MarketplacePatchReq) HasDocs() bool`

HasDocs returns a boolean if a field has been set.

### GetId

`func (o *MarketplacePatchReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplacePatchReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplacePatchReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplacePatchReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *MarketplacePatchReq) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketplacePatchReq) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketplacePatchReq) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarketplacePatchReq) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPublic

`func (o *MarketplacePatchReq) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *MarketplacePatchReq) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *MarketplacePatchReq) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *MarketplacePatchReq) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRecipient

`func (o *MarketplacePatchReq) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *MarketplacePatchReq) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *MarketplacePatchReq) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *MarketplacePatchReq) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetTitle

`func (o *MarketplacePatchReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketplacePatchReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketplacePatchReq) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketplacePatchReq) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


