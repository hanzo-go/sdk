# Catalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorCount** | Pointer to **int64** |  | [optional] 
**Connectors** | Pointer to [**[]ConnectorMetadata**](ConnectorMetadata.md) |  | [optional] 

## Methods

### NewCatalog

`func NewCatalog() *Catalog`

NewCatalog instantiates a new Catalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogWithDefaults

`func NewCatalogWithDefaults() *Catalog`

NewCatalogWithDefaults instantiates a new Catalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorCount

`func (o *Catalog) GetConnectorCount() int64`

GetConnectorCount returns the ConnectorCount field if non-nil, zero value otherwise.

### GetConnectorCountOk

`func (o *Catalog) GetConnectorCountOk() (*int64, bool)`

GetConnectorCountOk returns a tuple with the ConnectorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCount

`func (o *Catalog) SetConnectorCount(v int64)`

SetConnectorCount sets ConnectorCount field to given value.

### HasConnectorCount

`func (o *Catalog) HasConnectorCount() bool`

HasConnectorCount returns a boolean if a field has been set.

### GetConnectors

`func (o *Catalog) GetConnectors() []ConnectorMetadata`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *Catalog) GetConnectorsOk() (*[]ConnectorMetadata, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *Catalog) SetConnectors(v []ConnectorMetadata)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *Catalog) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


