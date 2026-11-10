# BillingAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | Pointer to [**[]BillingSubscriptionSeat**](BillingSubscriptionSeat.md) | Assigned is every seat assigned: the member, and the seat type (a plan the subscription holds seats of, \&quot;team\&quot; or \&quot;team-premium\&quot;) they hold. It replaces what was assigned; an empty list unassigns every seat. | [optional] 

## Methods

### NewBillingAssignment

`func NewBillingAssignment() *BillingAssignment`

NewBillingAssignment instantiates a new BillingAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAssignmentWithDefaults

`func NewBillingAssignmentWithDefaults() *BillingAssignment`

NewBillingAssignmentWithDefaults instantiates a new BillingAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *BillingAssignment) GetAssigned() []BillingSubscriptionSeat`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *BillingAssignment) GetAssignedOk() (*[]BillingSubscriptionSeat, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *BillingAssignment) SetAssigned(v []BillingSubscriptionSeat)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *BillingAssignment) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


