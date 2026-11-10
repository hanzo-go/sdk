# TaskComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is who wrote it: their name, else their forge login. | [optional] 
**Body** | Pointer to **string** | Body is what was written, as markdown. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when it was written, RFC 3339. | [optional] 
**Id** | Pointer to **int64** | ID is the forge&#39;s id for the comment. | [optional] 

## Methods

### NewTaskComment

`func NewTaskComment() *TaskComment`

NewTaskComment instantiates a new TaskComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCommentWithDefaults

`func NewTaskCommentWithDefaults() *TaskComment`

NewTaskCommentWithDefaults instantiates a new TaskComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *TaskComment) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TaskComment) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TaskComment) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *TaskComment) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBody

`func (o *TaskComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TaskComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TaskComment) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TaskComment) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskComment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskComment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskComment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskComment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *TaskComment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskComment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskComment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskComment) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


