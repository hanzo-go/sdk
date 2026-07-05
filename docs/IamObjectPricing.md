# IamObjectPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Plans** | Pointer to **[]string** |  | [optional] 
**TrialDuration** | Pointer to **int64** |  | [optional] 

## Methods

### NewIamObjectPricing

`func NewIamObjectPricing() *IamObjectPricing`

NewIamObjectPricing instantiates a new IamObjectPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectPricingWithDefaults

`func NewIamObjectPricingWithDefaults() *IamObjectPricing`

NewIamObjectPricingWithDefaults instantiates a new IamObjectPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *IamObjectPricing) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamObjectPricing) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamObjectPricing) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamObjectPricing) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectPricing) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectPricing) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectPricing) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectPricing) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *IamObjectPricing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamObjectPricing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamObjectPricing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamObjectPricing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectPricing) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectPricing) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectPricing) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectPricing) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamObjectPricing) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamObjectPricing) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamObjectPricing) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamObjectPricing) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *IamObjectPricing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectPricing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectPricing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectPricing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectPricing) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectPricing) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectPricing) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectPricing) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPlans

`func (o *IamObjectPricing) GetPlans() []string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *IamObjectPricing) GetPlansOk() (*[]string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *IamObjectPricing) SetPlans(v []string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *IamObjectPricing) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetTrialDuration

`func (o *IamObjectPricing) GetTrialDuration() int64`

GetTrialDuration returns the TrialDuration field if non-nil, zero value otherwise.

### GetTrialDurationOk

`func (o *IamObjectPricing) GetTrialDurationOk() (*int64, bool)`

GetTrialDurationOk returns a tuple with the TrialDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDuration

`func (o *IamObjectPricing) SetTrialDuration(v int64)`

SetTrialDuration sets TrialDuration field to given value.

### HasTrialDuration

`func (o *IamObjectPricing) HasTrialDuration() bool`

HasTrialDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


