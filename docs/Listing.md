# Listing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **interface{}** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**PublisherOrg** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to **string** | seller payout WALLET ID, in PublisherOrg. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Tool** | Pointer to **string** |  | [optional] 

## Methods

### NewListing

`func NewListing() *Listing`

NewListing instantiates a new Listing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingWithDefaults

`func NewListingWithDefaults() *Listing`

NewListingWithDefaults instantiates a new Listing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Listing) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Listing) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Listing) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Listing) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Listing) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Listing) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Listing) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Listing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Listing) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Listing) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Listing) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Listing) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *Listing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Listing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Listing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Listing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Listing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Listing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Listing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Listing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *Listing) GetPrice() interface{}`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Listing) GetPriceOk() (*interface{}, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Listing) SetPrice(v interface{})`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Listing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *Listing) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *Listing) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPublic

`func (o *Listing) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Listing) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Listing) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *Listing) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPublisherOrg

`func (o *Listing) GetPublisherOrg() string`

GetPublisherOrg returns the PublisherOrg field if non-nil, zero value otherwise.

### GetPublisherOrgOk

`func (o *Listing) GetPublisherOrgOk() (*string, bool)`

GetPublisherOrgOk returns a tuple with the PublisherOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherOrg

`func (o *Listing) SetPublisherOrg(v string)`

SetPublisherOrg sets PublisherOrg field to given value.

### HasPublisherOrg

`func (o *Listing) HasPublisherOrg() bool`

HasPublisherOrg returns a boolean if a field has been set.

### GetRecipient

`func (o *Listing) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Listing) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Listing) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *Listing) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetTitle

`func (o *Listing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Listing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Listing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Listing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *Listing) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *Listing) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *Listing) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *Listing) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


