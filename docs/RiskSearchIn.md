# RiskSearchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int64** | Days is how much of the organisation&#39;s own history to replay, 1 to 400. Zero takes thirty. | [optional] 

## Methods

### NewRiskSearchIn

`func NewRiskSearchIn() *RiskSearchIn`

NewRiskSearchIn instantiates a new RiskSearchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSearchInWithDefaults

`func NewRiskSearchInWithDefaults() *RiskSearchIn`

NewRiskSearchInWithDefaults instantiates a new RiskSearchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *RiskSearchIn) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *RiskSearchIn) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *RiskSearchIn) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *RiskSearchIn) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


