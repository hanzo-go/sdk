# GraphGraphSupport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertion** | Pointer to **string** | Assertion is the ID of the statement in force the atom rests on. Absent when the atom is itself derived: it is then an entry of Derived or Lemmas, with a proof of its own. | [optional] 
**Object** | Pointer to **string** | Object is its second argument. Absent for a predicate the rules use with one argument. | [optional] 
**Predicate** | Pointer to **string** | Predicate is the relation the supporting atom is of. | [optional] 
**Subject** | Pointer to **string** | Subject is its first argument. | [optional] 

## Methods

### NewGraphGraphSupport

`func NewGraphGraphSupport() *GraphGraphSupport`

NewGraphGraphSupport instantiates a new GraphGraphSupport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphSupportWithDefaults

`func NewGraphGraphSupportWithDefaults() *GraphGraphSupport`

NewGraphGraphSupportWithDefaults instantiates a new GraphGraphSupport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertion

`func (o *GraphGraphSupport) GetAssertion() string`

GetAssertion returns the Assertion field if non-nil, zero value otherwise.

### GetAssertionOk

`func (o *GraphGraphSupport) GetAssertionOk() (*string, bool)`

GetAssertionOk returns a tuple with the Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertion

`func (o *GraphGraphSupport) SetAssertion(v string)`

SetAssertion sets Assertion field to given value.

### HasAssertion

`func (o *GraphGraphSupport) HasAssertion() bool`

HasAssertion returns a boolean if a field has been set.

### GetObject

`func (o *GraphGraphSupport) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GraphGraphSupport) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GraphGraphSupport) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GraphGraphSupport) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPredicate

`func (o *GraphGraphSupport) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *GraphGraphSupport) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *GraphGraphSupport) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *GraphGraphSupport) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.

### GetSubject

`func (o *GraphGraphSupport) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GraphGraphSupport) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GraphGraphSupport) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GraphGraphSupport) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


