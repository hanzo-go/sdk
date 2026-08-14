# AskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is one or two sentences answering the question, every number in it taken from Figures. | [optional] 
**Figures** | Pointer to [**[]Figure**](Figure.md) | Figures are the grounded numbers the answer states, each already formatted. | [optional] 
**Followups** | Pointer to **[]string** | Followups are sharper questions to ask next, chosen from the same intent. | [optional] 
**Sources** | Pointer to **[]string** | Sources name the books reports the figures were computed from — \&quot;pnl\&quot;, \&quot;position\&quot;, \&quot;trial\&quot;. | [optional] 

## Methods

### NewAskResponse

`func NewAskResponse() *AskResponse`

NewAskResponse instantiates a new AskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskResponseWithDefaults

`func NewAskResponseWithDefaults() *AskResponse`

NewAskResponseWithDefaults instantiates a new AskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *AskResponse) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *AskResponse) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *AskResponse) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *AskResponse) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetFigures

`func (o *AskResponse) GetFigures() []Figure`

GetFigures returns the Figures field if non-nil, zero value otherwise.

### GetFiguresOk

`func (o *AskResponse) GetFiguresOk() (*[]Figure, bool)`

GetFiguresOk returns a tuple with the Figures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigures

`func (o *AskResponse) SetFigures(v []Figure)`

SetFigures sets Figures field to given value.

### HasFigures

`func (o *AskResponse) HasFigures() bool`

HasFigures returns a boolean if a field has been set.

### GetFollowups

`func (o *AskResponse) GetFollowups() []string`

GetFollowups returns the Followups field if non-nil, zero value otherwise.

### GetFollowupsOk

`func (o *AskResponse) GetFollowupsOk() (*[]string, bool)`

GetFollowupsOk returns a tuple with the Followups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowups

`func (o *AskResponse) SetFollowups(v []string)`

SetFollowups sets Followups field to given value.

### HasFollowups

`func (o *AskResponse) HasFollowups() bool`

HasFollowups returns a boolean if a field has been set.

### GetSources

`func (o *AskResponse) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AskResponse) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AskResponse) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AskResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


