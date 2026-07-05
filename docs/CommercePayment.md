# CommercePayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**OrderId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**AmountRefunded** | Pointer to **int32** |  | [optional] 
**Fee** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**CommercePaymentStatus**](CommercePaymentStatus.md) |  | [optional] 
**Captured** | Pointer to **bool** |  | [optional] 
**Live** | Pointer to **bool** |  | [optional] 
**Buyer** | Pointer to [**CommerceBuyer**](CommerceBuyer.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCommercePayment

`func NewCommercePayment() *CommercePayment`

NewCommercePayment instantiates a new CommercePayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommercePaymentWithDefaults

`func NewCommercePaymentWithDefaults() *CommercePayment`

NewCommercePaymentWithDefaults instantiates a new CommercePayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommercePayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommercePayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommercePayment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommercePayment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *CommercePayment) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CommercePayment) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CommercePayment) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CommercePayment) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetUserId

`func (o *CommercePayment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommercePayment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommercePayment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommercePayment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCurrency

`func (o *CommercePayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommercePayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommercePayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommercePayment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *CommercePayment) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CommercePayment) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CommercePayment) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CommercePayment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountRefunded

`func (o *CommercePayment) GetAmountRefunded() int32`

GetAmountRefunded returns the AmountRefunded field if non-nil, zero value otherwise.

### GetAmountRefundedOk

`func (o *CommercePayment) GetAmountRefundedOk() (*int32, bool)`

GetAmountRefundedOk returns a tuple with the AmountRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefunded

`func (o *CommercePayment) SetAmountRefunded(v int32)`

SetAmountRefunded sets AmountRefunded field to given value.

### HasAmountRefunded

`func (o *CommercePayment) HasAmountRefunded() bool`

HasAmountRefunded returns a boolean if a field has been set.

### GetFee

`func (o *CommercePayment) GetFee() int32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CommercePayment) GetFeeOk() (*int32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CommercePayment) SetFee(v int32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CommercePayment) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetStatus

`func (o *CommercePayment) GetStatus() CommercePaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommercePayment) GetStatusOk() (*CommercePaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommercePayment) SetStatus(v CommercePaymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommercePayment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCaptured

`func (o *CommercePayment) GetCaptured() bool`

GetCaptured returns the Captured field if non-nil, zero value otherwise.

### GetCapturedOk

`func (o *CommercePayment) GetCapturedOk() (*bool, bool)`

GetCapturedOk returns a tuple with the Captured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptured

`func (o *CommercePayment) SetCaptured(v bool)`

SetCaptured sets Captured field to given value.

### HasCaptured

`func (o *CommercePayment) HasCaptured() bool`

HasCaptured returns a boolean if a field has been set.

### GetLive

`func (o *CommercePayment) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *CommercePayment) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *CommercePayment) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *CommercePayment) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetBuyer

`func (o *CommercePayment) GetBuyer() CommerceBuyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *CommercePayment) GetBuyerOk() (*CommerceBuyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *CommercePayment) SetBuyer(v CommerceBuyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *CommercePayment) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetMetadata

`func (o *CommercePayment) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommercePayment) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommercePayment) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommercePayment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommercePayment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommercePayment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommercePayment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommercePayment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommercePayment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommercePayment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommercePayment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommercePayment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


