# NetworkNetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]NetworkNetworkView**](NetworkNetworkView.md) | Networks holds the org&#39;s overlay network, or is empty when the org has no edge-routers on the fabric (no nodes → no network, never a fabricated one). | [optional] 

## Methods

### NewNetworkNetworkList

`func NewNetworkNetworkList() *NetworkNetworkList`

NewNetworkNetworkList instantiates a new NetworkNetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkListWithDefaults

`func NewNetworkNetworkListWithDefaults() *NetworkNetworkList`

NewNetworkNetworkListWithDefaults instantiates a new NetworkNetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *NetworkNetworkList) GetNetworks() []NetworkNetworkView`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NetworkNetworkList) GetNetworksOk() (*[]NetworkNetworkView, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NetworkNetworkList) SetNetworks(v []NetworkNetworkView)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NetworkNetworkList) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


