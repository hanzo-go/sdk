# AutoCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorCount** | Pointer to **int64** |  | [optional] 
**Connectors** | Pointer to [**[]AutoConnectorMetadata**](AutoConnectorMetadata.md) |  | [optional] 
**RunnableCount** | Pointer to **int64** | RunnableCount is how many of them this deployment can actually invoke. It is the number that matters and it was absent: a caller read connectorCount and reasonably believed all of it worked. | [optional] 

## Methods

### NewAutoCatalog

`func NewAutoCatalog() *AutoCatalog`

NewAutoCatalog instantiates a new AutoCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCatalogWithDefaults

`func NewAutoCatalogWithDefaults() *AutoCatalog`

NewAutoCatalogWithDefaults instantiates a new AutoCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorCount

`func (o *AutoCatalog) GetConnectorCount() int64`

GetConnectorCount returns the ConnectorCount field if non-nil, zero value otherwise.

### GetConnectorCountOk

`func (o *AutoCatalog) GetConnectorCountOk() (*int64, bool)`

GetConnectorCountOk returns a tuple with the ConnectorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCount

`func (o *AutoCatalog) SetConnectorCount(v int64)`

SetConnectorCount sets ConnectorCount field to given value.

### HasConnectorCount

`func (o *AutoCatalog) HasConnectorCount() bool`

HasConnectorCount returns a boolean if a field has been set.

### GetConnectors

`func (o *AutoCatalog) GetConnectors() []AutoConnectorMetadata`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *AutoCatalog) GetConnectorsOk() (*[]AutoConnectorMetadata, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *AutoCatalog) SetConnectors(v []AutoConnectorMetadata)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *AutoCatalog) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetRunnableCount

`func (o *AutoCatalog) GetRunnableCount() int64`

GetRunnableCount returns the RunnableCount field if non-nil, zero value otherwise.

### GetRunnableCountOk

`func (o *AutoCatalog) GetRunnableCountOk() (*int64, bool)`

GetRunnableCountOk returns a tuple with the RunnableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnableCount

`func (o *AutoCatalog) SetRunnableCount(v int64)`

SetRunnableCount sets RunnableCount field to given value.

### HasRunnableCount

`func (o *AutoCatalog) HasRunnableCount() bool`

HasRunnableCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


