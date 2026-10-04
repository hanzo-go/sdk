# TeamTeamMessageEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the message, from the path. | [optional] 
**Space** | Pointer to **string** | Space names the space holding it. Body-only. | [optional] 
**Text** | Pointer to **string** | Text is the message&#39;s new text, with the same &#x60;&lt;@account-uuid&gt;&#x60; mentions a new message takes. Somebody newly mentioned is notified; somebody already mentioned is not notified twice. | [optional] 

## Methods

### NewTeamTeamMessageEdit

`func NewTeamTeamMessageEdit() *TeamTeamMessageEdit`

NewTeamTeamMessageEdit instantiates a new TeamTeamMessageEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamMessageEditWithDefaults

`func NewTeamTeamMessageEditWithDefaults() *TeamTeamMessageEdit`

NewTeamTeamMessageEditWithDefaults instantiates a new TeamTeamMessageEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamTeamMessageEdit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamMessageEdit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamMessageEdit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamMessageEdit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamMessageEdit) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamMessageEdit) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamMessageEdit) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamMessageEdit) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetText

`func (o *TeamTeamMessageEdit) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TeamTeamMessageEdit) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TeamTeamMessageEdit) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TeamTeamMessageEdit) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


