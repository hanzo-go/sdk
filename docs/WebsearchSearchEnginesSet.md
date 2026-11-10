# WebsearchSearchEnginesSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **[]string** | Enabled are the engines searches run, in order. At least one, each a registered engine named once, and at least one of them must be available. | [optional] 
**Options** | Pointer to [**map[string]WebsearchSearchEngineOptions**](WebsearchSearchEngineOptions.md) | Options are each engine&#39;s values, by engine name — for any registered engine, enabled or not. An engine or a field left out takes its default: the setting replaces the stored one whole. | [optional] 

## Methods

### NewWebsearchSearchEnginesSet

`func NewWebsearchSearchEnginesSet() *WebsearchSearchEnginesSet`

NewWebsearchSearchEnginesSet instantiates a new WebsearchSearchEnginesSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchSearchEnginesSetWithDefaults

`func NewWebsearchSearchEnginesSetWithDefaults() *WebsearchSearchEnginesSet`

NewWebsearchSearchEnginesSetWithDefaults instantiates a new WebsearchSearchEnginesSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WebsearchSearchEnginesSet) GetEnabled() []string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebsearchSearchEnginesSet) GetEnabledOk() (*[]string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebsearchSearchEnginesSet) SetEnabled(v []string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebsearchSearchEnginesSet) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOptions

`func (o *WebsearchSearchEnginesSet) GetOptions() map[string]WebsearchSearchEngineOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WebsearchSearchEnginesSet) GetOptionsOk() (*map[string]WebsearchSearchEngineOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WebsearchSearchEnginesSet) SetOptions(v map[string]WebsearchSearchEngineOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *WebsearchSearchEnginesSet) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


