# NetworkNetworkView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the org-derived id of the overlay network — the key GET /v1/network/{id} addresses. | [optional] 
**Name** | Pointer to **string** | Name is the org the overlay belongs to. | [optional] 
**Nodes** | Pointer to **int64** | Nodes is how many edge-routers the org has on the fabric. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;connected\&quot; once at least one of the org&#39;s edge-routers is online, else \&quot;provisioning\&quot; (routers exist but none has dialed home). | [optional] 

## Methods

### NewNetworkNetworkView

`func NewNetworkNetworkView() *NetworkNetworkView`

NewNetworkNetworkView instantiates a new NetworkNetworkView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkViewWithDefaults

`func NewNetworkNetworkViewWithDefaults() *NetworkNetworkView`

NewNetworkNetworkViewWithDefaults instantiates a new NetworkNetworkView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkNetworkView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkNetworkView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkNetworkView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkNetworkView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkNetworkView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkNetworkView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkNetworkView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkNetworkView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *NetworkNetworkView) GetNodes() int64`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NetworkNetworkView) GetNodesOk() (*int64, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NetworkNetworkView) SetNodes(v int64)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NetworkNetworkView) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkNetworkView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkNetworkView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkNetworkView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkNetworkView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


