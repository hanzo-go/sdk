# TaskIssueEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee hands the work to somebody — a person or an agent, by the name they are known by on the forge. \&quot;\&quot; TAKES IT OFF whoever holds it, which is why this is a pointer: absent leaves the holder alone.  It is the other half of &#x60;claim&#x60;, which that handler already named: a claim takes work for the CALLER and refuses to name anyone else, because giving work away is a different act with different authority. This is that act, and until it existed a board could only be worked by whoever clicked first — an agent could never be given anything.  &#x60;me&#x60; is you, by the GitHub account IAM links to you (or your forge login). &#x60;agent&#x60; hands the issue to the coding agent: it is labelled &#x60;agent&#x60;, and a coding run starts on its repository with the issue as the task, narrating in the thread the hand-off was asked in; its pull request closes the issue. Labelling an issue &#x60;agent&#x60; on GitHub does the same. | [optional] 
**Comment** | Pointer to **string** | Comment is posted on the issue, after every other change — what a person in the issue&#39;s thread wants said on it. | [optional] 
**Description** | Pointer to **string** | Description rewrites the body. | [optional] 
**Key** | Pointer to **string** | Key is the board — the repository name, from the path. | [optional] 
**Num** | Pointer to **int64** | Num is the issue number on that repository, from the path. | [optional] 
**Priority** | Pointer to **string** | Priority re-prioritises it. | [optional] 
**Reply** | Pointer to [**TaskReply**](TaskReply.md) | Reply is the conversation the edit was asked for in, written by a chat turn and never chosen (see newIssue). An issue bound to no thread is bound to this one, and a run the edit starts narrates here. | [optional] 
**Status** | Pointer to **string** | Status moves the card to another column. | [optional] 
**Title** | Pointer to **string** | Title renames the work item. | [optional] 

## Methods

### NewTaskIssueEdit

`func NewTaskIssueEdit() *TaskIssueEdit`

NewTaskIssueEdit instantiates a new TaskIssueEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskIssueEditWithDefaults

`func NewTaskIssueEditWithDefaults() *TaskIssueEdit`

NewTaskIssueEditWithDefaults instantiates a new TaskIssueEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *TaskIssueEdit) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TaskIssueEdit) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TaskIssueEdit) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *TaskIssueEdit) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetComment

`func (o *TaskIssueEdit) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TaskIssueEdit) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TaskIssueEdit) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TaskIssueEdit) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *TaskIssueEdit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskIssueEdit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskIssueEdit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskIssueEdit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *TaskIssueEdit) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaskIssueEdit) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaskIssueEdit) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TaskIssueEdit) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNum

`func (o *TaskIssueEdit) GetNum() int64`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *TaskIssueEdit) GetNumOk() (*int64, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *TaskIssueEdit) SetNum(v int64)`

SetNum sets Num field to given value.

### HasNum

`func (o *TaskIssueEdit) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetPriority

`func (o *TaskIssueEdit) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskIssueEdit) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskIssueEdit) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskIssueEdit) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReply

`func (o *TaskIssueEdit) GetReply() TaskReply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *TaskIssueEdit) GetReplyOk() (*TaskReply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *TaskIssueEdit) SetReply(v TaskReply)`

SetReply sets Reply field to given value.

### HasReply

`func (o *TaskIssueEdit) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetStatus

`func (o *TaskIssueEdit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskIssueEdit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskIssueEdit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskIssueEdit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *TaskIssueEdit) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskIssueEdit) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskIssueEdit) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskIssueEdit) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


