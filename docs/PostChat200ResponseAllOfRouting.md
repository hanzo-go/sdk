# PostChat200ResponseAllOfRouting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | [**[]PostChat200ResponseAllOfRoutingCandidatesInner**](PostChat200ResponseAllOfRoutingCandidatesInner.md) | Each model the router&#39;s table names, with the reasoning depths that read it, shallow to deep. | 
**Chosen** | **string** | The model the request is served at. | 
**Inputs** | [**PostChat200ResponseAllOfRoutingInputs**](PostChat200ResponseAllOfRoutingInputs.md) |  | 
**Reason** | **string** | How the inputs led to chosen, in a sentence. | 
**Router** | **string** | The router that decided: enso. | 

## Methods

### NewPostChat200ResponseAllOfRouting

`func NewPostChat200ResponseAllOfRouting(candidates []PostChat200ResponseAllOfRoutingCandidatesInner, chosen string, inputs PostChat200ResponseAllOfRoutingInputs, reason string, router string, ) *PostChat200ResponseAllOfRouting`

NewPostChat200ResponseAllOfRouting instantiates a new PostChat200ResponseAllOfRouting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostChat200ResponseAllOfRoutingWithDefaults

`func NewPostChat200ResponseAllOfRoutingWithDefaults() *PostChat200ResponseAllOfRouting`

NewPostChat200ResponseAllOfRoutingWithDefaults instantiates a new PostChat200ResponseAllOfRouting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *PostChat200ResponseAllOfRouting) GetCandidates() []PostChat200ResponseAllOfRoutingCandidatesInner`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *PostChat200ResponseAllOfRouting) GetCandidatesOk() (*[]PostChat200ResponseAllOfRoutingCandidatesInner, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *PostChat200ResponseAllOfRouting) SetCandidates(v []PostChat200ResponseAllOfRoutingCandidatesInner)`

SetCandidates sets Candidates field to given value.


### GetChosen

`func (o *PostChat200ResponseAllOfRouting) GetChosen() string`

GetChosen returns the Chosen field if non-nil, zero value otherwise.

### GetChosenOk

`func (o *PostChat200ResponseAllOfRouting) GetChosenOk() (*string, bool)`

GetChosenOk returns a tuple with the Chosen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosen

`func (o *PostChat200ResponseAllOfRouting) SetChosen(v string)`

SetChosen sets Chosen field to given value.


### GetInputs

`func (o *PostChat200ResponseAllOfRouting) GetInputs() PostChat200ResponseAllOfRoutingInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *PostChat200ResponseAllOfRouting) GetInputsOk() (*PostChat200ResponseAllOfRoutingInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *PostChat200ResponseAllOfRouting) SetInputs(v PostChat200ResponseAllOfRoutingInputs)`

SetInputs sets Inputs field to given value.


### GetReason

`func (o *PostChat200ResponseAllOfRouting) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostChat200ResponseAllOfRouting) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostChat200ResponseAllOfRouting) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRouter

`func (o *PostChat200ResponseAllOfRouting) GetRouter() string`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *PostChat200ResponseAllOfRouting) GetRouterOk() (*string, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *PostChat200ResponseAllOfRouting) SetRouter(v string)`

SetRouter sets Router field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


