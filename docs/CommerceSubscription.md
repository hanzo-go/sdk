# CommerceSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Number** | Pointer to **int32** |  | [optional] [readonly] 
**PlanId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**PeriodStart** | Pointer to **time.Time** |  | [optional] 
**PeriodEnd** | Pointer to **time.Time** |  | [optional] 
**TrialStart** | Pointer to **time.Time** |  | [optional] 
**TrialEnd** | Pointer to **time.Time** |  | [optional] 
**CanceledAt** | Pointer to **time.Time** |  | [optional] 
**Buyer** | Pointer to [**CommerceBuyer**](CommerceBuyer.md) |  | [optional] 
**Plan** | Pointer to [**CommercePlan**](CommercePlan.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCommerceSubscription

`func NewCommerceSubscription() *CommerceSubscription`

NewCommerceSubscription instantiates a new CommerceSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceSubscriptionWithDefaults

`func NewCommerceSubscriptionWithDefaults() *CommerceSubscription`

NewCommerceSubscriptionWithDefaults instantiates a new CommerceSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *CommerceSubscription) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CommerceSubscription) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CommerceSubscription) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CommerceSubscription) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPlanId

`func (o *CommerceSubscription) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CommerceSubscription) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CommerceSubscription) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CommerceSubscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetUserId

`func (o *CommerceSubscription) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommerceSubscription) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommerceSubscription) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommerceSubscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStatus

`func (o *CommerceSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommerceSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommerceSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommerceSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQuantity

`func (o *CommerceSubscription) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CommerceSubscription) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CommerceSubscription) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CommerceSubscription) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPeriodStart

`func (o *CommerceSubscription) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *CommerceSubscription) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *CommerceSubscription) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *CommerceSubscription) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *CommerceSubscription) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *CommerceSubscription) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *CommerceSubscription) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *CommerceSubscription) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetTrialStart

`func (o *CommerceSubscription) GetTrialStart() time.Time`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *CommerceSubscription) GetTrialStartOk() (*time.Time, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *CommerceSubscription) SetTrialStart(v time.Time)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *CommerceSubscription) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### GetTrialEnd

`func (o *CommerceSubscription) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *CommerceSubscription) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *CommerceSubscription) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *CommerceSubscription) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetCanceledAt

`func (o *CommerceSubscription) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *CommerceSubscription) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *CommerceSubscription) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *CommerceSubscription) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetBuyer

`func (o *CommerceSubscription) GetBuyer() CommerceBuyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *CommerceSubscription) GetBuyerOk() (*CommerceBuyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *CommerceSubscription) SetBuyer(v CommerceBuyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *CommerceSubscription) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetPlan

`func (o *CommerceSubscription) GetPlan() CommercePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CommerceSubscription) GetPlanOk() (*CommercePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CommerceSubscription) SetPlan(v CommercePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CommerceSubscription) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetMetadata

`func (o *CommerceSubscription) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommerceSubscription) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommerceSubscription) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommerceSubscription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommerceSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommerceSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommerceSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommerceSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommerceSubscription) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommerceSubscription) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommerceSubscription) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommerceSubscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


