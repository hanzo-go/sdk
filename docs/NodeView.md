# NodeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caps** | Pointer to **[]string** | Caps is the capability list the node reported. It is a self-report, useful to SHOW and never load-bearing: what a node may actually be asked to do is decided at the socket by the deployment&#39;s allowlist. | [optional] 
**Commands** | Pointer to **[]string** | Commands is the command list the node reported. Same standing as Caps: a self-report, checked again at the socket before anything runs. | [optional] 
**ConnectedAt** | Pointer to **string** | ConnectedAt is when this node&#39;s socket was established, RFC3339 UTC. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is the human name the node reported for itself. | [optional] 
**Id** | Pointer to **string** | ID is the node&#39;s own identifier within the org — the value POST /v1/bot/nodes/{id}/invoke addresses it by. | [optional] 
**Platform** | Pointer to **string** | Platform is the operating system and architecture the node reported. | [optional] 
**Version** | Pointer to **string** | Version is the node agent&#39;s own version string. | [optional] 

## Methods

### NewNodeView

`func NewNodeView() *NodeView`

NewNodeView instantiates a new NodeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeViewWithDefaults

`func NewNodeViewWithDefaults() *NodeView`

NewNodeViewWithDefaults instantiates a new NodeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaps

`func (o *NodeView) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *NodeView) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *NodeView) SetCaps(v []string)`

SetCaps sets Caps field to given value.

### HasCaps

`func (o *NodeView) HasCaps() bool`

HasCaps returns a boolean if a field has been set.

### GetCommands

`func (o *NodeView) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *NodeView) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *NodeView) SetCommands(v []string)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *NodeView) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetConnectedAt

`func (o *NodeView) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *NodeView) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *NodeView) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *NodeView) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *NodeView) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NodeView) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NodeView) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NodeView) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *NodeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *NodeView) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NodeView) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NodeView) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *NodeView) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *NodeView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


