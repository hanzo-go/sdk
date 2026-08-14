# RiskDatasetDisposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to **int32** |  | [optional] 
**Versions** | Pointer to **int32** | Versions is how many versions went, and Rows how many rows they held between them, as the register recorded them. | [optional] 

## Methods

### NewRiskDatasetDisposal

`func NewRiskDatasetDisposal() *RiskDatasetDisposal`

NewRiskDatasetDisposal instantiates a new RiskDatasetDisposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDatasetDisposalWithDefaults

`func NewRiskDatasetDisposalWithDefaults() *RiskDatasetDisposal`

NewRiskDatasetDisposalWithDefaults instantiates a new RiskDatasetDisposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *RiskDatasetDisposal) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *RiskDatasetDisposal) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *RiskDatasetDisposal) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *RiskDatasetDisposal) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetRows

`func (o *RiskDatasetDisposal) GetRows() int32`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *RiskDatasetDisposal) GetRowsOk() (*int32, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *RiskDatasetDisposal) SetRows(v int32)`

SetRows sets Rows field to given value.

### HasRows

`func (o *RiskDatasetDisposal) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetVersions

`func (o *RiskDatasetDisposal) GetVersions() int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *RiskDatasetDisposal) GetVersionsOk() (*int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *RiskDatasetDisposal) SetVersions(v int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *RiskDatasetDisposal) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


