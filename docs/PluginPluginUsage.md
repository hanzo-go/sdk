# PluginPluginUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuNs** | Pointer to **int64** | CPU is total user+system time this instance has consumed since it started, in nanoseconds. It only ever climbs, so a rate is the interesting derivative.  | [optional] 
**RssBytes** | Pointer to **int64** | RSS is resident memory in bytes — the honest number for \&quot;what does this plugin cost\&quot;, as opposed to virtual size.  | [optional] 
**Threads** | Pointer to **int32** | Threads and FDs are the two limits a busy service hits first, and both are leaks when they climb without bound.  | [optional] 
**Fds** | Pointer to **int32** |  | [optional] 

## Methods

### NewPluginPluginUsage

`func NewPluginPluginUsage() *PluginPluginUsage`

NewPluginPluginUsage instantiates a new PluginPluginUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginPluginUsageWithDefaults

`func NewPluginPluginUsageWithDefaults() *PluginPluginUsage`

NewPluginPluginUsageWithDefaults instantiates a new PluginPluginUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuNs

`func (o *PluginPluginUsage) GetCpuNs() int64`

GetCpuNs returns the CpuNs field if non-nil, zero value otherwise.

### GetCpuNsOk

`func (o *PluginPluginUsage) GetCpuNsOk() (*int64, bool)`

GetCpuNsOk returns a tuple with the CpuNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuNs

`func (o *PluginPluginUsage) SetCpuNs(v int64)`

SetCpuNs sets CpuNs field to given value.

### HasCpuNs

`func (o *PluginPluginUsage) HasCpuNs() bool`

HasCpuNs returns a boolean if a field has been set.

### GetRssBytes

`func (o *PluginPluginUsage) GetRssBytes() int64`

GetRssBytes returns the RssBytes field if non-nil, zero value otherwise.

### GetRssBytesOk

`func (o *PluginPluginUsage) GetRssBytesOk() (*int64, bool)`

GetRssBytesOk returns a tuple with the RssBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssBytes

`func (o *PluginPluginUsage) SetRssBytes(v int64)`

SetRssBytes sets RssBytes field to given value.

### HasRssBytes

`func (o *PluginPluginUsage) HasRssBytes() bool`

HasRssBytes returns a boolean if a field has been set.

### GetThreads

`func (o *PluginPluginUsage) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *PluginPluginUsage) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *PluginPluginUsage) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *PluginPluginUsage) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetFds

`func (o *PluginPluginUsage) GetFds() int32`

GetFds returns the Fds field if non-nil, zero value otherwise.

### GetFdsOk

`func (o *PluginPluginUsage) GetFdsOk() (*int32, bool)`

GetFdsOk returns a tuple with the Fds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFds

`func (o *PluginPluginUsage) SetFds(v int32)`

SetFds sets Fds field to given value.

### HasFds

`func (o *PluginPluginUsage) HasFds() bool`

HasFds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


