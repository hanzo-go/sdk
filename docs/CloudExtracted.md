# CloudExtracted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | proposed slug (software|cloud|office|…) | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**IssuedAt** | Pointer to **string** | YYYY-MM-DD | [optional] 
**LineItems** | Pointer to [**[]CloudLineItem**](CloudLineItem.md) |  | [optional] 
**Merchant** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**TaxCents** | Pointer to **int32** |  | [optional] 
**TotalCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudExtracted

`func NewCloudExtracted() *CloudExtracted`

NewCloudExtracted instantiates a new CloudExtracted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudExtractedWithDefaults

`func NewCloudExtractedWithDefaults() *CloudExtracted`

NewCloudExtractedWithDefaults instantiates a new CloudExtracted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CloudExtracted) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudExtracted) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudExtracted) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudExtracted) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudExtracted) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudExtracted) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudExtracted) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudExtracted) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIssuedAt

`func (o *CloudExtracted) GetIssuedAt() string`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *CloudExtracted) GetIssuedAtOk() (*string, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *CloudExtracted) SetIssuedAt(v string)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *CloudExtracted) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetLineItems

`func (o *CloudExtracted) GetLineItems() []CloudLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CloudExtracted) GetLineItemsOk() (*[]CloudLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CloudExtracted) SetLineItems(v []CloudLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CloudExtracted) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMerchant

`func (o *CloudExtracted) GetMerchant() string`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *CloudExtracted) GetMerchantOk() (*string, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *CloudExtracted) SetMerchant(v string)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *CloudExtracted) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetNote

`func (o *CloudExtracted) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CloudExtracted) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CloudExtracted) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CloudExtracted) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTaxCents

`func (o *CloudExtracted) GetTaxCents() int32`

GetTaxCents returns the TaxCents field if non-nil, zero value otherwise.

### GetTaxCentsOk

`func (o *CloudExtracted) GetTaxCentsOk() (*int32, bool)`

GetTaxCentsOk returns a tuple with the TaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCents

`func (o *CloudExtracted) SetTaxCents(v int32)`

SetTaxCents sets TaxCents field to given value.

### HasTaxCents

`func (o *CloudExtracted) HasTaxCents() bool`

HasTaxCents returns a boolean if a field has been set.

### GetTotalCents

`func (o *CloudExtracted) GetTotalCents() int32`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *CloudExtracted) GetTotalCentsOk() (*int32, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *CloudExtracted) SetTotalCents(v int32)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *CloudExtracted) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


