# TaskReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the room inside that workspace, and it is what decides whether there is an address at all. Empty is the ordinary case: a caller of /v1/agent/coding holds a session id and reads the stream instead. | [optional] 
**Provider** | Pointer to **string** | Provider is the transport the conversation is on — slack, discord, telegram, teams — and it is what decides who can carry the text at all. An address without it names a room in no particular workspace. | [optional] 
**Thread** | Pointer to **string** | Thread narrows the address to one thread inside the channel — on Slack the parent message&#39;s ts, the same value a reply carries as thread_ts. Empty puts the text at the top level of the channel. | [optional] 
**Turn** | Pointer to **string** | Turn is the turn that renders a run started here: the session whose one message the run&#39;s progress and its outcome are shown in. A run with a Turn posts nothing of its own — one question, one message. The dispatcher writes it (apps/agents/replyto.go); a model never does. | [optional] 
**User** | Pointer to **string** | User is the platform&#39;s id for the person the conversation is with — whose run it is. A run&#39;s buttons name them to anyone else who presses one; it is never what decides who may (that is the chat identity). | [optional] 

## Methods

### NewTaskReply

`func NewTaskReply() *TaskReply`

NewTaskReply instantiates a new TaskReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskReplyWithDefaults

`func NewTaskReplyWithDefaults() *TaskReply`

NewTaskReplyWithDefaults instantiates a new TaskReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *TaskReply) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TaskReply) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TaskReply) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TaskReply) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetProvider

`func (o *TaskReply) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TaskReply) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TaskReply) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *TaskReply) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetThread

`func (o *TaskReply) GetThread() string`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *TaskReply) GetThreadOk() (*string, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *TaskReply) SetThread(v string)`

SetThread sets Thread field to given value.

### HasThread

`func (o *TaskReply) HasThread() bool`

HasThread returns a boolean if a field has been set.

### GetTurn

`func (o *TaskReply) GetTurn() string`

GetTurn returns the Turn field if non-nil, zero value otherwise.

### GetTurnOk

`func (o *TaskReply) GetTurnOk() (*string, bool)`

GetTurnOk returns a tuple with the Turn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurn

`func (o *TaskReply) SetTurn(v string)`

SetTurn sets Turn field to given value.

### HasTurn

`func (o *TaskReply) HasTurn() bool`

HasTurn returns a boolean if a field has been set.

### GetUser

`func (o *TaskReply) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TaskReply) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TaskReply) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *TaskReply) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


