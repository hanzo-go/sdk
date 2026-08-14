# VerificationDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the verification to decide, from the path. | [optional] 
**Status** | Pointer to **string** | Status is the reviewer&#39;s decision: \&quot;reviewer_confirmed\&quot; (a pass) or \&quot;manual_review\&quot; (withheld for review) — never a provider status. | [optional] 

## Methods

### NewVerificationDecision

`func NewVerificationDecision() *VerificationDecision`

NewVerificationDecision instantiates a new VerificationDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationDecisionWithDefaults

`func NewVerificationDecisionWithDefaults() *VerificationDecision`

NewVerificationDecisionWithDefaults instantiates a new VerificationDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerificationDecision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerificationDecision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerificationDecision) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerificationDecision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *VerificationDecision) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerificationDecision) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerificationDecision) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerificationDecision) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


