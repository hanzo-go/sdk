# WebResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the ENGINE&#39;s snippet — the few lines shown under the title, not the page&#39;s text. Read the page itself with POST /v1/crawl. | [optional] 
**Engine** | Pointer to **string** | Engine names the backend that found this hit, so one engine&#39;s view of a query can be told from another&#39;s. | [optional] 
**Title** | Pointer to **string** | Title is the page&#39;s title. | [optional] 
**Url** | Pointer to **string** | URL is the page&#39;s address, as the engine reported it. | [optional] 

## Methods

### NewWebResult

`func NewWebResult() *WebResult`

NewWebResult instantiates a new WebResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebResultWithDefaults

`func NewWebResultWithDefaults() *WebResult`

NewWebResultWithDefaults instantiates a new WebResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *WebResult) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebResult) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebResult) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebResult) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEngine

`func (o *WebResult) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *WebResult) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *WebResult) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *WebResult) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetTitle

`func (o *WebResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *WebResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


