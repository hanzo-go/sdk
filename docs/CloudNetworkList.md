# CloudNetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]CloudNetworkView**](CloudNetworkView.md) | Networks holds the org&#39;s overlay network, or is empty when the org has no edge-routers on the fabric (no nodes → no network, never a fabricated one). | [optional] 

## Methods

### NewCloudNetworkList

`func NewCloudNetworkList() *CloudNetworkList`

NewCloudNetworkList instantiates a new CloudNetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkListWithDefaults

`func NewCloudNetworkListWithDefaults() *CloudNetworkList`

NewCloudNetworkListWithDefaults instantiates a new CloudNetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *CloudNetworkList) GetNetworks() []CloudNetworkView`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CloudNetworkList) GetNetworksOk() (*[]CloudNetworkView, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CloudNetworkList) SetNetworks(v []CloudNetworkView)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *CloudNetworkList) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


