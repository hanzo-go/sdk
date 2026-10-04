# ToolKit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Excess** | Pointer to **[]string** | Excess are servers left out, by id, because a run carries only so many. | [optional] 
**Muted** | Pointer to **[]string** | Muted are the names the run&#39;s person took out of their own runs — a skill, a server&#39;s tool, or a server by its id. None of them is above. | [optional] 
**Omitted** | Pointer to **[]string** | Omitted are activated skills left out because a run carries only so many. | [optional] 
**Servers** | Pointer to [**[]ToolRemote**](ToolRemote.md) | Servers are the MCP servers an admin registered and activated tools of, at most KitServers of them. | [optional] 
**Skills** | Pointer to [**[]ToolSkillDoc**](ToolSkillDoc.md) | Skills are the documents of the skills an admin activated, by name. | [optional] 

## Methods

### NewToolKit

`func NewToolKit() *ToolKit`

NewToolKit instantiates a new ToolKit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolKitWithDefaults

`func NewToolKitWithDefaults() *ToolKit`

NewToolKitWithDefaults instantiates a new ToolKit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcess

`func (o *ToolKit) GetExcess() []string`

GetExcess returns the Excess field if non-nil, zero value otherwise.

### GetExcessOk

`func (o *ToolKit) GetExcessOk() (*[]string, bool)`

GetExcessOk returns a tuple with the Excess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcess

`func (o *ToolKit) SetExcess(v []string)`

SetExcess sets Excess field to given value.

### HasExcess

`func (o *ToolKit) HasExcess() bool`

HasExcess returns a boolean if a field has been set.

### GetMuted

`func (o *ToolKit) GetMuted() []string`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ToolKit) GetMutedOk() (*[]string, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ToolKit) SetMuted(v []string)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ToolKit) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetOmitted

`func (o *ToolKit) GetOmitted() []string`

GetOmitted returns the Omitted field if non-nil, zero value otherwise.

### GetOmittedOk

`func (o *ToolKit) GetOmittedOk() (*[]string, bool)`

GetOmittedOk returns a tuple with the Omitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitted

`func (o *ToolKit) SetOmitted(v []string)`

SetOmitted sets Omitted field to given value.

### HasOmitted

`func (o *ToolKit) HasOmitted() bool`

HasOmitted returns a boolean if a field has been set.

### GetServers

`func (o *ToolKit) GetServers() []ToolRemote`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ToolKit) GetServersOk() (*[]ToolRemote, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ToolKit) SetServers(v []ToolRemote)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ToolKit) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSkills

`func (o *ToolKit) GetSkills() []ToolSkillDoc`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ToolKit) GetSkillsOk() (*[]ToolSkillDoc, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ToolKit) SetSkills(v []ToolSkillDoc)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ToolKit) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


