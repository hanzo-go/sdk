# CloudPricingProviderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to **map[string]map[string]interface{}** | Providers maps a provider name to its opaque info object. A provider hidden for the caller&#39;s org is absent entirely. | [optional] 
**Updated** | Pointer to **map[string]interface{}** | Updated is when the catalog was last refreshed, as the pricing source recorded it. | [optional] 

## Methods

### NewCloudPricingProviderList

`func NewCloudPricingProviderList() *CloudPricingProviderList`

NewCloudPricingProviderList instantiates a new CloudPricingProviderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPricingProviderListWithDefaults

`func NewCloudPricingProviderListWithDefaults() *CloudPricingProviderList`

NewCloudPricingProviderListWithDefaults instantiates a new CloudPricingProviderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *CloudPricingProviderList) GetProviders() map[string]map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CloudPricingProviderList) GetProvidersOk() (*map[string]map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CloudPricingProviderList) SetProviders(v map[string]map[string]interface{})`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CloudPricingProviderList) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudPricingProviderList) GetUpdated() map[string]interface{}`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudPricingProviderList) GetUpdatedOk() (*map[string]interface{}, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudPricingProviderList) SetUpdated(v map[string]interface{})`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudPricingProviderList) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


