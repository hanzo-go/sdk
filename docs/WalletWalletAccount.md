# WalletWalletAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the account was opened, Unix seconds. Listings order by it, newest first. | [optional] 
**Id** | Pointer to **string** | ID is the account id, minted by the server as \&quot;acct_\&quot; + 24 hex. Wallets name it as their accountId, and it becomes a segment of each of their key refs — so it addresses key material and cannot be reassigned. | [optional] 
**Name** | Pointer to **string** | Name is the label given at creation, trimmed and required. It groups wallets: it is not a key, holds no balance, and is not unique in the org. | [optional] 
**Org** | Pointer to **string** | Org is the tenant that owns the account, stamped from the validated principal rather than taken from the request. Every read is physically scoped to it, so another tenant&#39;s accounts are not reachable at all. | [optional] 

## Methods

### NewWalletWalletAccount

`func NewWalletWalletAccount() *WalletWalletAccount`

NewWalletWalletAccount instantiates a new WalletWalletAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWalletAccountWithDefaults

`func NewWalletWalletAccountWithDefaults() *WalletWalletAccount`

NewWalletWalletAccountWithDefaults instantiates a new WalletWalletAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WalletWalletAccount) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletWalletAccount) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletWalletAccount) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WalletWalletAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *WalletWalletAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletWalletAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletWalletAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WalletWalletAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WalletWalletAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletWalletAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletWalletAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WalletWalletAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *WalletWalletAccount) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *WalletWalletAccount) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *WalletWalletAccount) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *WalletWalletAccount) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


