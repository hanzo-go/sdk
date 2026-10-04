# TeamTeamMessageWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]TeamTeamFileIn**](TeamTeamFileIn.md) | Files attaches blobs already uploaded to the space. A message may be files alone, with no text. | [optional] 
**Id** | Pointer to **string** | ID is the room to say it in, from the path. | [optional] 
**Space** | Pointer to **string** | Space names the space holding the room. Body-only: a query string may not redirect a write. | [optional] 
**Text** | Pointer to **string** | Text is what to say, as plain text. It is wrapped in the client&#39;s markup on the way in, so a caller writes words rather than HTML. &#x60;&lt;@account-uuid&gt;&#x60; mentions a member of the space: it is stored as the platform&#39;s mention, so they are notified, and an agent mentioned this way answers. | [optional] 

## Methods

### NewTeamTeamMessageWrite

`func NewTeamTeamMessageWrite() *TeamTeamMessageWrite`

NewTeamTeamMessageWrite instantiates a new TeamTeamMessageWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamMessageWriteWithDefaults

`func NewTeamTeamMessageWriteWithDefaults() *TeamTeamMessageWrite`

NewTeamTeamMessageWriteWithDefaults instantiates a new TeamTeamMessageWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *TeamTeamMessageWrite) GetFiles() []TeamTeamFileIn`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TeamTeamMessageWrite) GetFilesOk() (*[]TeamTeamFileIn, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TeamTeamMessageWrite) SetFiles(v []TeamTeamFileIn)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TeamTeamMessageWrite) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamMessageWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamMessageWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamMessageWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamMessageWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamMessageWrite) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamMessageWrite) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamMessageWrite) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamMessageWrite) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetText

`func (o *TeamTeamMessageWrite) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TeamTeamMessageWrite) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TeamTeamMessageWrite) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TeamTeamMessageWrite) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


