# TrackerIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** | &#39;KEY-&lt;number&gt;, the human handle&#39; | [optional] 
**ProjectKey** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TrackerIssueStatus**](TrackerIssueStatus.md) |  | [optional] 
**Priority** | Pointer to [**TrackerIssuePriority**](TrackerIssuePriority.md) |  | [optional] 
**Assignee** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewTrackerIssue

`func NewTrackerIssue() *TrackerIssue`

NewTrackerIssue instantiates a new TrackerIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackerIssueWithDefaults

`func NewTrackerIssueWithDefaults() *TrackerIssue`

NewTrackerIssueWithDefaults instantiates a new TrackerIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackerIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackerIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackerIssue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackerIssue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *TrackerIssue) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TrackerIssue) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TrackerIssue) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *TrackerIssue) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetProjectKey

`func (o *TrackerIssue) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *TrackerIssue) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *TrackerIssue) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *TrackerIssue) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetNumber

`func (o *TrackerIssue) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TrackerIssue) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TrackerIssue) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *TrackerIssue) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetTitle

`func (o *TrackerIssue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrackerIssue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrackerIssue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrackerIssue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TrackerIssue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrackerIssue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrackerIssue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrackerIssue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *TrackerIssue) GetStatus() TrackerIssueStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrackerIssue) GetStatusOk() (*TrackerIssueStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrackerIssue) SetStatus(v TrackerIssueStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrackerIssue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *TrackerIssue) GetPriority() TrackerIssuePriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TrackerIssue) GetPriorityOk() (*TrackerIssuePriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TrackerIssue) SetPriority(v TrackerIssuePriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TrackerIssue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssignee

`func (o *TrackerIssue) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TrackerIssue) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TrackerIssue) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *TrackerIssue) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetLabels

`func (o *TrackerIssue) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TrackerIssue) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TrackerIssue) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TrackerIssue) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TrackerIssue) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrackerIssue) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrackerIssue) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrackerIssue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TrackerIssue) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TrackerIssue) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TrackerIssue) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TrackerIssue) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


