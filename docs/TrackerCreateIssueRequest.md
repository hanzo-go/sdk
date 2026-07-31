# TrackerCreateIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TrackerIssueStatus**](TrackerIssueStatus.md) |  | [optional] 
**Priority** | Pointer to [**TrackerIssuePriority**](TrackerIssuePriority.md) |  | [optional] 
**Assignee** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTrackerCreateIssueRequest

`func NewTrackerCreateIssueRequest(title string, ) *TrackerCreateIssueRequest`

NewTrackerCreateIssueRequest instantiates a new TrackerCreateIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackerCreateIssueRequestWithDefaults

`func NewTrackerCreateIssueRequestWithDefaults() *TrackerCreateIssueRequest`

NewTrackerCreateIssueRequestWithDefaults instantiates a new TrackerCreateIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TrackerCreateIssueRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrackerCreateIssueRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrackerCreateIssueRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *TrackerCreateIssueRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrackerCreateIssueRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrackerCreateIssueRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrackerCreateIssueRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *TrackerCreateIssueRequest) GetStatus() TrackerIssueStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrackerCreateIssueRequest) GetStatusOk() (*TrackerIssueStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrackerCreateIssueRequest) SetStatus(v TrackerIssueStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrackerCreateIssueRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *TrackerCreateIssueRequest) GetPriority() TrackerIssuePriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TrackerCreateIssueRequest) GetPriorityOk() (*TrackerIssuePriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TrackerCreateIssueRequest) SetPriority(v TrackerIssuePriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TrackerCreateIssueRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssignee

`func (o *TrackerCreateIssueRequest) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TrackerCreateIssueRequest) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TrackerCreateIssueRequest) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *TrackerCreateIssueRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetLabels

`func (o *TrackerCreateIssueRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TrackerCreateIssueRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TrackerCreateIssueRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TrackerCreateIssueRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


