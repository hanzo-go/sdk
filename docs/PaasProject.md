# PaasProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Iid** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Team** | Pointer to [**[]PaasProjectTeamInner**](PaasProjectTeamInner.md) |  | [optional] 
**Environments** | Pointer to [**[]PaasEnvironment**](PaasEnvironment.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaasProject

`func NewPaasProject() *PaasProject`

NewPaasProject instantiates a new PaasProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasProjectWithDefaults

`func NewPaasProjectWithDefaults() *PaasProject`

NewPaasProjectWithDefaults instantiates a new PaasProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaasProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaasProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaasProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaasProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIid

`func (o *PaasProject) GetIid() string`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *PaasProject) GetIidOk() (*string, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *PaasProject) SetIid(v string)`

SetIid sets Iid field to given value.

### HasIid

`func (o *PaasProject) HasIid() bool`

HasIid returns a boolean if a field has been set.

### GetOrgId

`func (o *PaasProject) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PaasProject) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PaasProject) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PaasProject) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *PaasProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaasProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaasProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaasProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *PaasProject) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PaasProject) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PaasProject) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PaasProject) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetTeam

`func (o *PaasProject) GetTeam() []PaasProjectTeamInner`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PaasProject) GetTeamOk() (*[]PaasProjectTeamInner, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PaasProject) SetTeam(v []PaasProjectTeamInner)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PaasProject) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetEnvironments

`func (o *PaasProject) GetEnvironments() []PaasEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PaasProject) GetEnvironmentsOk() (*[]PaasEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PaasProject) SetEnvironments(v []PaasEnvironment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PaasProject) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaasProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaasProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaasProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaasProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


