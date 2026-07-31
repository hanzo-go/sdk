# CloudCustomerDetailData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | Pointer to **int32** |  | [optional] 
**BalanceCents** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**MrrCents** | Pointer to **int32** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**SpendCents** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Transactions** | Pointer to [**[]CloudCustomerTxn**](CloudCustomerTxn.md) |  | [optional] 
**Users** | Pointer to [**[]CloudCustomerUser**](CloudCustomerUser.md) |  | [optional] 

## Methods

### NewCloudCustomerDetailData

`func NewCloudCustomerDetailData() *CloudCustomerDetailData`

NewCloudCustomerDetailData instantiates a new CloudCustomerDetailData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCustomerDetailDataWithDefaults

`func NewCloudCustomerDetailDataWithDefaults() *CloudCustomerDetailData`

NewCloudCustomerDetailDataWithDefaults instantiates a new CloudCustomerDetailData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *CloudCustomerDetailData) GetApiKeys() int32`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *CloudCustomerDetailData) GetApiKeysOk() (*int32, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *CloudCustomerDetailData) SetApiKeys(v int32)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *CloudCustomerDetailData) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetBalanceCents

`func (o *CloudCustomerDetailData) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *CloudCustomerDetailData) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *CloudCustomerDetailData) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *CloudCustomerDetailData) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetCreated

`func (o *CloudCustomerDetailData) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudCustomerDetailData) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudCustomerDetailData) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudCustomerDetailData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDisplay

`func (o *CloudCustomerDetailData) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CloudCustomerDetailData) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CloudCustomerDetailData) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CloudCustomerDetailData) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetMrrCents

`func (o *CloudCustomerDetailData) GetMrrCents() int32`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *CloudCustomerDetailData) GetMrrCentsOk() (*int32, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *CloudCustomerDetailData) SetMrrCents(v int32)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *CloudCustomerDetailData) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetOrg

`func (o *CloudCustomerDetailData) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudCustomerDetailData) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudCustomerDetailData) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudCustomerDetailData) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *CloudCustomerDetailData) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *CloudCustomerDetailData) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *CloudCustomerDetailData) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *CloudCustomerDetailData) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetPlan

`func (o *CloudCustomerDetailData) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CloudCustomerDetailData) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CloudCustomerDetailData) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CloudCustomerDetailData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSpendCents

`func (o *CloudCustomerDetailData) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *CloudCustomerDetailData) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *CloudCustomerDetailData) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *CloudCustomerDetailData) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCustomerDetailData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCustomerDetailData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCustomerDetailData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCustomerDetailData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactions

`func (o *CloudCustomerDetailData) GetTransactions() []CloudCustomerTxn`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *CloudCustomerDetailData) GetTransactionsOk() (*[]CloudCustomerTxn, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *CloudCustomerDetailData) SetTransactions(v []CloudCustomerTxn)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *CloudCustomerDetailData) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetUsers

`func (o *CloudCustomerDetailData) GetUsers() []CloudCustomerUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CloudCustomerDetailData) GetUsersOk() (*[]CloudCustomerUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CloudCustomerDetailData) SetUsers(v []CloudCustomerUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CloudCustomerDetailData) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


