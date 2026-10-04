# ComputeWorkerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workers** | Pointer to [**[]ComputeByoWorker**](ComputeByoWorker.md) | Workers is one row per connected BYO machine, each carrying the host&#39;s own report (GPUs, driver versions, capabilities) rather than a normalized view. | [optional] 

## Methods

### NewComputeWorkerList

`func NewComputeWorkerList() *ComputeWorkerList`

NewComputeWorkerList instantiates a new ComputeWorkerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeWorkerListWithDefaults

`func NewComputeWorkerListWithDefaults() *ComputeWorkerList`

NewComputeWorkerListWithDefaults instantiates a new ComputeWorkerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkers

`func (o *ComputeWorkerList) GetWorkers() []ComputeByoWorker`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *ComputeWorkerList) GetWorkersOk() (*[]ComputeByoWorker, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *ComputeWorkerList) SetWorkers(v []ComputeByoWorker)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *ComputeWorkerList) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


