# BillingSubscriptionLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the seat type&#39;s display name, e.g. \&quot;Team Premium\&quot;. | [optional] 
**PendingQuantity** | Pointer to **int64** | PendingQuantity is the seats of it the next period opens at, when a change waits for the renewal that charges it; absent when none does. | [optional] 
**Plan** | Pointer to **string** | Plan is the seat type&#39;s catalog plan, e.g. \&quot;team-premium\&quot;. | [optional] 
**Price** | Pointer to **int64** | Price is one seat of it for the row&#39;s period, in cents. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is the seats of it the row holds. | [optional] 
**Seats** | Pointer to **int64** | Seats is the seats of it the period in hand was paid for; 0 when unpaid. | [optional] 

## Methods

### NewBillingSubscriptionLine

`func NewBillingSubscriptionLine() *BillingSubscriptionLine`

NewBillingSubscriptionLine instantiates a new BillingSubscriptionLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionLineWithDefaults

`func NewBillingSubscriptionLineWithDefaults() *BillingSubscriptionLine`

NewBillingSubscriptionLineWithDefaults instantiates a new BillingSubscriptionLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingSubscriptionLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingSubscriptionLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingSubscriptionLine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingSubscriptionLine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPendingQuantity

`func (o *BillingSubscriptionLine) GetPendingQuantity() int64`

GetPendingQuantity returns the PendingQuantity field if non-nil, zero value otherwise.

### GetPendingQuantityOk

`func (o *BillingSubscriptionLine) GetPendingQuantityOk() (*int64, bool)`

GetPendingQuantityOk returns a tuple with the PendingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingQuantity

`func (o *BillingSubscriptionLine) SetPendingQuantity(v int64)`

SetPendingQuantity sets PendingQuantity field to given value.

### HasPendingQuantity

`func (o *BillingSubscriptionLine) HasPendingQuantity() bool`

HasPendingQuantity returns a boolean if a field has been set.

### GetPlan

`func (o *BillingSubscriptionLine) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *BillingSubscriptionLine) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *BillingSubscriptionLine) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *BillingSubscriptionLine) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPrice

`func (o *BillingSubscriptionLine) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingSubscriptionLine) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingSubscriptionLine) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingSubscriptionLine) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingSubscriptionLine) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingSubscriptionLine) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingSubscriptionLine) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingSubscriptionLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSeats

`func (o *BillingSubscriptionLine) GetSeats() int64`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *BillingSubscriptionLine) GetSeatsOk() (*int64, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *BillingSubscriptionLine) SetSeats(v int64)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *BillingSubscriptionLine) HasSeats() bool`

HasSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


