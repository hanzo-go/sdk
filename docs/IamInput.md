# IamInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | Name addresses the team on update and names it on create; every other field is content and binds from the BODY, never the URL. | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Teams** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamInput

`func NewIamInput() *IamInput`

NewIamInput instantiates a new IamInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamInputWithDefaults

`func NewIamInputWithDefaults() *IamInput`

NewIamInputWithDefaults instantiates a new IamInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *IamInput) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamInput) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamInput) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamInput) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *IamInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamInput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamInput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamInput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamInput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomains

`func (o *IamInput) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IamInput) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IamInput) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IamInput) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamInput) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamInput) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamInput) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamInput) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *IamInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamInput) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamInput) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamInput) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamInput) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRoles

`func (o *IamInput) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamInput) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamInput) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamInput) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTeams

`func (o *IamInput) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *IamInput) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *IamInput) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *IamInput) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetUsers

`func (o *IamInput) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamInput) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamInput) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamInput) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


