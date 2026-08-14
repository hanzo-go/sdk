# O11yO11yReductionSeriesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**[]O11yO11yAggregation**](O11yO11yAggregation.md) | Aggregations are the query&#39;s aggregation buckets. | [optional] 
**QueryName** | Pointer to **string** | QueryName names the query the result answers. | [optional] 

## Methods

### NewO11yO11yReductionSeriesResult

`func NewO11yO11yReductionSeriesResult() *O11yO11yReductionSeriesResult`

NewO11yO11yReductionSeriesResult instantiates a new O11yO11yReductionSeriesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yReductionSeriesResultWithDefaults

`func NewO11yO11yReductionSeriesResultWithDefaults() *O11yO11yReductionSeriesResult`

NewO11yO11yReductionSeriesResultWithDefaults instantiates a new O11yO11yReductionSeriesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *O11yO11yReductionSeriesResult) GetAggregations() []O11yO11yAggregation`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *O11yO11yReductionSeriesResult) GetAggregationsOk() (*[]O11yO11yAggregation, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *O11yO11yReductionSeriesResult) SetAggregations(v []O11yO11yAggregation)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *O11yO11yReductionSeriesResult) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetQueryName

`func (o *O11yO11yReductionSeriesResult) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *O11yO11yReductionSeriesResult) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *O11yO11yReductionSeriesResult) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.

### HasQueryName

`func (o *O11yO11yReductionSeriesResult) HasQueryName() bool`

HasQueryName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


