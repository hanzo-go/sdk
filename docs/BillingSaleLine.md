# BillingSaleLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** | PlanID is the seat type, e.g. \&quot;team-premium\&quot;. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is the seats of it bought, at least 1. | [optional] 

## Methods

### NewBillingSaleLine

`func NewBillingSaleLine() *BillingSaleLine`

NewBillingSaleLine instantiates a new BillingSaleLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSaleLineWithDefaults

`func NewBillingSaleLineWithDefaults() *BillingSaleLine`

NewBillingSaleLineWithDefaults instantiates a new BillingSaleLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *BillingSaleLine) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *BillingSaleLine) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *BillingSaleLine) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *BillingSaleLine) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingSaleLine) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingSaleLine) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingSaleLine) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingSaleLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


