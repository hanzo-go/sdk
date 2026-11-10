# WebsearchSearchEngineOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count is how many hits the engine is asked for per search. 0 or absent on a write means the engine&#39;s default; a read reports the value in force. | [optional] 
**Language** | Pointer to **string** | Language is the default language the engine is asked in, as a tag like &#x60;en&#x60; or &#x60;en-US&#x60;. A search that names its own language overrides it. | [optional] 
**Region** | Pointer to **string** | Region is a two-letter country code the engine favours, like &#x60;GB&#x60;. Stored upper case. | [optional] 

## Methods

### NewWebsearchSearchEngineOptions

`func NewWebsearchSearchEngineOptions() *WebsearchSearchEngineOptions`

NewWebsearchSearchEngineOptions instantiates a new WebsearchSearchEngineOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchSearchEngineOptionsWithDefaults

`func NewWebsearchSearchEngineOptionsWithDefaults() *WebsearchSearchEngineOptions`

NewWebsearchSearchEngineOptionsWithDefaults instantiates a new WebsearchSearchEngineOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WebsearchSearchEngineOptions) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WebsearchSearchEngineOptions) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WebsearchSearchEngineOptions) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *WebsearchSearchEngineOptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLanguage

`func (o *WebsearchSearchEngineOptions) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WebsearchSearchEngineOptions) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WebsearchSearchEngineOptions) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *WebsearchSearchEngineOptions) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *WebsearchSearchEngineOptions) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *WebsearchSearchEngineOptions) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *WebsearchSearchEngineOptions) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *WebsearchSearchEngineOptions) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


