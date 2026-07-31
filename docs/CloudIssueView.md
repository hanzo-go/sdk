# CloudIssueView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
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
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudIssueView

`func NewCloudIssueView() *CloudIssueView`

NewCloudIssueView instantiates a new CloudIssueView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIssueViewWithDefaults

`func NewCloudIssueViewWithDefaults() *CloudIssueView`

NewCloudIssueViewWithDefaults instantiates a new CloudIssueView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *CloudIssueView) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *CloudIssueView) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *CloudIssueView) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *CloudIssueView) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudIssueView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudIssueView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudIssueView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudIssueView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CloudIssueView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudIssueView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudIssueView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudIssueView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtRef

`func (o *CloudIssueView) GetExtRef() string`

GetExtRef returns the ExtRef field if non-nil, zero value otherwise.

### GetExtRefOk

`func (o *CloudIssueView) GetExtRefOk() (*string, bool)`

GetExtRefOk returns a tuple with the ExtRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtRef

`func (o *CloudIssueView) SetExtRef(v string)`

SetExtRef sets ExtRef field to given value.

### HasExtRef

`func (o *CloudIssueView) HasExtRef() bool`

HasExtRef returns a boolean if a field has been set.

### GetId

`func (o *CloudIssueView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudIssueView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudIssueView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudIssueView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *CloudIssueView) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CloudIssueView) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CloudIssueView) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CloudIssueView) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetKind

`func (o *CloudIssueView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudIssueView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudIssueView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudIssueView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabels

`func (o *CloudIssueView) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CloudIssueView) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CloudIssueView) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CloudIssueView) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNumber

`func (o *CloudIssueView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CloudIssueView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CloudIssueView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CloudIssueView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPriority

`func (o *CloudIssueView) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CloudIssueView) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CloudIssueView) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CloudIssueView) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectKey

`func (o *CloudIssueView) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *CloudIssueView) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *CloudIssueView) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *CloudIssueView) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetRepo

`func (o *CloudIssueView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudIssueView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudIssueView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudIssueView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSource

`func (o *CloudIssueView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudIssueView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudIssueView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudIssueView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *CloudIssueView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudIssueView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudIssueView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudIssueView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CloudIssueView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudIssueView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudIssueView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudIssueView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudIssueView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudIssueView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudIssueView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudIssueView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


