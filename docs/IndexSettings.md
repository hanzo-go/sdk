# IndexSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterableAttributes** | Pointer to **[]string** | FilterableAttributes are the document fields a search &#x60;filter&#x60; may constrain. A field not listed here cannot be filtered on. | [optional] 

## Methods

### NewIndexSettings

`func NewIndexSettings() *IndexSettings`

NewIndexSettings instantiates a new IndexSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexSettingsWithDefaults

`func NewIndexSettingsWithDefaults() *IndexSettings`

NewIndexSettingsWithDefaults instantiates a new IndexSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterableAttributes

`func (o *IndexSettings) GetFilterableAttributes() []string`

GetFilterableAttributes returns the FilterableAttributes field if non-nil, zero value otherwise.

### GetFilterableAttributesOk

`func (o *IndexSettings) GetFilterableAttributesOk() (*[]string, bool)`

GetFilterableAttributesOk returns a tuple with the FilterableAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterableAttributes

`func (o *IndexSettings) SetFilterableAttributes(v []string)`

SetFilterableAttributes sets FilterableAttributes field to given value.

### HasFilterableAttributes

`func (o *IndexSettings) HasFilterableAttributes() bool`

HasFilterableAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


