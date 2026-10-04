# DatasetRiskDatasetDisposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** | Dataset is the dataset that was disposed of. The NAME survives: declaring it again continues the version count rather than starting over at 1. | [optional] 
**Rows** | Pointer to **int64** | Rows is how many rows they held between them, as the REGISTER recorded them when each was materialised — not a count of what the drop deleted, which is gone by the time this answers. | [optional] 
**Versions** | Pointer to **int64** | Versions is how many versions went. | [optional] 

## Methods

### NewDatasetRiskDatasetDisposal

`func NewDatasetRiskDatasetDisposal() *DatasetRiskDatasetDisposal`

NewDatasetRiskDatasetDisposal instantiates a new DatasetRiskDatasetDisposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetRiskDatasetDisposalWithDefaults

`func NewDatasetRiskDatasetDisposalWithDefaults() *DatasetRiskDatasetDisposal`

NewDatasetRiskDatasetDisposalWithDefaults instantiates a new DatasetRiskDatasetDisposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *DatasetRiskDatasetDisposal) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *DatasetRiskDatasetDisposal) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *DatasetRiskDatasetDisposal) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *DatasetRiskDatasetDisposal) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetRows

`func (o *DatasetRiskDatasetDisposal) GetRows() int64`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *DatasetRiskDatasetDisposal) GetRowsOk() (*int64, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *DatasetRiskDatasetDisposal) SetRows(v int64)`

SetRows sets Rows field to given value.

### HasRows

`func (o *DatasetRiskDatasetDisposal) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetVersions

`func (o *DatasetRiskDatasetDisposal) GetVersions() int64`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *DatasetRiskDatasetDisposal) GetVersionsOk() (*int64, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *DatasetRiskDatasetDisposal) SetVersions(v int64)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *DatasetRiskDatasetDisposal) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


