# BillingSubscriptionSeat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to **string** | Member is the IAM user who holds the seat, as the org&#39;s members list names them. | [optional] 
**Plan** | Pointer to **string** | Plan is the seat type held: \&quot;team\&quot; (standard) or \&quot;team-premium\&quot;. | [optional] 

## Methods

### NewBillingSubscriptionSeat

`func NewBillingSubscriptionSeat() *BillingSubscriptionSeat`

NewBillingSubscriptionSeat instantiates a new BillingSubscriptionSeat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionSeatWithDefaults

`func NewBillingSubscriptionSeatWithDefaults() *BillingSubscriptionSeat`

NewBillingSubscriptionSeatWithDefaults instantiates a new BillingSubscriptionSeat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *BillingSubscriptionSeat) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *BillingSubscriptionSeat) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *BillingSubscriptionSeat) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *BillingSubscriptionSeat) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetPlan

`func (o *BillingSubscriptionSeat) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *BillingSubscriptionSeat) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *BillingSubscriptionSeat) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *BillingSubscriptionSeat) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


