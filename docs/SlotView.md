# SlotView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlsPubkey** | Pointer to **string** | BLSPubkey is the node&#39;s BLS public key, hex. | [optional] 
**CrName** | Pointer to **string** | CRName is the LuxNetwork custom resource that materializes the node. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the slot was first claimed, as a Unix timestamp. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the Kubernetes namespace the node&#39;s CR lives in. | [optional] 
**Network** | Pointer to **string** | Network is the luxd network slug the node joins. | [optional] 
**NodeID** | Pointer to **string** | NodeID is the luxd node id derived from the sealed staking identity. It is stable across re-claims of the same slot. | [optional] 
**NodeStatus** | Pointer to **string** | NodeStatus is the provisioning state of the node: \&quot;node_created\&quot; once the CR is applied, \&quot;node_pending\&quot; when no cluster is reachable (the slot is still claimed and the keys are still sealed). | [optional] 
**Registration** | Pointer to [**RegistrationView**](RegistrationView.md) | Registration is the queued owner-gated registration, absent until one exists. | [optional] 
**Slot** | Pointer to **int32** | Slot is the validator slot number — the same value as tokenId, under the name the portal reads. | [optional] 
**TokenId** | Pointer to **int32** | TokenID is the GenesisNFT token id that IS this slot. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the slot last changed, as a Unix timestamp. | [optional] 
**Wallet** | Pointer to **string** | Wallet is the lowercase Ethereum address that proved ownership of the NFT. | [optional] 

## Methods

### NewSlotView

`func NewSlotView() *SlotView`

NewSlotView instantiates a new SlotView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlotViewWithDefaults

`func NewSlotViewWithDefaults() *SlotView`

NewSlotViewWithDefaults instantiates a new SlotView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlsPubkey

`func (o *SlotView) GetBlsPubkey() string`

GetBlsPubkey returns the BlsPubkey field if non-nil, zero value otherwise.

### GetBlsPubkeyOk

`func (o *SlotView) GetBlsPubkeyOk() (*string, bool)`

GetBlsPubkeyOk returns a tuple with the BlsPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlsPubkey

`func (o *SlotView) SetBlsPubkey(v string)`

SetBlsPubkey sets BlsPubkey field to given value.

### HasBlsPubkey

`func (o *SlotView) HasBlsPubkey() bool`

HasBlsPubkey returns a boolean if a field has been set.

### GetCrName

`func (o *SlotView) GetCrName() string`

GetCrName returns the CrName field if non-nil, zero value otherwise.

### GetCrNameOk

`func (o *SlotView) GetCrNameOk() (*string, bool)`

GetCrNameOk returns a tuple with the CrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrName

`func (o *SlotView) SetCrName(v string)`

SetCrName sets CrName field to given value.

### HasCrName

`func (o *SlotView) HasCrName() bool`

HasCrName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SlotView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SlotView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SlotView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SlotView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNamespace

`func (o *SlotView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SlotView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SlotView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SlotView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNetwork

`func (o *SlotView) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SlotView) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SlotView) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SlotView) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNodeID

`func (o *SlotView) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *SlotView) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *SlotView) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *SlotView) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetNodeStatus

`func (o *SlotView) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *SlotView) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *SlotView) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *SlotView) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.

### GetRegistration

`func (o *SlotView) GetRegistration() RegistrationView`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SlotView) GetRegistrationOk() (*RegistrationView, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SlotView) SetRegistration(v RegistrationView)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SlotView) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetSlot

`func (o *SlotView) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *SlotView) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *SlotView) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *SlotView) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetTokenId

`func (o *SlotView) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SlotView) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SlotView) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *SlotView) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SlotView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SlotView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SlotView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SlotView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWallet

`func (o *SlotView) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *SlotView) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *SlotView) SetWallet(v string)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *SlotView) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


