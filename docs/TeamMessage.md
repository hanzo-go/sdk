# TeamMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is the team account uuid that wrote it. It is an ACCOUNT and not a display name: what to call somebody is the roster&#39;s answer, and copying it onto every message is how the two come to disagree. An agent&#39;s messages carry the account derived from its id, so the same field answers for both. | [optional] 
**CreatedOn** | Pointer to **int64** | CreatedOn is unix MILLIseconds, which is what the platform stamps. | [optional] 
**Id** | Pointer to **string** | ID is the message document&#39;s own id. | [optional] 
**Room** | Pointer to **string** | Room is the room it was said in — the same id the room listing answers with, so a caller holding a message can name its room without a second read. | [optional] 
**Text** | Pointer to **string** | Text is the message as PLAIN TEXT. The document stores markup; this is the same &#x60;plainText&#x60; reduction the agent responder reads a prompt with, so a caller never has to parse the client&#39;s markup to know what was said. | [optional] 

## Methods

### NewTeamMessage

`func NewTeamMessage() *TeamMessage`

NewTeamMessage instantiates a new TeamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMessageWithDefaults

`func NewTeamMessageWithDefaults() *TeamMessage`

NewTeamMessageWithDefaults instantiates a new TeamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *TeamMessage) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TeamMessage) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TeamMessage) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *TeamMessage) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreatedOn

`func (o *TeamMessage) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TeamMessage) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TeamMessage) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TeamMessage) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetId

`func (o *TeamMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoom

`func (o *TeamMessage) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *TeamMessage) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *TeamMessage) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *TeamMessage) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetText

`func (o *TeamMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TeamMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TeamMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TeamMessage) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


