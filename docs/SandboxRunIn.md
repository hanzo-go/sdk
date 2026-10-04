# SandboxRunIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Argv** | Pointer to **[]string** | Argv is the program and its arguments, already split — the form no shell can misread. Give this or Command, not both. | [optional] 
**Blind** | Pointer to **[]string** | Blind is the set of secrets this command must never publish.  It exists because output is redacted where it is PRODUCED or not at all. A caller that scrubbed the returned result would still have streamed the unredacted bytes into the session as they were written — to a durable event store, an SSE feed and a chat thread — because the narration leaves the sandbox by a different path from the result. Nothing downstream can take a secret back out of a message that has already been delivered.  The sandbox holds these only for the life of the one command, applies them to every stream leaving it, and never logs or stores them. | [optional] 
**Command** | Pointer to **string** | Command is a shell line, run by &#x60;sh -c&#x60;. Use it when a pipeline or a redirection is the point, and Argv when it is not. | [optional] 
**Dir** | Pointer to **string** | Dir runs the command somewhere other than the sandbox&#39;s working directory, which Leased.Workdir names. | [optional] 
**Events** | Pointer to **bool** | Events says the command writes JSON lines to stdout, one event each, as an agent harness&#39;s --json does: each is appended to Session as one &#x60;event&#x60;, in the shape the command wrote it, and stderr is narrated as output. | [optional] 
**Id** | Pointer to **string** | ID is the sandbox to run in, from an earlier lease. | [optional] 
**Session** | Pointer to **string** | Session is the live agent session this command narrates into: its output is appended there AS IT IS PRODUCED, so every surface watching that session watches the command work instead of a blank pause. A long run is otherwise a silence with a verdict at the end.  It names a SESSION and never a tenant. The org is the one the caller already proved, so a session belonging to somebody else is simply absent from the org this call acts for and the append is refused there. Empty means nothing is watching, and then nothing is sent. | [optional] 
**Stdin** | Pointer to **string** | Stdin is fed to the program on standard input. This is how bytes reach a file without being quoted into a shell line: &#x60;cat &gt; path&#x60; with the contents here writes them exactly. | [optional] 
**TimeoutSec** | Pointer to **int64** | TimeoutSec bounds this ONE command, so a wedged program holds the caller for its own timeout rather than for the whole lease. | [optional] 

## Methods

### NewSandboxRunIn

`func NewSandboxRunIn() *SandboxRunIn`

NewSandboxRunIn instantiates a new SandboxRunIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxRunInWithDefaults

`func NewSandboxRunInWithDefaults() *SandboxRunIn`

NewSandboxRunInWithDefaults instantiates a new SandboxRunIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgv

`func (o *SandboxRunIn) GetArgv() []string`

GetArgv returns the Argv field if non-nil, zero value otherwise.

### GetArgvOk

`func (o *SandboxRunIn) GetArgvOk() (*[]string, bool)`

GetArgvOk returns a tuple with the Argv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgv

`func (o *SandboxRunIn) SetArgv(v []string)`

SetArgv sets Argv field to given value.

### HasArgv

`func (o *SandboxRunIn) HasArgv() bool`

HasArgv returns a boolean if a field has been set.

### GetBlind

`func (o *SandboxRunIn) GetBlind() []string`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *SandboxRunIn) GetBlindOk() (*[]string, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *SandboxRunIn) SetBlind(v []string)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *SandboxRunIn) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetCommand

`func (o *SandboxRunIn) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SandboxRunIn) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SandboxRunIn) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *SandboxRunIn) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDir

`func (o *SandboxRunIn) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *SandboxRunIn) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *SandboxRunIn) SetDir(v string)`

SetDir sets Dir field to given value.

### HasDir

`func (o *SandboxRunIn) HasDir() bool`

HasDir returns a boolean if a field has been set.

### GetEvents

`func (o *SandboxRunIn) GetEvents() bool`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SandboxRunIn) GetEventsOk() (*bool, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SandboxRunIn) SetEvents(v bool)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SandboxRunIn) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *SandboxRunIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxRunIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxRunIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxRunIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSession

`func (o *SandboxRunIn) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SandboxRunIn) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SandboxRunIn) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *SandboxRunIn) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStdin

`func (o *SandboxRunIn) GetStdin() string`

GetStdin returns the Stdin field if non-nil, zero value otherwise.

### GetStdinOk

`func (o *SandboxRunIn) GetStdinOk() (*string, bool)`

GetStdinOk returns a tuple with the Stdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdin

`func (o *SandboxRunIn) SetStdin(v string)`

SetStdin sets Stdin field to given value.

### HasStdin

`func (o *SandboxRunIn) HasStdin() bool`

HasStdin returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *SandboxRunIn) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *SandboxRunIn) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *SandboxRunIn) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *SandboxRunIn) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


