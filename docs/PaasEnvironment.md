# PaasEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Iid** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to [**PaasEnvironmentVersion**](PaasEnvironmentVersion.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaasEnvironment

`func NewPaasEnvironment() *PaasEnvironment`

NewPaasEnvironment instantiates a new PaasEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasEnvironmentWithDefaults

`func NewPaasEnvironmentWithDefaults() *PaasEnvironment`

NewPaasEnvironmentWithDefaults instantiates a new PaasEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaasEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaasEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaasEnvironment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaasEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIid

`func (o *PaasEnvironment) GetIid() string`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *PaasEnvironment) GetIidOk() (*string, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *PaasEnvironment) SetIid(v string)`

SetIid sets Iid field to given value.

### HasIid

`func (o *PaasEnvironment) HasIid() bool`

HasIid returns a boolean if a field has been set.

### GetOrgId

`func (o *PaasEnvironment) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PaasEnvironment) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PaasEnvironment) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PaasEnvironment) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *PaasEnvironment) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PaasEnvironment) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PaasEnvironment) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PaasEnvironment) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *PaasEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaasEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaasEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaasEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *PaasEnvironment) GetVersion() PaasEnvironmentVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PaasEnvironment) GetVersionOk() (*PaasEnvironmentVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PaasEnvironment) SetVersion(v PaasEnvironmentVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PaasEnvironment) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetReadOnly

`func (o *PaasEnvironment) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PaasEnvironment) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PaasEnvironment) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PaasEnvironment) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSuspended

`func (o *PaasEnvironment) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *PaasEnvironment) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *PaasEnvironment) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *PaasEnvironment) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaasEnvironment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaasEnvironment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaasEnvironment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaasEnvironment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


