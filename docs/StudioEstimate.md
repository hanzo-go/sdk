# StudioEstimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cents** | Pointer to **int64** |  | [optional] 
**Steps** | Pointer to [**[]StudioStepPrice**](StudioStepPrice.md) |  | [optional] 

## Methods

### NewStudioEstimate

`func NewStudioEstimate() *StudioEstimate`

NewStudioEstimate instantiates a new StudioEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioEstimateWithDefaults

`func NewStudioEstimateWithDefaults() *StudioEstimate`

NewStudioEstimateWithDefaults instantiates a new StudioEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCents

`func (o *StudioEstimate) GetCents() int64`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *StudioEstimate) GetCentsOk() (*int64, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *StudioEstimate) SetCents(v int64)`

SetCents sets Cents field to given value.

### HasCents

`func (o *StudioEstimate) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetSteps

`func (o *StudioEstimate) GetSteps() []StudioStepPrice`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *StudioEstimate) GetStepsOk() (*[]StudioStepPrice, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *StudioEstimate) SetSteps(v []StudioStepPrice)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *StudioEstimate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


