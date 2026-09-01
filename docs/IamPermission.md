# IamPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** |  | [optional] 
**Adapter** | Pointer to **string** |  | [optional] 
**ApproveTime** | Pointer to **string** |  | [optional] 
**Approver** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** | Descriptive metadata. | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Effect** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Model** | Pointer to **string** | Authorization model, targets, and decision. AuthzModel carries the v1 &#x60;model&#x60; column (the named authz model); it is not the Go identifier &#x60;Model&#x60; because that name is taken by the embedded orm.Model[Permission] mixin. The HTTP contract is unchanged — json:\&quot;model\&quot;. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** | Identity — the (owner, name) natural key. | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Submitter** | Pointer to **string** | Submission / approval workflow. | [optional] 
**Teams** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Users** | Pointer to **[]string** | Subjects the grant is evaluated for. | [optional] 

## Methods

### NewIamPermission

`func NewIamPermission() *IamPermission`

NewIamPermission instantiates a new IamPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPermissionWithDefaults

`func NewIamPermissionWithDefaults() *IamPermission`

NewIamPermissionWithDefaults instantiates a new IamPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *IamPermission) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IamPermission) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IamPermission) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *IamPermission) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAdapter

`func (o *IamPermission) GetAdapter() string`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *IamPermission) GetAdapterOk() (*string, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *IamPermission) SetAdapter(v string)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *IamPermission) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetApproveTime

`func (o *IamPermission) GetApproveTime() string`

GetApproveTime returns the ApproveTime field if non-nil, zero value otherwise.

### GetApproveTimeOk

`func (o *IamPermission) GetApproveTimeOk() (*string, bool)`

GetApproveTimeOk returns a tuple with the ApproveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveTime

`func (o *IamPermission) SetApproveTime(v string)`

SetApproveTime sets ApproveTime field to given value.

### HasApproveTime

`func (o *IamPermission) HasApproveTime() bool`

HasApproveTime returns a boolean if a field has been set.

### GetApprover

`func (o *IamPermission) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *IamPermission) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *IamPermission) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *IamPermission) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamPermission) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamPermission) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamPermission) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamPermission) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamPermission) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamPermission) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamPermission) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamPermission) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamPermission) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamPermission) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamPermission) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamPermission) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *IamPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamPermission) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamPermission) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamPermission) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamPermission) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomains

`func (o *IamPermission) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IamPermission) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IamPermission) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IamPermission) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEffect

`func (o *IamPermission) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *IamPermission) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *IamPermission) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *IamPermission) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetId

`func (o *IamPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamPermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamPermission) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamPermission) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamPermission) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamPermission) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetModel

`func (o *IamPermission) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IamPermission) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IamPermission) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *IamPermission) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *IamPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamPermission) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamPermission) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamPermission) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamPermission) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamPermission) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamPermission) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResourceType

`func (o *IamPermission) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *IamPermission) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *IamPermission) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *IamPermission) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResources

`func (o *IamPermission) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IamPermission) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IamPermission) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *IamPermission) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRoles

`func (o *IamPermission) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamPermission) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamPermission) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamPermission) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetState

`func (o *IamPermission) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamPermission) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamPermission) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamPermission) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubmitter

`func (o *IamPermission) GetSubmitter() string`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *IamPermission) GetSubmitterOk() (*string, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *IamPermission) SetSubmitter(v string)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *IamPermission) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetTeams

`func (o *IamPermission) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *IamPermission) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *IamPermission) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *IamPermission) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamPermission) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamPermission) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamPermission) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamPermission) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsers

`func (o *IamPermission) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamPermission) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamPermission) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamPermission) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


