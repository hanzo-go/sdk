# ExperimentDecideBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Winner** | **string** | Winner is the variant to promote. It must name one of this experiment&#39;s own arms. | 

## Methods

### NewExperimentDecideBody

`func NewExperimentDecideBody(winner string, ) *ExperimentDecideBody`

NewExperimentDecideBody instantiates a new ExperimentDecideBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentDecideBodyWithDefaults

`func NewExperimentDecideBodyWithDefaults() *ExperimentDecideBody`

NewExperimentDecideBodyWithDefaults instantiates a new ExperimentDecideBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWinner

`func (o *ExperimentDecideBody) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ExperimentDecideBody) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ExperimentDecideBody) SetWinner(v string)`

SetWinner sets Winner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


