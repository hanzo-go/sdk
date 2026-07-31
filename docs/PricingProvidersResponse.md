# PricingProvidersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **time.Time** |  | [optional] 
**Providers** | Pointer to **map[string]int32** | Map of provider name to model count | [optional] 

## Methods

### NewPricingProvidersResponse

`func NewPricingProvidersResponse() *PricingProvidersResponse`

NewPricingProvidersResponse instantiates a new PricingProvidersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingProvidersResponseWithDefaults

`func NewPricingProvidersResponseWithDefaults() *PricingProvidersResponse`

NewPricingProvidersResponseWithDefaults instantiates a new PricingProvidersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *PricingProvidersResponse) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PricingProvidersResponse) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PricingProvidersResponse) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PricingProvidersResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetProviders

`func (o *PricingProvidersResponse) GetProviders() map[string]int32`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *PricingProvidersResponse) GetProvidersOk() (*map[string]int32, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *PricingProvidersResponse) SetProviders(v map[string]int32)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *PricingProvidersResponse) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


