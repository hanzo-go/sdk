# NewIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description becomes the issue body. | [optional] 
**Key** | Pointer to **string** | Key is the board — the repository name, from the path. | [optional] 
**Priority** | Pointer to **string** | Priority is one of none, urgent, high, medium or low. | [optional] 
**Status** | Pointer to **string** | Status is the board column to open into: backlog, todo, in_progress, done or canceled. Empty opens into backlog. | [optional] 
**Title** | Pointer to **string** | Title is required. | [optional] 

## Methods

### NewNewIssue

`func NewNewIssue() *NewIssue`

NewNewIssue instantiates a new NewIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewIssueWithDefaults

`func NewNewIssueWithDefaults() *NewIssue`

NewNewIssueWithDefaults instantiates a new NewIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NewIssue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewIssue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewIssue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewIssue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *NewIssue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NewIssue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NewIssue) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *NewIssue) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPriority

`func (o *NewIssue) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NewIssue) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NewIssue) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NewIssue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *NewIssue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewIssue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewIssue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewIssue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *NewIssue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewIssue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewIssue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NewIssue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


