# ContentStateGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initial** | Pointer to **string** | Initial is the state a fresh document starts in — \&quot;draft\&quot;. A stored document with no status at all is read as this too. | [optional] 
**Live** | Pointer to **string** | Live is the ONE state that is publicly readable — \&quot;published\&quot;. The site pulls only documents in it, so reaching Live IS site-publish; every other state is invisible to a reader. | [optional] 
**States** | Pointer to **[]string** | States is every lifecycle state in canonical order: draft, in_review, approved, queued, published, archived. The console lays its board columns out in exactly this order, so the order is part of the answer. | [optional] 
**Transitions** | Pointer to **map[string][]string** | Transitions maps each state to the states it may move to. A target absent from a state&#39;s list is REFUSED, at the endpoint and again at the storage boundary — this is the whole rule, not a hint for the UI. A state never lists itself; a move that changes nothing is always legal. | [optional] 

## Methods

### NewContentStateGraph

`func NewContentStateGraph() *ContentStateGraph`

NewContentStateGraph instantiates a new ContentStateGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentStateGraphWithDefaults

`func NewContentStateGraphWithDefaults() *ContentStateGraph`

NewContentStateGraphWithDefaults instantiates a new ContentStateGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitial

`func (o *ContentStateGraph) GetInitial() string`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *ContentStateGraph) GetInitialOk() (*string, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *ContentStateGraph) SetInitial(v string)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *ContentStateGraph) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetLive

`func (o *ContentStateGraph) GetLive() string`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *ContentStateGraph) GetLiveOk() (*string, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *ContentStateGraph) SetLive(v string)`

SetLive sets Live field to given value.

### HasLive

`func (o *ContentStateGraph) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetStates

`func (o *ContentStateGraph) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *ContentStateGraph) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *ContentStateGraph) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *ContentStateGraph) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetTransitions

`func (o *ContentStateGraph) GetTransitions() map[string][]string`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *ContentStateGraph) GetTransitionsOk() (*map[string][]string, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *ContentStateGraph) SetTransitions(v map[string][]string)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *ContentStateGraph) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


