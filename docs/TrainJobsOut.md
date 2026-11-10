# TrainJobsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TrainJob**](TrainJob.md) | Data is the org&#39;s jobs, newest first. | [optional] 

## Methods

### NewTrainJobsOut

`func NewTrainJobsOut() *TrainJobsOut`

NewTrainJobsOut instantiates a new TrainJobsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainJobsOutWithDefaults

`func NewTrainJobsOutWithDefaults() *TrainJobsOut`

NewTrainJobsOutWithDefaults instantiates a new TrainJobsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TrainJobsOut) GetData() []TrainJob`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrainJobsOut) GetDataOk() (*[]TrainJob, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrainJobsOut) SetData(v []TrainJob)`

SetData sets Data field to given value.

### HasData

`func (o *TrainJobsOut) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


