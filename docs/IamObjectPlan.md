# IamObjectPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PaymentProviders** | Pointer to **[]string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectPlan

`func NewIamObjectPlan() *IamObjectPlan`

NewIamObjectPlan instantiates a new IamObjectPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectPlanWithDefaults

`func NewIamObjectPlanWithDefaults() *IamObjectPlan`

NewIamObjectPlanWithDefaults instantiates a new IamObjectPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *IamObjectPlan) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectPlan) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectPlan) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectPlan) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *IamObjectPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IamObjectPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IamObjectPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IamObjectPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *IamObjectPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamObjectPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamObjectPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamObjectPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectPlan) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectPlan) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectPlan) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectPlan) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamObjectPlan) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamObjectPlan) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamObjectPlan) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamObjectPlan) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *IamObjectPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *IamObjectPlan) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IamObjectPlan) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IamObjectPlan) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IamObjectPlan) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectPlan) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectPlan) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectPlan) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectPlan) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPaymentProviders

`func (o *IamObjectPlan) GetPaymentProviders() []string`

GetPaymentProviders returns the PaymentProviders field if non-nil, zero value otherwise.

### GetPaymentProvidersOk

`func (o *IamObjectPlan) GetPaymentProvidersOk() (*[]string, bool)`

GetPaymentProvidersOk returns a tuple with the PaymentProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProviders

`func (o *IamObjectPlan) SetPaymentProviders(v []string)`

SetPaymentProviders sets PaymentProviders field to given value.

### HasPaymentProviders

`func (o *IamObjectPlan) HasPaymentProviders() bool`

HasPaymentProviders returns a boolean if a field has been set.

### GetPeriod

`func (o *IamObjectPlan) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *IamObjectPlan) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *IamObjectPlan) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *IamObjectPlan) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPrice

`func (o *IamObjectPlan) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *IamObjectPlan) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *IamObjectPlan) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *IamObjectPlan) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProduct

`func (o *IamObjectPlan) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *IamObjectPlan) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *IamObjectPlan) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *IamObjectPlan) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRole

`func (o *IamObjectPlan) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IamObjectPlan) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IamObjectPlan) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IamObjectPlan) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


