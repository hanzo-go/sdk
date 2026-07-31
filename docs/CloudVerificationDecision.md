# CloudVerificationDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the verification to decide, from the path. | [optional] 
**Status** | Pointer to **string** | Status is the reviewer&#39;s decision: \&quot;reviewer_confirmed\&quot; (a pass) or \&quot;manual_review\&quot; (withheld for review) — never a provider status. | [optional] 

## Methods

### NewCloudVerificationDecision

`func NewCloudVerificationDecision() *CloudVerificationDecision`

NewCloudVerificationDecision instantiates a new CloudVerificationDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVerificationDecisionWithDefaults

`func NewCloudVerificationDecisionWithDefaults() *CloudVerificationDecision`

NewCloudVerificationDecisionWithDefaults instantiates a new CloudVerificationDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudVerificationDecision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudVerificationDecision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudVerificationDecision) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudVerificationDecision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudVerificationDecision) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudVerificationDecision) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudVerificationDecision) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudVerificationDecision) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


