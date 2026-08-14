# O11yMetricsResponseSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]O11yPoint**](O11yPoint.md) |  | [optional] 
**LatencyP50Ms** | Pointer to [**[]O11yPoint**](O11yPoint.md) |  | [optional] 
**LatencyP95Ms** | Pointer to [**[]O11yPoint**](O11yPoint.md) |  | [optional] 
**Requests** | Pointer to [**[]O11yPoint**](O11yPoint.md) |  | [optional] 

## Methods

### NewO11yMetricsResponseSeries

`func NewO11yMetricsResponseSeries() *O11yMetricsResponseSeries`

NewO11yMetricsResponseSeries instantiates a new O11yMetricsResponseSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMetricsResponseSeriesWithDefaults

`func NewO11yMetricsResponseSeriesWithDefaults() *O11yMetricsResponseSeries`

NewO11yMetricsResponseSeriesWithDefaults instantiates a new O11yMetricsResponseSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *O11yMetricsResponseSeries) GetErrors() []O11yPoint`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *O11yMetricsResponseSeries) GetErrorsOk() (*[]O11yPoint, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *O11yMetricsResponseSeries) SetErrors(v []O11yPoint)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *O11yMetricsResponseSeries) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLatencyP50Ms

`func (o *O11yMetricsResponseSeries) GetLatencyP50Ms() []O11yPoint`

GetLatencyP50Ms returns the LatencyP50Ms field if non-nil, zero value otherwise.

### GetLatencyP50MsOk

`func (o *O11yMetricsResponseSeries) GetLatencyP50MsOk() (*[]O11yPoint, bool)`

GetLatencyP50MsOk returns a tuple with the LatencyP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP50Ms

`func (o *O11yMetricsResponseSeries) SetLatencyP50Ms(v []O11yPoint)`

SetLatencyP50Ms sets LatencyP50Ms field to given value.

### HasLatencyP50Ms

`func (o *O11yMetricsResponseSeries) HasLatencyP50Ms() bool`

HasLatencyP50Ms returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *O11yMetricsResponseSeries) GetLatencyP95Ms() []O11yPoint`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *O11yMetricsResponseSeries) GetLatencyP95MsOk() (*[]O11yPoint, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *O11yMetricsResponseSeries) SetLatencyP95Ms(v []O11yPoint)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *O11yMetricsResponseSeries) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.

### GetRequests

`func (o *O11yMetricsResponseSeries) GetRequests() []O11yPoint`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *O11yMetricsResponseSeries) GetRequestsOk() (*[]O11yPoint, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *O11yMetricsResponseSeries) SetRequests(v []O11yPoint)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *O11yMetricsResponseSeries) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


