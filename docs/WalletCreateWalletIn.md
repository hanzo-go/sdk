# WalletCreateWalletIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AccountID is the account this wallet belongs to. Required, and it must be an account of the caller&#39;s own org — an unknown one is a 404. | [optional] 
**Agent** | Pointer to **string** | Agent optionally narrows the wallet to one agent within the org. It becomes a segment of the key ref, so it must be a url-safe segment with no slash. | [optional] 
**Chain** | Pointer to **string** | Chain is the EVM chain this wallet is for, as \&quot;eip155:&lt;n&gt;\&quot; or a bare decimal chain id. Optional; a Safe defaults to the Hanzo L1 (36963). | [optional] 
**Custody** | Pointer to **string** | Custody selects the signing backend: \&quot;kms\&quot; (in-process, always available), \&quot;mpc\&quot; or \&quot;treasury\&quot; (the deployed MPC ring), or \&quot;safe\&quot; (a Safe smart wallet owned by an MPC key). Empty uses the deployment&#39;s default. A backend that is not configured fails CLOSED with 503 rather than fabricating a signature. | [optional] 
**Name** | Pointer to **string** | Name is the wallet&#39;s display label. Optional. | [optional] 
**Tier** | Pointer to **string** | Tier is the MPC wallet tier: hot, warm, cold, gas, bridge, contract_admin, validator, quarantine or disaster_recovery. Empty defaults to hot. | [optional] 

## Methods

### NewWalletCreateWalletIn

`func NewWalletCreateWalletIn() *WalletCreateWalletIn`

NewWalletCreateWalletIn instantiates a new WalletCreateWalletIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletCreateWalletInWithDefaults

`func NewWalletCreateWalletInWithDefaults() *WalletCreateWalletIn`

NewWalletCreateWalletInWithDefaults instantiates a new WalletCreateWalletIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *WalletCreateWalletIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WalletCreateWalletIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WalletCreateWalletIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *WalletCreateWalletIn) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAgent

`func (o *WalletCreateWalletIn) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *WalletCreateWalletIn) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *WalletCreateWalletIn) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *WalletCreateWalletIn) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetChain

`func (o *WalletCreateWalletIn) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *WalletCreateWalletIn) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *WalletCreateWalletIn) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *WalletCreateWalletIn) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetCustody

`func (o *WalletCreateWalletIn) GetCustody() string`

GetCustody returns the Custody field if non-nil, zero value otherwise.

### GetCustodyOk

`func (o *WalletCreateWalletIn) GetCustodyOk() (*string, bool)`

GetCustodyOk returns a tuple with the Custody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustody

`func (o *WalletCreateWalletIn) SetCustody(v string)`

SetCustody sets Custody field to given value.

### HasCustody

`func (o *WalletCreateWalletIn) HasCustody() bool`

HasCustody returns a boolean if a field has been set.

### GetName

`func (o *WalletCreateWalletIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletCreateWalletIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletCreateWalletIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WalletCreateWalletIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTier

`func (o *WalletCreateWalletIn) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *WalletCreateWalletIn) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *WalletCreateWalletIn) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *WalletCreateWalletIn) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


