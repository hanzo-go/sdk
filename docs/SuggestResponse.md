# SuggestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funnel** | Pointer to [**Funnel**](Funnel.md) | Funnel is the org&#39;s trailing-window traffic → signups → orders. | [optional] 
**Narrative** | Pointer to **string** | Narrative is the AI&#39;s grounded prose over those quests and numbers. Absent when no AI plane is wired or the completion failed — never fabricated. | [optional] 
**Next** | Pointer to **string** | Next is the id of the single next step the static journey names — the linear answer the ranked Suggestions refine. | [optional] 
**Recommendations** | Pointer to **[]string** | Recommendations are the next-best GTM actions derived from that funnel. | [optional] 
**Suggestions** | Pointer to [**[]Suggestion**](Suggestion.md) | Suggestions are the available, non-terminal quests ranked best-first by how much downstream work each unblocks. | [optional] 

## Methods

### NewSuggestResponse

`func NewSuggestResponse() *SuggestResponse`

NewSuggestResponse instantiates a new SuggestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestResponseWithDefaults

`func NewSuggestResponseWithDefaults() *SuggestResponse`

NewSuggestResponseWithDefaults instantiates a new SuggestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnel

`func (o *SuggestResponse) GetFunnel() Funnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *SuggestResponse) GetFunnelOk() (*Funnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *SuggestResponse) SetFunnel(v Funnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *SuggestResponse) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetNarrative

`func (o *SuggestResponse) GetNarrative() string`

GetNarrative returns the Narrative field if non-nil, zero value otherwise.

### GetNarrativeOk

`func (o *SuggestResponse) GetNarrativeOk() (*string, bool)`

GetNarrativeOk returns a tuple with the Narrative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrative

`func (o *SuggestResponse) SetNarrative(v string)`

SetNarrative sets Narrative field to given value.

### HasNarrative

`func (o *SuggestResponse) HasNarrative() bool`

HasNarrative returns a boolean if a field has been set.

### GetNext

`func (o *SuggestResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SuggestResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SuggestResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *SuggestResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetRecommendations

`func (o *SuggestResponse) GetRecommendations() []string`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *SuggestResponse) GetRecommendationsOk() (*[]string, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *SuggestResponse) SetRecommendations(v []string)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *SuggestResponse) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetSuggestions

`func (o *SuggestResponse) GetSuggestions() []Suggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *SuggestResponse) GetSuggestionsOk() (*[]Suggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *SuggestResponse) SetSuggestions(v []Suggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *SuggestResponse) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


