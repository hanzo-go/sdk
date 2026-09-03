# GraphVocabularyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bound** | Pointer to **int64** | Bound is the ceiling on one walk. | [optional] 
**Relations** | Pointer to **[]string** | Relations is what this organization has actually asserted, which is the only vocabulary there is: this plane declares none of its own. | [optional] 
**Rule** | Pointer to **[]string** | Rule names the terms of the precedence order, in the order they apply. A reader who is told a winner without the rule cannot check it. | [optional] 

## Methods

### NewGraphVocabularyOut

`func NewGraphVocabularyOut() *GraphVocabularyOut`

NewGraphVocabularyOut instantiates a new GraphVocabularyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVocabularyOutWithDefaults

`func NewGraphVocabularyOutWithDefaults() *GraphVocabularyOut`

NewGraphVocabularyOutWithDefaults instantiates a new GraphVocabularyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBound

`func (o *GraphVocabularyOut) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *GraphVocabularyOut) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *GraphVocabularyOut) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *GraphVocabularyOut) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetRelations

`func (o *GraphVocabularyOut) GetRelations() []string`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *GraphVocabularyOut) GetRelationsOk() (*[]string, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *GraphVocabularyOut) SetRelations(v []string)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *GraphVocabularyOut) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetRule

`func (o *GraphVocabularyOut) GetRule() []string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GraphVocabularyOut) GetRuleOk() (*[]string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GraphVocabularyOut) SetRule(v []string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *GraphVocabularyOut) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


