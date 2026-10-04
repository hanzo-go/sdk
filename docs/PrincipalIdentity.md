# PrincipalIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason says which founders, and what is outstanding. | [optional] 
**Status** | Pointer to **string** | Status is verified (every founder of its formation passed identity verification through the licensed provider, or a platform reviewer confirmed them), pending, failed, or none (no formation, or one naming no founder). | [optional] 

## Methods

### NewPrincipalIdentity

`func NewPrincipalIdentity() *PrincipalIdentity`

NewPrincipalIdentity instantiates a new PrincipalIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalIdentityWithDefaults

`func NewPrincipalIdentityWithDefaults() *PrincipalIdentity`

NewPrincipalIdentityWithDefaults instantiates a new PrincipalIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *PrincipalIdentity) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PrincipalIdentity) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PrincipalIdentity) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PrincipalIdentity) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalIdentity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalIdentity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalIdentity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrincipalIdentity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


