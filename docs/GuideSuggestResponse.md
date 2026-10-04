# GuideSuggestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funnel** | Pointer to [**GuideFunnel**](GuideFunnel.md) | Funnel is the org&#39;s trailing-window traffic → signups → orders. | [optional] 
**Narrative** | Pointer to **string** | Narrative is the AI&#39;s grounded prose over those quests and numbers. Absent when no AI plane is wired or the completion failed — never fabricated. | [optional] 
**Next** | Pointer to **string** | Next is the id of the single next step the static journey names — the linear answer the ranked Suggestions refine. | [optional] 
**Recommendations** | Pointer to **[]string** | Recommendations are the next-best GTM actions derived from that funnel. | [optional] 
**Suggestions** | Pointer to [**[]GuideSuggestion**](GuideSuggestion.md) | Suggestions are the available, non-terminal quests ranked best-first by how much downstream work each unblocks. | [optional] 

## Methods

### NewGuideSuggestResponse

`func NewGuideSuggestResponse() *GuideSuggestResponse`

NewGuideSuggestResponse instantiates a new GuideSuggestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuideSuggestResponseWithDefaults

`func NewGuideSuggestResponseWithDefaults() *GuideSuggestResponse`

NewGuideSuggestResponseWithDefaults instantiates a new GuideSuggestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnel

`func (o *GuideSuggestResponse) GetFunnel() GuideFunnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *GuideSuggestResponse) GetFunnelOk() (*GuideFunnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *GuideSuggestResponse) SetFunnel(v GuideFunnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *GuideSuggestResponse) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetNarrative

`func (o *GuideSuggestResponse) GetNarrative() string`

GetNarrative returns the Narrative field if non-nil, zero value otherwise.

### GetNarrativeOk

`func (o *GuideSuggestResponse) GetNarrativeOk() (*string, bool)`

GetNarrativeOk returns a tuple with the Narrative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrative

`func (o *GuideSuggestResponse) SetNarrative(v string)`

SetNarrative sets Narrative field to given value.

### HasNarrative

`func (o *GuideSuggestResponse) HasNarrative() bool`

HasNarrative returns a boolean if a field has been set.

### GetNext

`func (o *GuideSuggestResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GuideSuggestResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GuideSuggestResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *GuideSuggestResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetRecommendations

`func (o *GuideSuggestResponse) GetRecommendations() []string`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *GuideSuggestResponse) GetRecommendationsOk() (*[]string, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *GuideSuggestResponse) SetRecommendations(v []string)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *GuideSuggestResponse) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetSuggestions

`func (o *GuideSuggestResponse) GetSuggestions() []GuideSuggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *GuideSuggestResponse) GetSuggestionsOk() (*[]GuideSuggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *GuideSuggestResponse) SetSuggestions(v []GuideSuggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *GuideSuggestResponse) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


