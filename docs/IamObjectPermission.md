# IamObjectPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** |  | [optional] 
**Adapter** | Pointer to **string** |  | [optional] 
**ApproveTime** | Pointer to **string** |  | [optional] 
**Approver** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Effect** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Submitter** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamObjectPermission

`func NewIamObjectPermission() *IamObjectPermission`

NewIamObjectPermission instantiates a new IamObjectPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectPermissionWithDefaults

`func NewIamObjectPermissionWithDefaults() *IamObjectPermission`

NewIamObjectPermissionWithDefaults instantiates a new IamObjectPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *IamObjectPermission) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IamObjectPermission) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IamObjectPermission) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *IamObjectPermission) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAdapter

`func (o *IamObjectPermission) GetAdapter() string`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *IamObjectPermission) GetAdapterOk() (*string, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *IamObjectPermission) SetAdapter(v string)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *IamObjectPermission) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetApproveTime

`func (o *IamObjectPermission) GetApproveTime() string`

GetApproveTime returns the ApproveTime field if non-nil, zero value otherwise.

### GetApproveTimeOk

`func (o *IamObjectPermission) GetApproveTimeOk() (*string, bool)`

GetApproveTimeOk returns a tuple with the ApproveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveTime

`func (o *IamObjectPermission) SetApproveTime(v string)`

SetApproveTime sets ApproveTime field to given value.

### HasApproveTime

`func (o *IamObjectPermission) HasApproveTime() bool`

HasApproveTime returns a boolean if a field has been set.

### GetApprover

`func (o *IamObjectPermission) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *IamObjectPermission) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *IamObjectPermission) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *IamObjectPermission) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectPermission) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectPermission) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectPermission) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectPermission) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *IamObjectPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamObjectPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamObjectPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamObjectPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectPermission) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectPermission) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectPermission) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectPermission) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomains

`func (o *IamObjectPermission) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IamObjectPermission) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IamObjectPermission) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IamObjectPermission) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEffect

`func (o *IamObjectPermission) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *IamObjectPermission) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *IamObjectPermission) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *IamObjectPermission) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetGroups

`func (o *IamObjectPermission) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IamObjectPermission) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IamObjectPermission) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IamObjectPermission) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamObjectPermission) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamObjectPermission) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamObjectPermission) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamObjectPermission) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetModel

`func (o *IamObjectPermission) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IamObjectPermission) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IamObjectPermission) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *IamObjectPermission) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *IamObjectPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectPermission) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectPermission) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectPermission) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectPermission) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectPermission) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectPermission) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResourceType

`func (o *IamObjectPermission) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *IamObjectPermission) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *IamObjectPermission) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *IamObjectPermission) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResources

`func (o *IamObjectPermission) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IamObjectPermission) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IamObjectPermission) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *IamObjectPermission) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRoles

`func (o *IamObjectPermission) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamObjectPermission) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamObjectPermission) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamObjectPermission) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetState

`func (o *IamObjectPermission) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamObjectPermission) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamObjectPermission) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamObjectPermission) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubmitter

`func (o *IamObjectPermission) GetSubmitter() string`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *IamObjectPermission) GetSubmitterOk() (*string, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *IamObjectPermission) SetSubmitter(v string)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *IamObjectPermission) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetUsers

`func (o *IamObjectPermission) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamObjectPermission) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamObjectPermission) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamObjectPermission) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


