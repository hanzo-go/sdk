# BillingTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to [**BillingTierBalance**](BillingTierBalance.md) |  | [optional] 
**Plan** | Pointer to **string** | Plan is the catalog rung the subject is served as (\&quot;\&quot; with no plan). A tier is a class of rungs; a bound that differs within one class keys on this. | [optional] 
**Subscription** | Pointer to **string** | Subscription is the id of the subscription row whose plan is the served rung, \&quot;\&quot; when the tier is not a subscription&#39;s. | [optional] 
**Tier** | Pointer to [**BillingTierLimits**](BillingTierLimits.md) |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingTier

`func NewBillingTier() *BillingTier`

NewBillingTier instantiates a new BillingTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingTierWithDefaults

`func NewBillingTierWithDefaults() *BillingTier`

NewBillingTierWithDefaults instantiates a new BillingTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BillingTier) GetBalance() BillingTierBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BillingTier) GetBalanceOk() (*BillingTierBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BillingTier) SetBalance(v BillingTierBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BillingTier) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetPlan

`func (o *BillingTier) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *BillingTier) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *BillingTier) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *BillingTier) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *BillingTier) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *BillingTier) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *BillingTier) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *BillingTier) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTier

`func (o *BillingTier) GetTier() BillingTierLimits`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *BillingTier) GetTierOk() (*BillingTierLimits, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *BillingTier) SetTier(v BillingTierLimits)`

SetTier sets Tier field to given value.

### HasTier

`func (o *BillingTier) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUser

`func (o *BillingTier) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BillingTier) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BillingTier) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *BillingTier) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


