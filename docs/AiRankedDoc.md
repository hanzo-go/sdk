# AiRankedDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**AiRankedText**](AiRankedText.md) |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**RelevanceScore** | Pointer to **float32** |  | [optional] 

## Methods

### NewAiRankedDoc

`func NewAiRankedDoc() *AiRankedDoc`

NewAiRankedDoc instantiates a new AiRankedDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRankedDocWithDefaults

`func NewAiRankedDocWithDefaults() *AiRankedDoc`

NewAiRankedDocWithDefaults instantiates a new AiRankedDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *AiRankedDoc) GetDocument() AiRankedText`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *AiRankedDoc) GetDocumentOk() (*AiRankedText, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *AiRankedDoc) SetDocument(v AiRankedText)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *AiRankedDoc) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetIndex

`func (o *AiRankedDoc) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AiRankedDoc) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AiRankedDoc) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AiRankedDoc) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetRelevanceScore

`func (o *AiRankedDoc) GetRelevanceScore() float32`

GetRelevanceScore returns the RelevanceScore field if non-nil, zero value otherwise.

### GetRelevanceScoreOk

`func (o *AiRankedDoc) GetRelevanceScoreOk() (*float32, bool)`

GetRelevanceScoreOk returns a tuple with the RelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceScore

`func (o *AiRankedDoc) SetRelevanceScore(v float32)`

SetRelevanceScore sets RelevanceScore field to given value.

### HasRelevanceScore

`func (o *AiRankedDoc) HasRelevanceScore() bool`

HasRelevanceScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


