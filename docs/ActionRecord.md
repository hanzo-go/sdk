# ActionRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **string** | Args is the JSON the tool was called with, recorded as TEXT exactly as sent — including whatever the AI drafted into it — so a run can be read back and reproduced. It is a string, not an object. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the run was recorded, as Unix seconds. The ledger is read newest-first on this column. | [optional] 
**Err** | Pointer to **string** | Err is why the run failed, when it did. Empty on a successful run. | [optional] 
**Id** | Pointer to **string** | ID identifies this one execution. The ledger is append-only, so an id is never reused and never updated. | [optional] 
**Ok** | Pointer to **bool** | OK is whether the tool ran to completion. It is the ledger&#39;s own verdict, not the tool&#39;s opinion of the outcome — a tool that succeeded at reporting bad news is ok. | [optional] 
**Result** | Pointer to **string** | Result is the tool&#39;s own answer, likewise recorded as JSON text. Present on a failed run too, where the tool answered but the answer was a refusal. | [optional] 
**StepId** | Pointer to **string** | StepID is the checklist step the Business AI was acting on. | [optional] 
**Tool** | Pointer to **string** | Tool is the MCP tool that was dispatched, by name. | [optional] 

## Methods

### NewActionRecord

`func NewActionRecord() *ActionRecord`

NewActionRecord instantiates a new ActionRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRecordWithDefaults

`func NewActionRecordWithDefaults() *ActionRecord`

NewActionRecordWithDefaults instantiates a new ActionRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *ActionRecord) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ActionRecord) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ActionRecord) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ActionRecord) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ActionRecord) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActionRecord) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActionRecord) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ActionRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetErr

`func (o *ActionRecord) GetErr() string`

GetErr returns the Err field if non-nil, zero value otherwise.

### GetErrOk

`func (o *ActionRecord) GetErrOk() (*string, bool)`

GetErrOk returns a tuple with the Err field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErr

`func (o *ActionRecord) SetErr(v string)`

SetErr sets Err field to given value.

### HasErr

`func (o *ActionRecord) HasErr() bool`

HasErr returns a boolean if a field has been set.

### GetId

`func (o *ActionRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOk

`func (o *ActionRecord) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ActionRecord) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ActionRecord) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *ActionRecord) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetResult

`func (o *ActionRecord) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ActionRecord) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ActionRecord) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ActionRecord) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStepId

`func (o *ActionRecord) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *ActionRecord) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *ActionRecord) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *ActionRecord) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetTool

`func (o *ActionRecord) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *ActionRecord) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *ActionRecord) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *ActionRecord) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


