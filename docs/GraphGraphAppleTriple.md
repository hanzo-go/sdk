# GraphGraphAppleTriple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to **float64** | Confidence is the daemon&#39;s certainty in [0,1]. Absent defaults to 1 for copied facts and 0 for unscored rows. | [optional] 
**Object** | **string** | Object is the target entity identifier, class code or scalar literal. Required. | 
**Predicate** | **string** | Predicate is the predicate code, such as PS1, PS33, PS500 or PS598. Required. | 
**RelationshipId** | Pointer to **string** | RelationshipID groups the two rows of a reified edge in stable_graph: one row with RelationshipPredicate PS1 naming the CS*_/nm_* relationship class, and a companion row naming the target entity or value. Absent on single-row triples. | [optional] 
**RelationshipPredicate** | Pointer to **string** | RelationshipPredicate is the sub-predicate on a reified edge row, such as PS1 for the relationship class or PS45 for the target entity. Absent on single-row triples. | [optional] 
**SourceDuplicates** | Pointer to **int64** | SourceDuplicates is how many duplicate source records backed the row. | [optional] 
**Sources** | Pointer to **string** | Sources names the OS stores that contributed the row, such as Contacts, Calendar, Photos or Family Sharing. Absent defaults to apple/graph.db. | [optional] 
**Subject** | **string** | Subject is the source entity identifier in graph.db, such as 9100000000007919. Required. | 
**Timestamp** | Pointer to **string** | Timestamp is when the daemon recorded the row, RFC 3339. Absent defaults to the batch&#39;s At instant. | [optional] 
**Until** | Pointer to **string** | Until is when a withdrawn row in expired_stable_graph stopped holding, RFC 3339. Absent on expired rows defaults to the batch&#39;s At instant. | [optional] 

## Methods

### NewGraphGraphAppleTriple

`func NewGraphGraphAppleTriple(object string, predicate string, subject string, ) *GraphGraphAppleTriple`

NewGraphGraphAppleTriple instantiates a new GraphGraphAppleTriple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphAppleTripleWithDefaults

`func NewGraphGraphAppleTripleWithDefaults() *GraphGraphAppleTriple`

NewGraphGraphAppleTripleWithDefaults instantiates a new GraphGraphAppleTriple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *GraphGraphAppleTriple) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GraphGraphAppleTriple) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GraphGraphAppleTriple) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *GraphGraphAppleTriple) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetObject

`func (o *GraphGraphAppleTriple) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GraphGraphAppleTriple) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GraphGraphAppleTriple) SetObject(v string)`

SetObject sets Object field to given value.


### GetPredicate

`func (o *GraphGraphAppleTriple) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *GraphGraphAppleTriple) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *GraphGraphAppleTriple) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.


### GetRelationshipId

`func (o *GraphGraphAppleTriple) GetRelationshipId() string`

GetRelationshipId returns the RelationshipId field if non-nil, zero value otherwise.

### GetRelationshipIdOk

`func (o *GraphGraphAppleTriple) GetRelationshipIdOk() (*string, bool)`

GetRelationshipIdOk returns a tuple with the RelationshipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipId

`func (o *GraphGraphAppleTriple) SetRelationshipId(v string)`

SetRelationshipId sets RelationshipId field to given value.

### HasRelationshipId

`func (o *GraphGraphAppleTriple) HasRelationshipId() bool`

HasRelationshipId returns a boolean if a field has been set.

### GetRelationshipPredicate

`func (o *GraphGraphAppleTriple) GetRelationshipPredicate() string`

GetRelationshipPredicate returns the RelationshipPredicate field if non-nil, zero value otherwise.

### GetRelationshipPredicateOk

`func (o *GraphGraphAppleTriple) GetRelationshipPredicateOk() (*string, bool)`

GetRelationshipPredicateOk returns a tuple with the RelationshipPredicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipPredicate

`func (o *GraphGraphAppleTriple) SetRelationshipPredicate(v string)`

SetRelationshipPredicate sets RelationshipPredicate field to given value.

### HasRelationshipPredicate

`func (o *GraphGraphAppleTriple) HasRelationshipPredicate() bool`

HasRelationshipPredicate returns a boolean if a field has been set.

### GetSourceDuplicates

`func (o *GraphGraphAppleTriple) GetSourceDuplicates() int64`

GetSourceDuplicates returns the SourceDuplicates field if non-nil, zero value otherwise.

### GetSourceDuplicatesOk

`func (o *GraphGraphAppleTriple) GetSourceDuplicatesOk() (*int64, bool)`

GetSourceDuplicatesOk returns a tuple with the SourceDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDuplicates

`func (o *GraphGraphAppleTriple) SetSourceDuplicates(v int64)`

SetSourceDuplicates sets SourceDuplicates field to given value.

### HasSourceDuplicates

`func (o *GraphGraphAppleTriple) HasSourceDuplicates() bool`

HasSourceDuplicates returns a boolean if a field has been set.

### GetSources

`func (o *GraphGraphAppleTriple) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GraphGraphAppleTriple) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GraphGraphAppleTriple) SetSources(v string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GraphGraphAppleTriple) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSubject

`func (o *GraphGraphAppleTriple) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GraphGraphAppleTriple) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GraphGraphAppleTriple) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTimestamp

`func (o *GraphGraphAppleTriple) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GraphGraphAppleTriple) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GraphGraphAppleTriple) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GraphGraphAppleTriple) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUntil

`func (o *GraphGraphAppleTriple) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *GraphGraphAppleTriple) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *GraphGraphAppleTriple) SetUntil(v string)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *GraphGraphAppleTriple) HasUntil() bool`

HasUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


