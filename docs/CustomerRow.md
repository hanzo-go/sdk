# CustomerRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceCents** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**LastActive** | Pointer to **string** |  | [optional] 
**MrrCents** | Pointer to **int32** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**SpendCents** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | \&quot;active\&quot; | \&quot;suspended\&quot; | [optional] 
**Users** | Pointer to **int32** |  | [optional] 

## Methods

### NewCustomerRow

`func NewCustomerRow() *CustomerRow`

NewCustomerRow instantiates a new CustomerRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRowWithDefaults

`func NewCustomerRowWithDefaults() *CustomerRow`

NewCustomerRowWithDefaults instantiates a new CustomerRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceCents

`func (o *CustomerRow) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *CustomerRow) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *CustomerRow) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *CustomerRow) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetCreated

`func (o *CustomerRow) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerRow) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerRow) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CustomerRow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDisplay

`func (o *CustomerRow) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CustomerRow) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CustomerRow) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CustomerRow) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetLastActive

`func (o *CustomerRow) GetLastActive() string`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *CustomerRow) GetLastActiveOk() (*string, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *CustomerRow) SetLastActive(v string)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *CustomerRow) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetMrrCents

`func (o *CustomerRow) GetMrrCents() int32`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *CustomerRow) GetMrrCentsOk() (*int32, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *CustomerRow) SetMrrCents(v int32)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *CustomerRow) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetOrg

`func (o *CustomerRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CustomerRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CustomerRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CustomerRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *CustomerRow) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *CustomerRow) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *CustomerRow) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *CustomerRow) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetPlan

`func (o *CustomerRow) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CustomerRow) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CustomerRow) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CustomerRow) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSpendCents

`func (o *CustomerRow) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *CustomerRow) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *CustomerRow) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *CustomerRow) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsers

`func (o *CustomerRow) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CustomerRow) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CustomerRow) SetUsers(v int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CustomerRow) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


