# LeaderboardStandingNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is the points at which it unlocks. | [optional] 
**More** | Pointer to **int64** | More is At minus the reader&#39;s points: what is left to earn. | [optional] 
**Reveal** | Pointer to **int64** | Reveal is how many more ranks above the reader that rung unlocks. | [optional] 

## Methods

### NewLeaderboardStandingNext

`func NewLeaderboardStandingNext() *LeaderboardStandingNext`

NewLeaderboardStandingNext instantiates a new LeaderboardStandingNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardStandingNextWithDefaults

`func NewLeaderboardStandingNextWithDefaults() *LeaderboardStandingNext`

NewLeaderboardStandingNextWithDefaults instantiates a new LeaderboardStandingNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *LeaderboardStandingNext) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *LeaderboardStandingNext) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *LeaderboardStandingNext) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *LeaderboardStandingNext) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetMore

`func (o *LeaderboardStandingNext) GetMore() int64`

GetMore returns the More field if non-nil, zero value otherwise.

### GetMoreOk

`func (o *LeaderboardStandingNext) GetMoreOk() (*int64, bool)`

GetMoreOk returns a tuple with the More field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMore

`func (o *LeaderboardStandingNext) SetMore(v int64)`

SetMore sets More field to given value.

### HasMore

`func (o *LeaderboardStandingNext) HasMore() bool`

HasMore returns a boolean if a field has been set.

### GetReveal

`func (o *LeaderboardStandingNext) GetReveal() int64`

GetReveal returns the Reveal field if non-nil, zero value otherwise.

### GetRevealOk

`func (o *LeaderboardStandingNext) GetRevealOk() (*int64, bool)`

GetRevealOk returns a tuple with the Reveal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReveal

`func (o *LeaderboardStandingNext) SetReveal(v int64)`

SetReveal sets Reveal field to given value.

### HasReveal

`func (o *LeaderboardStandingNext) HasReveal() bool`

HasReveal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


