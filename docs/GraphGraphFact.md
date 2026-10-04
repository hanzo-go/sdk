# GraphGraphFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | **string** | At is when the thing was so, RFC 3339. Required, and refused when it sits more than five minutes ahead of the server clock — an assertion dated further out would never mature and would skew every read until it did. | 
**Confidence** | Pointer to **float64** | Confidence in [0,1]. A tie-breaker within the order, never a substitute for it. Absent is 0, the weakest an assertion can be. | [optional] 
**Entity** | **string** | Entity is the thing being described, in the organization&#39;s own namespace. It is not created: an entity exists because something was asserted about it. It is a key, folded to Unicode NFC with control characters removed, so two spellings of one name are one entity. An entity named &#x60;relation:&lt;name&gt;&#x60; is that relation&#39;s declaration and one named &#x60;rule:&lt;name&gt;&#x60; is a rule (graphDerive), and only an admin of the organization may assert anything of either. Required, 512 bytes at most. | 
**Evidence** | Pointer to **string** | Evidence points at the record this claim came from, 512 bytes at most. An assertion without one is admitted and carries no defence. | [optional] 
**Names** | Pointer to **bool** | Names says the value is an entity. A walk reads only the edges, so this is a declaration and never a guess about the value&#39;s shape. | [optional] 
**Relation** | **string** | Relation is what is being asserted — &#x60;depends&#x60;, &#x60;owner&#x60;, &#x60;same&#x60;, &#x60;title&#x60; — and a key, folded as Entity is. It is open until the organization declares it: once &#x60;relation:&lt;name&gt;&#x60; has a &#x60;domain&#x60; or &#x60;range&#x60; in force, the entity&#39;s &#x60;type&#x60; (and, for range, the value&#39;s) must match it. Required, 128 bytes at most. | 
**Seen** | Pointer to **string** | Seen is when this assertion became knowable, RFC 3339. Defaults to At and may not precede it. It is provenance and it decides nothing: the instant every read uses is derived as the later of Seen and the server&#39;s own clock. | [optional] 
**Source** | **string** | Source names who asserted. Required, because an assertion nobody is named for cannot be weighed against one that is. Open text: this plane ranks no source above another. | 
**Until** | Pointer to **string** | Until is when the thing stopped being so, RFC 3339, and must follow At. Absent is open: the statement holds until something ends it. A term of office, an employment, a membership are At and Until on one statement. | [optional] 
**Value** | Pointer to **string** | Value is what the relation points at. When Names is true it is another entity&#39;s key, folded as Entity is, and the assertion is an EDGE; otherwise it is a scalar, kept byte for byte, and the assertion is a property. 2048 bytes at most, or 512 when it names an entity. | [optional] 

## Methods

### NewGraphGraphFact

`func NewGraphGraphFact(at string, entity string, relation string, source string, ) *GraphGraphFact`

NewGraphGraphFact instantiates a new GraphGraphFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphFactWithDefaults

`func NewGraphGraphFactWithDefaults() *GraphGraphFact`

NewGraphGraphFactWithDefaults instantiates a new GraphGraphFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *GraphGraphFact) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *GraphGraphFact) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *GraphGraphFact) SetAt(v string)`

SetAt sets At field to given value.


### GetConfidence

`func (o *GraphGraphFact) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GraphGraphFact) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GraphGraphFact) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *GraphGraphFact) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetEntity

`func (o *GraphGraphFact) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphGraphFact) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphGraphFact) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetEvidence

`func (o *GraphGraphFact) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *GraphGraphFact) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *GraphGraphFact) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *GraphGraphFact) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetNames

`func (o *GraphGraphFact) GetNames() bool`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GraphGraphFact) GetNamesOk() (*bool, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GraphGraphFact) SetNames(v bool)`

SetNames sets Names field to given value.

### HasNames

`func (o *GraphGraphFact) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetRelation

`func (o *GraphGraphFact) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphGraphFact) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphGraphFact) SetRelation(v string)`

SetRelation sets Relation field to given value.


### GetSeen

`func (o *GraphGraphFact) GetSeen() string`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *GraphGraphFact) GetSeenOk() (*string, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *GraphGraphFact) SetSeen(v string)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *GraphGraphFact) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetSource

`func (o *GraphGraphFact) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphGraphFact) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphGraphFact) SetSource(v string)`

SetSource sets Source field to given value.


### GetUntil

`func (o *GraphGraphFact) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *GraphGraphFact) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *GraphGraphFact) SetUntil(v string)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *GraphGraphFact) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetValue

`func (o *GraphGraphFact) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GraphGraphFact) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GraphGraphFact) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GraphGraphFact) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


