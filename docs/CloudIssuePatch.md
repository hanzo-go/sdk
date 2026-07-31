# CloudIssuePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee is who owns the issue, at most 256 characters. Empty unassigns it. | [optional] 
**Description** | Pointer to **string** | Description is the issue body, at most 32768 characters. | [optional] 
**Key** | Pointer to **string** | Key is the issue&#39;s project, from the path. | [optional] 
**Labels** | Pointer to **[]string** | Labels REPLACES the issue&#39;s labels with exactly this set. Each label is at most 48 characters and may not contain a comma (the storage separator); empty entries are dropped. | [optional] 
**Num** | Pointer to **int32** | Num is the issue&#39;s number within that project, from the path. | [optional] 
**Priority** | Pointer to **string** | Priority is none, urgent, high, medium or low. Empty resets it to none. | [optional] 
**Status** | Pointer to **string** | Status moves the issue between board columns: backlog, todo, in_progress, done or canceled. Empty resets it to backlog. | [optional] 
**Title** | Pointer to **string** | Title is the issue&#39;s one-line summary. Non-empty, at most 512 characters. | [optional] 

## Methods

### NewCloudIssuePatch

`func NewCloudIssuePatch() *CloudIssuePatch`

NewCloudIssuePatch instantiates a new CloudIssuePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIssuePatchWithDefaults

`func NewCloudIssuePatchWithDefaults() *CloudIssuePatch`

NewCloudIssuePatchWithDefaults instantiates a new CloudIssuePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *CloudIssuePatch) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *CloudIssuePatch) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *CloudIssuePatch) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *CloudIssuePatch) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetDescription

`func (o *CloudIssuePatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudIssuePatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudIssuePatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudIssuePatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *CloudIssuePatch) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudIssuePatch) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudIssuePatch) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudIssuePatch) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabels

`func (o *CloudIssuePatch) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CloudIssuePatch) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CloudIssuePatch) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CloudIssuePatch) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNum

`func (o *CloudIssuePatch) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *CloudIssuePatch) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *CloudIssuePatch) SetNum(v int32)`

SetNum sets Num field to given value.

### HasNum

`func (o *CloudIssuePatch) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetPriority

`func (o *CloudIssuePatch) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CloudIssuePatch) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CloudIssuePatch) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CloudIssuePatch) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *CloudIssuePatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudIssuePatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudIssuePatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudIssuePatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CloudIssuePatch) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudIssuePatch) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudIssuePatch) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudIssuePatch) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


