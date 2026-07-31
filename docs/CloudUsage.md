# CloudUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuNs** | Pointer to **int32** | CPU is total user+system time this instance has consumed since it started. It only ever climbs, so a rate is the interesting derivative. | [optional] 
**Fds** | Pointer to **int32** |  | [optional] 
**RssBytes** | Pointer to **int32** | RSS is resident memory in bytes — the honest number for \&quot;what does this plugin cost\&quot;, as opposed to virtual size. | [optional] 
**Threads** | Pointer to **int32** | Threads and FDs are the two limits a busy service hits first, and both are leaks when they climb without bound. | [optional] 

## Methods

### NewCloudUsage

`func NewCloudUsage() *CloudUsage`

NewCloudUsage instantiates a new CloudUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUsageWithDefaults

`func NewCloudUsageWithDefaults() *CloudUsage`

NewCloudUsageWithDefaults instantiates a new CloudUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuNs

`func (o *CloudUsage) GetCpuNs() int32`

GetCpuNs returns the CpuNs field if non-nil, zero value otherwise.

### GetCpuNsOk

`func (o *CloudUsage) GetCpuNsOk() (*int32, bool)`

GetCpuNsOk returns a tuple with the CpuNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuNs

`func (o *CloudUsage) SetCpuNs(v int32)`

SetCpuNs sets CpuNs field to given value.

### HasCpuNs

`func (o *CloudUsage) HasCpuNs() bool`

HasCpuNs returns a boolean if a field has been set.

### GetFds

`func (o *CloudUsage) GetFds() int32`

GetFds returns the Fds field if non-nil, zero value otherwise.

### GetFdsOk

`func (o *CloudUsage) GetFdsOk() (*int32, bool)`

GetFdsOk returns a tuple with the Fds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFds

`func (o *CloudUsage) SetFds(v int32)`

SetFds sets Fds field to given value.

### HasFds

`func (o *CloudUsage) HasFds() bool`

HasFds returns a boolean if a field has been set.

### GetRssBytes

`func (o *CloudUsage) GetRssBytes() int32`

GetRssBytes returns the RssBytes field if non-nil, zero value otherwise.

### GetRssBytesOk

`func (o *CloudUsage) GetRssBytesOk() (*int32, bool)`

GetRssBytesOk returns a tuple with the RssBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssBytes

`func (o *CloudUsage) SetRssBytes(v int32)`

SetRssBytes sets RssBytes field to given value.

### HasRssBytes

`func (o *CloudUsage) HasRssBytes() bool`

HasRssBytes returns a boolean if a field has been set.

### GetThreads

`func (o *CloudUsage) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *CloudUsage) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *CloudUsage) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *CloudUsage) HasThreads() bool`

HasThreads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


