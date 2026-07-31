# CloudMcpCatalogSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to **int32** | Added is how many listings the catalog did not have before. | [optional] 
**Registry** | Pointer to **string** | Registry is the upstream this pass read. | [optional] 
**Total** | Pointer to **int32** | Total is how many listings the catalog holds now. | [optional] 
**Updated** | Pointer to **int32** | Updated is how many the publisher has changed since we last looked. | [optional] 

## Methods

### NewCloudMcpCatalogSync

`func NewCloudMcpCatalogSync() *CloudMcpCatalogSync`

NewCloudMcpCatalogSync instantiates a new CloudMcpCatalogSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMcpCatalogSyncWithDefaults

`func NewCloudMcpCatalogSyncWithDefaults() *CloudMcpCatalogSync`

NewCloudMcpCatalogSyncWithDefaults instantiates a new CloudMcpCatalogSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *CloudMcpCatalogSync) GetAdded() int32`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *CloudMcpCatalogSync) GetAddedOk() (*int32, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *CloudMcpCatalogSync) SetAdded(v int32)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *CloudMcpCatalogSync) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetRegistry

`func (o *CloudMcpCatalogSync) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *CloudMcpCatalogSync) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *CloudMcpCatalogSync) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *CloudMcpCatalogSync) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetTotal

`func (o *CloudMcpCatalogSync) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudMcpCatalogSync) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudMcpCatalogSync) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudMcpCatalogSync) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudMcpCatalogSync) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudMcpCatalogSync) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudMcpCatalogSync) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudMcpCatalogSync) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


