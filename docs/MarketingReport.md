# MarketingReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bases** | Pointer to **map[string]int64** | Bases count the rows written by the consent basis recorded for them. | [optional] 
**Columns** | Pointer to **[]string** | Columns are the file&#39;s columns, as its header named them. | [optional] 
**Imported** | Pointer to **int64** | Imported is how many rows became new subscribers. | [optional] 
**Invalid** | Pointer to **int64** | Invalid is how many could not be read as an address, for the reasons in Reasons. | [optional] 
**Map** | Pointer to [**MarketingMapping**](MarketingMapping.md) | Map is the column map the import used, with what it found by header. | [optional] 
**Reasons** | Pointer to **map[string]int64** | Reasons count the skipped and invalid rows by why. | [optional] 
**Rows** | Pointer to **int64** | Rows is how many rows were read. | [optional] 
**Samples** | Pointer to [**[]MarketingSample**](MarketingSample.md) | Samples are the first rejected rows: their row number and the reason. | [optional] 
**Skipped** | Pointer to **int64** | Skipped is how many were left as they were, for the reasons in Reasons. | [optional] 
**Statuses** | Pointer to **map[string]int64** | Statuses count the rows written by the status they were written with. | [optional] 
**Updated** | Pointer to **int64** | Updated is how many were already on the list and gained a name, tags or fields, or a status that leaves the list. | [optional] 

## Methods

### NewMarketingReport

`func NewMarketingReport() *MarketingReport`

NewMarketingReport instantiates a new MarketingReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingReportWithDefaults

`func NewMarketingReportWithDefaults() *MarketingReport`

NewMarketingReportWithDefaults instantiates a new MarketingReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBases

`func (o *MarketingReport) GetBases() map[string]int64`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *MarketingReport) GetBasesOk() (*map[string]int64, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *MarketingReport) SetBases(v map[string]int64)`

SetBases sets Bases field to given value.

### HasBases

`func (o *MarketingReport) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetColumns

`func (o *MarketingReport) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *MarketingReport) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *MarketingReport) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *MarketingReport) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetImported

`func (o *MarketingReport) GetImported() int64`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *MarketingReport) GetImportedOk() (*int64, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *MarketingReport) SetImported(v int64)`

SetImported sets Imported field to given value.

### HasImported

`func (o *MarketingReport) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetInvalid

`func (o *MarketingReport) GetInvalid() int64`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *MarketingReport) GetInvalidOk() (*int64, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *MarketingReport) SetInvalid(v int64)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *MarketingReport) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetMap

`func (o *MarketingReport) GetMap() MarketingMapping`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *MarketingReport) GetMapOk() (*MarketingMapping, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *MarketingReport) SetMap(v MarketingMapping)`

SetMap sets Map field to given value.

### HasMap

`func (o *MarketingReport) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetReasons

`func (o *MarketingReport) GetReasons() map[string]int64`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *MarketingReport) GetReasonsOk() (*map[string]int64, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *MarketingReport) SetReasons(v map[string]int64)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *MarketingReport) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetRows

`func (o *MarketingReport) GetRows() int64`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *MarketingReport) GetRowsOk() (*int64, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *MarketingReport) SetRows(v int64)`

SetRows sets Rows field to given value.

### HasRows

`func (o *MarketingReport) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSamples

`func (o *MarketingReport) GetSamples() []MarketingSample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MarketingReport) GetSamplesOk() (*[]MarketingSample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MarketingReport) SetSamples(v []MarketingSample)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *MarketingReport) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSkipped

`func (o *MarketingReport) GetSkipped() int64`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *MarketingReport) GetSkippedOk() (*int64, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *MarketingReport) SetSkipped(v int64)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *MarketingReport) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetStatuses

`func (o *MarketingReport) GetStatuses() map[string]int64`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *MarketingReport) GetStatusesOk() (*map[string]int64, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *MarketingReport) SetStatuses(v map[string]int64)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *MarketingReport) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetUpdated

`func (o *MarketingReport) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *MarketingReport) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *MarketingReport) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *MarketingReport) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


