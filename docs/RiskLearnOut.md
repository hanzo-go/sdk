# RiskLearnOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Learned** | Pointer to **int64** | Learned is how many of the events the model actually learned from, and is also what the call is metered at: one screen per event learned from. It is the batch minus the events already in this organisation&#39;s record, so a retried batch reports — and is charged — zero. | [optional] 

## Methods

### NewRiskLearnOut

`func NewRiskLearnOut() *RiskLearnOut`

NewRiskLearnOut instantiates a new RiskLearnOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLearnOutWithDefaults

`func NewRiskLearnOutWithDefaults() *RiskLearnOut`

NewRiskLearnOutWithDefaults instantiates a new RiskLearnOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLearned

`func (o *RiskLearnOut) GetLearned() int64`

GetLearned returns the Learned field if non-nil, zero value otherwise.

### GetLearnedOk

`func (o *RiskLearnOut) GetLearnedOk() (*int64, bool)`

GetLearnedOk returns a tuple with the Learned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearned

`func (o *RiskLearnOut) SetLearned(v int64)`

SetLearned sets Learned field to given value.

### HasLearned

`func (o *RiskLearnOut) HasLearned() bool`

HasLearned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


