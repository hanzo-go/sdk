# BillingPaypalOrderIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | AmountCents is the top-up amount in whole cents when PlanID is empty. | [optional] 
**BalanceCents** | Pointer to **int64** | BalanceCents is an organization&#39;s optional starting usage balance. | [optional] 
**Coupon** | Pointer to **string** | Coupon is an optional plan coupon code for a monthly subscription. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO-4217 code (default usd). | [optional] 
**Interval** | Pointer to **string** | Interval selects the billing term (\&quot;month\&quot; or \&quot;year\&quot;). | [optional] 
**Level** | Pointer to **int64** | Level selects which published price level of the plan to buy. | [optional] 
**Lines** | Pointer to [**[]BillingSaleLine**](BillingSaleLine.md) | Lines are a team&#39;s additional seat types bought on the same subscription. | [optional] 
**PlanId** | Pointer to **string** | PlanID is the plan slug when buying a plan&#39;s first period; empty for a top-up. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is the seat count for a per-seat plan. | [optional] 
**StoreId** | Pointer to **string** | StoreID scopes the plan catalog when set. | [optional] 

## Methods

### NewBillingPaypalOrderIn

`func NewBillingPaypalOrderIn() *BillingPaypalOrderIn`

NewBillingPaypalOrderIn instantiates a new BillingPaypalOrderIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPaypalOrderInWithDefaults

`func NewBillingPaypalOrderInWithDefaults() *BillingPaypalOrderIn`

NewBillingPaypalOrderInWithDefaults instantiates a new BillingPaypalOrderIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *BillingPaypalOrderIn) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BillingPaypalOrderIn) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BillingPaypalOrderIn) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BillingPaypalOrderIn) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetBalanceCents

`func (o *BillingPaypalOrderIn) GetBalanceCents() int64`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *BillingPaypalOrderIn) GetBalanceCentsOk() (*int64, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *BillingPaypalOrderIn) SetBalanceCents(v int64)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *BillingPaypalOrderIn) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetCoupon

`func (o *BillingPaypalOrderIn) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *BillingPaypalOrderIn) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *BillingPaypalOrderIn) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *BillingPaypalOrderIn) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingPaypalOrderIn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingPaypalOrderIn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingPaypalOrderIn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingPaypalOrderIn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInterval

`func (o *BillingPaypalOrderIn) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BillingPaypalOrderIn) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BillingPaypalOrderIn) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BillingPaypalOrderIn) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLevel

`func (o *BillingPaypalOrderIn) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *BillingPaypalOrderIn) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *BillingPaypalOrderIn) SetLevel(v int64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *BillingPaypalOrderIn) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLines

`func (o *BillingPaypalOrderIn) GetLines() []BillingSaleLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *BillingPaypalOrderIn) GetLinesOk() (*[]BillingSaleLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *BillingPaypalOrderIn) SetLines(v []BillingSaleLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *BillingPaypalOrderIn) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetPlanId

`func (o *BillingPaypalOrderIn) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *BillingPaypalOrderIn) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *BillingPaypalOrderIn) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *BillingPaypalOrderIn) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingPaypalOrderIn) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingPaypalOrderIn) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingPaypalOrderIn) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingPaypalOrderIn) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStoreId

`func (o *BillingPaypalOrderIn) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *BillingPaypalOrderIn) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *BillingPaypalOrderIn) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *BillingPaypalOrderIn) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


