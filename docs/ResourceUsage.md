# ResourceUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **float64** | CostCents would be the window&#39;s spend in cents. Always null here — the money a run costs is the metering ledger&#39;s, joined by the run id, and repeating it from this side would be a second number that could disagree with the bill. | [optional] 
**CpuVcpuHours** | Pointer to **float64** | CPUVcpuHours would be vCPU-hours over the window. Always null: this store holds agent definitions and run I/O, and nothing here meters a CPU. Null is the honest answer and 0 would be a claim. | [optional] 
**MemGbHours** | Pointer to **float64** | MemGbHours would be gigabyte-hours of memory. Always null, same reason. | [optional] 
**StorageIoBytes** | Pointer to **float64** | StorageIoBytes would be bytes moved to and from storage. Always null, same reason. | [optional] 

## Methods

### NewResourceUsage

`func NewResourceUsage() *ResourceUsage`

NewResourceUsage instantiates a new ResourceUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsageWithDefaults

`func NewResourceUsageWithDefaults() *ResourceUsage`

NewResourceUsageWithDefaults instantiates a new ResourceUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *ResourceUsage) GetCostCents() float64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ResourceUsage) GetCostCentsOk() (*float64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ResourceUsage) SetCostCents(v float64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ResourceUsage) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCpuVcpuHours

`func (o *ResourceUsage) GetCpuVcpuHours() float64`

GetCpuVcpuHours returns the CpuVcpuHours field if non-nil, zero value otherwise.

### GetCpuVcpuHoursOk

`func (o *ResourceUsage) GetCpuVcpuHoursOk() (*float64, bool)`

GetCpuVcpuHoursOk returns a tuple with the CpuVcpuHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuVcpuHours

`func (o *ResourceUsage) SetCpuVcpuHours(v float64)`

SetCpuVcpuHours sets CpuVcpuHours field to given value.

### HasCpuVcpuHours

`func (o *ResourceUsage) HasCpuVcpuHours() bool`

HasCpuVcpuHours returns a boolean if a field has been set.

### GetMemGbHours

`func (o *ResourceUsage) GetMemGbHours() float64`

GetMemGbHours returns the MemGbHours field if non-nil, zero value otherwise.

### GetMemGbHoursOk

`func (o *ResourceUsage) GetMemGbHoursOk() (*float64, bool)`

GetMemGbHoursOk returns a tuple with the MemGbHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemGbHours

`func (o *ResourceUsage) SetMemGbHours(v float64)`

SetMemGbHours sets MemGbHours field to given value.

### HasMemGbHours

`func (o *ResourceUsage) HasMemGbHours() bool`

HasMemGbHours returns a boolean if a field has been set.

### GetStorageIoBytes

`func (o *ResourceUsage) GetStorageIoBytes() float64`

GetStorageIoBytes returns the StorageIoBytes field if non-nil, zero value otherwise.

### GetStorageIoBytesOk

`func (o *ResourceUsage) GetStorageIoBytesOk() (*float64, bool)`

GetStorageIoBytesOk returns a tuple with the StorageIoBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageIoBytes

`func (o *ResourceUsage) SetStorageIoBytes(v float64)`

SetStorageIoBytes sets StorageIoBytes field to given value.

### HasStorageIoBytes

`func (o *ResourceUsage) HasStorageIoBytes() bool`

HasStorageIoBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


