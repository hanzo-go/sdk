# O11yO11yErrorItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message is the detail. | [optional] 
**Suggestions** | Pointer to **[]string** | Suggestions say what to try about this detail. | [optional] 

## Methods

### NewO11yO11yErrorItem

`func NewO11yO11yErrorItem() *O11yO11yErrorItem`

NewO11yO11yErrorItem instantiates a new O11yO11yErrorItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yErrorItemWithDefaults

`func NewO11yO11yErrorItemWithDefaults() *O11yO11yErrorItem`

NewO11yO11yErrorItemWithDefaults instantiates a new O11yO11yErrorItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *O11yO11yErrorItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yO11yErrorItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yO11yErrorItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yO11yErrorItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuggestions

`func (o *O11yO11yErrorItem) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *O11yO11yErrorItem) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *O11yO11yErrorItem) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *O11yO11yErrorItem) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


