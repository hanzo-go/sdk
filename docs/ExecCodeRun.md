# ExecCodeRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | Args become the PROGRAM&#39;s argv, never the compiler&#39;s. For the compiled languages the toolchain builds first and these are passed to the binary it produced. | [optional] 
**Code** | **string** | Code is the WHOLE program, not a fragment: it is written to a single file and that file is what runs, so a compiled language needs its entry point and an interpreted one runs top to bottom. | 
**Files** | Pointer to [**[]ExecCodeFile**](ExecCodeFile.md) | Files are inputs the host already put in some session. Each names the session its bytes live in, which is usually — and ideally — the session this run wants. | [optional] 
**Lang** | **string** | Lang selects the toolchain, and with it the filename the code is written to and the line that runs it: py, js, ts, bash, r, php, go, rs, c, cpp, java, d, f90. Anything else is refused rather than guessed at — a run in the wrong language fails somewhere deep in a compiler, which reads as an outage. | 
**RuntimeSessionHint** | Pointer to **string** | RuntimeSessionHint is the stateful-session hint. It is carried so a client that sends it is not silently misread, and it selects nothing here: every session in this implementation is already a warm sandbox, so there is no second kind of runtime for a hint to choose between. | [optional] 
**SessionId** | Pointer to **string** | SessionID continues an EXISTING sandbox, which is what makes runs stateful: the same filesystem, so one run&#39;s output file is the next run&#39;s input. Empty leases a fresh sandbox and the id it got comes back on the result. | [optional] 
**UserId** | Pointer to **string** | UserID attributes the run inside the caller&#39;s org. It is a label, never a tenant: the org is resolved from the validated principal and a value here cannot widen what the run may reach. | [optional] 

## Methods

### NewExecCodeRun

`func NewExecCodeRun(code string, lang string, ) *ExecCodeRun`

NewExecCodeRun instantiates a new ExecCodeRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecCodeRunWithDefaults

`func NewExecCodeRunWithDefaults() *ExecCodeRun`

NewExecCodeRunWithDefaults instantiates a new ExecCodeRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *ExecCodeRun) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ExecCodeRun) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ExecCodeRun) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ExecCodeRun) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCode

`func (o *ExecCodeRun) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExecCodeRun) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExecCodeRun) SetCode(v string)`

SetCode sets Code field to given value.


### GetFiles

`func (o *ExecCodeRun) GetFiles() []ExecCodeFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExecCodeRun) GetFilesOk() (*[]ExecCodeFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExecCodeRun) SetFiles(v []ExecCodeFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExecCodeRun) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLang

`func (o *ExecCodeRun) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ExecCodeRun) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ExecCodeRun) SetLang(v string)`

SetLang sets Lang field to given value.


### GetRuntimeSessionHint

`func (o *ExecCodeRun) GetRuntimeSessionHint() string`

GetRuntimeSessionHint returns the RuntimeSessionHint field if non-nil, zero value otherwise.

### GetRuntimeSessionHintOk

`func (o *ExecCodeRun) GetRuntimeSessionHintOk() (*string, bool)`

GetRuntimeSessionHintOk returns a tuple with the RuntimeSessionHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeSessionHint

`func (o *ExecCodeRun) SetRuntimeSessionHint(v string)`

SetRuntimeSessionHint sets RuntimeSessionHint field to given value.

### HasRuntimeSessionHint

`func (o *ExecCodeRun) HasRuntimeSessionHint() bool`

HasRuntimeSessionHint returns a boolean if a field has been set.

### GetSessionId

`func (o *ExecCodeRun) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExecCodeRun) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExecCodeRun) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ExecCodeRun) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetUserId

`func (o *ExecCodeRun) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExecCodeRun) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExecCodeRun) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExecCodeRun) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


