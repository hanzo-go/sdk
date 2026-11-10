# WebsearchSearchEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptsKey** | Pointer to **bool** | AcceptsKey is whether the engine takes an API key at all. | [optional] 
**Available** | Pointer to **bool** | Available is whether the engine answers from this cluster as things stand: it needs no key there (bing, ddg), or its key is held. Mojeek&#39;s keyless page still runs as a fallback but does not count, because it does not answer from this cluster&#39;s egress. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is whether searches run it. | [optional] 
**Keyed** | Pointer to **bool** | Keyed is whether the platform holds a key for it in KMS. The key itself is never returned by any operation. | [optional] 
**Name** | Pointer to **string** | Name is the engine: &#x60;mojeek&#x60;, &#x60;bing&#x60;, &#x60;brave&#x60; or &#x60;ddg&#x60;. | [optional] 
**NeedsKey** | Pointer to **bool** | NeedsKey is whether the engine answers only with a key. | [optional] 
**Options** | Pointer to [**WebsearchSearchEngineOptions**](WebsearchSearchEngineOptions.md) | Options are the values the engine is asked with, defaults filled in. | [optional] 
**Order** | Pointer to **int64** | Order is its place among the enabled engines, from 1; 0 when disabled. | [optional] 
**Supports** | Pointer to [**[]WebsearchSearchEngineOption**](WebsearchSearchEngineOption.md) | Supports lists the options the engine reads, with each one&#39;s bounds. | [optional] 

## Methods

### NewWebsearchSearchEngine

`func NewWebsearchSearchEngine() *WebsearchSearchEngine`

NewWebsearchSearchEngine instantiates a new WebsearchSearchEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchSearchEngineWithDefaults

`func NewWebsearchSearchEngineWithDefaults() *WebsearchSearchEngine`

NewWebsearchSearchEngineWithDefaults instantiates a new WebsearchSearchEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptsKey

`func (o *WebsearchSearchEngine) GetAcceptsKey() bool`

GetAcceptsKey returns the AcceptsKey field if non-nil, zero value otherwise.

### GetAcceptsKeyOk

`func (o *WebsearchSearchEngine) GetAcceptsKeyOk() (*bool, bool)`

GetAcceptsKeyOk returns a tuple with the AcceptsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsKey

`func (o *WebsearchSearchEngine) SetAcceptsKey(v bool)`

SetAcceptsKey sets AcceptsKey field to given value.

### HasAcceptsKey

`func (o *WebsearchSearchEngine) HasAcceptsKey() bool`

HasAcceptsKey returns a boolean if a field has been set.

### GetAvailable

`func (o *WebsearchSearchEngine) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *WebsearchSearchEngine) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *WebsearchSearchEngine) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *WebsearchSearchEngine) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetEnabled

`func (o *WebsearchSearchEngine) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebsearchSearchEngine) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebsearchSearchEngine) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebsearchSearchEngine) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKeyed

`func (o *WebsearchSearchEngine) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *WebsearchSearchEngine) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *WebsearchSearchEngine) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *WebsearchSearchEngine) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.

### GetName

`func (o *WebsearchSearchEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsearchSearchEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsearchSearchEngine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebsearchSearchEngine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedsKey

`func (o *WebsearchSearchEngine) GetNeedsKey() bool`

GetNeedsKey returns the NeedsKey field if non-nil, zero value otherwise.

### GetNeedsKeyOk

`func (o *WebsearchSearchEngine) GetNeedsKeyOk() (*bool, bool)`

GetNeedsKeyOk returns a tuple with the NeedsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsKey

`func (o *WebsearchSearchEngine) SetNeedsKey(v bool)`

SetNeedsKey sets NeedsKey field to given value.

### HasNeedsKey

`func (o *WebsearchSearchEngine) HasNeedsKey() bool`

HasNeedsKey returns a boolean if a field has been set.

### GetOptions

`func (o *WebsearchSearchEngine) GetOptions() WebsearchSearchEngineOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WebsearchSearchEngine) GetOptionsOk() (*WebsearchSearchEngineOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WebsearchSearchEngine) SetOptions(v WebsearchSearchEngineOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *WebsearchSearchEngine) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOrder

`func (o *WebsearchSearchEngine) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WebsearchSearchEngine) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WebsearchSearchEngine) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WebsearchSearchEngine) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSupports

`func (o *WebsearchSearchEngine) GetSupports() []WebsearchSearchEngineOption`

GetSupports returns the Supports field if non-nil, zero value otherwise.

### GetSupportsOk

`func (o *WebsearchSearchEngine) GetSupportsOk() (*[]WebsearchSearchEngineOption, bool)`

GetSupportsOk returns a tuple with the Supports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupports

`func (o *WebsearchSearchEngine) SetSupports(v []WebsearchSearchEngineOption)`

SetSupports sets Supports field to given value.

### HasSupports

`func (o *WebsearchSearchEngine) HasSupports() bool`

HasSupports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


