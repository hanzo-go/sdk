# CloudAccessChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **[]string** | Affected lists the usernames that were updated. | [optional] 
**Failed** | Pointer to **[]string** | Failed lists the usernames that were NOT updated. Non-empty means the org is in a mixed state and the action should be retried. | [optional] 
**Org** | Pointer to **string** | Org is the tenant acted on. | [optional] 
**Suspended** | Pointer to **bool** | Suspended is the state applied: true for suspend, false for reactivate. | [optional] 

## Methods

### NewCloudAccessChange

`func NewCloudAccessChange() *CloudAccessChange`

NewCloudAccessChange instantiates a new CloudAccessChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccessChangeWithDefaults

`func NewCloudAccessChangeWithDefaults() *CloudAccessChange`

NewCloudAccessChangeWithDefaults instantiates a new CloudAccessChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *CloudAccessChange) GetAffected() []string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *CloudAccessChange) GetAffectedOk() (*[]string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *CloudAccessChange) SetAffected(v []string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *CloudAccessChange) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetFailed

`func (o *CloudAccessChange) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *CloudAccessChange) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *CloudAccessChange) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *CloudAccessChange) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetOrg

`func (o *CloudAccessChange) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudAccessChange) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudAccessChange) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudAccessChange) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSuspended

`func (o *CloudAccessChange) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *CloudAccessChange) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *CloudAccessChange) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *CloudAccessChange) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


