# BooksLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | AmountCents is that line&#39;s amount in whole cents. The scanner is instructed to return integer cents rather than a decimal, so no float rounding can enter the ledger through here. | [optional] 
**Description** | Pointer to **string** | Description is the line as it appears on the document. | [optional] 

## Methods

### NewBooksLineItem

`func NewBooksLineItem() *BooksLineItem`

NewBooksLineItem instantiates a new BooksLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksLineItemWithDefaults

`func NewBooksLineItemWithDefaults() *BooksLineItem`

NewBooksLineItemWithDefaults instantiates a new BooksLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *BooksLineItem) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BooksLineItem) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BooksLineItem) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BooksLineItem) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetDescription

`func (o *BooksLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BooksLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BooksLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BooksLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


