# CloudNodeView

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

### NewCloudNodeView

`func NewCloudNodeView() *CloudNodeView`

NewCloudNodeView instantiates a new CloudNodeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNodeViewWithDefaults

`func NewCloudNodeViewWithDefaults() *CloudNodeView`

NewCloudNodeViewWithDefaults instantiates a new CloudNodeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaps

`func (o *CloudNodeView) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *CloudNodeView) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *CloudNodeView) SetCaps(v []string)`

SetCaps sets Caps field to given value.

### HasCaps

`func (o *CloudNodeView) HasCaps() bool`

HasCaps returns a boolean if a field has been set.

### GetCommands

`func (o *CloudNodeView) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *CloudNodeView) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *CloudNodeView) SetCommands(v []string)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *CloudNodeView) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetConnectedAt

`func (o *CloudNodeView) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *CloudNodeView) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *CloudNodeView) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *CloudNodeView) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudNodeView) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudNodeView) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudNodeView) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudNodeView) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *CloudNodeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudNodeView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudNodeView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudNodeView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudNodeView) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudNodeView) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudNodeView) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudNodeView) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *CloudNodeView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudNodeView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudNodeView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudNodeView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


