# GraphGraphFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | Pointer to **[]string** | Assertions are the IDs of the assertions it cites, each a row a read of this graph returns. | [optional] 
**Communities** | Pointer to **[]int64** | Communities are the communities of the facts it cites, at level 0 of the partition at AsOf — the ids POST /v1/graph/communities answers with. | [optional] 
**Score** | Pointer to **int64** | Score is how much the model judged it helps, 1 to 100. | [optional] 
**Text** | Pointer to **string** | Text is what the model found, in its words. | [optional] 

## Methods

### NewGraphGraphFinding

`func NewGraphGraphFinding() *GraphGraphFinding`

NewGraphGraphFinding instantiates a new GraphGraphFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphFindingWithDefaults

`func NewGraphGraphFindingWithDefaults() *GraphGraphFinding`

NewGraphGraphFindingWithDefaults instantiates a new GraphGraphFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *GraphGraphFinding) GetAssertions() []string`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *GraphGraphFinding) GetAssertionsOk() (*[]string, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *GraphGraphFinding) SetAssertions(v []string)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *GraphGraphFinding) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.

### GetCommunities

`func (o *GraphGraphFinding) GetCommunities() []int64`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *GraphGraphFinding) GetCommunitiesOk() (*[]int64, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *GraphGraphFinding) SetCommunities(v []int64)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *GraphGraphFinding) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.

### GetScore

`func (o *GraphGraphFinding) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GraphGraphFinding) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GraphGraphFinding) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *GraphGraphFinding) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetText

`func (o *GraphGraphFinding) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GraphGraphFinding) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GraphGraphFinding) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GraphGraphFinding) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


