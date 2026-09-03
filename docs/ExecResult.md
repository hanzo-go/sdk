# ExecResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitCode** | Pointer to **int64** | ExitCode is the command&#39;s own exit status. A non-zero one is a SUCCESSFUL call carrying a failed command — the HTTP status stays 200, because \&quot;the command failed\&quot; and \&quot;the call failed\&quot; are different facts and a caller has to be able to tell them apart. | [optional] 
**Stderr** | Pointer to **string** | Stderr is everything it wrote to standard error. It is populated on a successful run too — plenty of tools report progress there — so it is not a signal that anything went wrong; ExitCode is. | [optional] 
**Stdout** | Pointer to **string** | Stdout is everything the command wrote to standard output, as text. | [optional] 

## Methods

### NewExecResult

`func NewExecResult() *ExecResult`

NewExecResult instantiates a new ExecResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecResultWithDefaults

`func NewExecResultWithDefaults() *ExecResult`

NewExecResultWithDefaults instantiates a new ExecResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitCode

`func (o *ExecResult) GetExitCode() int64`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ExecResult) GetExitCodeOk() (*int64, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ExecResult) SetExitCode(v int64)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ExecResult) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetStderr

`func (o *ExecResult) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *ExecResult) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *ExecResult) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *ExecResult) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdout

`func (o *ExecResult) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *ExecResult) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *ExecResult) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *ExecResult) HasStdout() bool`

HasStdout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


