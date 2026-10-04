# ProviderConnectionsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]ProviderConnectionView**](ProviderConnectionView.md) | Connectors is the caller&#39;s own set. Never null; [] when they have none. | [optional] 

## Methods

### NewProviderConnectionsOut

`func NewProviderConnectionsOut() *ProviderConnectionsOut`

NewProviderConnectionsOut instantiates a new ProviderConnectionsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderConnectionsOutWithDefaults

`func NewProviderConnectionsOutWithDefaults() *ProviderConnectionsOut`

NewProviderConnectionsOutWithDefaults instantiates a new ProviderConnectionsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *ProviderConnectionsOut) GetConnections() []ProviderConnectionView`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ProviderConnectionsOut) GetConnectionsOk() (*[]ProviderConnectionView, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ProviderConnectionsOut) SetConnections(v []ProviderConnectionView)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *ProviderConnectionsOut) HasConnections() bool`

HasConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


