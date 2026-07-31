# GatewayCreateTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamAlias** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Admins** | Pointer to **[]string** |  | [optional] 
**Members** | Pointer to **[]string** |  | [optional] 
**MaxBudget** | Pointer to **float32** |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGatewayCreateTeamRequest

`func NewGatewayCreateTeamRequest() *GatewayCreateTeamRequest`

NewGatewayCreateTeamRequest instantiates a new GatewayCreateTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateTeamRequestWithDefaults

`func NewGatewayCreateTeamRequestWithDefaults() *GatewayCreateTeamRequest`

NewGatewayCreateTeamRequestWithDefaults instantiates a new GatewayCreateTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamAlias

`func (o *GatewayCreateTeamRequest) GetTeamAlias() string`

GetTeamAlias returns the TeamAlias field if non-nil, zero value otherwise.

### GetTeamAliasOk

`func (o *GatewayCreateTeamRequest) GetTeamAliasOk() (*string, bool)`

GetTeamAliasOk returns a tuple with the TeamAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamAlias

`func (o *GatewayCreateTeamRequest) SetTeamAlias(v string)`

SetTeamAlias sets TeamAlias field to given value.

### HasTeamAlias

`func (o *GatewayCreateTeamRequest) HasTeamAlias() bool`

HasTeamAlias returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GatewayCreateTeamRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GatewayCreateTeamRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GatewayCreateTeamRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GatewayCreateTeamRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAdmins

`func (o *GatewayCreateTeamRequest) GetAdmins() []string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *GatewayCreateTeamRequest) GetAdminsOk() (*[]string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *GatewayCreateTeamRequest) SetAdmins(v []string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *GatewayCreateTeamRequest) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetMembers

`func (o *GatewayCreateTeamRequest) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GatewayCreateTeamRequest) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GatewayCreateTeamRequest) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GatewayCreateTeamRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMaxBudget

`func (o *GatewayCreateTeamRequest) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *GatewayCreateTeamRequest) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *GatewayCreateTeamRequest) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.

### HasMaxBudget

`func (o *GatewayCreateTeamRequest) HasMaxBudget() bool`

HasMaxBudget returns a boolean if a field has been set.

### GetModels

`func (o *GatewayCreateTeamRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *GatewayCreateTeamRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *GatewayCreateTeamRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *GatewayCreateTeamRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


