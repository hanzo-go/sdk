# TaskBoardView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the board was created, in unix seconds. 0 on a forge board for the same reason Description is absent. | [optional] 
**Description** | Pointer to **string** | Description is whatever an index board was created with. Absent on a forge board: this projection takes the repository&#39;s name and nothing else about the repository. | [optional] 
**Id** | Pointer to **string** | ID is the board&#39;s opaque handle, and it is NOT how you address it — Key is. Its shape says which source answered: a forge board&#39;s is the repository&#39;s full name (\&quot;hanzoai/cloud\&quot;), an index board&#39;s a minted \&quot;prj_\&quot; id. | [optional] 
**Key** | Pointer to **string** | Key addresses the board everywhere else — /v1/task/projects/&lt;key&gt;/issues — and prefixes every issue identifier on it. An index board&#39;s key is 2-8 uppercase alphanumerics starting with a letter (\&quot;ENG\&quot;, \&quot;OPS2\&quot;) and is matched case-insensitively; a forge board&#39;s is the repository name as the forge spells it. | [optional] 
**Name** | Pointer to **string** | Name is the board&#39;s display name. For a forge board it is the repository name, so it equals Key; an index board carries its own. | [optional] 
**Org** | Pointer to **string** | Org is the IAM org the board belongs to, taken from the validated principal and never from the request. Every board a caller can see is in it. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the board record last changed, in unix seconds — the BOARD, not the work on it, so filing an issue does not move it. 0 on a forge board. | [optional] 

## Methods

### NewTaskBoardView

`func NewTaskBoardView() *TaskBoardView`

NewTaskBoardView instantiates a new TaskBoardView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskBoardViewWithDefaults

`func NewTaskBoardViewWithDefaults() *TaskBoardView`

NewTaskBoardViewWithDefaults instantiates a new TaskBoardView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TaskBoardView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskBoardView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskBoardView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskBoardView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *TaskBoardView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskBoardView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskBoardView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskBoardView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TaskBoardView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskBoardView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskBoardView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskBoardView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *TaskBoardView) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaskBoardView) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaskBoardView) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TaskBoardView) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *TaskBoardView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskBoardView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskBoardView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskBoardView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *TaskBoardView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TaskBoardView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TaskBoardView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TaskBoardView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskBoardView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskBoardView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskBoardView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskBoardView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


