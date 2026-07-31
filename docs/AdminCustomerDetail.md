# AdminCustomerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**BalanceCents** | Pointer to **int64** |  | [optional] 
**SpendCents** | Pointer to **int64** |  | [optional] 
**MrrCents** | Pointer to **int64** |  | [optional] 
**ApiKeys** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to [**[]AdminCustomerUser**](AdminCustomerUser.md) |  | [optional] 
**Transactions** | Pointer to [**[]AdminCustomerTxn**](AdminCustomerTxn.md) |  | [optional] 

## Methods

### NewAdminCustomerDetail

`func NewAdminCustomerDetail() *AdminCustomerDetail`

NewAdminCustomerDetail instantiates a new AdminCustomerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCustomerDetailWithDefaults

`func NewAdminCustomerDetailWithDefaults() *AdminCustomerDetail`

NewAdminCustomerDetailWithDefaults instantiates a new AdminCustomerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *AdminCustomerDetail) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AdminCustomerDetail) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AdminCustomerDetail) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AdminCustomerDetail) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetDisplay

`func (o *AdminCustomerDetail) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *AdminCustomerDetail) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *AdminCustomerDetail) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *AdminCustomerDetail) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *AdminCustomerDetail) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *AdminCustomerDetail) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *AdminCustomerDetail) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *AdminCustomerDetail) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetPlan

`func (o *AdminCustomerDetail) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AdminCustomerDetail) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AdminCustomerDetail) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AdminCustomerDetail) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetStatus

`func (o *AdminCustomerDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminCustomerDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminCustomerDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminCustomerDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *AdminCustomerDetail) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AdminCustomerDetail) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AdminCustomerDetail) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AdminCustomerDetail) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetBalanceCents

`func (o *AdminCustomerDetail) GetBalanceCents() int64`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *AdminCustomerDetail) GetBalanceCentsOk() (*int64, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *AdminCustomerDetail) SetBalanceCents(v int64)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *AdminCustomerDetail) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetSpendCents

`func (o *AdminCustomerDetail) GetSpendCents() int64`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *AdminCustomerDetail) GetSpendCentsOk() (*int64, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *AdminCustomerDetail) SetSpendCents(v int64)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *AdminCustomerDetail) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetMrrCents

`func (o *AdminCustomerDetail) GetMrrCents() int64`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *AdminCustomerDetail) GetMrrCentsOk() (*int64, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *AdminCustomerDetail) SetMrrCents(v int64)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *AdminCustomerDetail) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetApiKeys

`func (o *AdminCustomerDetail) GetApiKeys() int32`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *AdminCustomerDetail) GetApiKeysOk() (*int32, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *AdminCustomerDetail) SetApiKeys(v int32)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *AdminCustomerDetail) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetUsers

`func (o *AdminCustomerDetail) GetUsers() []AdminCustomerUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AdminCustomerDetail) GetUsersOk() (*[]AdminCustomerUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AdminCustomerDetail) SetUsers(v []AdminCustomerUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AdminCustomerDetail) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetTransactions

`func (o *AdminCustomerDetail) GetTransactions() []AdminCustomerTxn`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *AdminCustomerDetail) GetTransactionsOk() (*[]AdminCustomerTxn, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *AdminCustomerDetail) SetTransactions(v []AdminCustomerTxn)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *AdminCustomerDetail) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


