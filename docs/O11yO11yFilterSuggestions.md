# O11yO11yFilterSuggestions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]O11yO11yAttributeKey**](O11yO11yAttributeKey.md) | Attributes are the suggested attribute keys. | [optional] 
**ExampleQueries** | Pointer to [**[]O11yO11yFilterSet**](O11yO11yFilterSet.md) | ExampleQueries are ready-to-run filter sets. | [optional] 

## Methods

### NewO11yO11yFilterSuggestions

`func NewO11yO11yFilterSuggestions() *O11yO11yFilterSuggestions`

NewO11yO11yFilterSuggestions instantiates a new O11yO11yFilterSuggestions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFilterSuggestionsWithDefaults

`func NewO11yO11yFilterSuggestionsWithDefaults() *O11yO11yFilterSuggestions`

NewO11yO11yFilterSuggestionsWithDefaults instantiates a new O11yO11yFilterSuggestions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *O11yO11yFilterSuggestions) GetAttributes() []O11yO11yAttributeKey`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *O11yO11yFilterSuggestions) GetAttributesOk() (*[]O11yO11yAttributeKey, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *O11yO11yFilterSuggestions) SetAttributes(v []O11yO11yAttributeKey)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *O11yO11yFilterSuggestions) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetExampleQueries

`func (o *O11yO11yFilterSuggestions) GetExampleQueries() []O11yO11yFilterSet`

GetExampleQueries returns the ExampleQueries field if non-nil, zero value otherwise.

### GetExampleQueriesOk

`func (o *O11yO11yFilterSuggestions) GetExampleQueriesOk() (*[]O11yO11yFilterSet, bool)`

GetExampleQueriesOk returns a tuple with the ExampleQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleQueries

`func (o *O11yO11yFilterSuggestions) SetExampleQueries(v []O11yO11yFilterSet)`

SetExampleQueries sets ExampleQueries field to given value.

### HasExampleQueries

`func (o *O11yO11yFilterSuggestions) HasExampleQueries() bool`

HasExampleQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


