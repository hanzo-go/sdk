# AskReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is the grounded prose, with inline markdown citations. Every link in it points at a page in Sources: the citation check runs on the text before it leaves the engine, so a cited URL is one THIS call fetched. | [optional] 
**FollowUps** | Pointer to **[]string** | FollowUps are the questions worth asking next. Best-effort — an empty list is a normal outcome, not a fault. | [optional] 
**Mode** | Pointer to **string** | Mode is the profile that ran: search, news, research or deep. | [optional] 
**Model** | Pointer to **string** | Model is the model that synthesized the answer. | [optional] 
**Sources** | Pointer to [**[]AskSource**](AskSource.md) | Sources are the pages the answer was written from, deduplicated and ranked. Always an array, never null. | [optional] 

## Methods

### NewAskReport

`func NewAskReport() *AskReport`

NewAskReport instantiates a new AskReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskReportWithDefaults

`func NewAskReportWithDefaults() *AskReport`

NewAskReportWithDefaults instantiates a new AskReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *AskReport) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *AskReport) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *AskReport) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *AskReport) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetFollowUps

`func (o *AskReport) GetFollowUps() []string`

GetFollowUps returns the FollowUps field if non-nil, zero value otherwise.

### GetFollowUpsOk

`func (o *AskReport) GetFollowUpsOk() (*[]string, bool)`

GetFollowUpsOk returns a tuple with the FollowUps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUps

`func (o *AskReport) SetFollowUps(v []string)`

SetFollowUps sets FollowUps field to given value.

### HasFollowUps

`func (o *AskReport) HasFollowUps() bool`

HasFollowUps returns a boolean if a field has been set.

### GetMode

`func (o *AskReport) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AskReport) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AskReport) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AskReport) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModel

`func (o *AskReport) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AskReport) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AskReport) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AskReport) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSources

`func (o *AskReport) GetSources() []AskSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AskReport) GetSourcesOk() (*[]AskSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AskReport) SetSources(v []AskSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AskReport) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


