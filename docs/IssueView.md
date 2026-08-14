# IssueView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DueAt** | Pointer to **int32** | unix seconds; absent &#x3D; no due date | [optional] 
**ExtRef** | Pointer to **string** | external anchor | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** | KEY-&lt;number&gt;, the human handle | [optional] 
**Kind** | Pointer to **string** | issue | pr | epic | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**ProjectKey** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** | git repo binding | [optional] 
**Source** | Pointer to **string** | team | git | crm | helpdesk | cms | agent | [optional] 
**StartAt** | Pointer to **int32** | unix seconds; absent &#x3D; unscheduled | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

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


