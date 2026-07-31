# CloudObjectPrometheusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLatency** | Pointer to [**[]CloudObjectHistogramVecInfo**](CloudObjectHistogramVecInfo.md) |  | [optional] 
**ApiThroughput** | Pointer to [**[]CloudObjectGaugeVecInfo**](CloudObjectGaugeVecInfo.md) |  | [optional] 
**TotalThroughput** | Pointer to **float64** |  | [optional] 

## Methods

### NewCloudObjectPrometheusInfo

`func NewCloudObjectPrometheusInfo() *CloudObjectPrometheusInfo`

NewCloudObjectPrometheusInfo instantiates a new CloudObjectPrometheusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectPrometheusInfoWithDefaults

`func NewCloudObjectPrometheusInfoWithDefaults() *CloudObjectPrometheusInfo`

NewCloudObjectPrometheusInfoWithDefaults instantiates a new CloudObjectPrometheusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiLatency

`func (o *CloudObjectPrometheusInfo) GetApiLatency() []CloudObjectHistogramVecInfo`

GetApiLatency returns the ApiLatency field if non-nil, zero value otherwise.

### GetApiLatencyOk

`func (o *CloudObjectPrometheusInfo) GetApiLatencyOk() (*[]CloudObjectHistogramVecInfo, bool)`

GetApiLatencyOk returns a tuple with the ApiLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiLatency

`func (o *CloudObjectPrometheusInfo) SetApiLatency(v []CloudObjectHistogramVecInfo)`

SetApiLatency sets ApiLatency field to given value.

### HasApiLatency

`func (o *CloudObjectPrometheusInfo) HasApiLatency() bool`

HasApiLatency returns a boolean if a field has been set.

### GetApiThroughput

`func (o *CloudObjectPrometheusInfo) GetApiThroughput() []CloudObjectGaugeVecInfo`

GetApiThroughput returns the ApiThroughput field if non-nil, zero value otherwise.

### GetApiThroughputOk

`func (o *CloudObjectPrometheusInfo) GetApiThroughputOk() (*[]CloudObjectGaugeVecInfo, bool)`

GetApiThroughputOk returns a tuple with the ApiThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiThroughput

`func (o *CloudObjectPrometheusInfo) SetApiThroughput(v []CloudObjectGaugeVecInfo)`

SetApiThroughput sets ApiThroughput field to given value.

### HasApiThroughput

`func (o *CloudObjectPrometheusInfo) HasApiThroughput() bool`

HasApiThroughput returns a boolean if a field has been set.

### GetTotalThroughput

`func (o *CloudObjectPrometheusInfo) GetTotalThroughput() float64`

GetTotalThroughput returns the TotalThroughput field if non-nil, zero value otherwise.

### GetTotalThroughputOk

`func (o *CloudObjectPrometheusInfo) GetTotalThroughputOk() (*float64, bool)`

GetTotalThroughputOk returns a tuple with the TotalThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalThroughput

`func (o *CloudObjectPrometheusInfo) SetTotalThroughput(v float64)`

SetTotalThroughput sets TotalThroughput field to given value.

### HasTotalThroughput

`func (o *CloudObjectPrometheusInfo) HasTotalThroughput() bool`

HasTotalThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


