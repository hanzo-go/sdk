# CodeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]CodeFile**](CodeFile.md) | Files are what this run CREATED OR CHANGED, decided by mtime against a marker taken before the program started — so it is the run&#39;s output, not a listing of the directory. Fetch each from GET /v1/download/{session}/{id}. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the sandbox this run used — the one that was passed in, or the fresh one that was leased. Pass it to the next run to keep the filesystem. | [optional] 
**Stderr** | Pointer to **string** | Stderr is what the program wrote to standard error, INCLUDING a compiler&#39;s diagnostics and the trace of a program that exited non-zero. Its presence is not a failed call. | [optional] 
**Stdout** | Pointer to **string** | Stdout is what the program wrote to standard output. | [optional] 

## Methods

### NewCodeResult

`func NewCodeResult() *CodeResult`

NewCodeResult instantiates a new CodeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeResultWithDefaults

`func NewCodeResultWithDefaults() *CodeResult`

NewCodeResultWithDefaults instantiates a new CodeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *CodeResult) GetFiles() []CodeFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CodeResult) GetFilesOk() (*[]CodeFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CodeResult) SetFiles(v []CodeFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CodeResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetSessionId

`func (o *CodeResult) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CodeResult) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CodeResult) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CodeResult) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStderr

`func (o *CodeResult) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *CodeResult) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *CodeResult) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *CodeResult) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdout

`func (o *CodeResult) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *CodeResult) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *CodeResult) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *CodeResult) HasStdout() bool`

HasStdout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


