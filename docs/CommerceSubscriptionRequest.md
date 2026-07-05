# CommerceSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** |  | 
**Quantity** | Pointer to **int32** |  | [optional] [default to 1]
**Buyer** | Pointer to [**CommerceBuyer**](CommerceBuyer.md) |  | [optional] 
**PaymentMethod** | Pointer to [**CommercePaymentRequest**](CommercePaymentRequest.md) |  | [optional] 

## Methods

### NewCommerceSubscriptionRequest

`func NewCommerceSubscriptionRequest(planId string, ) *CommerceSubscriptionRequest`

NewCommerceSubscriptionRequest instantiates a new CommerceSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceSubscriptionRequestWithDefaults

`func NewCommerceSubscriptionRequestWithDefaults() *CommerceSubscriptionRequest`

NewCommerceSubscriptionRequestWithDefaults instantiates a new CommerceSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CommerceSubscriptionRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CommerceSubscriptionRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CommerceSubscriptionRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetQuantity

`func (o *CommerceSubscriptionRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CommerceSubscriptionRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CommerceSubscriptionRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CommerceSubscriptionRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetBuyer

`func (o *CommerceSubscriptionRequest) GetBuyer() CommerceBuyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *CommerceSubscriptionRequest) GetBuyerOk() (*CommerceBuyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *CommerceSubscriptionRequest) SetBuyer(v CommerceBuyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *CommerceSubscriptionRequest) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CommerceSubscriptionRequest) GetPaymentMethod() CommercePaymentRequest`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CommerceSubscriptionRequest) GetPaymentMethodOk() (*CommercePaymentRequest, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CommerceSubscriptionRequest) SetPaymentMethod(v CommercePaymentRequest)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CommerceSubscriptionRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


