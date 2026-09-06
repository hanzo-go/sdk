# AgentBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentName** | Pointer to **string** | AgentName is the cloud Agent (/v1/agent) this machine runs — the agent a message to the bot is actually run against. It is the one field that decides what the bot DOES. | [optional] 
**BotVersion** | Pointer to **string** | BotVersion pins the @hanzo/bot runtime version the machine runs. Empty means the machine took the default in force when it was bound. | [optional] 
**CreatedTime** | Pointer to **string** | CreatedTime is when the binding was first made. | [optional] 
**MachineId** | Pointer to **string** | MachineId is the bound machine as vm addresses it, owner-qualified (\&quot;&lt;org&gt;/&lt;machine&gt;\&quot;). The unqualified half is what this surface&#39;s :id routes take. | [optional] 
**Message** | Pointer to **string** | Message is vm&#39;s human-readable detail on Status (\&quot;machine provisioning; @hanzo/bot runtime not yet confirmed\&quot;) — the reason behind the state, not a second state. | [optional] 
**Name** | Pointer to **string** | Name is the binding&#39;s own key, which is the machine&#39;s id: a machine hosts at most one agent, so the binding is named for it. This is the key a bots list joins bindings onto machines by. | [optional] 
**Org** | Pointer to **string** | Org is the Hanzo tenant the binding belongs to. | [optional] 
**Owner** | Pointer to **string** | Owner is the tenant vm filed the binding under, resolved from the ?owner it was called with — which is the caller&#39;s validated org and never a body field. | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud the bound machine runs on, carried here so a bindings list says where each bot lives without a second read per machine. | [optional] 
**PublicIp** | Pointer to **string** | PublicIp is the bound machine&#39;s public address as vm recorded it on the binding. Empty while the machine has none yet. | [optional] 
**Status** | Pointer to **string** | Status is the binding&#39;s lifecycle in VM&#39;s OWN words — \&quot;Pending\&quot; while the machine provisions and the runtime is unconfirmed, \&quot;running\&quot; once vm has confirmed it. The vocabulary is vm&#39;s and passes through unmapped, which is why its capitalization does not match the machine states beside it, and it is vm&#39;s reconciled reading rather than anything asserted here. | [optional] 
**UpdatedTime** | Pointer to **string** | UpdatedTime is when vm last reconciled it — the age of Status. | [optional] 

## Methods

### NewAgentBinding

`func NewAgentBinding() *AgentBinding`

NewAgentBinding instantiates a new AgentBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentBindingWithDefaults

`func NewAgentBindingWithDefaults() *AgentBinding`

NewAgentBindingWithDefaults instantiates a new AgentBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentName

`func (o *AgentBinding) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *AgentBinding) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *AgentBinding) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *AgentBinding) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetBotVersion

`func (o *AgentBinding) GetBotVersion() string`

GetBotVersion returns the BotVersion field if non-nil, zero value otherwise.

### GetBotVersionOk

`func (o *AgentBinding) GetBotVersionOk() (*string, bool)`

GetBotVersionOk returns a tuple with the BotVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotVersion

`func (o *AgentBinding) SetBotVersion(v string)`

SetBotVersion sets BotVersion field to given value.

### HasBotVersion

`func (o *AgentBinding) HasBotVersion() bool`

HasBotVersion returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AgentBinding) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AgentBinding) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AgentBinding) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AgentBinding) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetMachineId

`func (o *AgentBinding) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *AgentBinding) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *AgentBinding) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *AgentBinding) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetMessage

`func (o *AgentBinding) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AgentBinding) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AgentBinding) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AgentBinding) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *AgentBinding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentBinding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentBinding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentBinding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *AgentBinding) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AgentBinding) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AgentBinding) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AgentBinding) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOwner

`func (o *AgentBinding) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AgentBinding) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AgentBinding) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AgentBinding) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProvider

`func (o *AgentBinding) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AgentBinding) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AgentBinding) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AgentBinding) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *AgentBinding) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *AgentBinding) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *AgentBinding) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *AgentBinding) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetStatus

`func (o *AgentBinding) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentBinding) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentBinding) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentBinding) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AgentBinding) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AgentBinding) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AgentBinding) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AgentBinding) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


