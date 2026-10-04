# ExecCodeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]ExecCodeFile**](ExecCodeFile.md) | Files are what this run CREATED OR CHANGED, decided by mtime against a marker taken before the program started — so it is the run&#39;s output, not a listing of the directory. Fetch each from GET /v1/exec/download/{session}/{id}. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the sandbox this run used — the one that was passed in, or the fresh one that was leased. Pass it to the next run to keep the filesystem. | [optional] 
**Stderr** | Pointer to **string** | Stderr is what the program wrote to standard error, INCLUDING a compiler&#39;s diagnostics and the trace of a program that exited non-zero. Its presence is not a failed call. | [optional] 
**Stdout** | Pointer to **string** | Stdout is what the program wrote to standard output. | [optional] 

## Methods

### NewExecCodeResult

`func NewExecCodeResult() *ExecCodeResult`

NewExecCodeResult instantiates a new ExecCodeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecCodeResultWithDefaults

`func NewExecCodeResultWithDefaults() *ExecCodeResult`

NewExecCodeResultWithDefaults instantiates a new ExecCodeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ExecCodeResult) GetFiles() []ExecCodeFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExecCodeResult) GetFilesOk() (*[]ExecCodeFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExecCodeResult) SetFiles(v []ExecCodeFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExecCodeResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetSessionId

`func (o *ExecCodeResult) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExecCodeResult) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExecCodeResult) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ExecCodeResult) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStderr

`func (o *ExecCodeResult) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *ExecCodeResult) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *ExecCodeResult) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *ExecCodeResult) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdout

`func (o *ExecCodeResult) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *ExecCodeResult) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *ExecCodeResult) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *ExecCodeResult) HasStdout() bool`

HasStdout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


