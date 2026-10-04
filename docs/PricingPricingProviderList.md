# PricingPricingProviderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to **map[string]interface{}** | Providers maps a provider name to its opaque info object. A provider hidden for the caller&#39;s org is absent entirely. | [optional] 
**Updated** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPricingPricingProviderList

`func NewPricingPricingProviderList() *PricingPricingProviderList`

NewPricingPricingProviderList instantiates a new PricingPricingProviderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPricingProviderListWithDefaults

`func NewPricingPricingProviderListWithDefaults() *PricingPricingProviderList`

NewPricingPricingProviderListWithDefaults instantiates a new PricingPricingProviderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *PricingPricingProviderList) GetProviders() map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *PricingPricingProviderList) GetProvidersOk() (*map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *PricingPricingProviderList) SetProviders(v map[string]interface{})`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *PricingPricingProviderList) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetUpdated

`func (o *PricingPricingProviderList) GetUpdated() interface{}`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PricingPricingProviderList) GetUpdatedOk() (*interface{}, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PricingPricingProviderList) SetUpdated(v interface{})`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PricingPricingProviderList) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *PricingPricingProviderList) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *PricingPricingProviderList) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


