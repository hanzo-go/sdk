# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelAtPeriodEnd** | Pointer to **bool** |  | [optional] 
**CanceledAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CurrentPeriodEnd** | Pointer to **string** |  | [optional] 
**CurrentPeriodStart** | Pointer to **string** |  | [optional] 
**DefaultPaymentMethod** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MrrCents** | Pointer to **int32** | MRRCents is what this subscription contributes per month — commerce&#39;s own figure, interval-normalized and multiplied by its seats, so no reader re-derives it from price and interval. | [optional] 
**Plan** | Pointer to [**SubscriptionPlan**](SubscriptionPlan.md) |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**ProviderType** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TrialEnd** | Pointer to **string** |  | [optional] 
**TrialStart** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelAtPeriodEnd

`func (o *Subscription) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *Subscription) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *Subscription) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *Subscription) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCanceledAt

`func (o *Subscription) GetCanceledAt() string`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *Subscription) GetCanceledAtOk() (*string, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *Subscription) SetCanceledAt(v string)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *Subscription) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscription) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *Subscription) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *Subscription) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *Subscription) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *Subscription) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *Subscription) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *Subscription) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *Subscription) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *Subscription) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetDefaultPaymentMethod

`func (o *Subscription) GetDefaultPaymentMethod() string`

GetDefaultPaymentMethod returns the DefaultPaymentMethod field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodOk

`func (o *Subscription) GetDefaultPaymentMethodOk() (*string, bool)`

GetDefaultPaymentMethodOk returns a tuple with the DefaultPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethod

`func (o *Subscription) SetDefaultPaymentMethod(v string)`

SetDefaultPaymentMethod sets DefaultPaymentMethod field to given value.

### HasDefaultPaymentMethod

`func (o *Subscription) HasDefaultPaymentMethod() bool`

HasDefaultPaymentMethod returns a boolean if a field has been set.

### GetEndedAt

`func (o *Subscription) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Subscription) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Subscription) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Subscription) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMrrCents

`func (o *Subscription) GetMrrCents() int32`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *Subscription) GetMrrCentsOk() (*int32, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *Subscription) SetMrrCents(v int32)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *Subscription) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetPlan

`func (o *Subscription) GetPlan() SubscriptionPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Subscription) GetPlanOk() (*SubscriptionPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Subscription) SetPlan(v SubscriptionPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Subscription) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *Subscription) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *Subscription) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *Subscription) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *Subscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProviderType

`func (o *Subscription) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *Subscription) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *Subscription) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *Subscription) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetQuantity

`func (o *Subscription) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Subscription) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Subscription) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Subscription) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrialEnd

`func (o *Subscription) GetTrialEnd() string`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *Subscription) GetTrialEndOk() (*string, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *Subscription) SetTrialEnd(v string)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *Subscription) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetTrialStart

`func (o *Subscription) GetTrialStart() string`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *Subscription) GetTrialStartOk() (*string, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *Subscription) SetTrialStart(v string)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *Subscription) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Subscription) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subscription) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subscription) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Subscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *Subscription) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Subscription) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Subscription) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Subscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


