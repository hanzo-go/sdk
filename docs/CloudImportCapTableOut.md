# CloudImportCapTableOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**CloudFormation**](CloudFormation.md) | Formation is the org&#39;s incorporation record, now marked cap-table-imported. | [optional] 
**Rows** | Pointer to **int32** | Rows is how many rows were read from the sheet, header included. | [optional] 
**StakeholdersImported** | Pointer to **int32** | StakeholdersImported is how many stakeholders the cap table accepted. | [optional] 

## Methods

### NewCloudImportCapTableOut

`func NewCloudImportCapTableOut() *CloudImportCapTableOut`

NewCloudImportCapTableOut instantiates a new CloudImportCapTableOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudImportCapTableOutWithDefaults

`func NewCloudImportCapTableOutWithDefaults() *CloudImportCapTableOut`

NewCloudImportCapTableOutWithDefaults instantiates a new CloudImportCapTableOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *CloudImportCapTableOut) GetFormation() CloudFormation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *CloudImportCapTableOut) GetFormationOk() (*CloudFormation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *CloudImportCapTableOut) SetFormation(v CloudFormation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *CloudImportCapTableOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetRows

`func (o *CloudImportCapTableOut) GetRows() int32`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CloudImportCapTableOut) GetRowsOk() (*int32, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CloudImportCapTableOut) SetRows(v int32)`

SetRows sets Rows field to given value.

### HasRows

`func (o *CloudImportCapTableOut) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetStakeholdersImported

`func (o *CloudImportCapTableOut) GetStakeholdersImported() int32`

GetStakeholdersImported returns the StakeholdersImported field if non-nil, zero value otherwise.

### GetStakeholdersImportedOk

`func (o *CloudImportCapTableOut) GetStakeholdersImportedOk() (*int32, bool)`

GetStakeholdersImportedOk returns a tuple with the StakeholdersImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholdersImported

`func (o *CloudImportCapTableOut) SetStakeholdersImported(v int32)`

SetStakeholdersImported sets StakeholdersImported field to given value.

### HasStakeholdersImported

`func (o *CloudImportCapTableOut) HasStakeholdersImported() bool`

HasStakeholdersImported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


