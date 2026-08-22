# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the account&#39;s human name, for a statement&#39;s line label. | [optional] 
**Number** | Pointer to **string** | Number is the posting key every voucher leg, rule and report references — stable, and the reason the chart is a fixed value rather than a table anybody can edit. It looks numeric and is a string: \&quot;1000\&quot; sorts and compares as text. | [optional] 
**Party** | Pointer to **string** | Party marks the account as carrying a SUBLEDGER — receivable is money owed to us, payable money we owe — so a leg posted here also writes a payment-ledger row against a counterparty. Absent means no subledger: a bank, wallet, revenue or cost account tracks no counterparty at all. | [optional] 
**Type** | Pointer to **string** | Type is the account&#39;s fundamental class, which is also its NORMAL balance side: asset and expense are debit-normal, liability, income and equity credit-normal. | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *Account) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Account) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Account) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Account) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetParty

`func (o *Account) GetParty() string`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *Account) GetPartyOk() (*string, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *Account) SetParty(v string)`

SetParty sets Party field to given value.

### HasParty

`func (o *Account) HasParty() bool`

HasParty returns a boolean if a field has been set.

### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Account) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


