# LinkRoutePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | Pointer to [**[]LinkRouteCandidate**](LinkRouteCandidate.md) | Candidates is every linked account in preference order: subscriptions first, then metered api-key accounts as the backstop. | [optional] 
**GeneratedAt** | Pointer to **string** | GeneratedAt is when the plan was computed, RFC 3339 UTC. | [optional] 
**Primary** | Pointer to [**LinkRouteCandidate**](LinkRouteCandidate.md) | Primary is the first available candidate; absent when every account is rate-limited. | [optional] 

## Methods

### NewLinkRoutePlan

`func NewLinkRoutePlan() *LinkRoutePlan`

NewLinkRoutePlan instantiates a new LinkRoutePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRoutePlanWithDefaults

`func NewLinkRoutePlanWithDefaults() *LinkRoutePlan`

NewLinkRoutePlanWithDefaults instantiates a new LinkRoutePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *LinkRoutePlan) GetCandidates() []LinkRouteCandidate`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *LinkRoutePlan) GetCandidatesOk() (*[]LinkRouteCandidate, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *LinkRoutePlan) SetCandidates(v []LinkRouteCandidate)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *LinkRoutePlan) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *LinkRoutePlan) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *LinkRoutePlan) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *LinkRoutePlan) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *LinkRoutePlan) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetPrimary

`func (o *LinkRoutePlan) GetPrimary() LinkRouteCandidate`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *LinkRoutePlan) GetPrimaryOk() (*LinkRouteCandidate, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *LinkRoutePlan) SetPrimary(v LinkRouteCandidate)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *LinkRoutePlan) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


