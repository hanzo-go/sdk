# WalletWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** | Address is the on-chain address. For kms/mpc/treasury it is the EOA a signature from this wallet recovers to; for safe it is the CREATE2 address of the Safe CONTRACT, which holds no key — its approvals recover to the MPC owner instead. Rotating a kms wallet mints a new key and therefore a NEW address, and funds and approvals at the old one do not follow; mpc, treasury and safe addresses are invariant under rotation. | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to **string** | Chain is the EVM chain the wallet is bound to, CAIP-2 \&quot;eip155:&lt;n&gt;\&quot; or a bare decimal chain id. Empty is chain-agnostic: the ring signs an unbound digest, and a Safe falls back to the Hanzo L1 (36963) because a Safe and its EIP-712 domain must be chain-bound. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the wallet was provisioned, Unix seconds. Listings order by it, newest first. | [optional] 
**Custody** | Pointer to **string** | Custody is the backend holding the signing material, fixed at creation: \&quot;kms\&quot; (a secp256k1 key sealed under KMS and opened in-process), \&quot;mpc\&quot; or \&quot;treasury\&quot; (an m-of-n threshold key on the deployed ring, which differ by governance and not by signing mechanics), or \&quot;safe\&quot; (a Safe contract owned by an MPC key). A kind the deployment has not wired refuses with 503 rather than fabricating a signature. | [optional] 
**FinanceAccount** | Pointer to **string** | FinanceAccount is the finance ledger account bound to this wallet — the lookup that turns a ledger account back into an on-chain signer. Absent is the normal state and means unbound; the column is NULL until something binds it. | [optional] 
**Id** | Pointer to **string** | ID is the wallet id, minted by the server as \&quot;wal_\&quot; + 24 hex. It is the last segment of the key ref, and it is the LEDGER SUBJECT an x402 payment into this wallet credits — so it names money as well as key material. | [optional] 
**Name** | Pointer to **string** | Name is the display label given at creation. It addresses nothing: the key ref is derived from the scope and the id, so renaming moves no material. | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **string** | Tier is the wallet tier the ring keys its TierPolicy on: hot, warm, cold, gas, bridge, contract_admin, validator, quarantine or disaster_recovery. It defaults to hot and is refused at the boundary if it is none of the nine. | [optional] 

## Methods

### NewWalletWallet

`func NewWalletWallet() *WalletWallet`

NewWalletWallet instantiates a new WalletWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWalletWithDefaults

`func NewWalletWalletWithDefaults() *WalletWallet`

NewWalletWalletWithDefaults instantiates a new WalletWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *WalletWallet) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WalletWallet) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WalletWallet) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *WalletWallet) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAddress

`func (o *WalletWallet) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WalletWallet) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WalletWallet) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WalletWallet) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAgent

`func (o *WalletWallet) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *WalletWallet) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *WalletWallet) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *WalletWallet) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetChain

`func (o *WalletWallet) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *WalletWallet) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *WalletWallet) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *WalletWallet) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WalletWallet) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletWallet) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletWallet) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WalletWallet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustody

`func (o *WalletWallet) GetCustody() string`

GetCustody returns the Custody field if non-nil, zero value otherwise.

### GetCustodyOk

`func (o *WalletWallet) GetCustodyOk() (*string, bool)`

GetCustodyOk returns a tuple with the Custody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustody

`func (o *WalletWallet) SetCustody(v string)`

SetCustody sets Custody field to given value.

### HasCustody

`func (o *WalletWallet) HasCustody() bool`

HasCustody returns a boolean if a field has been set.

### GetFinanceAccount

`func (o *WalletWallet) GetFinanceAccount() string`

GetFinanceAccount returns the FinanceAccount field if non-nil, zero value otherwise.

### GetFinanceAccountOk

`func (o *WalletWallet) GetFinanceAccountOk() (*string, bool)`

GetFinanceAccountOk returns a tuple with the FinanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinanceAccount

`func (o *WalletWallet) SetFinanceAccount(v string)`

SetFinanceAccount sets FinanceAccount field to given value.

### HasFinanceAccount

`func (o *WalletWallet) HasFinanceAccount() bool`

HasFinanceAccount returns a boolean if a field has been set.

### GetId

`func (o *WalletWallet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletWallet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletWallet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WalletWallet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WalletWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletWallet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WalletWallet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *WalletWallet) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *WalletWallet) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *WalletWallet) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *WalletWallet) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *WalletWallet) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *WalletWallet) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *WalletWallet) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *WalletWallet) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTier

`func (o *WalletWallet) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *WalletWallet) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *WalletWallet) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *WalletWallet) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


