# PaasContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Iid** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**EnvId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PaasContainerStatus**](PaasContainerStatus.md) |  | [optional] 
**Template** | Pointer to **string** | Template slug (e.g. nodejs, python, go, rust) | [optional] 
**Repo** | Pointer to [**PaasContainerRepo**](PaasContainerRepo.md) |  | [optional] 
**Networking** | Pointer to [**PaasContainerNetworking**](PaasContainerNetworking.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaasContainer

`func NewPaasContainer() *PaasContainer`

NewPaasContainer instantiates a new PaasContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasContainerWithDefaults

`func NewPaasContainerWithDefaults() *PaasContainer`

NewPaasContainerWithDefaults instantiates a new PaasContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaasContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaasContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaasContainer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaasContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIid

`func (o *PaasContainer) GetIid() string`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *PaasContainer) GetIidOk() (*string, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *PaasContainer) SetIid(v string)`

SetIid sets Iid field to given value.

### HasIid

`func (o *PaasContainer) HasIid() bool`

HasIid returns a boolean if a field has been set.

### GetOrgId

`func (o *PaasContainer) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PaasContainer) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PaasContainer) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PaasContainer) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *PaasContainer) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PaasContainer) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PaasContainer) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PaasContainer) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetEnvId

`func (o *PaasContainer) GetEnvId() string`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *PaasContainer) GetEnvIdOk() (*string, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *PaasContainer) SetEnvId(v string)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *PaasContainer) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### GetName

`func (o *PaasContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaasContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaasContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaasContainer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PaasContainer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaasContainer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaasContainer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaasContainer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PaasContainer) GetStatus() PaasContainerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaasContainer) GetStatusOk() (*PaasContainerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaasContainer) SetStatus(v PaasContainerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaasContainer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplate

`func (o *PaasContainer) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PaasContainer) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PaasContainer) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PaasContainer) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRepo

`func (o *PaasContainer) GetRepo() PaasContainerRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PaasContainer) GetRepoOk() (*PaasContainerRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PaasContainer) SetRepo(v PaasContainerRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PaasContainer) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetNetworking

`func (o *PaasContainer) GetNetworking() PaasContainerNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *PaasContainer) GetNetworkingOk() (*PaasContainerNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *PaasContainer) SetNetworking(v PaasContainerNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *PaasContainer) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaasContainer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaasContainer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaasContainer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaasContainer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


