# WebsearchSearchEngineKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the API key. Sealed in KMS and never returned, logged or echoed. Never read from the query string. | [optional] 
**Name** | Pointer to **string** | Name is the engine, from the path. | [optional] 

## Methods

### NewWebsearchSearchEngineKey

`func NewWebsearchSearchEngineKey() *WebsearchSearchEngineKey`

NewWebsearchSearchEngineKey instantiates a new WebsearchSearchEngineKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchSearchEngineKeyWithDefaults

`func NewWebsearchSearchEngineKeyWithDefaults() *WebsearchSearchEngineKey`

NewWebsearchSearchEngineKeyWithDefaults instantiates a new WebsearchSearchEngineKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WebsearchSearchEngineKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebsearchSearchEngineKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebsearchSearchEngineKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WebsearchSearchEngineKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *WebsearchSearchEngineKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsearchSearchEngineKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsearchSearchEngineKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebsearchSearchEngineKey) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


