# AccessChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **[]string** | Affected lists the usernames that were updated. | [optional] 
**Failed** | Pointer to **[]string** | Failed lists the usernames that were NOT updated. Non-empty means the org is in a mixed state and the action should be retried. | [optional] 
**Org** | Pointer to **string** | Org is the tenant acted on. | [optional] 
**Suspended** | Pointer to **bool** | Suspended is the state applied: true for suspend, false for reactivate. | [optional] 

## Methods

### NewAccessChange

`func NewAccessChange() *AccessChange`

NewAccessChange instantiates a new AccessChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessChangeWithDefaults

`func NewAccessChangeWithDefaults() *AccessChange`

NewAccessChangeWithDefaults instantiates a new AccessChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AccessChange) GetAffected() []string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AccessChange) GetAffectedOk() (*[]string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AccessChange) SetAffected(v []string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AccessChange) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetFailed

`func (o *AccessChange) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *AccessChange) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *AccessChange) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *AccessChange) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetOrg

`func (o *AccessChange) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AccessChange) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AccessChange) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AccessChange) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSuspended

`func (o *AccessChange) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *AccessChange) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *AccessChange) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *AccessChange) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


