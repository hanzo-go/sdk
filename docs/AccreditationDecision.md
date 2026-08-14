# AccreditationDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the accreditation record to decide, from the path. | [optional] 
**Status** | Pointer to **string** | Status is the decision being recorded: reviewer_confirmed, provider_verified, rejected, or expired. | [optional] 

## Methods

### NewAccreditationDecision

`func NewAccreditationDecision() *AccreditationDecision`

NewAccreditationDecision instantiates a new AccreditationDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccreditationDecisionWithDefaults

`func NewAccreditationDecisionWithDefaults() *AccreditationDecision`

NewAccreditationDecisionWithDefaults instantiates a new AccreditationDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccreditationDecision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccreditationDecision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccreditationDecision) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccreditationDecision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AccreditationDecision) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccreditationDecision) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccreditationDecision) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccreditationDecision) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


