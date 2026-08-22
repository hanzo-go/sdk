# ExecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Argv** | Pointer to **[]string** | Argv is the command as an argument vector, which is the honest form: it cannot be word-split by accident. Send this OR Command, not both. | [optional] 
**Command** | Pointer to **string** | Command is a shell line, for a caller that holds one. It is a convenience over Argv and is the only input that ever reaches a shell. | [optional] 
**Dir** | Pointer to **string** | Dir is the working directory to run in. Empty runs in the class&#39;s own workdir — /mnt/data for exec, /work for dev. | [optional] 
**Id** | Pointer to **string** | ID is the sandbox to run in, from the path. | [optional] 
**Stdin** | Pointer to **string** | Stdin is fed to the command on its standard input. | [optional] 
**TimeoutSec** | Pointer to **int32** | TimeoutSec bounds the run in seconds. Zero takes the default. | [optional] 

## Methods

### NewExecRequest

`func NewExecRequest() *ExecRequest`

NewExecRequest instantiates a new ExecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecRequestWithDefaults

`func NewExecRequestWithDefaults() *ExecRequest`

NewExecRequestWithDefaults instantiates a new ExecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgv

`func (o *ExecRequest) GetArgv() []string`

GetArgv returns the Argv field if non-nil, zero value otherwise.

### GetArgvOk

`func (o *ExecRequest) GetArgvOk() (*[]string, bool)`

GetArgvOk returns a tuple with the Argv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgv

`func (o *ExecRequest) SetArgv(v []string)`

SetArgv sets Argv field to given value.

### HasArgv

`func (o *ExecRequest) HasArgv() bool`

HasArgv returns a boolean if a field has been set.

### GetCommand

`func (o *ExecRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ExecRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ExecRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ExecRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDir

`func (o *ExecRequest) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *ExecRequest) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *ExecRequest) SetDir(v string)`

SetDir sets Dir field to given value.

### HasDir

`func (o *ExecRequest) HasDir() bool`

HasDir returns a boolean if a field has been set.

### GetId

`func (o *ExecRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStdin

`func (o *ExecRequest) GetStdin() string`

GetStdin returns the Stdin field if non-nil, zero value otherwise.

### GetStdinOk

`func (o *ExecRequest) GetStdinOk() (*string, bool)`

GetStdinOk returns a tuple with the Stdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdin

`func (o *ExecRequest) SetStdin(v string)`

SetStdin sets Stdin field to given value.

### HasStdin

`func (o *ExecRequest) HasStdin() bool`

HasStdin returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *ExecRequest) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *ExecRequest) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *ExecRequest) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *ExecRequest) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


