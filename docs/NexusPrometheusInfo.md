# NexusPrometheusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLatency** | Pointer to [**[]NexusHistogramVecInfo**](NexusHistogramVecInfo.md) |  | [optional] 
**ApiThroughput** | Pointer to [**[]NexusGaugeVecInfo**](NexusGaugeVecInfo.md) |  | [optional] 
**TotalThroughput** | Pointer to **float64** |  | [optional] 

## Methods

### NewNexusPrometheusInfo

`func NewNexusPrometheusInfo() *NexusPrometheusInfo`

NewNexusPrometheusInfo instantiates a new NexusPrometheusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusPrometheusInfoWithDefaults

`func NewNexusPrometheusInfoWithDefaults() *NexusPrometheusInfo`

NewNexusPrometheusInfoWithDefaults instantiates a new NexusPrometheusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiLatency

`func (o *NexusPrometheusInfo) GetApiLatency() []NexusHistogramVecInfo`

GetApiLatency returns the ApiLatency field if non-nil, zero value otherwise.

### GetApiLatencyOk

`func (o *NexusPrometheusInfo) GetApiLatencyOk() (*[]NexusHistogramVecInfo, bool)`

GetApiLatencyOk returns a tuple with the ApiLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiLatency

`func (o *NexusPrometheusInfo) SetApiLatency(v []NexusHistogramVecInfo)`

SetApiLatency sets ApiLatency field to given value.

### HasApiLatency

`func (o *NexusPrometheusInfo) HasApiLatency() bool`

HasApiLatency returns a boolean if a field has been set.

### GetApiThroughput

`func (o *NexusPrometheusInfo) GetApiThroughput() []NexusGaugeVecInfo`

GetApiThroughput returns the ApiThroughput field if non-nil, zero value otherwise.

### GetApiThroughputOk

`func (o *NexusPrometheusInfo) GetApiThroughputOk() (*[]NexusGaugeVecInfo, bool)`

GetApiThroughputOk returns a tuple with the ApiThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiThroughput

`func (o *NexusPrometheusInfo) SetApiThroughput(v []NexusGaugeVecInfo)`

SetApiThroughput sets ApiThroughput field to given value.

### HasApiThroughput

`func (o *NexusPrometheusInfo) HasApiThroughput() bool`

HasApiThroughput returns a boolean if a field has been set.

### GetTotalThroughput

`func (o *NexusPrometheusInfo) GetTotalThroughput() float64`

GetTotalThroughput returns the TotalThroughput field if non-nil, zero value otherwise.

### GetTotalThroughputOk

`func (o *NexusPrometheusInfo) GetTotalThroughputOk() (*float64, bool)`

GetTotalThroughputOk returns a tuple with the TotalThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalThroughput

`func (o *NexusPrometheusInfo) SetTotalThroughput(v float64)`

SetTotalThroughput sets TotalThroughput field to given value.

### HasTotalThroughput

`func (o *NexusPrometheusInfo) HasTotalThroughput() bool`

HasTotalThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


