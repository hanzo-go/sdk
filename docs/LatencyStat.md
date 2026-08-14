# LatencyStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | false when no GenAI spans carry timing; the percentiles are then null | [optional] 
**P50Ms** | Pointer to **float32** | median latency over the window | [optional] 
**P95Ms** | Pointer to **float32** | 95th-percentile latency | [optional] 
**P99Ms** | Pointer to **float32** | 99th-percentile latency | [optional] 

## Methods

### NewLatencyStat

`func NewLatencyStat() *LatencyStat`

NewLatencyStat instantiates a new LatencyStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatencyStatWithDefaults

`func NewLatencyStatWithDefaults() *LatencyStat`

NewLatencyStatWithDefaults instantiates a new LatencyStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LatencyStat) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LatencyStat) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LatencyStat) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LatencyStat) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetP50Ms

`func (o *LatencyStat) GetP50Ms() float32`

GetP50Ms returns the P50Ms field if non-nil, zero value otherwise.

### GetP50MsOk

`func (o *LatencyStat) GetP50MsOk() (*float32, bool)`

GetP50MsOk returns a tuple with the P50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50Ms

`func (o *LatencyStat) SetP50Ms(v float32)`

SetP50Ms sets P50Ms field to given value.

### HasP50Ms

`func (o *LatencyStat) HasP50Ms() bool`

HasP50Ms returns a boolean if a field has been set.

### GetP95Ms

`func (o *LatencyStat) GetP95Ms() float32`

GetP95Ms returns the P95Ms field if non-nil, zero value otherwise.

### GetP95MsOk

`func (o *LatencyStat) GetP95MsOk() (*float32, bool)`

GetP95MsOk returns a tuple with the P95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95Ms

`func (o *LatencyStat) SetP95Ms(v float32)`

SetP95Ms sets P95Ms field to given value.

### HasP95Ms

`func (o *LatencyStat) HasP95Ms() bool`

HasP95Ms returns a boolean if a field has been set.

### GetP99Ms

`func (o *LatencyStat) GetP99Ms() float32`

GetP99Ms returns the P99Ms field if non-nil, zero value otherwise.

### GetP99MsOk

`func (o *LatencyStat) GetP99MsOk() (*float32, bool)`

GetP99MsOk returns a tuple with the P99Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99Ms

`func (o *LatencyStat) SetP99Ms(v float32)`

SetP99Ms sets P99Ms field to given value.

### HasP99Ms

`func (o *LatencyStat) HasP99Ms() bool`

HasP99Ms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


