# OperativeEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | Editor command to execute | 
**Path** | **string** | Absolute file path | 
**FileText** | Pointer to **string** | File content (for create command) | [optional] 
**OldStr** | Pointer to **string** | String to find (for str_replace command) | [optional] 
**NewStr** | Pointer to **string** | Replacement string (for str_replace and insert commands) | [optional] 
**InsertLine** | Pointer to **int32** | Line number to insert at (for insert command) | [optional] 
**ViewRange** | Pointer to **[]int32** | [startLine, endLine] range to view | [optional] 

## Methods

### NewOperativeEditRequest

`func NewOperativeEditRequest(command string, path string, ) *OperativeEditRequest`

NewOperativeEditRequest instantiates a new OperativeEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeEditRequestWithDefaults

`func NewOperativeEditRequestWithDefaults() *OperativeEditRequest`

NewOperativeEditRequestWithDefaults instantiates a new OperativeEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *OperativeEditRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *OperativeEditRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *OperativeEditRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetPath

`func (o *OperativeEditRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OperativeEditRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OperativeEditRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetFileText

`func (o *OperativeEditRequest) GetFileText() string`

GetFileText returns the FileText field if non-nil, zero value otherwise.

### GetFileTextOk

`func (o *OperativeEditRequest) GetFileTextOk() (*string, bool)`

GetFileTextOk returns a tuple with the FileText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileText

`func (o *OperativeEditRequest) SetFileText(v string)`

SetFileText sets FileText field to given value.

### HasFileText

`func (o *OperativeEditRequest) HasFileText() bool`

HasFileText returns a boolean if a field has been set.

### GetOldStr

`func (o *OperativeEditRequest) GetOldStr() string`

GetOldStr returns the OldStr field if non-nil, zero value otherwise.

### GetOldStrOk

`func (o *OperativeEditRequest) GetOldStrOk() (*string, bool)`

GetOldStrOk returns a tuple with the OldStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldStr

`func (o *OperativeEditRequest) SetOldStr(v string)`

SetOldStr sets OldStr field to given value.

### HasOldStr

`func (o *OperativeEditRequest) HasOldStr() bool`

HasOldStr returns a boolean if a field has been set.

### GetNewStr

`func (o *OperativeEditRequest) GetNewStr() string`

GetNewStr returns the NewStr field if non-nil, zero value otherwise.

### GetNewStrOk

`func (o *OperativeEditRequest) GetNewStrOk() (*string, bool)`

GetNewStrOk returns a tuple with the NewStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStr

`func (o *OperativeEditRequest) SetNewStr(v string)`

SetNewStr sets NewStr field to given value.

### HasNewStr

`func (o *OperativeEditRequest) HasNewStr() bool`

HasNewStr returns a boolean if a field has been set.

### GetInsertLine

`func (o *OperativeEditRequest) GetInsertLine() int32`

GetInsertLine returns the InsertLine field if non-nil, zero value otherwise.

### GetInsertLineOk

`func (o *OperativeEditRequest) GetInsertLineOk() (*int32, bool)`

GetInsertLineOk returns a tuple with the InsertLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertLine

`func (o *OperativeEditRequest) SetInsertLine(v int32)`

SetInsertLine sets InsertLine field to given value.

### HasInsertLine

`func (o *OperativeEditRequest) HasInsertLine() bool`

HasInsertLine returns a boolean if a field has been set.

### GetViewRange

`func (o *OperativeEditRequest) GetViewRange() []int32`

GetViewRange returns the ViewRange field if non-nil, zero value otherwise.

### GetViewRangeOk

`func (o *OperativeEditRequest) GetViewRangeOk() (*[]int32, bool)`

GetViewRangeOk returns a tuple with the ViewRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRange

`func (o *OperativeEditRequest) SetViewRange(v []int32)`

SetViewRange sets ViewRange field to given value.

### HasViewRange

`func (o *OperativeEditRequest) HasViewRange() bool`

HasViewRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


