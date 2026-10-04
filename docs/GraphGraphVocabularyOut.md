# GraphGraphVocabularyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bound** | Pointer to **int64** | Bound is the ceiling on one walk. | [optional] 
**Relations** | Pointer to **[]string** | Relations is what this organization has actually asserted. This plane declares none of its own. | [optional] 
**Rule** | Pointer to **[]string** | Rule names the terms of the precedence order, in the order they apply. A reader who is told a winner without the rule cannot check it. | [optional] 
**Schema** | Pointer to [**[]GraphGraphDeclared**](GraphGraphDeclared.md) | Schema is what the organization has declared, in force now: each relation with a domain or range, asserted as (&#x60;relation:&lt;name&gt;&#x60;, &#x60;domain&#x60; | &#x60;range&#x60;, &#x60;&lt;type&gt;&#x60;) and checked against the entities&#39; (&#x60;&lt;entity&gt;&#x60;, &#x60;type&#x60;, &#x60;&lt;type&gt;&#x60;). A relation absent here is open.  A decision is an entity of type &#x60;decision&#x60; by convention, never enforced: &#x60;scenario&#x60;, &#x60;reasoning&#x60;, &#x60;outcome&#x60; and &#x60;confidence&#x60; are its properties; &#x60;decided_by&#x60;, &#x60;caused_by&#x60;, &#x60;influenced&#x60; and &#x60;precedent&#x60; are its edges. Declare their domain and range to make that convention checked; find a precedent with search, then neighbors over &#x60;precedent&#x60;. | [optional] 

## Methods

### NewGraphGraphVocabularyOut

`func NewGraphGraphVocabularyOut() *GraphGraphVocabularyOut`

NewGraphGraphVocabularyOut instantiates a new GraphGraphVocabularyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphVocabularyOutWithDefaults

`func NewGraphGraphVocabularyOutWithDefaults() *GraphGraphVocabularyOut`

NewGraphGraphVocabularyOutWithDefaults instantiates a new GraphGraphVocabularyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBound

`func (o *GraphGraphVocabularyOut) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *GraphGraphVocabularyOut) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *GraphGraphVocabularyOut) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *GraphGraphVocabularyOut) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetRelations

`func (o *GraphGraphVocabularyOut) GetRelations() []string`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *GraphGraphVocabularyOut) GetRelationsOk() (*[]string, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *GraphGraphVocabularyOut) SetRelations(v []string)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *GraphGraphVocabularyOut) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetRule

`func (o *GraphGraphVocabularyOut) GetRule() []string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GraphGraphVocabularyOut) GetRuleOk() (*[]string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GraphGraphVocabularyOut) SetRule(v []string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *GraphGraphVocabularyOut) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSchema

`func (o *GraphGraphVocabularyOut) GetSchema() []GraphGraphDeclared`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *GraphGraphVocabularyOut) GetSchemaOk() (*[]GraphGraphDeclared, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *GraphGraphVocabularyOut) SetSchema(v []GraphGraphDeclared)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *GraphGraphVocabularyOut) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


