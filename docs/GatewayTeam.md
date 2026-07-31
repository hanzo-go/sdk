# GatewayTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | Pointer to **string** |  | [optional] 
**TeamAlias** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Admins** | Pointer to **[]string** |  | [optional] 
**Members** | Pointer to **[]string** |  | [optional] 
**MaxBudget** | Pointer to **float32** |  | [optional] 
**Spend** | Pointer to **float32** |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGatewayTeam

`func NewGatewayTeam() *GatewayTeam`

NewGatewayTeam instantiates a new GatewayTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTeamWithDefaults

`func NewGatewayTeamWithDefaults() *GatewayTeam`

NewGatewayTeamWithDefaults instantiates a new GatewayTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *GatewayTeam) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GatewayTeam) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GatewayTeam) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *GatewayTeam) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamAlias

`func (o *GatewayTeam) GetTeamAlias() string`

GetTeamAlias returns the TeamAlias field if non-nil, zero value otherwise.

### GetTeamAliasOk

`func (o *GatewayTeam) GetTeamAliasOk() (*string, bool)`

GetTeamAliasOk returns a tuple with the TeamAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamAlias

`func (o *GatewayTeam) SetTeamAlias(v string)`

SetTeamAlias sets TeamAlias field to given value.

### HasTeamAlias

`func (o *GatewayTeam) HasTeamAlias() bool`

HasTeamAlias returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GatewayTeam) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GatewayTeam) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GatewayTeam) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GatewayTeam) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAdmins

`func (o *GatewayTeam) GetAdmins() []string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *GatewayTeam) GetAdminsOk() (*[]string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *GatewayTeam) SetAdmins(v []string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *GatewayTeam) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetMembers

`func (o *GatewayTeam) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GatewayTeam) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GatewayTeam) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GatewayTeam) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMaxBudget

`func (o *GatewayTeam) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *GatewayTeam) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *GatewayTeam) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.

### HasMaxBudget

`func (o *GatewayTeam) HasMaxBudget() bool`

HasMaxBudget returns a boolean if a field has been set.

### GetSpend

`func (o *GatewayTeam) GetSpend() float32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *GatewayTeam) GetSpendOk() (*float32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *GatewayTeam) SetSpend(v float32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *GatewayTeam) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetModels

`func (o *GatewayTeam) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *GatewayTeam) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *GatewayTeam) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *GatewayTeam) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetMetadata

`func (o *GatewayTeam) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GatewayTeam) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GatewayTeam) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GatewayTeam) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


