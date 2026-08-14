# IssueEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description rewrites the body. | [optional] 
**Key** | Pointer to **string** | Key is the board — the repository name, from the path. | [optional] 
**Num** | Pointer to **int32** | Num is the issue number on that repository, from the path. | [optional] 
**Priority** | Pointer to **string** | Priority re-prioritises it. | [optional] 
**Status** | Pointer to **string** | Status moves the card to another column. | [optional] 
**Title** | Pointer to **string** | Title renames the work item. | [optional] 

## Methods

### NewIssueEdit

`func NewIssueEdit() *IssueEdit`

NewIssueEdit instantiates a new IssueEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEditWithDefaults

`func NewIssueEditWithDefaults() *IssueEdit`

NewIssueEditWithDefaults instantiates a new IssueEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IssueEdit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueEdit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueEdit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueEdit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *IssueEdit) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueEdit) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueEdit) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IssueEdit) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNum

`func (o *IssueEdit) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *IssueEdit) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *IssueEdit) SetNum(v int32)`

SetNum sets Num field to given value.

### HasNum

`func (o *IssueEdit) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetPriority

`func (o *IssueEdit) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IssueEdit) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IssueEdit) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IssueEdit) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *IssueEdit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IssueEdit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IssueEdit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IssueEdit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *IssueEdit) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssueEdit) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssueEdit) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IssueEdit) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


