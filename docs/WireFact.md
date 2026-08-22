# WireFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when the thing was so, RFC 3339, as the asserter gave it. | [optional] 
**By** | Pointer to **string** | By is the identity that filed it — &#x60;owner&#x60; or &#x60;owner/user&#x60; — stamped from the validated principal at the write, never from the body. | [optional] 
**Confidence** | Pointer to **float32** | Confidence in [0,1] as the asserter gave it; absent is 0. It breaks a tie between two assertions equally knowable and decides nothing else. | [optional] 
**Entity** | Pointer to **string** | Entity is the thing described, in the organization&#39;s own namespace. | [optional] 
**Evidence** | Pointer to **string** | Evidence points at the record the claim came from. Absent when the asserter gave none. | [optional] 
**Id** | Pointer to **string** | ID is the assertion&#39;s content address, minted by the server from what was asserted. Two callers who assert the identical thing land on one ID and one row; changing any asserted field makes a different ID and a second row. | [optional] 
**Knowable** | Pointer to **string** | Knowable is the first instant this plane could have answered with the assertion, RFC 3339: the later of Seen and the server&#39;s clock at the write. Derived and never supplied, which is what stops history filed today from being backdated into a past read. | [optional] 
**Names** | Pointer to **bool** | Names true means the assertion is an edge and Value is an entity. A walk reads only these. | [optional] 
**Relation** | Pointer to **string** | Relation is what was asserted of it. | [optional] 
**Seen** | Pointer to **string** | Seen is when the asserter says it became knowable, RFC 3339. Provenance only — Knowable is what an as-of read is bounded by. | [optional] 
**Source** | Pointer to **string** | Source names who asserted, as the caller gave it. This plane ranks no source above another, so it never outweighs a later Knowable. | [optional] 
**Value** | Pointer to **string** | Value is what the relation points at: another entity&#39;s key when Names is true, otherwise a scalar. | [optional] 

## Methods

### NewWireFact

`func NewWireFact() *WireFact`

NewWireFact instantiates a new WireFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireFactWithDefaults

`func NewWireFactWithDefaults() *WireFact`

NewWireFactWithDefaults instantiates a new WireFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *WireFact) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *WireFact) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *WireFact) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *WireFact) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *WireFact) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *WireFact) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *WireFact) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *WireFact) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetConfidence

`func (o *WireFact) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *WireFact) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *WireFact) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *WireFact) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetEntity

`func (o *WireFact) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *WireFact) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *WireFact) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *WireFact) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEvidence

`func (o *WireFact) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *WireFact) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *WireFact) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *WireFact) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetId

`func (o *WireFact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireFact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireFact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireFact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKnowable

`func (o *WireFact) GetKnowable() string`

GetKnowable returns the Knowable field if non-nil, zero value otherwise.

### GetKnowableOk

`func (o *WireFact) GetKnowableOk() (*string, bool)`

GetKnowableOk returns a tuple with the Knowable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowable

`func (o *WireFact) SetKnowable(v string)`

SetKnowable sets Knowable field to given value.

### HasKnowable

`func (o *WireFact) HasKnowable() bool`

HasKnowable returns a boolean if a field has been set.

### GetNames

`func (o *WireFact) GetNames() bool`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *WireFact) GetNamesOk() (*bool, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *WireFact) SetNames(v bool)`

SetNames sets Names field to given value.

### HasNames

`func (o *WireFact) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetRelation

`func (o *WireFact) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *WireFact) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *WireFact) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *WireFact) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSeen

`func (o *WireFact) GetSeen() string`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *WireFact) GetSeenOk() (*string, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *WireFact) SetSeen(v string)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *WireFact) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetSource

`func (o *WireFact) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WireFact) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WireFact) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *WireFact) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *WireFact) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WireFact) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WireFact) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WireFact) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


