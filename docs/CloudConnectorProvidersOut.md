# CloudConnectorProvidersOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]CloudConnectorProviderView**](CloudConnectorProviderView.md) | Providers is every USER-scoped provider, sorted by id. Never null. | [optional] 

## Methods

### NewCloudConnectorProvidersOut

`func NewCloudConnectorProvidersOut() *CloudConnectorProvidersOut`

NewCloudConnectorProvidersOut instantiates a new CloudConnectorProvidersOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectorProvidersOutWithDefaults

`func NewCloudConnectorProvidersOutWithDefaults() *CloudConnectorProvidersOut`

NewCloudConnectorProvidersOutWithDefaults instantiates a new CloudConnectorProvidersOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *CloudConnectorProvidersOut) GetProviders() []CloudConnectorProviderView`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CloudConnectorProvidersOut) GetProvidersOk() (*[]CloudConnectorProviderView, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CloudConnectorProvidersOut) SetProviders(v []CloudConnectorProviderView)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CloudConnectorProvidersOut) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


