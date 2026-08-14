# RiskScoreIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**RiskEvent**](RiskEvent.md) | Event is the thing to judge. It is judged against the caller&#39;s OWN model and nothing is learned from it. | [optional] 

## Methods

### NewRiskScoreIn

`func NewRiskScoreIn() *RiskScoreIn`

NewRiskScoreIn instantiates a new RiskScoreIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskScoreInWithDefaults

`func NewRiskScoreInWithDefaults() *RiskScoreIn`

NewRiskScoreInWithDefaults instantiates a new RiskScoreIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *RiskScoreIn) GetEvent() RiskEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RiskScoreIn) GetEventOk() (*RiskEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RiskScoreIn) SetEvent(v RiskEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *RiskScoreIn) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


