# McpCatalogSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to **int32** | Added is how many listings the catalog did not have before. | [optional] 
**Registry** | Pointer to **string** | Registry is the upstream this pass read. | [optional] 
**Total** | Pointer to **int32** | Total is how many listings the catalog holds now. | [optional] 
**Updated** | Pointer to **int32** | Updated is how many the publisher has changed since we last looked. | [optional] 

## Methods

### NewMcpCatalogSync

`func NewMcpCatalogSync() *McpCatalogSync`

NewMcpCatalogSync instantiates a new McpCatalogSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpCatalogSyncWithDefaults

`func NewMcpCatalogSyncWithDefaults() *McpCatalogSync`

NewMcpCatalogSyncWithDefaults instantiates a new McpCatalogSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *McpCatalogSync) GetAdded() int32`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *McpCatalogSync) GetAddedOk() (*int32, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *McpCatalogSync) SetAdded(v int32)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *McpCatalogSync) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetRegistry

`func (o *McpCatalogSync) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *McpCatalogSync) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *McpCatalogSync) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *McpCatalogSync) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetTotal

`func (o *McpCatalogSync) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *McpCatalogSync) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *McpCatalogSync) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *McpCatalogSync) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpdated

`func (o *McpCatalogSync) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *McpCatalogSync) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *McpCatalogSync) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *McpCatalogSync) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


