# BillingSubscriptionAct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtPeriodEnd** | Pointer to **bool** | AtPeriodEnd cancels at the end of the paid period rather than at once. It defaults TRUE, because a customer who cancels has already paid for the period they are in. | [optional] 
**Id** | Pointer to **string** | ID is the subscription, from the path. | [optional] 

## Methods

### NewBillingSubscriptionAct

`func NewBillingSubscriptionAct() *BillingSubscriptionAct`

NewBillingSubscriptionAct instantiates a new BillingSubscriptionAct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionActWithDefaults

`func NewBillingSubscriptionActWithDefaults() *BillingSubscriptionAct`

NewBillingSubscriptionActWithDefaults instantiates a new BillingSubscriptionAct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtPeriodEnd

`func (o *BillingSubscriptionAct) GetAtPeriodEnd() bool`

GetAtPeriodEnd returns the AtPeriodEnd field if non-nil, zero value otherwise.

### GetAtPeriodEndOk

`func (o *BillingSubscriptionAct) GetAtPeriodEndOk() (*bool, bool)`

GetAtPeriodEndOk returns a tuple with the AtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtPeriodEnd

`func (o *BillingSubscriptionAct) SetAtPeriodEnd(v bool)`

SetAtPeriodEnd sets AtPeriodEnd field to given value.

### HasAtPeriodEnd

`func (o *BillingSubscriptionAct) HasAtPeriodEnd() bool`

HasAtPeriodEnd returns a boolean if a field has been set.

### GetId

`func (o *BillingSubscriptionAct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSubscriptionAct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSubscriptionAct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingSubscriptionAct) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


