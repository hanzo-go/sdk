# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount is what actually moved, as an exact 18-decimal-place USD string. It is NOT the atomic-unit figure the client signed: the challenge quotes the asset&#39;s own units (USDC&#39;s 6 dp) and truncates to fit them, while the ledger moves this exact value. | [optional] 
**From** | Pointer to **string** | From is the payer&#39;s EVM address: the account that signed the EIP-3009 authorization, recovered from the signature rather than taken on trust. | [optional] 
**Id** | Pointer to **string** | ID is the settle-once key: \&quot;x402_\&quot; + keccak(from|nonce) in hex. It is DERIVED, not minted, so a client that re-submits the same authorization addresses the same settlement and is served again for free rather than charged twice. It is also the id GET /v1/x402/settlements/:id takes. | [optional] 
**Network** | Pointer to **string** | Network is the CAIP-2 identifier the payment was settled under, e.g. \&quot;eip155:36963\&quot;. Its eip155 reference is the chain id in the EIP-712 domain the payer signed, so it is not a label — changing it invalidates the signature. | [optional] 
**Nonce** | Pointer to **string** | Nonce is the client-chosen nonce from the authorization, hex — up to 32 bytes, left-padded to the contract&#39;s bytes32. It is the replay anchor: the token contract refuses a second on-chain transfer for one (from, nonce), and this rail refuses a second settlement for the same pair, so a ledger settlement inherits the identical guarantee. | [optional] 
**Payee** | Pointer to **string** | Payee is the recipient&#39;s EVM address — the &#x60;payTo&#x60; the challenge advertised and the authorization named. A payment to any other address never settles. | [optional] 
**PayeeOrg** | Pointer to **string** | PayeeOrg is the tenant that owns the recipient wallet, resolved at settlement. It is who got PAID, as Payer is who paid. | [optional] 
**Payer** | Pointer to **string** | Payer is the payer ORG — the tenant whose ledger was debited — and not an address. It is the org the request was authenticated as, so it answers who is billed, which the payer address alone cannot. | [optional] 
**Resource** | Pointer to **string** | Resource is what was paid for, in the same spelling the price table and the challenge used: the request path for a priced route, \&quot;tool:&lt;id&gt;\&quot; for a priced tool. | [optional] 
**SettledAt** | Pointer to **int32** | SettledAt is when this settlement was CLAIMED, in unix seconds — the moment the authorization was accepted, which is also the moment the time window it carried stopped applying. A settlement finished later by reconciliation keeps this instant. | [optional] 
**SettledVia** | Pointer to **string** | SettledVia is which rail moved the money: \&quot;ledger\&quot;, the live default, or \&quot;chain\&quot; when the authorization is broadcast. Those two values and no others. | [optional] 
**TxHash** | Pointer to **string** | TxHash is the chain transaction hash, present only for a \&quot;chain\&quot; settlement. Empty on a ledger settlement — that is the normal case today, and it means the money moved without a chain, not that it failed. The wire&#39;s PAYMENT-RESPONSE &#x60;transaction&#x60; falls back to ID when this is empty. | [optional] 

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


