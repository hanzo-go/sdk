# BillingResize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | Pointer to [**[]BillingSaleLine**](BillingSaleLine.md) | Lines are the team&#39;s other seat types and the seats of each: [{\&quot;planId\&quot;:\&quot;team-premium\&quot;,\&quot;quantity\&quot;:2}]. A seat type of the team the subscription does not hold yet is added. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is the seats of the subscription&#39;s own seat type; 0 keeps them. | [optional] 

## Methods

### NewBillingResize

`func NewBillingResize() *BillingResize`

NewBillingResize instantiates a new BillingResize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingResizeWithDefaults

`func NewBillingResizeWithDefaults() *BillingResize`

NewBillingResizeWithDefaults instantiates a new BillingResize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *BillingResize) GetLines() []BillingSaleLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *BillingResize) GetLinesOk() (*[]BillingSaleLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *BillingResize) SetLines(v []BillingSaleLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *BillingResize) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingResize) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingResize) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingResize) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingResize) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


