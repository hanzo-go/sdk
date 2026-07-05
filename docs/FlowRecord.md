# FlowRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Cells** | Pointer to **map[string]interface{}** |  | [optional] 
**TableId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlowRecord

`func NewFlowRecord() *FlowRecord`

NewFlowRecord instantiates a new FlowRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowRecordWithDefaults

`func NewFlowRecordWithDefaults() *FlowRecord`

NewFlowRecordWithDefaults instantiates a new FlowRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCells

`func (o *FlowRecord) GetCells() map[string]interface{}`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *FlowRecord) GetCellsOk() (*map[string]interface{}, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *FlowRecord) SetCells(v map[string]interface{})`

SetCells sets Cells field to given value.

### HasCells

`func (o *FlowRecord) HasCells() bool`

HasCells returns a boolean if a field has been set.

### GetTableId

`func (o *FlowRecord) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *FlowRecord) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *FlowRecord) SetTableId(v string)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *FlowRecord) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetProjectId

`func (o *FlowRecord) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FlowRecord) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FlowRecord) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FlowRecord) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCreated

`func (o *FlowRecord) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowRecord) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowRecord) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowRecord) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


