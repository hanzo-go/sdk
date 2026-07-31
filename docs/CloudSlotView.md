# CloudSlotView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlsPubkey** | Pointer to **string** | BLSPubkey is the node&#39;s BLS public key, hex. | [optional] 
**CrName** | Pointer to **string** | CRName is the LuxNetwork custom resource that materializes the node. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the slot was first claimed, as a Unix timestamp. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the Kubernetes namespace the node&#39;s CR lives in. | [optional] 
**Network** | Pointer to **string** | Network is the luxd network slug the node joins. | [optional] 
**NodeID** | Pointer to **string** | NodeID is the luxd node id derived from the sealed staking identity. It is stable across re-claims of the same slot. | [optional] 
**NodeStatus** | Pointer to **string** | NodeStatus is the provisioning state of the node: \&quot;node_created\&quot; once the CR is applied, \&quot;node_pending\&quot; when no cluster is reachable (the slot is still claimed and the keys are still sealed). | [optional] 
**Registration** | Pointer to [**CloudRegistrationView**](CloudRegistrationView.md) | Registration is the queued owner-gated registration, absent until one exists. | [optional] 
**Slot** | Pointer to **int32** | Slot is the validator slot number — the same value as tokenId, under the name the portal reads. | [optional] 
**TokenId** | Pointer to **int32** | TokenID is the GenesisNFT token id that IS this slot. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the slot last changed, as a Unix timestamp. | [optional] 
**Wallet** | Pointer to **string** | Wallet is the lowercase Ethereum address that proved ownership of the NFT. | [optional] 

## Methods

### NewCloudSlotView

`func NewCloudSlotView() *CloudSlotView`

NewCloudSlotView instantiates a new CloudSlotView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSlotViewWithDefaults

`func NewCloudSlotViewWithDefaults() *CloudSlotView`

NewCloudSlotViewWithDefaults instantiates a new CloudSlotView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlsPubkey

`func (o *CloudSlotView) GetBlsPubkey() string`

GetBlsPubkey returns the BlsPubkey field if non-nil, zero value otherwise.

### GetBlsPubkeyOk

`func (o *CloudSlotView) GetBlsPubkeyOk() (*string, bool)`

GetBlsPubkeyOk returns a tuple with the BlsPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlsPubkey

`func (o *CloudSlotView) SetBlsPubkey(v string)`

SetBlsPubkey sets BlsPubkey field to given value.

### HasBlsPubkey

`func (o *CloudSlotView) HasBlsPubkey() bool`

HasBlsPubkey returns a boolean if a field has been set.

### GetCrName

`func (o *CloudSlotView) GetCrName() string`

GetCrName returns the CrName field if non-nil, zero value otherwise.

### GetCrNameOk

`func (o *CloudSlotView) GetCrNameOk() (*string, bool)`

GetCrNameOk returns a tuple with the CrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrName

`func (o *CloudSlotView) SetCrName(v string)`

SetCrName sets CrName field to given value.

### HasCrName

`func (o *CloudSlotView) HasCrName() bool`

HasCrName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudSlotView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSlotView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSlotView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSlotView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNamespace

`func (o *CloudSlotView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CloudSlotView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CloudSlotView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CloudSlotView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNetwork

`func (o *CloudSlotView) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CloudSlotView) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CloudSlotView) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CloudSlotView) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNodeID

`func (o *CloudSlotView) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *CloudSlotView) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *CloudSlotView) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *CloudSlotView) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetNodeStatus

`func (o *CloudSlotView) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *CloudSlotView) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *CloudSlotView) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *CloudSlotView) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.

### GetRegistration

`func (o *CloudSlotView) GetRegistration() CloudRegistrationView`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *CloudSlotView) GetRegistrationOk() (*CloudRegistrationView, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *CloudSlotView) SetRegistration(v CloudRegistrationView)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *CloudSlotView) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetSlot

`func (o *CloudSlotView) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *CloudSlotView) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *CloudSlotView) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *CloudSlotView) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetTokenId

`func (o *CloudSlotView) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CloudSlotView) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CloudSlotView) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *CloudSlotView) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudSlotView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudSlotView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudSlotView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudSlotView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWallet

`func (o *CloudSlotView) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *CloudSlotView) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *CloudSlotView) SetWallet(v string)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *CloudSlotView) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


