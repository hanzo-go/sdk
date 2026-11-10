# GraphGraphKinshipScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertion** | Pointer to [**GraphWireFact**](GraphWireFact.md) | Assertion is the underlying scorecard assertion holding at the point. | [optional] 
**Relationship** | Pointer to **string** | Relationship is the scorecard column, such as mother, father, partner, family or familyAndFriends. | [optional] 
**Score** | Pointer to **float64** | Score is the classifier probability in [0,1]. | [optional] 

## Methods

### NewGraphGraphKinshipScore

`func NewGraphGraphKinshipScore() *GraphGraphKinshipScore`

NewGraphGraphKinshipScore instantiates a new GraphGraphKinshipScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphKinshipScoreWithDefaults

`func NewGraphGraphKinshipScoreWithDefaults() *GraphGraphKinshipScore`

NewGraphGraphKinshipScoreWithDefaults instantiates a new GraphGraphKinshipScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertion

`func (o *GraphGraphKinshipScore) GetAssertion() GraphWireFact`

GetAssertion returns the Assertion field if non-nil, zero value otherwise.

### GetAssertionOk

`func (o *GraphGraphKinshipScore) GetAssertionOk() (*GraphWireFact, bool)`

GetAssertionOk returns a tuple with the Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertion

`func (o *GraphGraphKinshipScore) SetAssertion(v GraphWireFact)`

SetAssertion sets Assertion field to given value.

### HasAssertion

`func (o *GraphGraphKinshipScore) HasAssertion() bool`

HasAssertion returns a boolean if a field has been set.

### GetRelationship

`func (o *GraphGraphKinshipScore) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *GraphGraphKinshipScore) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *GraphGraphKinshipScore) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *GraphGraphKinshipScore) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetScore

`func (o *GraphGraphKinshipScore) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GraphGraphKinshipScore) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GraphGraphKinshipScore) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *GraphGraphKinshipScore) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


