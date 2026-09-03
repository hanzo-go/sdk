# O11yO11yMetricTreemapIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** | End is the end of the window as a Unix timestamp in milliseconds. Required. | 
**Filter** | Pointer to [**O11yO11yMetricFilter**](O11yO11yMetricFilter.md) | Filter narrows the metrics counted. | [optional] 
**Limit** | **int64** | Limit caps how many entries come back, between 1 and 5000. Required. | 
**Mode** | **string** | Mode picks the measure: timeseries or samples. Required. | 
**Start** | **int64** | Start is the start of the window as a Unix timestamp in milliseconds. Required. | 

## Methods

### NewO11yO11yMetricTreemapIn

`func NewO11yO11yMetricTreemapIn(end int64, limit int64, mode string, start int64, ) *O11yO11yMetricTreemapIn`

NewO11yO11yMetricTreemapIn instantiates a new O11yO11yMetricTreemapIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricTreemapInWithDefaults

`func NewO11yO11yMetricTreemapInWithDefaults() *O11yO11yMetricTreemapIn`

NewO11yO11yMetricTreemapInWithDefaults instantiates a new O11yO11yMetricTreemapIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yMetricTreemapIn) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yMetricTreemapIn) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yMetricTreemapIn) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetFilter

`func (o *O11yO11yMetricTreemapIn) GetFilter() O11yO11yMetricFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *O11yO11yMetricTreemapIn) GetFilterOk() (*O11yO11yMetricFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *O11yO11yMetricTreemapIn) SetFilter(v O11yO11yMetricFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *O11yO11yMetricTreemapIn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yMetricTreemapIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yMetricTreemapIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yMetricTreemapIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetMode

`func (o *O11yO11yMetricTreemapIn) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *O11yO11yMetricTreemapIn) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *O11yO11yMetricTreemapIn) SetMode(v string)`

SetMode sets Mode field to given value.


### GetStart

`func (o *O11yO11yMetricTreemapIn) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yMetricTreemapIn) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yMetricTreemapIn) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


