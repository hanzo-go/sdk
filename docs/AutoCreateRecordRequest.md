# AutoCreateRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableId** | **string** |  | 
**Cells** | **map[string]interface{}** |  | 

## Methods

### NewAutoCreateRecordRequest

`func NewAutoCreateRecordRequest(tableId string, cells map[string]interface{}, ) *AutoCreateRecordRequest`

NewAutoCreateRecordRequest instantiates a new AutoCreateRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCreateRecordRequestWithDefaults

`func NewAutoCreateRecordRequestWithDefaults() *AutoCreateRecordRequest`

NewAutoCreateRecordRequestWithDefaults instantiates a new AutoCreateRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableId

`func (o *AutoCreateRecordRequest) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *AutoCreateRecordRequest) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *AutoCreateRecordRequest) SetTableId(v string)`

SetTableId sets TableId field to given value.


### GetCells

`func (o *AutoCreateRecordRequest) GetCells() map[string]interface{}`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *AutoCreateRecordRequest) GetCellsOk() (*map[string]interface{}, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *AutoCreateRecordRequest) SetCells(v map[string]interface{})`

SetCells sets Cells field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


