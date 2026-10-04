# Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clause** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApproval

`func NewApproval() *Approval`

NewApproval instantiates a new Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWithDefaults

`func NewApprovalWithDefaults() *Approval`

NewApprovalWithDefaults instantiates a new Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClause

`func (o *Approval) GetClause() string`

GetClause returns the Clause field if non-nil, zero value otherwise.

### GetClauseOk

`func (o *Approval) GetClauseOk() (*string, bool)`

GetClauseOk returns a tuple with the Clause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClause

`func (o *Approval) SetClause(v string)`

SetClause sets Clause field to given value.

### HasClause

`func (o *Approval) HasClause() bool`

HasClause returns a boolean if a field has been set.

### GetReason

`func (o *Approval) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Approval) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Approval) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Approval) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *Approval) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Approval) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Approval) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Approval) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


