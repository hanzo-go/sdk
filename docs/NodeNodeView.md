# NodeNodeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caps** | Pointer to **[]string** | Caps is the capability list the node reported. It is a self-report, useful to SHOW and never load-bearing: what a node may actually be asked to do is decided at the socket by the deployment&#39;s allowlist. | [optional] 
**Commands** | Pointer to **[]string** | Commands is the command list the node reported. Same standing as Caps: a self-report, checked again at the socket before anything runs. | [optional] 
**ConnectedAt** | Pointer to **string** | ConnectedAt is when this node&#39;s socket was established, RFC3339 UTC. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is the human name the node reported for itself. | [optional] 
**Id** | Pointer to **string** | ID is the node&#39;s own identifier within the org — the value POST /v1/node/{id}/invoke addresses it by. | [optional] 
**Platform** | Pointer to **string** | Platform is the operating system and architecture the node reported. | [optional] 
**Version** | Pointer to **string** | Version is the node agent&#39;s own version string. | [optional] 

## Methods

### NewNodeNodeView

`func NewNodeNodeView() *NodeNodeView`

NewNodeNodeView instantiates a new NodeNodeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeNodeViewWithDefaults

`func NewNodeNodeViewWithDefaults() *NodeNodeView`

NewNodeNodeViewWithDefaults instantiates a new NodeNodeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaps

`func (o *NodeNodeView) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *NodeNodeView) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *NodeNodeView) SetCaps(v []string)`

SetCaps sets Caps field to given value.

### HasCaps

`func (o *NodeNodeView) HasCaps() bool`

HasCaps returns a boolean if a field has been set.

### GetCommands

`func (o *NodeNodeView) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *NodeNodeView) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *NodeNodeView) SetCommands(v []string)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *NodeNodeView) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetConnectedAt

`func (o *NodeNodeView) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *NodeNodeView) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *NodeNodeView) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *NodeNodeView) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *NodeNodeView) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NodeNodeView) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NodeNodeView) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NodeNodeView) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *NodeNodeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeNodeView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeNodeView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeNodeView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *NodeNodeView) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NodeNodeView) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NodeNodeView) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *NodeNodeView) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *NodeNodeView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeNodeView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeNodeView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeNodeView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


