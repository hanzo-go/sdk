# Ran

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitCode** | Pointer to **int32** | ExitCode is the PROGRAM&#39;s own status — 0 succeeded, anything else is what it returned, and a Command runs under &#x60;sh -c&#x60; so its shell&#39;s conventions apply. A command that never reached an exit does not arrive here at all: a timeout or a stop cancels the channel, and that is an error on the call rather than a code of ours invented to fill this field. | [optional] 
**Stderr** | Pointer to **string** | Stderr is standard error, kept apart from Stdout so a caller reading a program&#39;s OUTPUT is not reading its diagnostics as data. Same 1 MiB cap, same redaction. A program that failed usually says why here and nowhere else. | [optional] 
**Stdout** | Pointer to **string** | Stdout is what the program wrote to standard output, collected whole rather than streamed — to watch it arrive instead, name a RunIn.Session and read that session&#39;s feed. Capped at 1 MiB, past which it ends in \&quot;[truncated at 1MiB]\&quot;. Every string named in RunIn.Blind is replaced by \&quot;[redacted]\&quot; before it gets here, and before it reaches the session. | [optional] 

## Methods

### NewRan

`func NewRan() *Ran`

NewRan instantiates a new Ran object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRanWithDefaults

`func NewRanWithDefaults() *Ran`

NewRanWithDefaults instantiates a new Ran object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitCode

`func (o *Ran) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Ran) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Ran) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Ran) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetStderr

`func (o *Ran) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *Ran) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *Ran) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *Ran) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdout

`func (o *Ran) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *Ran) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *Ran) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *Ran) HasStdout() bool`

HasStdout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


