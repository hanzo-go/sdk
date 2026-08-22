# IssueView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee is who holds the work — an IAM username, or the login of the FIRST assignee when a forge issue has several. Absent when nobody holds it, which is exactly the state a claim needs. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the item was opened, in unix seconds. 0 when the source gave no parseable timestamp. | [optional] 
**Description** | Pointer to **string** | Description is the body, markdown as its author wrote it. Absent when empty. | [optional] 
**DueAt** | Pointer to **int32** | DueAt is when the work is due, in unix seconds; absent means no due date. A forge row takes it from its MILESTONE&#39;s due date, since a forge issue has no deadline of its own. Never before StartAt, and never past 2200-01-01. | [optional] 
**ExtRef** | Pointer to **string** | ExtRef anchors the item to something outside the todo — a mirrored issue (\&quot;github:owner/repo#123\&quot;), a pushed PR branch, or a record on another plane. It is the idempotency key the mirror upsert matches on. Absent when the item has no external origin. | [optional] 
**Id** | Pointer to **string** | ID is the work item&#39;s opaque handle, and it is NOT how you address it — ProjectKey plus Number is. Its shape says which source answered: a forge issue&#39;s is the forge&#39;s own numeric id in decimal, an index row&#39;s a minted \&quot;issue_\&quot; id. | [optional] 
**Identifier** | Pointer to **string** | Identifier is the human handle, \&quot;&lt;key&gt;#&lt;number&gt;\&quot; — the board and the number on it, joined. ONE spelling whichever source answered, because a list where forge rows read cli#1 and index rows read OPS-3 is two products in one list. | [optional] 
**Kind** | Pointer to **string** | Kind is what the item IS: issue, pr or epic. Set once at create and never changed, so a row does not migrate between surfaces. Deliberately not \&quot;task\&quot; — that word is the async plane (contract.go). | [optional] 
**Labels** | Pointer to **[]string** | Labels are the item&#39;s remaining tags, with the status and priority labels lifted OUT — a column that stayed here would render twice, once as the card&#39;s column and once as a chip on the card. Always present; empty is []. | [optional] 
**Number** | Pointer to **int32** | Number is the item&#39;s number ON ITS BOARD, from 1 and monotonic there — the forge&#39;s own issue number for a forge row, allocated inside the create transaction for an index row so it cannot race. Unique per board, never across the org. | [optional] 
**Priority** | Pointer to **string** | Priority is urgent, high, medium, low or none. Also a label on a forge row. Never empty: \&quot;none\&quot; when nothing names one, so callers compare a value rather than test for absence. | [optional] 
**ProjectKey** | Pointer to **string** | ProjectKey is the board this item is on: the repository name for a forge issue, the index board&#39;s key otherwise. With Number it is the item&#39;s address in every other route. | [optional] 
**Repo** | Pointer to **string** | Repo is the git repository the item is bound to, so a repository&#39;s Issues and PRs tabs are filters over this one table. Absent when the item is not repo-bound. | [optional] 
**Source** | Pointer to **string** | Source is which surface OPENED it: team, git, crm, helpdesk, cms or agent. Also set once. It is the ORIGIN, not the subject — source&#x3D;helpdesk is an engineering issue opened from a support escalation, not a support ticket. | [optional] 
**StartAt** | Pointer to **int32** | StartAt is when the work starts, in unix seconds; absent means unscheduled. A forge row takes it from when the issue was opened, but only once the issue has a due date — an interval needs both ends. | [optional] 
**Status** | Pointer to **string** | Status is the board column: backlog, todo, in_progress, done or canceled, and nothing else. On a forge row it is read off a LABEL, so relabelling in the forge web UI moves the card here and vice versa — and a CLOSED forge issue reads done whatever its labels say. Never empty: \&quot;backlog\&quot; when nothing names a column. | [optional] 
**Title** | Pointer to **string** | Title is the item&#39;s one-line summary. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when it last changed, in unix seconds. | [optional] 

## Methods

### NewIssueView

`func NewIssueView() *IssueView`

NewIssueView instantiates a new IssueView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueViewWithDefaults

`func NewIssueViewWithDefaults() *IssueView`

NewIssueViewWithDefaults instantiates a new IssueView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *IssueView) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *IssueView) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *IssueView) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *IssueView) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssueView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *IssueView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueAt

`func (o *IssueView) GetDueAt() int32`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *IssueView) GetDueAtOk() (*int32, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *IssueView) SetDueAt(v int32)`

SetDueAt sets DueAt field to given value.

### HasDueAt

`func (o *IssueView) HasDueAt() bool`

HasDueAt returns a boolean if a field has been set.

### GetExtRef

`func (o *IssueView) GetExtRef() string`

GetExtRef returns the ExtRef field if non-nil, zero value otherwise.

### GetExtRefOk

`func (o *IssueView) GetExtRefOk() (*string, bool)`

GetExtRefOk returns a tuple with the ExtRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtRef

`func (o *IssueView) SetExtRef(v string)`

SetExtRef sets ExtRef field to given value.

### HasExtRef

`func (o *IssueView) HasExtRef() bool`

HasExtRef returns a boolean if a field has been set.

### GetId

`func (o *IssueView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *IssueView) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IssueView) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IssueView) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *IssueView) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetKind

`func (o *IssueView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IssueView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IssueView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IssueView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabels

`func (o *IssueView) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IssueView) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IssueView) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IssueView) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNumber

`func (o *IssueView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IssueView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IssueView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *IssueView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPriority

`func (o *IssueView) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IssueView) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IssueView) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IssueView) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectKey

`func (o *IssueView) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *IssueView) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *IssueView) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *IssueView) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetRepo

`func (o *IssueView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *IssueView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *IssueView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *IssueView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSource

`func (o *IssueView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IssueView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IssueView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IssueView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStartAt

`func (o *IssueView) GetStartAt() int32`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *IssueView) GetStartAtOk() (*int32, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *IssueView) SetStartAt(v int32)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *IssueView) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetStatus

`func (o *IssueView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IssueView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IssueView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IssueView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *IssueView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssueView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssueView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IssueView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


