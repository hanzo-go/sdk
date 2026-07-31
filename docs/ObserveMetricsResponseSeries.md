# ObserveMetricsResponseSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]ObservePoint**](ObservePoint.md) |  | [optional] 
**Errors** | Pointer to [**[]ObservePoint**](ObservePoint.md) |  | [optional] 
**LatencyP50Ms** | Pointer to [**[]ObservePoint**](ObservePoint.md) |  | [optional] 
**LatencyP95Ms** | Pointer to [**[]ObservePoint**](ObservePoint.md) |  | [optional] 

## Methods

### NewObserveMetricsResponseSeries

`func NewObserveMetricsResponseSeries() *ObserveMetricsResponseSeries`

NewObserveMetricsResponseSeries instantiates a new ObserveMetricsResponseSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveMetricsResponseSeriesWithDefaults

`func NewObserveMetricsResponseSeriesWithDefaults() *ObserveMetricsResponseSeries`

NewObserveMetricsResponseSeriesWithDefaults instantiates a new ObserveMetricsResponseSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *ObserveMetricsResponseSeries) GetRequests() []ObservePoint`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ObserveMetricsResponseSeries) GetRequestsOk() (*[]ObservePoint, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ObserveMetricsResponseSeries) SetRequests(v []ObservePoint)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ObserveMetricsResponseSeries) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetErrors

`func (o *ObserveMetricsResponseSeries) GetErrors() []ObservePoint`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ObserveMetricsResponseSeries) GetErrorsOk() (*[]ObservePoint, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ObserveMetricsResponseSeries) SetErrors(v []ObservePoint)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ObserveMetricsResponseSeries) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLatencyP50Ms

`func (o *ObserveMetricsResponseSeries) GetLatencyP50Ms() []ObservePoint`

GetLatencyP50Ms returns the LatencyP50Ms field if non-nil, zero value otherwise.

### GetLatencyP50MsOk

`func (o *ObserveMetricsResponseSeries) GetLatencyP50MsOk() (*[]ObservePoint, bool)`

GetLatencyP50MsOk returns a tuple with the LatencyP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP50Ms

`func (o *ObserveMetricsResponseSeries) SetLatencyP50Ms(v []ObservePoint)`

SetLatencyP50Ms sets LatencyP50Ms field to given value.

### HasLatencyP50Ms

`func (o *ObserveMetricsResponseSeries) HasLatencyP50Ms() bool`

HasLatencyP50Ms returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *ObserveMetricsResponseSeries) GetLatencyP95Ms() []ObservePoint`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *ObserveMetricsResponseSeries) GetLatencyP95MsOk() (*[]ObservePoint, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *ObserveMetricsResponseSeries) SetLatencyP95Ms(v []ObservePoint)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *ObserveMetricsResponseSeries) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


