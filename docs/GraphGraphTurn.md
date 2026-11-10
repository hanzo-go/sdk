# GraphGraphTurn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when this turn was spoken, RFC 3339. Absent defaults to the conversation&#39;s At instant. | [optional] 
**Id** | Pointer to **string** | ID is the turn&#39;s identifier within the conversation, appended to Source as &#x60;#&lt;id&gt;&#x60; in Evidence. Absent defaults to its 1-based position. | [optional] 
**Partner** | Pointer to **string** | Partner is the entity key of the participant addressed in this turn. Second-person references (you, your) resolve to Partner. | [optional] 
**Speaker** | **string** | Speaker is the entity key of the participant speaking this turn. First- person references (I, me, my) resolve to Speaker. Required. | 
**Text** | **string** | Text is what the speaker said. Required. | 

## Methods

### NewGraphGraphTurn

`func NewGraphGraphTurn(speaker string, text string, ) *GraphGraphTurn`

NewGraphGraphTurn instantiates a new GraphGraphTurn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphTurnWithDefaults

`func NewGraphGraphTurnWithDefaults() *GraphGraphTurn`

NewGraphGraphTurnWithDefaults instantiates a new GraphGraphTurn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *GraphGraphTurn) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *GraphGraphTurn) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *GraphGraphTurn) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *GraphGraphTurn) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetId

`func (o *GraphGraphTurn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphGraphTurn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphGraphTurn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GraphGraphTurn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartner

`func (o *GraphGraphTurn) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *GraphGraphTurn) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *GraphGraphTurn) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *GraphGraphTurn) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetSpeaker

`func (o *GraphGraphTurn) GetSpeaker() string`

GetSpeaker returns the Speaker field if non-nil, zero value otherwise.

### GetSpeakerOk

`func (o *GraphGraphTurn) GetSpeakerOk() (*string, bool)`

GetSpeakerOk returns a tuple with the Speaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeaker

`func (o *GraphGraphTurn) SetSpeaker(v string)`

SetSpeaker sets Speaker field to given value.


### GetText

`func (o *GraphGraphTurn) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GraphGraphTurn) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GraphGraphTurn) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


