# AdminAccessResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**Affected** | Pointer to **[]string** |  | [optional] 
**Failed** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdminAccessResult

`func NewAdminAccessResult() *AdminAccessResult`

NewAdminAccessResult instantiates a new AdminAccessResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAccessResultWithDefaults

`func NewAdminAccessResultWithDefaults() *AdminAccessResult`

NewAdminAccessResultWithDefaults instantiates a new AdminAccessResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *AdminAccessResult) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AdminAccessResult) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AdminAccessResult) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AdminAccessResult) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSuspended

`func (o *AdminAccessResult) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *AdminAccessResult) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *AdminAccessResult) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *AdminAccessResult) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetAffected

`func (o *AdminAccessResult) GetAffected() []string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdminAccessResult) GetAffectedOk() (*[]string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdminAccessResult) SetAffected(v []string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdminAccessResult) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetFailed

`func (o *AdminAccessResult) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *AdminAccessResult) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *AdminAccessResult) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *AdminAccessResult) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


