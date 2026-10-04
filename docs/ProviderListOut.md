# ProviderListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]ProviderProviderView**](ProviderProviderView.md) | Providers is the whole catalog. Never null; [] when nothing is registered. | [optional] 

## Methods

### NewProviderListOut

`func NewProviderListOut() *ProviderListOut`

NewProviderListOut instantiates a new ProviderListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderListOutWithDefaults

`func NewProviderListOutWithDefaults() *ProviderListOut`

NewProviderListOutWithDefaults instantiates a new ProviderListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *ProviderListOut) GetProviders() []ProviderProviderView`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ProviderListOut) GetProvidersOk() (*[]ProviderProviderView, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ProviderListOut) SetProviders(v []ProviderProviderView)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ProviderListOut) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


