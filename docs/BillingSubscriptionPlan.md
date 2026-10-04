# BillingSubscriptionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency is the ISO code the price is in, lower-cased. | [optional] 
**Id** | Pointer to **string** | ID is the plan row the subscription was opened on. | [optional] 
**Interval** | Pointer to **string** | Interval is the billing period: \&quot;month\&quot; or \&quot;year\&quot;. | [optional] 
**Name** | Pointer to **string** | Name is the plan&#39;s display name, e.g. \&quot;Max 20x\&quot;. | [optional] 
**Price** | Pointer to **int64** | Price is what one period costs per seat, in cents, as the subscription snapshotted it: a later catalog price change does not move it. | [optional] 

## Methods

### NewBillingSubscriptionPlan

`func NewBillingSubscriptionPlan() *BillingSubscriptionPlan`

NewBillingSubscriptionPlan instantiates a new BillingSubscriptionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionPlanWithDefaults

`func NewBillingSubscriptionPlanWithDefaults() *BillingSubscriptionPlan`

NewBillingSubscriptionPlanWithDefaults instantiates a new BillingSubscriptionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *BillingSubscriptionPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingSubscriptionPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingSubscriptionPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingSubscriptionPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *BillingSubscriptionPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSubscriptionPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSubscriptionPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingSubscriptionPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *BillingSubscriptionPlan) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BillingSubscriptionPlan) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BillingSubscriptionPlan) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BillingSubscriptionPlan) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *BillingSubscriptionPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingSubscriptionPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingSubscriptionPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingSubscriptionPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *BillingSubscriptionPlan) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingSubscriptionPlan) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingSubscriptionPlan) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingSubscriptionPlan) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


