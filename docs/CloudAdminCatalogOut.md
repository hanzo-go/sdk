# CloudAdminCatalogOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to **[]map[string]map[string]interface{}** | Models is every model the catalog holds — disabled ones included — each carrying its enablement state under \&quot;_overlay\&quot;. | [optional] 
**Providers** | Pointer to **map[string]map[string]interface{}** | Providers is every provider the catalog holds, keyed by name, each carrying its enablement state under \&quot;_overlay\&quot;. | [optional] 
**Updated** | Pointer to **map[string]interface{}** | Updated is when the catalog was last refreshed, as the pricing source recorded it. | [optional] 

## Methods

### NewCloudAdminCatalogOut

`func NewCloudAdminCatalogOut() *CloudAdminCatalogOut`

NewCloudAdminCatalogOut instantiates a new CloudAdminCatalogOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdminCatalogOutWithDefaults

`func NewCloudAdminCatalogOutWithDefaults() *CloudAdminCatalogOut`

NewCloudAdminCatalogOutWithDefaults instantiates a new CloudAdminCatalogOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *CloudAdminCatalogOut) GetModels() []map[string]map[string]interface{}`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CloudAdminCatalogOut) GetModelsOk() (*[]map[string]map[string]interface{}, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CloudAdminCatalogOut) SetModels(v []map[string]map[string]interface{})`

SetModels sets Models field to given value.

### HasModels

`func (o *CloudAdminCatalogOut) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProviders

`func (o *CloudAdminCatalogOut) GetProviders() map[string]map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CloudAdminCatalogOut) GetProvidersOk() (*map[string]map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CloudAdminCatalogOut) SetProviders(v map[string]map[string]interface{})`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CloudAdminCatalogOut) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudAdminCatalogOut) GetUpdated() map[string]interface{}`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudAdminCatalogOut) GetUpdatedOk() (*map[string]interface{}, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudAdminCatalogOut) SetUpdated(v map[string]interface{})`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudAdminCatalogOut) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


