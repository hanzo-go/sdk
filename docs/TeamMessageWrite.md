# TeamMessageWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the room to say it in, from the path. | [optional] 
**Space** | Pointer to **string** | Space names the space holding the room. Body-only: a query string may not redirect a write. | [optional] 
**Text** | Pointer to **string** | Text is what to say, as plain text. It is wrapped in the client&#39;s markup on the way in, so a caller writes words rather than HTML. | [optional] 

## Methods

### NewTeamMessageWrite

`func NewTeamMessageWrite() *TeamMessageWrite`

NewTeamMessageWrite instantiates a new TeamMessageWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMessageWriteWithDefaults

`func NewTeamMessageWriteWithDefaults() *TeamMessageWrite`

NewTeamMessageWriteWithDefaults instantiates a new TeamMessageWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamMessageWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamMessageWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamMessageWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamMessageWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpace

`func (o *TeamMessageWrite) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamMessageWrite) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamMessageWrite) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamMessageWrite) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetText

`func (o *TeamMessageWrite) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TeamMessageWrite) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TeamMessageWrite) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TeamMessageWrite) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


