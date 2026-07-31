# CloudDecisionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email identifies the founder on the formation. | [optional] 
**Status** | Pointer to **string** | Status is the decision: reviewer_confirmed or failed. Nothing else is accepted. | [optional] 

## Methods

### NewCloudDecisionIn

`func NewCloudDecisionIn() *CloudDecisionIn`

NewCloudDecisionIn instantiates a new CloudDecisionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDecisionInWithDefaults

`func NewCloudDecisionInWithDefaults() *CloudDecisionIn`

NewCloudDecisionInWithDefaults instantiates a new CloudDecisionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CloudDecisionIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudDecisionIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudDecisionIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudDecisionIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *CloudDecisionIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudDecisionIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudDecisionIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudDecisionIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


