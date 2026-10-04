# WebsearchWebSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **string** | Language narrows the engines to a locale, BCP-47-ish (\&quot;en\&quot;, \&quot;ja\&quot;, \&quot;de\&quot;). Empty means no narrowing. | [optional] 
**Q** | Pointer to **string** | Q is the query. Required — an empty one is refused rather than answered with the whole web. | [optional] 

## Methods

### NewWebsearchWebSearchQuery

`func NewWebsearchWebSearchQuery() *WebsearchWebSearchQuery`

NewWebsearchWebSearchQuery instantiates a new WebsearchWebSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchWebSearchQueryWithDefaults

`func NewWebsearchWebSearchQueryWithDefaults() *WebsearchWebSearchQuery`

NewWebsearchWebSearchQueryWithDefaults instantiates a new WebsearchWebSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *WebsearchWebSearchQuery) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WebsearchWebSearchQuery) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WebsearchWebSearchQuery) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *WebsearchWebSearchQuery) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetQ

`func (o *WebsearchWebSearchQuery) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *WebsearchWebSearchQuery) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *WebsearchWebSearchQuery) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *WebsearchWebSearchQuery) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


