# IamObjectPrometheusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLatency** | Pointer to [**[]IamObjectHistogramVecInfo**](IamObjectHistogramVecInfo.md) |  | [optional] 
**ApiThroughput** | Pointer to [**[]IamObjectGaugeVecInfo**](IamObjectGaugeVecInfo.md) |  | [optional] 
**TotalThroughput** | Pointer to **float64** |  | [optional] 

## Methods

### NewIamObjectPrometheusInfo

`func NewIamObjectPrometheusInfo() *IamObjectPrometheusInfo`

NewIamObjectPrometheusInfo instantiates a new IamObjectPrometheusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectPrometheusInfoWithDefaults

`func NewIamObjectPrometheusInfoWithDefaults() *IamObjectPrometheusInfo`

NewIamObjectPrometheusInfoWithDefaults instantiates a new IamObjectPrometheusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiLatency

`func (o *IamObjectPrometheusInfo) GetApiLatency() []IamObjectHistogramVecInfo`

GetApiLatency returns the ApiLatency field if non-nil, zero value otherwise.

### GetApiLatencyOk

`func (o *IamObjectPrometheusInfo) GetApiLatencyOk() (*[]IamObjectHistogramVecInfo, bool)`

GetApiLatencyOk returns a tuple with the ApiLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiLatency

`func (o *IamObjectPrometheusInfo) SetApiLatency(v []IamObjectHistogramVecInfo)`

SetApiLatency sets ApiLatency field to given value.

### HasApiLatency

`func (o *IamObjectPrometheusInfo) HasApiLatency() bool`

HasApiLatency returns a boolean if a field has been set.

### GetApiThroughput

`func (o *IamObjectPrometheusInfo) GetApiThroughput() []IamObjectGaugeVecInfo`

GetApiThroughput returns the ApiThroughput field if non-nil, zero value otherwise.

### GetApiThroughputOk

`func (o *IamObjectPrometheusInfo) GetApiThroughputOk() (*[]IamObjectGaugeVecInfo, bool)`

GetApiThroughputOk returns a tuple with the ApiThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiThroughput

`func (o *IamObjectPrometheusInfo) SetApiThroughput(v []IamObjectGaugeVecInfo)`

SetApiThroughput sets ApiThroughput field to given value.

### HasApiThroughput

`func (o *IamObjectPrometheusInfo) HasApiThroughput() bool`

HasApiThroughput returns a boolean if a field has been set.

### GetTotalThroughput

`func (o *IamObjectPrometheusInfo) GetTotalThroughput() float64`

GetTotalThroughput returns the TotalThroughput field if non-nil, zero value otherwise.

### GetTotalThroughputOk

`func (o *IamObjectPrometheusInfo) GetTotalThroughputOk() (*float64, bool)`

GetTotalThroughputOk returns a tuple with the TotalThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalThroughput

`func (o *IamObjectPrometheusInfo) SetTotalThroughput(v float64)`

SetTotalThroughput sets TotalThroughput field to given value.

### HasTotalThroughput

`func (o *IamObjectPrometheusInfo) HasTotalThroughput() bool`

HasTotalThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


