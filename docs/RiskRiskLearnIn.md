# RiskRiskLearnIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]RiskRiskEvent**](RiskRiskEvent.md) | Events are the things that happened, oldest first. An empty batch is refused: learning nothing is not an operation. | [optional] 

## Methods

### NewRiskRiskLearnIn

`func NewRiskRiskLearnIn() *RiskRiskLearnIn`

NewRiskRiskLearnIn instantiates a new RiskRiskLearnIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskLearnInWithDefaults

`func NewRiskRiskLearnInWithDefaults() *RiskRiskLearnIn`

NewRiskRiskLearnInWithDefaults instantiates a new RiskRiskLearnIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *RiskRiskLearnIn) GetEvents() []RiskRiskEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RiskRiskLearnIn) GetEventsOk() (*[]RiskRiskEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RiskRiskLearnIn) SetEvents(v []RiskRiskEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RiskRiskLearnIn) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


