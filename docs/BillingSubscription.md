# BillingSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelAtPeriodEnd** | Pointer to **bool** | CancelAtPeriodEnd reports that the plan ends at CurrentPeriodEnd. | [optional] 
**CanceledAt** | Pointer to **string** | CanceledAt is when it was canceled; absent when it was not. | [optional] 
**ChargedCents** | Pointer to **int64** | ChargedCents is what the period in hand was actually paid, after any discount and net of refunds. 0 when the period is unpaid. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the row was written, RFC 3339. | [optional] 
**CurrentPeriodEnd** | Pointer to **string** | CurrentPeriodEnd is when it ends, RFC 3339. The plan is paid through it. | [optional] 
**CurrentPeriodStart** | Pointer to **string** | CurrentPeriodStart is when the period in hand began, RFC 3339. | [optional] 
**DefaultPaymentMethod** | Pointer to **string** | DefaultPaymentMethod is the saved method a renewal would use. Empty means there is none, so nothing can charge the plan. | [optional] 
**EndedAt** | Pointer to **string** | EndedAt is when it ended; absent while it runs. | [optional] 
**Id** | Pointer to **string** | ID is the subscription&#39;s own id, the one cancel and reactivate take. | [optional] 
**MrrCents** | Pointer to **int64** | MRRCents is what this subscription contributes per month — commerce&#39;s own figure, interval-normalized and multiplied by its seats, so no reader re-derives it from price and interval. | [optional] 
**Plan** | Pointer to [**BillingSubscriptionPlan**](BillingSubscriptionPlan.md) | Plan is the plan as this subscription snapshotted it. | [optional] 
**PlanId** | Pointer to **string** | PlanID is the plan slug, e.g. \&quot;max-20x\&quot;. | [optional] 
**ProviderType** | Pointer to **string** | ProviderType is who collects it: a processor (\&quot;square\&quot;), \&quot;credit\&quot; for a plan bought with credits, \&quot;internal\&quot; or \&quot;bundle\&quot; for a row nobody paid for. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is the seat count; a flat plan holds 1. | [optional] 
**Seats** | Pointer to **int64** | Seats is the seat count the period in hand was paid for — its paid invoice&#39;s quantity, or the quantity recorded with an external period — never the row&#39;s live Quantity, which the holder may change. 0 when the period is unpaid. | [optional] 
**Settled** | Pointer to **string** | Settled is how the period in hand was paid: \&quot;card\&quot; (a card paid all of it), \&quot;external:&lt;processor&gt;\&quot; (a payment recorded as collected outside commerce), \&quot;balance\&quot;, \&quot;credit\&quot; or \&quot;mixed\&quot; (prepaid money). Empty when unpaid or unknown. | [optional] 
**Status** | Pointer to **string** | Status is trialing, active, past_due, canceled or unpaid. Only active and trialing confer the plan. | [optional] 
**TrialEnd** | Pointer to **string** | TrialEnd is when that trial ends; absent when there was none. | [optional] 
**TrialStart** | Pointer to **string** | TrialStart is when a trial began; absent when there was none. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when the row last changed, RFC 3339. | [optional] 
**UserId** | Pointer to **string** | UserID is the billing account that holds it: the org slug for an org&#39;s own plan, \&quot;&lt;org&gt;/&lt;name&gt;\&quot; for a member&#39;s. | [optional] 

## Methods

### NewBillingSubscription

`func NewBillingSubscription() *BillingSubscription`

NewBillingSubscription instantiates a new BillingSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionWithDefaults

`func NewBillingSubscriptionWithDefaults() *BillingSubscription`

NewBillingSubscriptionWithDefaults instantiates a new BillingSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelAtPeriodEnd

`func (o *BillingSubscription) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *BillingSubscription) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *BillingSubscription) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *BillingSubscription) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCanceledAt

`func (o *BillingSubscription) GetCanceledAt() string`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *BillingSubscription) GetCanceledAtOk() (*string, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *BillingSubscription) SetCanceledAt(v string)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *BillingSubscription) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetChargedCents

`func (o *BillingSubscription) GetChargedCents() int64`

GetChargedCents returns the ChargedCents field if non-nil, zero value otherwise.

### GetChargedCentsOk

`func (o *BillingSubscription) GetChargedCentsOk() (*int64, bool)`

GetChargedCentsOk returns a tuple with the ChargedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargedCents

`func (o *BillingSubscription) SetChargedCents(v int64)`

SetChargedCents sets ChargedCents field to given value.

### HasChargedCents

`func (o *BillingSubscription) HasChargedCents() bool`

HasChargedCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingSubscription) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingSubscription) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingSubscription) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *BillingSubscription) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *BillingSubscription) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *BillingSubscription) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *BillingSubscription) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *BillingSubscription) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *BillingSubscription) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *BillingSubscription) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *BillingSubscription) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetDefaultPaymentMethod

`func (o *BillingSubscription) GetDefaultPaymentMethod() string`

GetDefaultPaymentMethod returns the DefaultPaymentMethod field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodOk

`func (o *BillingSubscription) GetDefaultPaymentMethodOk() (*string, bool)`

GetDefaultPaymentMethodOk returns a tuple with the DefaultPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethod

`func (o *BillingSubscription) SetDefaultPaymentMethod(v string)`

SetDefaultPaymentMethod sets DefaultPaymentMethod field to given value.

### HasDefaultPaymentMethod

`func (o *BillingSubscription) HasDefaultPaymentMethod() bool`

HasDefaultPaymentMethod returns a boolean if a field has been set.

### GetEndedAt

`func (o *BillingSubscription) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BillingSubscription) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BillingSubscription) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *BillingSubscription) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetId

`func (o *BillingSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMrrCents

`func (o *BillingSubscription) GetMrrCents() int64`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *BillingSubscription) GetMrrCentsOk() (*int64, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *BillingSubscription) SetMrrCents(v int64)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *BillingSubscription) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetPlan

`func (o *BillingSubscription) GetPlan() BillingSubscriptionPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *BillingSubscription) GetPlanOk() (*BillingSubscriptionPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *BillingSubscription) SetPlan(v BillingSubscriptionPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *BillingSubscription) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *BillingSubscription) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *BillingSubscription) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *BillingSubscription) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *BillingSubscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProviderType

`func (o *BillingSubscription) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *BillingSubscription) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *BillingSubscription) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *BillingSubscription) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingSubscription) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingSubscription) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingSubscription) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingSubscription) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSeats

`func (o *BillingSubscription) GetSeats() int64`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *BillingSubscription) GetSeatsOk() (*int64, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *BillingSubscription) SetSeats(v int64)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *BillingSubscription) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### GetSettled

`func (o *BillingSubscription) GetSettled() string`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *BillingSubscription) GetSettledOk() (*string, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *BillingSubscription) SetSettled(v string)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *BillingSubscription) HasSettled() bool`

HasSettled returns a boolean if a field has been set.

### GetStatus

`func (o *BillingSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrialEnd

`func (o *BillingSubscription) GetTrialEnd() string`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *BillingSubscription) GetTrialEndOk() (*string, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *BillingSubscription) SetTrialEnd(v string)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *BillingSubscription) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetTrialStart

`func (o *BillingSubscription) GetTrialStart() string`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *BillingSubscription) GetTrialStartOk() (*string, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *BillingSubscription) SetTrialStart(v string)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *BillingSubscription) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingSubscription) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingSubscription) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingSubscription) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillingSubscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *BillingSubscription) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BillingSubscription) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BillingSubscription) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BillingSubscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


