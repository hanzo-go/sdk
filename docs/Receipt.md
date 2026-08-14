# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | exact 18-dp USD (money.Amount string) | [optional] 
**From** | Pointer to **string** | payer address | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**Payee** | Pointer to **string** | recipient address | [optional] 
**PayeeOrg** | Pointer to **string** |  | [optional] 
**Payer** | Pointer to **string** | payer ORG (the debited ledger) | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**SettledAt** | Pointer to **int32** |  | [optional] 
**SettledVia** | Pointer to **string** | \&quot;ledger\&quot; (live) | \&quot;chain\&quot; (seam) | [optional] 
**TxHash** | Pointer to **string** |  | [optional] 

## Methods

### NewReceipt

`func NewReceipt() *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Receipt) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Receipt) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Receipt) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Receipt) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFrom

`func (o *Receipt) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Receipt) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Receipt) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Receipt) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *Receipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Receipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Receipt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Receipt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetwork

`func (o *Receipt) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Receipt) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Receipt) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Receipt) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNonce

`func (o *Receipt) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Receipt) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Receipt) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Receipt) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPayee

`func (o *Receipt) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *Receipt) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *Receipt) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *Receipt) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### GetPayeeOrg

`func (o *Receipt) GetPayeeOrg() string`

GetPayeeOrg returns the PayeeOrg field if non-nil, zero value otherwise.

### GetPayeeOrgOk

`func (o *Receipt) GetPayeeOrgOk() (*string, bool)`

GetPayeeOrgOk returns a tuple with the PayeeOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeOrg

`func (o *Receipt) SetPayeeOrg(v string)`

SetPayeeOrg sets PayeeOrg field to given value.

### HasPayeeOrg

`func (o *Receipt) HasPayeeOrg() bool`

HasPayeeOrg returns a boolean if a field has been set.

### GetPayer

`func (o *Receipt) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *Receipt) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *Receipt) SetPayer(v string)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *Receipt) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetResource

`func (o *Receipt) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Receipt) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Receipt) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Receipt) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSettledAt

`func (o *Receipt) GetSettledAt() int32`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *Receipt) GetSettledAtOk() (*int32, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *Receipt) SetSettledAt(v int32)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *Receipt) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### GetSettledVia

`func (o *Receipt) GetSettledVia() string`

GetSettledVia returns the SettledVia field if non-nil, zero value otherwise.

### GetSettledViaOk

`func (o *Receipt) GetSettledViaOk() (*string, bool)`

GetSettledViaOk returns a tuple with the SettledVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledVia

`func (o *Receipt) SetSettledVia(v string)`

SetSettledVia sets SettledVia field to given value.

### HasSettledVia

`func (o *Receipt) HasSettledVia() bool`

HasSettledVia returns a boolean if a field has been set.

### GetTxHash

`func (o *Receipt) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *Receipt) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *Receipt) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *Receipt) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


