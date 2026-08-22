# BillingAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** | Role is the caller&#39;s standing. Absent when nobody was named, so an anonymous read returns an account rather than an implied membership. | [optional] 

## Methods

### NewBillingAccount

`func NewBillingAccount() *BillingAccount`

NewBillingAccount instantiates a new BillingAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAccountWithDefaults

`func NewBillingAccountWithDefaults() *BillingAccount`

NewBillingAccountWithDefaults instantiates a new BillingAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BillingAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *BillingAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BillingAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *BillingAccount) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BillingAccount) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BillingAccount) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *BillingAccount) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *BillingAccount) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *BillingAccount) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *BillingAccount) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *BillingAccount) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetRole

`func (o *BillingAccount) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BillingAccount) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BillingAccount) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *BillingAccount) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


