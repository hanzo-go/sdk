# PricingSubscriptionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PriceMonthly** | Pointer to **float32** |  | [optional] 
**PriceAnnual** | Pointer to **float32** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Limits** | Pointer to [**PricingSubscriptionPlanLimits**](PricingSubscriptionPlanLimits.md) |  | [optional] 
**Payouts** | Pointer to [**PricingSubscriptionPlanPayouts**](PricingSubscriptionPlanPayouts.md) |  | [optional] 

## Methods

### NewPricingSubscriptionPlan

`func NewPricingSubscriptionPlan() *PricingSubscriptionPlan`

NewPricingSubscriptionPlan instantiates a new PricingSubscriptionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingSubscriptionPlanWithDefaults

`func NewPricingSubscriptionPlanWithDefaults() *PricingSubscriptionPlan`

NewPricingSubscriptionPlanWithDefaults instantiates a new PricingSubscriptionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PricingSubscriptionPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingSubscriptionPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingSubscriptionPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PricingSubscriptionPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PricingSubscriptionPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingSubscriptionPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingSubscriptionPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PricingSubscriptionPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PricingSubscriptionPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PricingSubscriptionPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PricingSubscriptionPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PricingSubscriptionPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriceMonthly

`func (o *PricingSubscriptionPlan) GetPriceMonthly() float32`

GetPriceMonthly returns the PriceMonthly field if non-nil, zero value otherwise.

### GetPriceMonthlyOk

`func (o *PricingSubscriptionPlan) GetPriceMonthlyOk() (*float32, bool)`

GetPriceMonthlyOk returns a tuple with the PriceMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonthly

`func (o *PricingSubscriptionPlan) SetPriceMonthly(v float32)`

SetPriceMonthly sets PriceMonthly field to given value.

### HasPriceMonthly

`func (o *PricingSubscriptionPlan) HasPriceMonthly() bool`

HasPriceMonthly returns a boolean if a field has been set.

### GetPriceAnnual

`func (o *PricingSubscriptionPlan) GetPriceAnnual() float32`

GetPriceAnnual returns the PriceAnnual field if non-nil, zero value otherwise.

### GetPriceAnnualOk

`func (o *PricingSubscriptionPlan) GetPriceAnnualOk() (*float32, bool)`

GetPriceAnnualOk returns a tuple with the PriceAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAnnual

`func (o *PricingSubscriptionPlan) SetPriceAnnual(v float32)`

SetPriceAnnual sets PriceAnnual field to given value.

### HasPriceAnnual

`func (o *PricingSubscriptionPlan) HasPriceAnnual() bool`

HasPriceAnnual returns a boolean if a field has been set.

### GetCategory

`func (o *PricingSubscriptionPlan) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PricingSubscriptionPlan) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PricingSubscriptionPlan) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PricingSubscriptionPlan) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFeatures

`func (o *PricingSubscriptionPlan) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PricingSubscriptionPlan) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PricingSubscriptionPlan) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PricingSubscriptionPlan) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetLimits

`func (o *PricingSubscriptionPlan) GetLimits() PricingSubscriptionPlanLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *PricingSubscriptionPlan) GetLimitsOk() (*PricingSubscriptionPlanLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *PricingSubscriptionPlan) SetLimits(v PricingSubscriptionPlanLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *PricingSubscriptionPlan) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetPayouts

`func (o *PricingSubscriptionPlan) GetPayouts() PricingSubscriptionPlanPayouts`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *PricingSubscriptionPlan) GetPayoutsOk() (*PricingSubscriptionPlanPayouts, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *PricingSubscriptionPlan) SetPayouts(v PricingSubscriptionPlanPayouts)`

SetPayouts sets Payouts field to given value.

### HasPayouts

`func (o *PricingSubscriptionPlan) HasPayouts() bool`

HasPayouts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


