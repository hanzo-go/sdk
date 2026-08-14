# ProvidersView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]ProviderCard**](ProviderCard.md) | Providers is every cloud this deployment can link, with what each needs. | [optional] 

## Methods

### NewProvidersView

`func NewProvidersView() *ProvidersView`

NewProvidersView instantiates a new ProvidersView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvidersViewWithDefaults

`func NewProvidersViewWithDefaults() *ProvidersView`

NewProvidersViewWithDefaults instantiates a new ProvidersView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *ProvidersView) GetProviders() []ProviderCard`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ProvidersView) GetProvidersOk() (*[]ProviderCard, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ProvidersView) SetProviders(v []ProviderCard)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ProvidersView) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


