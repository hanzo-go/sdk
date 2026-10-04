# BillingSubscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count is the row count beside the rows, which is the shape this address has always answered with. | [optional] 
**Subscriptions** | Pointer to [**[]BillingSubscription**](BillingSubscription.md) |  | [optional] 

## Methods

### NewBillingSubscriptions

`func NewBillingSubscriptions() *BillingSubscriptions`

NewBillingSubscriptions instantiates a new BillingSubscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionsWithDefaults

`func NewBillingSubscriptionsWithDefaults() *BillingSubscriptions`

NewBillingSubscriptionsWithDefaults instantiates a new BillingSubscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BillingSubscriptions) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BillingSubscriptions) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BillingSubscriptions) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BillingSubscriptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSubscriptions

`func (o *BillingSubscriptions) GetSubscriptions() []BillingSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *BillingSubscriptions) GetSubscriptionsOk() (*[]BillingSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *BillingSubscriptions) SetSubscriptions(v []BillingSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *BillingSubscriptions) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


