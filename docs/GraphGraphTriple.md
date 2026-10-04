# GraphGraphTriple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to **bool** | Names is the author&#39;s declaration that Object is an entity and the assertion is an EDGE, written &#x60;[[key]]&#x60;. Absent, the relation is a property of Subject. It is read from the notation and never guessed from the value&#39;s shape. | [optional] 
**Object** | Pointer to **string** | Object is what the relation points at. When Names is true it has been unwrapped from its &#x60;[[…]]&#x60; and is another entity&#39;s key. | [optional] 
**Predicate** | Pointer to **string** | Predicate is the relation, exactly as the line spells it before the &#x60;::&#x60;. This surface holds no vocabulary, so it renames nothing. | [optional] 
**Reason** | Pointer to **string** | Reason is why filing it would be refused: the declared schema, or a declaration from a caller who is not an admin of the organization. Present only on a member of Refused. | [optional] 
**Section** | Pointer to **int64** | Section is which section of the source stated it, counting from zero. It is the second half of every resulting assertion&#39;s evidence, &#x60;&lt;source&gt;#&lt;section&gt;&#x60;. | [optional] 
**Subject** | Pointer to **string** | Subject is the entity the statement is about: the nearest heading above the line, or the request&#39;s own subject where no heading has appeared yet. | [optional] 

## Methods

### NewGraphGraphTriple

`func NewGraphGraphTriple() *GraphGraphTriple`

NewGraphGraphTriple instantiates a new GraphGraphTriple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphTripleWithDefaults

`func NewGraphGraphTripleWithDefaults() *GraphGraphTriple`

NewGraphGraphTripleWithDefaults instantiates a new GraphGraphTriple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *GraphGraphTriple) GetNames() bool`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GraphGraphTriple) GetNamesOk() (*bool, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GraphGraphTriple) SetNames(v bool)`

SetNames sets Names field to given value.

### HasNames

`func (o *GraphGraphTriple) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetObject

`func (o *GraphGraphTriple) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GraphGraphTriple) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GraphGraphTriple) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GraphGraphTriple) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPredicate

`func (o *GraphGraphTriple) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *GraphGraphTriple) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *GraphGraphTriple) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *GraphGraphTriple) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.

### GetReason

`func (o *GraphGraphTriple) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GraphGraphTriple) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GraphGraphTriple) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GraphGraphTriple) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSection

`func (o *GraphGraphTriple) GetSection() int64`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *GraphGraphTriple) GetSectionOk() (*int64, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *GraphGraphTriple) SetSection(v int64)`

SetSection sets Section field to given value.

### HasSection

`func (o *GraphGraphTriple) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSubject

`func (o *GraphGraphTriple) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GraphGraphTriple) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GraphGraphTriple) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GraphGraphTriple) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


