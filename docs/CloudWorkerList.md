# CloudWorkerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workers** | Pointer to [**[]CloudByoWorker**](CloudByoWorker.md) | Workers is one row per connected BYO machine, each carrying the host&#39;s own report (GPUs, driver versions, capabilities) rather than a normalized view. | [optional] 

## Methods

### NewCloudWorkerList

`func NewCloudWorkerList() *CloudWorkerList`

NewCloudWorkerList instantiates a new CloudWorkerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWorkerListWithDefaults

`func NewCloudWorkerListWithDefaults() *CloudWorkerList`

NewCloudWorkerListWithDefaults instantiates a new CloudWorkerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkers

`func (o *CloudWorkerList) GetWorkers() []CloudByoWorker`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *CloudWorkerList) GetWorkersOk() (*[]CloudByoWorker, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *CloudWorkerList) SetWorkers(v []CloudByoWorker)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *CloudWorkerList) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


