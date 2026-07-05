# CommerceGetOrderStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CommerceOrderStatus**](CommerceOrderStatus.md) |  | [optional] 
**PaymentStatus** | Pointer to [**CommercePaymentStatus**](CommercePaymentStatus.md) |  | [optional] 

## Methods

### NewCommerceGetOrderStatus200Response

`func NewCommerceGetOrderStatus200Response() *CommerceGetOrderStatus200Response`

NewCommerceGetOrderStatus200Response instantiates a new CommerceGetOrderStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceGetOrderStatus200ResponseWithDefaults

`func NewCommerceGetOrderStatus200ResponseWithDefaults() *CommerceGetOrderStatus200Response`

NewCommerceGetOrderStatus200ResponseWithDefaults instantiates a new CommerceGetOrderStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CommerceGetOrderStatus200Response) GetStatus() CommerceOrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommerceGetOrderStatus200Response) GetStatusOk() (*CommerceOrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommerceGetOrderStatus200Response) SetStatus(v CommerceOrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommerceGetOrderStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *CommerceGetOrderStatus200Response) GetPaymentStatus() CommercePaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CommerceGetOrderStatus200Response) GetPaymentStatusOk() (*CommercePaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CommerceGetOrderStatus200Response) SetPaymentStatus(v CommercePaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *CommerceGetOrderStatus200Response) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


