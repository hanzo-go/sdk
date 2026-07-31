# CloudPublishReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category groups the listing in the shop window. | [optional] 
**Currency** | Pointer to **string** | Currency denominates Price. | [optional] 
**Description** | Pointer to **string** | Description is the long copy, clipped at 4096 characters. | [optional] 
**Price** | Pointer to **string** | Price is the per-call price as a decimal USD string, exact to 18 places — \&quot;0.0025\&quot; is a quarter of a cent and stays one. Empty or \&quot;0\&quot; (the default) publishes it free; any positive price makes the listing monetized and requires Recipient. | [optional] 
**Public** | Pointer to **bool** | Public makes the listing discoverable by other orgs. Private otherwise. | [optional] 
**Recipient** | Pointer to **string** | Recipient is the seller&#39;s payout wallet ID, in the publishing org — the wallet x402 pays. Required for a monetized listing. | [optional] 
**Title** | Pointer to **string** | Title is the shop-window name, 1-200 characters. Required. | [optional] 
**Tool** | Pointer to **string** | Tool is the registry name of the capability being offered. It must already resolve in the publisher&#39;s own scope — there are no phantom listings. | [optional] 

## Methods

### NewCloudPublishReq

`func NewCloudPublishReq() *CloudPublishReq`

NewCloudPublishReq instantiates a new CloudPublishReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPublishReqWithDefaults

`func NewCloudPublishReqWithDefaults() *CloudPublishReq`

NewCloudPublishReqWithDefaults instantiates a new CloudPublishReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CloudPublishReq) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudPublishReq) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudPublishReq) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudPublishReq) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudPublishReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudPublishReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudPublishReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudPublishReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *CloudPublishReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudPublishReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudPublishReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudPublishReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *CloudPublishReq) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CloudPublishReq) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CloudPublishReq) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CloudPublishReq) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPublic

`func (o *CloudPublishReq) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CloudPublishReq) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CloudPublishReq) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CloudPublishReq) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRecipient

`func (o *CloudPublishReq) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *CloudPublishReq) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *CloudPublishReq) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *CloudPublishReq) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetTitle

`func (o *CloudPublishReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudPublishReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudPublishReq) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudPublishReq) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *CloudPublishReq) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *CloudPublishReq) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *CloudPublishReq) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *CloudPublishReq) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


