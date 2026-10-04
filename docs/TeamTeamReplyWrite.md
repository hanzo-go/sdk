# TeamTeamReplyWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]TeamTeamFileIn**](TeamTeamFileIn.md) | Files attaches blobs already uploaded to the space. | [optional] 
**Id** | Pointer to **string** | ID is the message being answered, from the path. | [optional] 
**Space** | Pointer to **string** | Space names the space holding it. Body-only. | [optional] 
**Text** | Pointer to **string** | Text is the reply, with the same &#x60;&lt;@account-uuid&gt;&#x60; mentions a message takes. | [optional] 

## Methods

### NewTeamTeamReplyWrite

`func NewTeamTeamReplyWrite() *TeamTeamReplyWrite`

NewTeamTeamReplyWrite instantiates a new TeamTeamReplyWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamReplyWriteWithDefaults

`func NewTeamTeamReplyWriteWithDefaults() *TeamTeamReplyWrite`

NewTeamTeamReplyWriteWithDefaults instantiates a new TeamTeamReplyWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *TeamTeamReplyWrite) GetFiles() []TeamTeamFileIn`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TeamTeamReplyWrite) GetFilesOk() (*[]TeamTeamFileIn, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TeamTeamReplyWrite) SetFiles(v []TeamTeamFileIn)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TeamTeamReplyWrite) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamReplyWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamReplyWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamReplyWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamReplyWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamReplyWrite) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamReplyWrite) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamReplyWrite) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamReplyWrite) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetText

`func (o *TeamTeamReplyWrite) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TeamTeamReplyWrite) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TeamTeamReplyWrite) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TeamTeamReplyWrite) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


