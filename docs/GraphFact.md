# GraphFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when the thing was so, RFC 3339. Required, and refused when it sits more than five minutes ahead of the server clock — an assertion dated further out would never mature and would skew every read until it did. | [optional] 
**Confidence** | Pointer to **float32** | Confidence in [0,1]. A tie-breaker within the order, never a substitute for it. Absent is 0, the weakest an assertion can be. | [optional] 
**Entity** | Pointer to **string** | Entity is the thing being described, in the organization&#39;s own namespace. It is not created: an entity exists because something was asserted about it. Required, 512 bytes at most. | [optional] 
**Evidence** | Pointer to **string** | Evidence points at the record this claim came from, 512 bytes at most. An assertion without one is admitted and carries no defence. | [optional] 
**Names** | Pointer to **bool** | Names says the value is an entity. A walk reads only the edges, so this is a declaration and never a guess about the value&#39;s shape. | [optional] 
**Relation** | Pointer to **string** | Relation is what is being asserted — &#x60;depends&#x60;, &#x60;owner&#x60;, &#x60;same&#x60;, &#x60;title&#x60;. It is open: this plane holds no vocabulary of its own. Required, 128 bytes at most. | [optional] 
**Seen** | Pointer to **string** | Seen is when this assertion became knowable, RFC 3339. Defaults to At and may not precede it. It is provenance and it decides nothing: the instant every read uses is derived as the later of Seen and the server&#39;s own clock. | [optional] 
**Source** | Pointer to **string** | Source names who asserted. Required, because an assertion nobody is named for cannot be weighed against one that is. Open text: this plane ranks no source above another. | [optional] 
**Value** | Pointer to **string** | Value is what the relation points at. When Names is true it is another entity&#39;s key and the assertion is an EDGE; otherwise it is a scalar and the assertion is a property. 2048 bytes at most, or 512 when it names an entity. | [optional] 

## Methods

### NewGraphFact

`func NewGraphFact() *GraphFact`

NewGraphFact instantiates a new GraphFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphFactWithDefaults

`func NewGraphFactWithDefaults() *GraphFact`

NewGraphFactWithDefaults instantiates a new GraphFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *GraphFact) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *GraphFact) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *GraphFact) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *GraphFact) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetConfidence

`func (o *GraphFact) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GraphFact) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GraphFact) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *GraphFact) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetEntity

`func (o *GraphFact) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphFact) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphFact) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GraphFact) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEvidence

`func (o *GraphFact) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *GraphFact) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *GraphFact) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *GraphFact) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetNames

`func (o *GraphFact) GetNames() bool`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GraphFact) GetNamesOk() (*bool, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GraphFact) SetNames(v bool)`

SetNames sets Names field to given value.

### HasNames

`func (o *GraphFact) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetRelation

`func (o *GraphFact) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphFact) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphFact) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphFact) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSeen

`func (o *GraphFact) GetSeen() string`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *GraphFact) GetSeenOk() (*string, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *GraphFact) SetSeen(v string)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *GraphFact) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetSource

`func (o *GraphFact) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphFact) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphFact) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GraphFact) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *GraphFact) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GraphFact) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GraphFact) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GraphFact) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


