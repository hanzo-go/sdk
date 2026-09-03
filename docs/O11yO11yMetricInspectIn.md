# O11yO11yMetricInspectIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** | End is the end of the window as a Unix timestamp in milliseconds, at most thirty minutes after start. Required. | 
**Filter** | Pointer to [**O11yO11yMetricFilter**](O11yO11yMetricFilter.md) | Filter narrows the series returned. | [optional] 
**MetricName** | **string** | MetricName is the metric to inspect. Required. | 
**Start** | **int64** | Start is the start of the window as a Unix timestamp in milliseconds. Required. | 

## Methods

### NewO11yO11yMetricInspectIn

`func NewO11yO11yMetricInspectIn(end int64, metricName string, start int64, ) *O11yO11yMetricInspectIn`

NewO11yO11yMetricInspectIn instantiates a new O11yO11yMetricInspectIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricInspectInWithDefaults

`func NewO11yO11yMetricInspectInWithDefaults() *O11yO11yMetricInspectIn`

NewO11yO11yMetricInspectInWithDefaults instantiates a new O11yO11yMetricInspectIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yMetricInspectIn) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yMetricInspectIn) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yMetricInspectIn) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetFilter

`func (o *O11yO11yMetricInspectIn) GetFilter() O11yO11yMetricFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *O11yO11yMetricInspectIn) GetFilterOk() (*O11yO11yMetricFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *O11yO11yMetricInspectIn) SetFilter(v O11yO11yMetricFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *O11yO11yMetricInspectIn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMetricName

`func (o *O11yO11yMetricInspectIn) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *O11yO11yMetricInspectIn) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *O11yO11yMetricInspectIn) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetStart

`func (o *O11yO11yMetricInspectIn) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yMetricInspectIn) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yMetricInspectIn) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


