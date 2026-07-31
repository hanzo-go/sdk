# SearchEditDocumentsByFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** |  | [optional] 
**Function** | **string** | JavaScript function body to transform documents | 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSearchEditDocumentsByFunctionRequest

`func NewSearchEditDocumentsByFunctionRequest(function string, ) *SearchEditDocumentsByFunctionRequest`

NewSearchEditDocumentsByFunctionRequest instantiates a new SearchEditDocumentsByFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEditDocumentsByFunctionRequestWithDefaults

`func NewSearchEditDocumentsByFunctionRequestWithDefaults() *SearchEditDocumentsByFunctionRequest`

NewSearchEditDocumentsByFunctionRequestWithDefaults instantiates a new SearchEditDocumentsByFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *SearchEditDocumentsByFunctionRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchEditDocumentsByFunctionRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchEditDocumentsByFunctionRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchEditDocumentsByFunctionRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFunction

`func (o *SearchEditDocumentsByFunctionRequest) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *SearchEditDocumentsByFunctionRequest) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *SearchEditDocumentsByFunctionRequest) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetContext

`func (o *SearchEditDocumentsByFunctionRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SearchEditDocumentsByFunctionRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SearchEditDocumentsByFunctionRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *SearchEditDocumentsByFunctionRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


