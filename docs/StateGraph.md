# StateGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initial** | Pointer to **string** |  | [optional] 
**Live** | Pointer to **string** |  | [optional] 
**States** | Pointer to **[]string** |  | [optional] 
**Transitions** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewStateGraph

`func NewStateGraph() *StateGraph`

NewStateGraph instantiates a new StateGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateGraphWithDefaults

`func NewStateGraphWithDefaults() *StateGraph`

NewStateGraphWithDefaults instantiates a new StateGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitial

`func (o *StateGraph) GetInitial() string`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *StateGraph) GetInitialOk() (*string, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *StateGraph) SetInitial(v string)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *StateGraph) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetLive

`func (o *StateGraph) GetLive() string`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *StateGraph) GetLiveOk() (*string, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *StateGraph) SetLive(v string)`

SetLive sets Live field to given value.

### HasLive

`func (o *StateGraph) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetStates

`func (o *StateGraph) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *StateGraph) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *StateGraph) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *StateGraph) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetTransitions

`func (o *StateGraph) GetTransitions() map[string][]string`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *StateGraph) GetTransitionsOk() (*map[string][]string, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *StateGraph) SetTransitions(v map[string][]string)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *StateGraph) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


