# NetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]NetworkView**](NetworkView.md) | Networks holds the org&#39;s overlay network, or is empty when the org has no edge-routers on the fabric (no nodes → no network, never a fabricated one). | [optional] 

## Methods

### NewNetworkList

`func NewNetworkList() *NetworkList`

NewNetworkList instantiates a new NetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkListWithDefaults

`func NewNetworkListWithDefaults() *NetworkList`

NewNetworkListWithDefaults instantiates a new NetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *NetworkList) GetNetworks() []NetworkView`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NetworkList) GetNetworksOk() (*[]NetworkView, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NetworkList) SetNetworks(v []NetworkView)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NetworkList) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


