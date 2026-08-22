# StepList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Step**](Step.md) | Data is every step of the sequence, idx ascending — the order they send in. It is not paged: a sequence&#39;s steps are a handful, and a partial list would misstate the drip. An empty array for a sequence with no messages yet, which enrolls fine and completes immediately. | [optional] 

## Methods

### NewStepList

`func NewStepList() *StepList`

NewStepList instantiates a new StepList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepListWithDefaults

`func NewStepListWithDefaults() *StepList`

NewStepListWithDefaults instantiates a new StepList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *StepList) GetData() []Step`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StepList) GetDataOk() (*[]Step, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StepList) SetData(v []Step)`

SetData sets Data field to given value.

### HasData

`func (o *StepList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


