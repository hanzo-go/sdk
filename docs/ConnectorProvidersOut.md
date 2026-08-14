# ConnectorProvidersOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]ConnectorProviderView**](ConnectorProviderView.md) | Providers is every USER-scoped provider, sorted by id. Never null. | [optional] 

## Methods

### NewConnectorProvidersOut

`func NewConnectorProvidersOut() *ConnectorProvidersOut`

NewConnectorProvidersOut instantiates a new ConnectorProvidersOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorProvidersOutWithDefaults

`func NewConnectorProvidersOutWithDefaults() *ConnectorProvidersOut`

NewConnectorProvidersOutWithDefaults instantiates a new ConnectorProvidersOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *ConnectorProvidersOut) GetProviders() []ConnectorProviderView`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ConnectorProvidersOut) GetProvidersOk() (*[]ConnectorProviderView, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ConnectorProvidersOut) SetProviders(v []ConnectorProviderView)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ConnectorProvidersOut) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


