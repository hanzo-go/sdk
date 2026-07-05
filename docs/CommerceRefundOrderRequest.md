# CommerceRefundOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Amount to refund in cents (optional, full refund if omitted) | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewCommerceRefundOrderRequest

`func NewCommerceRefundOrderRequest() *CommerceRefundOrderRequest`

NewCommerceRefundOrderRequest instantiates a new CommerceRefundOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceRefundOrderRequestWithDefaults

`func NewCommerceRefundOrderRequestWithDefaults() *CommerceRefundOrderRequest`

NewCommerceRefundOrderRequestWithDefaults instantiates a new CommerceRefundOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CommerceRefundOrderRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CommerceRefundOrderRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CommerceRefundOrderRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CommerceRefundOrderRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetReason

`func (o *CommerceRefundOrderRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CommerceRefundOrderRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CommerceRefundOrderRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CommerceRefundOrderRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


