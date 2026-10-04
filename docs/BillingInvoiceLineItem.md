# BillingInvoiceLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** | The billed period. Both carry omitempty and neither is ever empty, which looks contradictory and is not: the shape they reproduce is a time value, and omitempty does nothing to a struct — so those keys render even for the zero instant. The adapter formats the zero instant rather than skipping it, which is what keeps the two wires the same. | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UnitPrice** | Pointer to **int64** |  | [optional] 

## Methods

### NewBillingInvoiceLineItem

`func NewBillingInvoiceLineItem() *BillingInvoiceLineItem`

NewBillingInvoiceLineItem instantiates a new BillingInvoiceLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceLineItemWithDefaults

`func NewBillingInvoiceLineItemWithDefaults() *BillingInvoiceLineItem`

NewBillingInvoiceLineItemWithDefaults instantiates a new BillingInvoiceLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingInvoiceLineItem) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingInvoiceLineItem) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingInvoiceLineItem) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingInvoiceLineItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingInvoiceLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInvoiceLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInvoiceLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInvoiceLineItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *BillingInvoiceLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingInvoiceLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingInvoiceLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingInvoiceLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *BillingInvoiceLineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoiceLineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoiceLineItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInvoiceLineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeterId

`func (o *BillingInvoiceLineItem) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *BillingInvoiceLineItem) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *BillingInvoiceLineItem) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *BillingInvoiceLineItem) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *BillingInvoiceLineItem) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BillingInvoiceLineItem) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BillingInvoiceLineItem) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *BillingInvoiceLineItem) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *BillingInvoiceLineItem) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BillingInvoiceLineItem) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BillingInvoiceLineItem) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *BillingInvoiceLineItem) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPlanId

`func (o *BillingInvoiceLineItem) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *BillingInvoiceLineItem) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *BillingInvoiceLineItem) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *BillingInvoiceLineItem) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *BillingInvoiceLineItem) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *BillingInvoiceLineItem) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *BillingInvoiceLineItem) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *BillingInvoiceLineItem) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingInvoiceLineItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingInvoiceLineItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingInvoiceLineItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingInvoiceLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetType

`func (o *BillingInvoiceLineItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingInvoiceLineItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingInvoiceLineItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingInvoiceLineItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnitPrice

`func (o *BillingInvoiceLineItem) GetUnitPrice() int64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *BillingInvoiceLineItem) GetUnitPriceOk() (*int64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *BillingInvoiceLineItem) SetUnitPrice(v int64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *BillingInvoiceLineItem) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


