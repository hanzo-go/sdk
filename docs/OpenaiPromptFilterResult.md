# OpenaiPromptFilterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentFilterResults** | Pointer to [**OpenaiContentFilterResults**](OpenaiContentFilterResults.md) |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 

## Methods

### NewOpenaiPromptFilterResult

`func NewOpenaiPromptFilterResult() *OpenaiPromptFilterResult`

NewOpenaiPromptFilterResult instantiates a new OpenaiPromptFilterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiPromptFilterResultWithDefaults

`func NewOpenaiPromptFilterResultWithDefaults() *OpenaiPromptFilterResult`

NewOpenaiPromptFilterResultWithDefaults instantiates a new OpenaiPromptFilterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentFilterResults

`func (o *OpenaiPromptFilterResult) GetContentFilterResults() OpenaiContentFilterResults`

GetContentFilterResults returns the ContentFilterResults field if non-nil, zero value otherwise.

### GetContentFilterResultsOk

`func (o *OpenaiPromptFilterResult) GetContentFilterResultsOk() (*OpenaiContentFilterResults, bool)`

GetContentFilterResultsOk returns a tuple with the ContentFilterResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilterResults

`func (o *OpenaiPromptFilterResult) SetContentFilterResults(v OpenaiContentFilterResults)`

SetContentFilterResults sets ContentFilterResults field to given value.

### HasContentFilterResults

`func (o *OpenaiPromptFilterResult) HasContentFilterResults() bool`

HasContentFilterResults returns a boolean if a field has been set.

### GetIndex

`func (o *OpenaiPromptFilterResult) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OpenaiPromptFilterResult) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OpenaiPromptFilterResult) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *OpenaiPromptFilterResult) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


