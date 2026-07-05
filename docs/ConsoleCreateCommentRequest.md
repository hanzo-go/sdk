# ConsoleCreateCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**ObjectType** | **string** |  | 
**ObjectId** | **string** |  | 
**Content** | **string** |  | 
**AuthorUserId** | Pointer to **string** |  | [optional] 

## Methods

### NewConsoleCreateCommentRequest

`func NewConsoleCreateCommentRequest(projectId string, objectType string, objectId string, content string, ) *ConsoleCreateCommentRequest`

NewConsoleCreateCommentRequest instantiates a new ConsoleCreateCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreateCommentRequestWithDefaults

`func NewConsoleCreateCommentRequestWithDefaults() *ConsoleCreateCommentRequest`

NewConsoleCreateCommentRequestWithDefaults instantiates a new ConsoleCreateCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ConsoleCreateCommentRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ConsoleCreateCommentRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ConsoleCreateCommentRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetObjectType

`func (o *ConsoleCreateCommentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConsoleCreateCommentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConsoleCreateCommentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ConsoleCreateCommentRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ConsoleCreateCommentRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ConsoleCreateCommentRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetContent

`func (o *ConsoleCreateCommentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConsoleCreateCommentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConsoleCreateCommentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetAuthorUserId

`func (o *ConsoleCreateCommentRequest) GetAuthorUserId() string`

GetAuthorUserId returns the AuthorUserId field if non-nil, zero value otherwise.

### GetAuthorUserIdOk

`func (o *ConsoleCreateCommentRequest) GetAuthorUserIdOk() (*string, bool)`

GetAuthorUserIdOk returns a tuple with the AuthorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUserId

`func (o *ConsoleCreateCommentRequest) SetAuthorUserId(v string)`

SetAuthorUserId sets AuthorUserId field to given value.

### HasAuthorUserId

`func (o *ConsoleCreateCommentRequest) HasAuthorUserId() bool`

HasAuthorUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


