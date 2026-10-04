# DatasetRiskDatasetVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DatasetRiskDataset**](DatasetRiskDataset.md) | Items is every version of it, newest first — including the disposed ones, whose record outlives their rows. Never null. | [optional] 
**Name** | Pointer to **string** | Name is the dataset these versions belong to, as the register holds it. | [optional] 

## Methods

### NewDatasetRiskDatasetVersions

`func NewDatasetRiskDatasetVersions() *DatasetRiskDatasetVersions`

NewDatasetRiskDatasetVersions instantiates a new DatasetRiskDatasetVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetRiskDatasetVersionsWithDefaults

`func NewDatasetRiskDatasetVersionsWithDefaults() *DatasetRiskDatasetVersions`

NewDatasetRiskDatasetVersionsWithDefaults instantiates a new DatasetRiskDatasetVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DatasetRiskDatasetVersions) GetItems() []DatasetRiskDataset`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DatasetRiskDatasetVersions) GetItemsOk() (*[]DatasetRiskDataset, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DatasetRiskDatasetVersions) SetItems(v []DatasetRiskDataset)`

SetItems sets Items field to given value.

### HasItems

`func (o *DatasetRiskDatasetVersions) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *DatasetRiskDatasetVersions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetRiskDatasetVersions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetRiskDatasetVersions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasetRiskDatasetVersions) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


