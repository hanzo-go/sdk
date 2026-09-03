# ImportCapTableOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**Formation**](Formation.md) | Formation is the org&#39;s incorporation record, now marked cap-table-imported. | [optional] 
**Rows** | Pointer to **int64** | Rows is how many rows were read from the sheet, header included. | [optional] 
**StakeholdersImported** | Pointer to **int64** | StakeholdersImported is how many stakeholders the cap table accepted. | [optional] 

## Methods

### NewImportCapTableOut

`func NewImportCapTableOut() *ImportCapTableOut`

NewImportCapTableOut instantiates a new ImportCapTableOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCapTableOutWithDefaults

`func NewImportCapTableOutWithDefaults() *ImportCapTableOut`

NewImportCapTableOutWithDefaults instantiates a new ImportCapTableOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *ImportCapTableOut) GetFormation() Formation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *ImportCapTableOut) GetFormationOk() (*Formation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *ImportCapTableOut) SetFormation(v Formation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *ImportCapTableOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetRows

`func (o *ImportCapTableOut) GetRows() int64`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ImportCapTableOut) GetRowsOk() (*int64, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ImportCapTableOut) SetRows(v int64)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ImportCapTableOut) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetStakeholdersImported

`func (o *ImportCapTableOut) GetStakeholdersImported() int64`

GetStakeholdersImported returns the StakeholdersImported field if non-nil, zero value otherwise.

### GetStakeholdersImportedOk

`func (o *ImportCapTableOut) GetStakeholdersImportedOk() (*int64, bool)`

GetStakeholdersImportedOk returns a tuple with the StakeholdersImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholdersImported

`func (o *ImportCapTableOut) SetStakeholdersImported(v int64)`

SetStakeholdersImported sets StakeholdersImported field to given value.

### HasStakeholdersImported

`func (o *ImportCapTableOut) HasStakeholdersImported() bool`

HasStakeholdersImported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


