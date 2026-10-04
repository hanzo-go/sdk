# BooksAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the account&#39;s human name, for a statement&#39;s line label. | [optional] 
**Number** | Pointer to **string** | Number is the posting key every voucher leg, rule and report references — stable, and the reason the chart is a fixed value rather than a table anybody can edit. It looks numeric and is a string: \&quot;1000\&quot; sorts and compares as text. | [optional] 
**Party** | Pointer to **string** | Party marks the account as carrying a SUBLEDGER — receivable is money owed to us, payable money we owe — so a leg posted here also writes a payment-ledger row against a counterparty. Absent means no subledger: a bank, wallet, revenue or cost account tracks no counterparty at all. | [optional] 
**Type** | Pointer to **string** | Type is the account&#39;s fundamental class, which is also its NORMAL balance side: asset and expense are debit-normal, liability, income and equity credit-normal. | [optional] 

## Methods

### NewBooksAccount

`func NewBooksAccount() *BooksAccount`

NewBooksAccount instantiates a new BooksAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksAccountWithDefaults

`func NewBooksAccountWithDefaults() *BooksAccount`

NewBooksAccountWithDefaults instantiates a new BooksAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BooksAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BooksAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BooksAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BooksAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *BooksAccount) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BooksAccount) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BooksAccount) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *BooksAccount) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetParty

`func (o *BooksAccount) GetParty() string`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *BooksAccount) GetPartyOk() (*string, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *BooksAccount) SetParty(v string)`

SetParty sets Party field to given value.

### HasParty

`func (o *BooksAccount) HasParty() bool`

HasParty returns a boolean if a field has been set.

### GetType

`func (o *BooksAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BooksAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BooksAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BooksAccount) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


