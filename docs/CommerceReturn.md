# CommerceReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]CommerceLineItem**](CommerceLineItem.md) |  | [optional] 
**RefundAmount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCommerceReturn

`func NewCommerceReturn() *CommerceReturn`

NewCommerceReturn instantiates a new CommerceReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceReturnWithDefaults

`func NewCommerceReturnWithDefaults() *CommerceReturn`

NewCommerceReturnWithDefaults instantiates a new CommerceReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceReturn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceReturn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceReturn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceReturn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *CommerceReturn) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CommerceReturn) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CommerceReturn) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CommerceReturn) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *CommerceReturn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommerceReturn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommerceReturn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommerceReturn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *CommerceReturn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CommerceReturn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CommerceReturn) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CommerceReturn) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetItems

`func (o *CommerceReturn) GetItems() []CommerceLineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CommerceReturn) GetItemsOk() (*[]CommerceLineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CommerceReturn) SetItems(v []CommerceLineItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CommerceReturn) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRefundAmount

`func (o *CommerceReturn) GetRefundAmount() int32`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *CommerceReturn) GetRefundAmountOk() (*int32, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *CommerceReturn) SetRefundAmount(v int32)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *CommerceReturn) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommerceReturn) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommerceReturn) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommerceReturn) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommerceReturn) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


