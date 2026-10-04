# TeamTeamCommentWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]TeamTeamFileIn**](TeamTeamFileIn.md) | Files attaches blobs already uploaded to the space. | [optional] 
**Id** | Pointer to **string** | ID is the document, from the path. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid holding it. Body-only. | [optional] 
**Text** | Pointer to **string** | Text is the comment, with the same &#x60;&lt;@account-uuid&gt;&#x60; mentions a message takes. | [optional] 

## Methods

### NewTeamTeamCommentWrite

`func NewTeamTeamCommentWrite() *TeamTeamCommentWrite`

NewTeamTeamCommentWrite instantiates a new TeamTeamCommentWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamCommentWriteWithDefaults

`func NewTeamTeamCommentWriteWithDefaults() *TeamTeamCommentWrite`

NewTeamTeamCommentWriteWithDefaults instantiates a new TeamTeamCommentWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *TeamTeamCommentWrite) GetFiles() []TeamTeamFileIn`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TeamTeamCommentWrite) GetFilesOk() (*[]TeamTeamFileIn, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TeamTeamCommentWrite) SetFiles(v []TeamTeamFileIn)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TeamTeamCommentWrite) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamCommentWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamCommentWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamCommentWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamCommentWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamCommentWrite) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamCommentWrite) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamCommentWrite) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamCommentWrite) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetText

`func (o *TeamTeamCommentWrite) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TeamTeamCommentWrite) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TeamTeamCommentWrite) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TeamTeamCommentWrite) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


