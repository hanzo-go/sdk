# TaskIssueHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee is who holds the work. EMPTY MEANS UNHELD, which is what makes the issue claimable: claiming one already held by someone else is refused with 409 rather than quietly taken. | [optional] 
**Kind** | Pointer to **string** | Kind is what the row IS: issue, pr or epic. | [optional] 
**Number** | Pointer to **int64** | Number is the issue&#39;s number on that board, from 1 and monotonic there. Unique per board, never across the org — so it addresses an issue only together with Project. | [optional] 
**Priority** | Pointer to **string** | Priority is urgent, high, medium, low or none. Never empty — an unset priority is the value \&quot;none\&quot;. | [optional] 
**Project** | Pointer to **string** | Project is the board key the issue is on. It and Number are the issue&#39;s address in every other route on this surface, which is why a hit carries it. | [optional] 
**Repo** | Pointer to **string** | Repo is the git repository the issue is bound to, empty when it is not repo-bound. | [optional] 
**Room** | Pointer to **string** | Room is the collaboration room the issue belongs to, spelled \&quot;&lt;space&gt;_&lt;room&gt;\&quot; — empty when it is not room-bound, which is most of them. It is here so an org-wide search says which channel each item came from without a second read. | [optional] 
**Source** | Pointer to **string** | Source is which surface opened it: team, git, crm, helpdesk, cms or agent. \&quot;git\&quot; is how the mirrored forge and GitHub rows are spelled. | [optional] 
**Status** | Pointer to **string** | Status is the board column: backlog, todo, in_progress, done or canceled. Claiming moves backlog and todo to in_progress and leaves the other three where they are. | [optional] 
**Title** | Pointer to **string** | Title is the issue&#39;s one-line summary — what the q filter matched, along with the description. | [optional] 
**Url** | Pointer to **string** | URL is the row&#39;s external anchor — its extRef — which is a link only when the feeder sent one. A mirrored GitHub issue carries \&quot;github:owner/repo#123\&quot; and an agent&#39;s PR row carries the pushed branch. Empty for a row opened here. | [optional] 

## Methods

### NewTaskIssueHit

`func NewTaskIssueHit() *TaskIssueHit`

NewTaskIssueHit instantiates a new TaskIssueHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskIssueHitWithDefaults

`func NewTaskIssueHitWithDefaults() *TaskIssueHit`

NewTaskIssueHitWithDefaults instantiates a new TaskIssueHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *TaskIssueHit) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TaskIssueHit) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TaskIssueHit) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *TaskIssueHit) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetKind

`func (o *TaskIssueHit) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TaskIssueHit) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TaskIssueHit) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TaskIssueHit) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetNumber

`func (o *TaskIssueHit) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TaskIssueHit) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TaskIssueHit) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *TaskIssueHit) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPriority

`func (o *TaskIssueHit) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskIssueHit) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskIssueHit) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskIssueHit) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProject

`func (o *TaskIssueHit) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TaskIssueHit) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TaskIssueHit) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *TaskIssueHit) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepo

`func (o *TaskIssueHit) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *TaskIssueHit) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *TaskIssueHit) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *TaskIssueHit) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRoom

`func (o *TaskIssueHit) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *TaskIssueHit) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *TaskIssueHit) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *TaskIssueHit) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetSource

`func (o *TaskIssueHit) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TaskIssueHit) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TaskIssueHit) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TaskIssueHit) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *TaskIssueHit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskIssueHit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskIssueHit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskIssueHit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *TaskIssueHit) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskIssueHit) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskIssueHit) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskIssueHit) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *TaskIssueHit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskIssueHit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskIssueHit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskIssueHit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


