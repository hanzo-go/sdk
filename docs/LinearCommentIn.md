# LinearCommentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the comment, Markdown. | [optional] 
**Issue** | Pointer to **string** | Issue is the issue&#39;s identifier (ENG-123) or its id. | [optional] 

## Methods

### NewLinearCommentIn

`func NewLinearCommentIn() *LinearCommentIn`

NewLinearCommentIn instantiates a new LinearCommentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinearCommentInWithDefaults

`func NewLinearCommentInWithDefaults() *LinearCommentIn`

NewLinearCommentInWithDefaults instantiates a new LinearCommentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *LinearCommentIn) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LinearCommentIn) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LinearCommentIn) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *LinearCommentIn) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetIssue

`func (o *LinearCommentIn) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *LinearCommentIn) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *LinearCommentIn) SetIssue(v string)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *LinearCommentIn) HasIssue() bool`

HasIssue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


