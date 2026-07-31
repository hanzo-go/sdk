# CloudCatalogOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to [**[]CloudCatalogEntry**](CloudCatalogEntry.md) | Connectors is every connectable source, sorted by provider. | [optional] 

## Methods

### NewCloudCatalogOut

`func NewCloudCatalogOut() *CloudCatalogOut`

NewCloudCatalogOut instantiates a new CloudCatalogOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCatalogOutWithDefaults

`func NewCloudCatalogOutWithDefaults() *CloudCatalogOut`

NewCloudCatalogOutWithDefaults instantiates a new CloudCatalogOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *CloudCatalogOut) GetConnectors() []CloudCatalogEntry`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *CloudCatalogOut) GetConnectorsOk() (*[]CloudCatalogEntry, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *CloudCatalogOut) SetConnectors(v []CloudCatalogEntry)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *CloudCatalogOut) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


