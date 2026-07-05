# CommercePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Payment processor token (e.g., Stripe token) | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** | Amount in cents (optional, defaults to order total) | [optional] 

## Methods

### NewCommercePaymentRequest

`func NewCommercePaymentRequest() *CommercePaymentRequest`

NewCommercePaymentRequest instantiates a new CommercePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommercePaymentRequestWithDefaults

`func NewCommercePaymentRequestWithDefaults() *CommercePaymentRequest`

NewCommercePaymentRequestWithDefaults instantiates a new CommercePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CommercePaymentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommercePaymentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommercePaymentRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommercePaymentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetToken

`func (o *CommercePaymentRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CommercePaymentRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CommercePaymentRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CommercePaymentRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *CommercePaymentRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CommercePaymentRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CommercePaymentRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *CommercePaymentRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetAmount

`func (o *CommercePaymentRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CommercePaymentRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CommercePaymentRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CommercePaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


