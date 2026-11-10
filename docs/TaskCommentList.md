# TaskCommentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TaskComment**](TaskComment.md) | Data is the conversation. Empty is an empty array, never null. | [optional] 
**Home** | Pointer to **string** | Home is the GitHub repository (owner/name) the issue lives on when its board mirrors one; its conversation is read there, and Data is then empty. | [optional] 

## Methods

### NewTaskCommentList

`func NewTaskCommentList() *TaskCommentList`

NewTaskCommentList instantiates a new TaskCommentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCommentListWithDefaults

`func NewTaskCommentListWithDefaults() *TaskCommentList`

NewTaskCommentListWithDefaults instantiates a new TaskCommentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TaskCommentList) GetData() []TaskComment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaskCommentList) GetDataOk() (*[]TaskComment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaskCommentList) SetData(v []TaskComment)`

SetData sets Data field to given value.

### HasData

`func (o *TaskCommentList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHome

`func (o *TaskCommentList) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *TaskCommentList) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *TaskCommentList) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *TaskCommentList) HasHome() bool`

HasHome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


