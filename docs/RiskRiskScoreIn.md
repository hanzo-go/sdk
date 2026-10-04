# RiskRiskScoreIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**RiskRiskEvent**](RiskRiskEvent.md) | Event is the thing to judge. It is judged against the caller&#39;s OWN model and nothing is learned from it. | [optional] 

## Methods

### NewRiskRiskScoreIn

`func NewRiskRiskScoreIn() *RiskRiskScoreIn`

NewRiskRiskScoreIn instantiates a new RiskRiskScoreIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskScoreInWithDefaults

`func NewRiskRiskScoreInWithDefaults() *RiskRiskScoreIn`

NewRiskRiskScoreInWithDefaults instantiates a new RiskRiskScoreIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *RiskRiskScoreIn) GetEvent() RiskRiskEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RiskRiskScoreIn) GetEventOk() (*RiskRiskEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RiskRiskScoreIn) SetEvent(v RiskRiskEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *RiskRiskScoreIn) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


