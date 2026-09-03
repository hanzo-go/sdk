# O11yO11yTreemapEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **string** | MetricName is the metric&#39;s name. | [optional] 
**Percentage** | Pointer to **float64** | Percentage is the metric&#39;s share, in percent. | [optional] 
**TotalValue** | Pointer to **int32** | TotalValue is the metric&#39;s absolute count. | [optional] 

## Methods

### NewO11yO11yTreemapEntry

`func NewO11yO11yTreemapEntry() *O11yO11yTreemapEntry`

NewO11yO11yTreemapEntry instantiates a new O11yO11yTreemapEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTreemapEntryWithDefaults

`func NewO11yO11yTreemapEntryWithDefaults() *O11yO11yTreemapEntry`

NewO11yO11yTreemapEntryWithDefaults instantiates a new O11yO11yTreemapEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *O11yO11yTreemapEntry) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *O11yO11yTreemapEntry) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *O11yO11yTreemapEntry) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *O11yO11yTreemapEntry) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetPercentage

`func (o *O11yO11yTreemapEntry) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *O11yO11yTreemapEntry) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *O11yO11yTreemapEntry) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *O11yO11yTreemapEntry) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetTotalValue

`func (o *O11yO11yTreemapEntry) GetTotalValue() int32`

GetTotalValue returns the TotalValue field if non-nil, zero value otherwise.

### GetTotalValueOk

`func (o *O11yO11yTreemapEntry) GetTotalValueOk() (*int32, bool)`

GetTotalValueOk returns a tuple with the TotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalValue

`func (o *O11yO11yTreemapEntry) SetTotalValue(v int32)`

SetTotalValue sets TotalValue field to given value.

### HasTotalValue

`func (o *O11yO11yTreemapEntry) HasTotalValue() bool`

HasTotalValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


