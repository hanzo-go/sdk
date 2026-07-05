# ExecExecResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** |  | [optional] 
**Stdout** | Pointer to **string** |  | [optional] 
**Stderr** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]ExecExecFile**](ExecExecFile.md) |  | [optional] 

## Methods

### NewExecExecResult

`func NewExecExecResult() *ExecExecResult`

NewExecExecResult instantiates a new ExecExecResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecExecResultWithDefaults

`func NewExecExecResultWithDefaults() *ExecExecResult`

NewExecExecResultWithDefaults instantiates a new ExecExecResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *ExecExecResult) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExecExecResult) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExecExecResult) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ExecExecResult) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStdout

`func (o *ExecExecResult) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *ExecExecResult) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *ExecExecResult) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *ExecExecResult) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetStderr

`func (o *ExecExecResult) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *ExecExecResult) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *ExecExecResult) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *ExecExecResult) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetFiles

`func (o *ExecExecResult) GetFiles() []ExecExecFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExecExecResult) GetFilesOk() (*[]ExecExecFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExecExecResult) SetFiles(v []ExecExecFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExecExecResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


