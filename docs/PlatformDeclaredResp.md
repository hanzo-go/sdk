# PlatformDeclaredResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]PlatformDeclared**](PlatformDeclared.md) |  | [optional] 
**CdUnavailable** | Pointer to [**PlatformUnreadable**](PlatformUnreadable.md) |  | [optional] 
**Org** | Pointer to **string** | Org is the directory read — the caller&#39;s own, or another when a SuperAdmin asked to act as it. | [optional] 
**Unreadable** | Pointer to [**[]PlatformUnreadableFile**](PlatformUnreadableFile.md) | Unreadable are the org&#39;s values files that do not parse, so they are not among apps. Absent when every file parsed. | [optional] 

## Methods

### NewPlatformDeclaredResp

`func NewPlatformDeclaredResp() *PlatformDeclaredResp`

NewPlatformDeclaredResp instantiates a new PlatformDeclaredResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDeclaredRespWithDefaults

`func NewPlatformDeclaredRespWithDefaults() *PlatformDeclaredResp`

NewPlatformDeclaredRespWithDefaults instantiates a new PlatformDeclaredResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *PlatformDeclaredResp) GetApps() []PlatformDeclared`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PlatformDeclaredResp) GetAppsOk() (*[]PlatformDeclared, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PlatformDeclaredResp) SetApps(v []PlatformDeclared)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PlatformDeclaredResp) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetCdUnavailable

`func (o *PlatformDeclaredResp) GetCdUnavailable() PlatformUnreadable`

GetCdUnavailable returns the CdUnavailable field if non-nil, zero value otherwise.

### GetCdUnavailableOk

`func (o *PlatformDeclaredResp) GetCdUnavailableOk() (*PlatformUnreadable, bool)`

GetCdUnavailableOk returns a tuple with the CdUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdUnavailable

`func (o *PlatformDeclaredResp) SetCdUnavailable(v PlatformUnreadable)`

SetCdUnavailable sets CdUnavailable field to given value.

### HasCdUnavailable

`func (o *PlatformDeclaredResp) HasCdUnavailable() bool`

HasCdUnavailable returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformDeclaredResp) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformDeclaredResp) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformDeclaredResp) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformDeclaredResp) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetUnreadable

`func (o *PlatformDeclaredResp) GetUnreadable() []PlatformUnreadableFile`

GetUnreadable returns the Unreadable field if non-nil, zero value otherwise.

### GetUnreadableOk

`func (o *PlatformDeclaredResp) GetUnreadableOk() (*[]PlatformUnreadableFile, bool)`

GetUnreadableOk returns a tuple with the Unreadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadable

`func (o *PlatformDeclaredResp) SetUnreadable(v []PlatformUnreadableFile)`

SetUnreadable sets Unreadable field to given value.

### HasUnreadable

`func (o *PlatformDeclaredResp) HasUnreadable() bool`

HasUnreadable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


