# O11yDaemonSetListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableNodes** | Pointer to **int32** |  | [optional] 
**CpuLimit** | Pointer to **float32** |  | [optional] 
**CpuRequest** | Pointer to **float32** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**DaemonSetName** | Pointer to **string** |  | [optional] 
**DesiredNodes** | Pointer to **int32** |  | [optional] 
**MemoryLimit** | Pointer to **float32** |  | [optional] 
**MemoryRequest** | Pointer to **float32** |  | [optional] 
**MemoryUsage** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**Restarts** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yDaemonSetListRecord

`func NewO11yDaemonSetListRecord() *O11yDaemonSetListRecord`

NewO11yDaemonSetListRecord instantiates a new O11yDaemonSetListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDaemonSetListRecordWithDefaults

`func NewO11yDaemonSetListRecordWithDefaults() *O11yDaemonSetListRecord`

NewO11yDaemonSetListRecordWithDefaults instantiates a new O11yDaemonSetListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableNodes

`func (o *O11yDaemonSetListRecord) GetAvailableNodes() int32`

GetAvailableNodes returns the AvailableNodes field if non-nil, zero value otherwise.

### GetAvailableNodesOk

`func (o *O11yDaemonSetListRecord) GetAvailableNodesOk() (*int32, bool)`

GetAvailableNodesOk returns a tuple with the AvailableNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNodes

`func (o *O11yDaemonSetListRecord) SetAvailableNodes(v int32)`

SetAvailableNodes sets AvailableNodes field to given value.

### HasAvailableNodes

`func (o *O11yDaemonSetListRecord) HasAvailableNodes() bool`

HasAvailableNodes returns a boolean if a field has been set.

### GetCpuLimit

`func (o *O11yDaemonSetListRecord) GetCpuLimit() float32`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *O11yDaemonSetListRecord) GetCpuLimitOk() (*float32, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *O11yDaemonSetListRecord) SetCpuLimit(v float32)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *O11yDaemonSetListRecord) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuRequest

`func (o *O11yDaemonSetListRecord) GetCpuRequest() float32`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *O11yDaemonSetListRecord) GetCpuRequestOk() (*float32, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *O11yDaemonSetListRecord) SetCpuRequest(v float32)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *O11yDaemonSetListRecord) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetCpuUsage

`func (o *O11yDaemonSetListRecord) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *O11yDaemonSetListRecord) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *O11yDaemonSetListRecord) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *O11yDaemonSetListRecord) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetDaemonSetName

`func (o *O11yDaemonSetListRecord) GetDaemonSetName() string`

GetDaemonSetName returns the DaemonSetName field if non-nil, zero value otherwise.

### GetDaemonSetNameOk

`func (o *O11yDaemonSetListRecord) GetDaemonSetNameOk() (*string, bool)`

GetDaemonSetNameOk returns a tuple with the DaemonSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetName

`func (o *O11yDaemonSetListRecord) SetDaemonSetName(v string)`

SetDaemonSetName sets DaemonSetName field to given value.

### HasDaemonSetName

`func (o *O11yDaemonSetListRecord) HasDaemonSetName() bool`

HasDaemonSetName returns a boolean if a field has been set.

### GetDesiredNodes

`func (o *O11yDaemonSetListRecord) GetDesiredNodes() int32`

GetDesiredNodes returns the DesiredNodes field if non-nil, zero value otherwise.

### GetDesiredNodesOk

`func (o *O11yDaemonSetListRecord) GetDesiredNodesOk() (*int32, bool)`

GetDesiredNodesOk returns a tuple with the DesiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodes

`func (o *O11yDaemonSetListRecord) SetDesiredNodes(v int32)`

SetDesiredNodes sets DesiredNodes field to given value.

### HasDesiredNodes

`func (o *O11yDaemonSetListRecord) HasDesiredNodes() bool`

HasDesiredNodes returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *O11yDaemonSetListRecord) GetMemoryLimit() float32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *O11yDaemonSetListRecord) GetMemoryLimitOk() (*float32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *O11yDaemonSetListRecord) SetMemoryLimit(v float32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *O11yDaemonSetListRecord) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *O11yDaemonSetListRecord) GetMemoryRequest() float32`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *O11yDaemonSetListRecord) GetMemoryRequestOk() (*float32, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *O11yDaemonSetListRecord) SetMemoryRequest(v float32)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *O11yDaemonSetListRecord) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *O11yDaemonSetListRecord) GetMemoryUsage() float32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *O11yDaemonSetListRecord) GetMemoryUsageOk() (*float32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *O11yDaemonSetListRecord) SetMemoryUsage(v float32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *O11yDaemonSetListRecord) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetMeta

`func (o *O11yDaemonSetListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yDaemonSetListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yDaemonSetListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yDaemonSetListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRestarts

`func (o *O11yDaemonSetListRecord) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *O11yDaemonSetListRecord) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *O11yDaemonSetListRecord) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *O11yDaemonSetListRecord) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


