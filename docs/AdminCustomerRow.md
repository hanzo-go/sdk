# AdminCustomerRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **int32** |  | [optional] 
**BalanceCents** | Pointer to **int64** |  | [optional] 
**SpendCents** | Pointer to **int64** |  | [optional] 
**MrrCents** | Pointer to **int64** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**LastActive** | Pointer to **string** |  | [optional] 

## Methods

### NewAdminCustomerRow

`func NewAdminCustomerRow() *AdminCustomerRow`

NewAdminCustomerRow instantiates a new AdminCustomerRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCustomerRowWithDefaults

`func NewAdminCustomerRowWithDefaults() *AdminCustomerRow`

NewAdminCustomerRowWithDefaults instantiates a new AdminCustomerRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *AdminCustomerRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AdminCustomerRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AdminCustomerRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AdminCustomerRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetDisplay

`func (o *AdminCustomerRow) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *AdminCustomerRow) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *AdminCustomerRow) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *AdminCustomerRow) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *AdminCustomerRow) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *AdminCustomerRow) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *AdminCustomerRow) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *AdminCustomerRow) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetPlan

`func (o *AdminCustomerRow) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AdminCustomerRow) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AdminCustomerRow) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AdminCustomerRow) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetStatus

`func (o *AdminCustomerRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminCustomerRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminCustomerRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminCustomerRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsers

`func (o *AdminCustomerRow) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AdminCustomerRow) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AdminCustomerRow) SetUsers(v int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AdminCustomerRow) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetBalanceCents

`func (o *AdminCustomerRow) GetBalanceCents() int64`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *AdminCustomerRow) GetBalanceCentsOk() (*int64, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *AdminCustomerRow) SetBalanceCents(v int64)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *AdminCustomerRow) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetSpendCents

`func (o *AdminCustomerRow) GetSpendCents() int64`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *AdminCustomerRow) GetSpendCentsOk() (*int64, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *AdminCustomerRow) SetSpendCents(v int64)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *AdminCustomerRow) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetMrrCents

`func (o *AdminCustomerRow) GetMrrCents() int64`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *AdminCustomerRow) GetMrrCentsOk() (*int64, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *AdminCustomerRow) SetMrrCents(v int64)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *AdminCustomerRow) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetCreated

`func (o *AdminCustomerRow) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AdminCustomerRow) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AdminCustomerRow) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AdminCustomerRow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastActive

`func (o *AdminCustomerRow) GetLastActive() string`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *AdminCustomerRow) GetLastActiveOk() (*string, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *AdminCustomerRow) SetLastActive(v string)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *AdminCustomerRow) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


