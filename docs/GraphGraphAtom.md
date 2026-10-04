# GraphGraphAtom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** | Object is its second argument. Absent for a predicate the rules use with one argument. | [optional] 
**Predicate** | Pointer to **string** | Predicate is the relation the tuple is of. | [optional] 
**Subject** | Pointer to **string** | Subject is its first argument. | [optional] 

## Methods

### NewGraphGraphAtom

`func NewGraphGraphAtom() *GraphGraphAtom`

NewGraphGraphAtom instantiates a new GraphGraphAtom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphAtomWithDefaults

`func NewGraphGraphAtomWithDefaults() *GraphGraphAtom`

NewGraphGraphAtomWithDefaults instantiates a new GraphGraphAtom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *GraphGraphAtom) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GraphGraphAtom) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GraphGraphAtom) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GraphGraphAtom) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPredicate

`func (o *GraphGraphAtom) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *GraphGraphAtom) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *GraphGraphAtom) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *GraphGraphAtom) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.

### GetSubject

`func (o *GraphGraphAtom) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GraphGraphAtom) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GraphGraphAtom) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GraphGraphAtom) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


