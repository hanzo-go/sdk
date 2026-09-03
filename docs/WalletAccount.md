# WalletAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the account was opened, Unix seconds. Listings order by it, newest first. | [optional] 
**Id** | Pointer to **string** | ID is the account id, minted by the server as \&quot;acct_\&quot; + 24 hex. Wallets name it as their accountId, and it becomes a segment of each of their key refs — so it addresses key material and cannot be reassigned. | [optional] 
**Name** | Pointer to **string** | Name is the label given at creation, trimmed and required. It groups wallets: it is not a key, holds no balance, and is not unique in the org. | [optional] 
**Org** | Pointer to **string** | Org is the tenant that owns the account, stamped from the validated principal rather than taken from the request. Every read is physically scoped to it, so another tenant&#39;s accounts are not reachable at all. | [optional] 

## Methods

### NewWalletAccount

`func NewWalletAccount() *WalletAccount`

NewWalletAccount instantiates a new WalletAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletAccountWithDefaults

`func NewWalletAccountWithDefaults() *WalletAccount`

NewWalletAccountWithDefaults instantiates a new WalletAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WalletAccount) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletAccount) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletAccount) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WalletAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *WalletAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WalletAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WalletAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WalletAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *WalletAccount) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *WalletAccount) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *WalletAccount) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *WalletAccount) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


