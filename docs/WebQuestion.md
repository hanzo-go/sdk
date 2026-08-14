# WebQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **string** | Language narrows the search to a locale, BCP-47-ish (\&quot;en\&quot;, \&quot;ja\&quot;). Empty means no narrowing. | [optional] 
**MaxSources** | Pointer to **int32** | MaxSources caps how many pages are read. Empty means the mode&#39;s own budget. | [optional] 
**Mode** | Pointer to **string** | Mode is how much work to do: &#x60;search&#x60; (fast, one pass), &#x60;news&#x60; (recency biased), &#x60;research&#x60; (a plan and several rounds) or &#x60;deep&#x60; (the widest survey). Empty means research. | [optional] 
**Q** | Pointer to **string** | Q is the question, in plain language. Required. | [optional] 
**Sources** | Pointer to **[]string** | Sources narrows where the evidence comes from: any of &#x60;web&#x60;, &#x60;news&#x60;, &#x60;academic&#x60;, &#x60;github&#x60;, &#x60;reddit&#x60;, &#x60;x&#x60;. Each becomes a site-scoped search, so &#x60;[\&quot;x\&quot;]&#x60; researches X/Twitter posts rather than the open web. | [optional] 

## Methods

### NewWebQuestion

`func NewWebQuestion() *WebQuestion`

NewWebQuestion instantiates a new WebQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebQuestionWithDefaults

`func NewWebQuestionWithDefaults() *WebQuestion`

NewWebQuestionWithDefaults instantiates a new WebQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *WebQuestion) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WebQuestion) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WebQuestion) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *WebQuestion) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMaxSources

`func (o *WebQuestion) GetMaxSources() int32`

GetMaxSources returns the MaxSources field if non-nil, zero value otherwise.

### GetMaxSourcesOk

`func (o *WebQuestion) GetMaxSourcesOk() (*int32, bool)`

GetMaxSourcesOk returns a tuple with the MaxSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSources

`func (o *WebQuestion) SetMaxSources(v int32)`

SetMaxSources sets MaxSources field to given value.

### HasMaxSources

`func (o *WebQuestion) HasMaxSources() bool`

HasMaxSources returns a boolean if a field has been set.

### GetMode

`func (o *WebQuestion) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WebQuestion) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WebQuestion) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WebQuestion) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetQ

`func (o *WebQuestion) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *WebQuestion) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *WebQuestion) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *WebQuestion) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetSources

`func (o *WebQuestion) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *WebQuestion) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *WebQuestion) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *WebQuestion) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


