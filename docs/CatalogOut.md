# CatalogOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to [**[]CatalogEntry**](CatalogEntry.md) | Connectors is every connectable source, sorted by provider. | [optional] 

## Methods

### NewCatalogOut

`func NewCatalogOut() *CatalogOut`

NewCatalogOut instantiates a new CatalogOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOutWithDefaults

`func NewCatalogOutWithDefaults() *CatalogOut`

NewCatalogOutWithDefaults instantiates a new CatalogOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *CatalogOut) GetConnectors() []CatalogEntry`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *CatalogOut) GetConnectorsOk() (*[]CatalogEntry, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *CatalogOut) SetConnectors(v []CatalogEntry)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *CatalogOut) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


