# BooksExtracted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the expense bucket the SCANNER guessed, as a slug — a hint only. Vendor rules override it whenever they know better, so this is the model&#39;s reading and not the account the entry will land on. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code the document is denominated in. | [optional] 
**IssuedAt** | Pointer to **string** | IssuedAt is the document&#39;s OWN date as YYYY-MM-DD — when the bill was issued, which is not when it was uploaded or when it will post. | [optional] 
**LineItems** | Pointer to [**[]BooksLineItem**](BooksLineItem.md) | LineItems are the individual lines read off the document, where it had any. They need not sum to totalCents: a document may carry lines the scanner could not read, and the total is taken from the total. | [optional] 
**Merchant** | Pointer to **string** | Merchant is the supplier as printed on the document. | [optional] 
**Note** | Pointer to **string** | Note is anything else worth carrying from the document that has no field of its own. | [optional] 
**TaxCents** | Pointer to **int64** | TaxCents is how much of that total is tax, in cents. It is part of totalCents, not additional to it. | [optional] 
**TotalCents** | Pointer to **int64** | TotalCents is the document total in whole cents, tax INCLUDED. | [optional] 

## Methods

### NewBooksExtracted

`func NewBooksExtracted() *BooksExtracted`

NewBooksExtracted instantiates a new BooksExtracted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksExtractedWithDefaults

`func NewBooksExtractedWithDefaults() *BooksExtracted`

NewBooksExtractedWithDefaults instantiates a new BooksExtracted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *BooksExtracted) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BooksExtracted) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BooksExtracted) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BooksExtracted) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCurrency

`func (o *BooksExtracted) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BooksExtracted) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BooksExtracted) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BooksExtracted) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIssuedAt

`func (o *BooksExtracted) GetIssuedAt() string`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *BooksExtracted) GetIssuedAtOk() (*string, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *BooksExtracted) SetIssuedAt(v string)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *BooksExtracted) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetLineItems

`func (o *BooksExtracted) GetLineItems() []BooksLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BooksExtracted) GetLineItemsOk() (*[]BooksLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BooksExtracted) SetLineItems(v []BooksLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *BooksExtracted) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMerchant

`func (o *BooksExtracted) GetMerchant() string`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *BooksExtracted) GetMerchantOk() (*string, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *BooksExtracted) SetMerchant(v string)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *BooksExtracted) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetNote

`func (o *BooksExtracted) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BooksExtracted) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BooksExtracted) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BooksExtracted) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTaxCents

`func (o *BooksExtracted) GetTaxCents() int64`

GetTaxCents returns the TaxCents field if non-nil, zero value otherwise.

### GetTaxCentsOk

`func (o *BooksExtracted) GetTaxCentsOk() (*int64, bool)`

GetTaxCentsOk returns a tuple with the TaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCents

`func (o *BooksExtracted) SetTaxCents(v int64)`

SetTaxCents sets TaxCents field to given value.

### HasTaxCents

`func (o *BooksExtracted) HasTaxCents() bool`

HasTaxCents returns a boolean if a field has been set.

### GetTotalCents

`func (o *BooksExtracted) GetTotalCents() int64`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *BooksExtracted) GetTotalCentsOk() (*int64, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *BooksExtracted) SetTotalCents(v int64)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *BooksExtracted) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


