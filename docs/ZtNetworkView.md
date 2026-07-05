# ZtNetworkView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | \&quot;connected\&quot; when at least one router is online, else \&quot;provisioning\&quot; | [optional] 
**Nodes** | Pointer to **int32** | Router count | [optional] 

## Methods

### NewZtNetworkView

`func NewZtNetworkView() *ZtNetworkView`

NewZtNetworkView instantiates a new ZtNetworkView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZtNetworkViewWithDefaults

`func NewZtNetworkViewWithDefaults() *ZtNetworkView`

NewZtNetworkViewWithDefaults instantiates a new ZtNetworkView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZtNetworkView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZtNetworkView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZtNetworkView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ZtNetworkView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ZtNetworkView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZtNetworkView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZtNetworkView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZtNetworkView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ZtNetworkView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZtNetworkView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZtNetworkView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ZtNetworkView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNodes

`func (o *ZtNetworkView) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ZtNetworkView) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ZtNetworkView) SetNodes(v int32)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ZtNetworkView) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


