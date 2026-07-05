# TrackerUpdateIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TrackerIssueStatus**](TrackerIssueStatus.md) |  | [optional] 
**Priority** | Pointer to [**TrackerIssuePriority**](TrackerIssuePriority.md) |  | [optional] 
**Assignee** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTrackerUpdateIssueRequest

`func NewTrackerUpdateIssueRequest() *TrackerUpdateIssueRequest`

NewTrackerUpdateIssueRequest instantiates a new TrackerUpdateIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackerUpdateIssueRequestWithDefaults

`func NewTrackerUpdateIssueRequestWithDefaults() *TrackerUpdateIssueRequest`

NewTrackerUpdateIssueRequestWithDefaults instantiates a new TrackerUpdateIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TrackerUpdateIssueRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrackerUpdateIssueRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrackerUpdateIssueRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrackerUpdateIssueRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TrackerUpdateIssueRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrackerUpdateIssueRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrackerUpdateIssueRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrackerUpdateIssueRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *TrackerUpdateIssueRequest) GetStatus() TrackerIssueStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrackerUpdateIssueRequest) GetStatusOk() (*TrackerIssueStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrackerUpdateIssueRequest) SetStatus(v TrackerIssueStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrackerUpdateIssueRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *TrackerUpdateIssueRequest) GetPriority() TrackerIssuePriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TrackerUpdateIssueRequest) GetPriorityOk() (*TrackerIssuePriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TrackerUpdateIssueRequest) SetPriority(v TrackerIssuePriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TrackerUpdateIssueRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssignee

`func (o *TrackerUpdateIssueRequest) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TrackerUpdateIssueRequest) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TrackerUpdateIssueRequest) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *TrackerUpdateIssueRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetLabels

`func (o *TrackerUpdateIssueRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TrackerUpdateIssueRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TrackerUpdateIssueRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TrackerUpdateIssueRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


