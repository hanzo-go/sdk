# TaxLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | AmountCents is its fair market value in U.S. dollars on the date paid, in cents — for a U.S. dollar payment, what was paid. | [optional] 
**Id** | Pointer to **string** | ID is the economic event&#39;s id. | [optional] 
**Memo** | Pointer to **string** | Memo is what was bought, in the rail&#39;s words. | [optional] 
**Paid** | Pointer to **int64** | Paid is when, unix seconds. | [optional] 
**Proof** | Pointer to **string** | Proof is the rail&#39;s own reference for the payment. | [optional] 
**Rail** | Pointer to **string** | Rail is how it moved: x402, ledger, chain, ach, wire, card, network or other. | [optional] 
**Verdict** | Pointer to [**TaxVerdict**](TaxVerdict.md) | Verdict is whether it is reportable, where, and the rules that decided it. | [optional] 

## Methods

### NewTaxLine

`func NewTaxLine() *TaxLine`

NewTaxLine instantiates a new TaxLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxLineWithDefaults

`func NewTaxLineWithDefaults() *TaxLine`

NewTaxLineWithDefaults instantiates a new TaxLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *TaxLine) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *TaxLine) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *TaxLine) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *TaxLine) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetId

`func (o *TaxLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxLine) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxLine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemo

`func (o *TaxLine) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TaxLine) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TaxLine) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TaxLine) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPaid

`func (o *TaxLine) GetPaid() int64`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *TaxLine) GetPaidOk() (*int64, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *TaxLine) SetPaid(v int64)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *TaxLine) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetProof

`func (o *TaxLine) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *TaxLine) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *TaxLine) SetProof(v string)`

SetProof sets Proof field to given value.

### HasProof

`func (o *TaxLine) HasProof() bool`

HasProof returns a boolean if a field has been set.

### GetRail

`func (o *TaxLine) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *TaxLine) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *TaxLine) SetRail(v string)`

SetRail sets Rail field to given value.

### HasRail

`func (o *TaxLine) HasRail() bool`

HasRail returns a boolean if a field has been set.

### GetVerdict

`func (o *TaxLine) GetVerdict() TaxVerdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *TaxLine) GetVerdictOk() (*TaxVerdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *TaxLine) SetVerdict(v TaxVerdict)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *TaxLine) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


