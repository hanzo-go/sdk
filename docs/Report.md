# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is the grounded prose, with inline markdown citations. Every link in it points at a page in Sources: the citation check runs on the text before it leaves the engine, so a cited URL is one THIS call fetched. | [optional] 
**FollowUps** | Pointer to **[]string** | FollowUps are the questions worth asking next. Best-effort — an empty list is a normal outcome, not a fault. | [optional] 
**Mode** | Pointer to **string** | Mode is the profile that ran: search, news, research or deep. | [optional] 
**Model** | Pointer to **string** | Model is the model that synthesized the answer. | [optional] 
**Sources** | Pointer to [**[]Source**](Source.md) | Sources are the pages the answer was written from, deduplicated and ranked. Always an array, never null. | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *Report) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *Report) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *Report) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *Report) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetFollowUps

`func (o *Report) GetFollowUps() []string`

GetFollowUps returns the FollowUps field if non-nil, zero value otherwise.

### GetFollowUpsOk

`func (o *Report) GetFollowUpsOk() (*[]string, bool)`

GetFollowUpsOk returns a tuple with the FollowUps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUps

`func (o *Report) SetFollowUps(v []string)`

SetFollowUps sets FollowUps field to given value.

### HasFollowUps

`func (o *Report) HasFollowUps() bool`

HasFollowUps returns a boolean if a field has been set.

### GetMode

`func (o *Report) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Report) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Report) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Report) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModel

`func (o *Report) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Report) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Report) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Report) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSources

`func (o *Report) GetSources() []Source`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Report) GetSourcesOk() (*[]Source, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Report) SetSources(v []Source)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Report) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


