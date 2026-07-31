# SearchBatchViewStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNbTasks** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **map[string]int32** |  | [optional] 
**Types** | Pointer to **map[string]int32** |  | [optional] 
**IndexUids** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewSearchBatchViewStats

`func NewSearchBatchViewStats() *SearchBatchViewStats`

NewSearchBatchViewStats instantiates a new SearchBatchViewStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBatchViewStatsWithDefaults

`func NewSearchBatchViewStatsWithDefaults() *SearchBatchViewStats`

NewSearchBatchViewStatsWithDefaults instantiates a new SearchBatchViewStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNbTasks

`func (o *SearchBatchViewStats) GetTotalNbTasks() int32`

GetTotalNbTasks returns the TotalNbTasks field if non-nil, zero value otherwise.

### GetTotalNbTasksOk

`func (o *SearchBatchViewStats) GetTotalNbTasksOk() (*int32, bool)`

GetTotalNbTasksOk returns a tuple with the TotalNbTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNbTasks

`func (o *SearchBatchViewStats) SetTotalNbTasks(v int32)`

SetTotalNbTasks sets TotalNbTasks field to given value.

### HasTotalNbTasks

`func (o *SearchBatchViewStats) HasTotalNbTasks() bool`

HasTotalNbTasks returns a boolean if a field has been set.

### GetStatus

`func (o *SearchBatchViewStats) GetStatus() map[string]int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchBatchViewStats) GetStatusOk() (*map[string]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchBatchViewStats) SetStatus(v map[string]int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchBatchViewStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypes

`func (o *SearchBatchViewStats) GetTypes() map[string]int32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchBatchViewStats) GetTypesOk() (*map[string]int32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchBatchViewStats) SetTypes(v map[string]int32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchBatchViewStats) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetIndexUids

`func (o *SearchBatchViewStats) GetIndexUids() map[string]int32`

GetIndexUids returns the IndexUids field if non-nil, zero value otherwise.

### GetIndexUidsOk

`func (o *SearchBatchViewStats) GetIndexUidsOk() (*map[string]int32, bool)`

GetIndexUidsOk returns a tuple with the IndexUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUids

`func (o *SearchBatchViewStats) SetIndexUids(v map[string]int32)`

SetIndexUids sets IndexUids field to given value.

### HasIndexUids

`func (o *SearchBatchViewStats) HasIndexUids() bool`

HasIndexUids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


