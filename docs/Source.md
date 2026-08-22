# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to **string** | Engine is the search backend the hit came from: bing, ddg, mojeek or brave. Omitted when the backend did not name itself. Results are merged across backends, so two sources in one answer can carry different engines. | [optional] 
**Favicon** | Pointer to **string** | Favicon is a 64px icon URL derived from the host for the client to render beside the citation. It is Google&#39;s s2 service, not something we host or fetched — an empty host yields the empty string. | [optional] 
**Snippet** | Pointer to **string** | Snippet is the engine&#39;s summary of the page, clipped to 600 runes. THIS IS WHAT THE CLIENT SHOWS. What the model reads is the fetched page, which is far larger and deliberately never on the wire. | [optional] 
**Title** | Pointer to **string** | Title is the page title the engine reported, stripped of the bracketed furniture engines staple on (\&quot;[PDF]\&quot;, \&quot;(Official Site)\&quot;). It falls back to the www-stripped host when the engine gave none, so it is never empty and is safe to use as link text. | [optional] 
**Url** | Pointer to **string** | URL is the page, absolute, exactly as the engine gave it. It is also the dedupe key — one source per URL, and at most hostCap per host — and what a markdown citation in the answer is checked against, so a link in the prose always matches a URL here. | [optional] 

## Methods

### NewSource

`func NewSource() *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *Source) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *Source) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *Source) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *Source) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetFavicon

`func (o *Source) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *Source) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *Source) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *Source) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetSnippet

`func (o *Source) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *Source) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *Source) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *Source) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetTitle

`func (o *Source) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Source) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Source) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Source) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *Source) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Source) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Source) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Source) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


