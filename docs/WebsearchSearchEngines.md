# WebsearchSearchEngines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engines** | Pointer to [**[]WebsearchSearchEngine**](WebsearchSearchEngine.md) | Engines are the enabled engines in the order they run, then the rest by name. | [optional] 
**Saved** | Pointer to **bool** | Saved is whether a setting is stored. False means the default is in force: Mojeek alone when its key is held, else bing, ddg and Mojeek. | [optional] 

## Methods

### NewWebsearchSearchEngines

`func NewWebsearchSearchEngines() *WebsearchSearchEngines`

NewWebsearchSearchEngines instantiates a new WebsearchSearchEngines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchSearchEnginesWithDefaults

`func NewWebsearchSearchEnginesWithDefaults() *WebsearchSearchEngines`

NewWebsearchSearchEnginesWithDefaults instantiates a new WebsearchSearchEngines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngines

`func (o *WebsearchSearchEngines) GetEngines() []WebsearchSearchEngine`

GetEngines returns the Engines field if non-nil, zero value otherwise.

### GetEnginesOk

`func (o *WebsearchSearchEngines) GetEnginesOk() (*[]WebsearchSearchEngine, bool)`

GetEnginesOk returns a tuple with the Engines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngines

`func (o *WebsearchSearchEngines) SetEngines(v []WebsearchSearchEngine)`

SetEngines sets Engines field to given value.

### HasEngines

`func (o *WebsearchSearchEngines) HasEngines() bool`

HasEngines returns a boolean if a field has been set.

### GetSaved

`func (o *WebsearchSearchEngines) GetSaved() bool`

GetSaved returns the Saved field if non-nil, zero value otherwise.

### GetSavedOk

`func (o *WebsearchSearchEngines) GetSavedOk() (*bool, bool)`

GetSavedOk returns a tuple with the Saved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaved

`func (o *WebsearchSearchEngines) SetSaved(v bool)`

SetSaved sets Saved field to given value.

### HasSaved

`func (o *WebsearchSearchEngines) HasSaved() bool`

HasSaved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


