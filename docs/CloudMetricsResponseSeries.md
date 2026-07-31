# CloudMetricsResponseSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]CloudPoint**](CloudPoint.md) |  | [optional] 
**LatencyP50Ms** | Pointer to [**[]CloudPoint**](CloudPoint.md) |  | [optional] 
**LatencyP95Ms** | Pointer to [**[]CloudPoint**](CloudPoint.md) |  | [optional] 
**Requests** | Pointer to [**[]CloudPoint**](CloudPoint.md) |  | [optional] 

## Methods

### NewCloudMetricsResponseSeries

`func NewCloudMetricsResponseSeries() *CloudMetricsResponseSeries`

NewCloudMetricsResponseSeries instantiates a new CloudMetricsResponseSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMetricsResponseSeriesWithDefaults

`func NewCloudMetricsResponseSeriesWithDefaults() *CloudMetricsResponseSeries`

NewCloudMetricsResponseSeriesWithDefaults instantiates a new CloudMetricsResponseSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *CloudMetricsResponseSeries) GetErrors() []CloudPoint`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CloudMetricsResponseSeries) GetErrorsOk() (*[]CloudPoint, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CloudMetricsResponseSeries) SetErrors(v []CloudPoint)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CloudMetricsResponseSeries) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLatencyP50Ms

`func (o *CloudMetricsResponseSeries) GetLatencyP50Ms() []CloudPoint`

GetLatencyP50Ms returns the LatencyP50Ms field if non-nil, zero value otherwise.

### GetLatencyP50MsOk

`func (o *CloudMetricsResponseSeries) GetLatencyP50MsOk() (*[]CloudPoint, bool)`

GetLatencyP50MsOk returns a tuple with the LatencyP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP50Ms

`func (o *CloudMetricsResponseSeries) SetLatencyP50Ms(v []CloudPoint)`

SetLatencyP50Ms sets LatencyP50Ms field to given value.

### HasLatencyP50Ms

`func (o *CloudMetricsResponseSeries) HasLatencyP50Ms() bool`

HasLatencyP50Ms returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *CloudMetricsResponseSeries) GetLatencyP95Ms() []CloudPoint`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *CloudMetricsResponseSeries) GetLatencyP95MsOk() (*[]CloudPoint, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *CloudMetricsResponseSeries) SetLatencyP95Ms(v []CloudPoint)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *CloudMetricsResponseSeries) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.

### GetRequests

`func (o *CloudMetricsResponseSeries) GetRequests() []CloudPoint`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudMetricsResponseSeries) GetRequestsOk() (*[]CloudPoint, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudMetricsResponseSeries) SetRequests(v []CloudPoint)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudMetricsResponseSeries) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


