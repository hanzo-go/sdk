# TaskNewIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description becomes the issue body. | [optional] 
**Key** | Pointer to **string** | Key is the board — the repository name, from the path. &#x60;owner/name&#x60;, or &#x60;github.com/owner/name&#x60;, names a repository on GitHub that no board on the forge stands for. | [optional] 
**Priority** | Pointer to **string** | Priority is one of none, urgent, high, medium or low. | [optional] 
**Reply** | Pointer to [**TaskReply**](TaskReply.md) | Reply is the conversation the issue was asked for in. A chat turn writes its own thread here, over anything else (apps/agents replyto.go); it is not a value to choose. The issue is bound to it, and what happens to the issue next — comments, closing, who holds it, a pull request that closes it — is said there. | [optional] 
**Status** | Pointer to **string** | Status is the board column to open into: backlog, todo, in_progress, done or canceled. Empty opens into backlog. | [optional] 
**Title** | Pointer to **string** | Title is the one line the card is read by on the board. Blank or whitespace is refused — an untitled card cannot be told apart from any other. | [optional] 

## Methods

### NewTaskNewIssue

`func NewTaskNewIssue() *TaskNewIssue`

NewTaskNewIssue instantiates a new TaskNewIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskNewIssueWithDefaults

`func NewTaskNewIssueWithDefaults() *TaskNewIssue`

NewTaskNewIssueWithDefaults instantiates a new TaskNewIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TaskNewIssue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskNewIssue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskNewIssue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskNewIssue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *TaskNewIssue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaskNewIssue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaskNewIssue) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TaskNewIssue) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPriority

`func (o *TaskNewIssue) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskNewIssue) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskNewIssue) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskNewIssue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReply

`func (o *TaskNewIssue) GetReply() TaskReply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *TaskNewIssue) GetReplyOk() (*TaskReply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *TaskNewIssue) SetReply(v TaskReply)`

SetReply sets Reply field to given value.

### HasReply

`func (o *TaskNewIssue) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetStatus

`func (o *TaskNewIssue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskNewIssue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskNewIssue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskNewIssue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *TaskNewIssue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskNewIssue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskNewIssue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskNewIssue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


