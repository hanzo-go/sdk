# BillingSignedPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**BillingGatewayAuthorization**](BillingGatewayAuthorization.md) | Authorization is the EIP-3009 TransferWithAuthorization as signed: every field a string, addresses and nonce 0x hex, value in USDC base units. | [optional] 
**Network** | Pointer to **string** | Network is the CAIP-2 network the payment was signed for, one of the terms&#39; accepts, e.g. \&quot;eip155:8453\&quot;. | [optional] 
**Salt** | Pointer to **string** | Salt is the salt the terms issued with the nonce. | [optional] 
**Signature** | Pointer to **string** | Signature is the 65-byte EIP-712 signature, 0x hex. | [optional] 

## Methods

### NewBillingSignedPayment

`func NewBillingSignedPayment() *BillingSignedPayment`

NewBillingSignedPayment instantiates a new BillingSignedPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSignedPaymentWithDefaults

`func NewBillingSignedPaymentWithDefaults() *BillingSignedPayment`

NewBillingSignedPaymentWithDefaults instantiates a new BillingSignedPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *BillingSignedPayment) GetAuthorization() BillingGatewayAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *BillingSignedPayment) GetAuthorizationOk() (*BillingGatewayAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *BillingSignedPayment) SetAuthorization(v BillingGatewayAuthorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *BillingSignedPayment) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetNetwork

`func (o *BillingSignedPayment) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BillingSignedPayment) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BillingSignedPayment) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *BillingSignedPayment) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSalt

`func (o *BillingSignedPayment) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *BillingSignedPayment) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *BillingSignedPayment) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *BillingSignedPayment) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetSignature

`func (o *BillingSignedPayment) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *BillingSignedPayment) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *BillingSignedPayment) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *BillingSignedPayment) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


