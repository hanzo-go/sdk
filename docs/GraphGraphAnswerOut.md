# GraphGraphAnswerOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is the answer, in prose. | [optional] 
**AsKnown** | Pointer to **string** | AsKnown is the knowledge instant it was taken at, RFC 3339, the same way. | [optional] 
**AsOf** | Pointer to **string** | AsOf is the instant the graph was read at, RFC 3339: the one asked for, or the server&#39;s clock when none was. | [optional] 
**Asked** | Pointer to **int64** | Asked is how many communities were read. | [optional] 
**Assertions** | Pointer to **[]string** | Assertions are the IDs of every assertion the answer rests on, ascending. | [optional] 
**Calls** | Pointer to **int64** | Calls is how many model calls the answer made, each billed to the calling organization. | [optional] 
**Clipped** | Pointer to **bool** | Clipped says the answer did not read all it gathered: a community held more facts than one call carries, or more members than one answer reads, and kept the ones the question named and then the most connected; a member held more than 64 assertions and kept its newest; or the findings held more than the answering call carries and kept the highest scored. | [optional] 
**Communities** | Pointer to **[]int64** | Communities are every community the answer rests on, ascending. | [optional] 
**Findings** | Pointer to [**[]GraphGraphFinding**](GraphGraphFinding.md) | Findings are the findings the answer cites, most relevant first. | [optional] 
**Grounded** | Pointer to **bool** | Grounded is false when no community held a fact bearing on the question, or the findings did not answer it. Answer is then empty: an answer resting on nothing is not one this plane gives. | [optional] 
**Skipped** | Pointer to **int64** | Skipped is how many were not, because one answer makes at most eight map calls and reads at most 2048 entities. They are the ones the question&#39;s words ranked last. | [optional] 

## Methods

### NewGraphGraphAnswerOut

`func NewGraphGraphAnswerOut() *GraphGraphAnswerOut`

NewGraphGraphAnswerOut instantiates a new GraphGraphAnswerOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphAnswerOutWithDefaults

`func NewGraphGraphAnswerOutWithDefaults() *GraphGraphAnswerOut`

NewGraphGraphAnswerOutWithDefaults instantiates a new GraphGraphAnswerOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *GraphGraphAnswerOut) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *GraphGraphAnswerOut) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *GraphGraphAnswerOut) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *GraphGraphAnswerOut) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetAsKnown

`func (o *GraphGraphAnswerOut) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphAnswerOut) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphAnswerOut) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphAnswerOut) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphAnswerOut) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphAnswerOut) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphAnswerOut) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphAnswerOut) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetAsked

`func (o *GraphGraphAnswerOut) GetAsked() int64`

GetAsked returns the Asked field if non-nil, zero value otherwise.

### GetAskedOk

`func (o *GraphGraphAnswerOut) GetAskedOk() (*int64, bool)`

GetAskedOk returns a tuple with the Asked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsked

`func (o *GraphGraphAnswerOut) SetAsked(v int64)`

SetAsked sets Asked field to given value.

### HasAsked

`func (o *GraphGraphAnswerOut) HasAsked() bool`

HasAsked returns a boolean if a field has been set.

### GetAssertions

`func (o *GraphGraphAnswerOut) GetAssertions() []string`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *GraphGraphAnswerOut) GetAssertionsOk() (*[]string, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *GraphGraphAnswerOut) SetAssertions(v []string)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *GraphGraphAnswerOut) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.

### GetCalls

`func (o *GraphGraphAnswerOut) GetCalls() int64`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *GraphGraphAnswerOut) GetCallsOk() (*int64, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *GraphGraphAnswerOut) SetCalls(v int64)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *GraphGraphAnswerOut) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetClipped

`func (o *GraphGraphAnswerOut) GetClipped() bool`

GetClipped returns the Clipped field if non-nil, zero value otherwise.

### GetClippedOk

`func (o *GraphGraphAnswerOut) GetClippedOk() (*bool, bool)`

GetClippedOk returns a tuple with the Clipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipped

`func (o *GraphGraphAnswerOut) SetClipped(v bool)`

SetClipped sets Clipped field to given value.

### HasClipped

`func (o *GraphGraphAnswerOut) HasClipped() bool`

HasClipped returns a boolean if a field has been set.

### GetCommunities

`func (o *GraphGraphAnswerOut) GetCommunities() []int64`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *GraphGraphAnswerOut) GetCommunitiesOk() (*[]int64, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *GraphGraphAnswerOut) SetCommunities(v []int64)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *GraphGraphAnswerOut) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.

### GetFindings

`func (o *GraphGraphAnswerOut) GetFindings() []GraphGraphFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *GraphGraphAnswerOut) GetFindingsOk() (*[]GraphGraphFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *GraphGraphAnswerOut) SetFindings(v []GraphGraphFinding)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *GraphGraphAnswerOut) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetGrounded

`func (o *GraphGraphAnswerOut) GetGrounded() bool`

GetGrounded returns the Grounded field if non-nil, zero value otherwise.

### GetGroundedOk

`func (o *GraphGraphAnswerOut) GetGroundedOk() (*bool, bool)`

GetGroundedOk returns a tuple with the Grounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrounded

`func (o *GraphGraphAnswerOut) SetGrounded(v bool)`

SetGrounded sets Grounded field to given value.

### HasGrounded

`func (o *GraphGraphAnswerOut) HasGrounded() bool`

HasGrounded returns a boolean if a field has been set.

### GetSkipped

`func (o *GraphGraphAnswerOut) GetSkipped() int64`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *GraphGraphAnswerOut) GetSkippedOk() (*int64, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *GraphGraphAnswerOut) SetSkipped(v int64)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *GraphGraphAnswerOut) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


