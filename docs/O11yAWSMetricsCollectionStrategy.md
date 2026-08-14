# O11yAWSMetricsCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamFilters** | Pointer to [**[]O11yAWSCloudWatchMetricStreamFilter**](O11yAWSCloudWatchMetricStreamFilter.md) | to be used as https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html#cfn-cloudwatch-metricstream-includefilters | [optional] 

## Methods

### NewO11yAWSMetricsCollectionStrategy

`func NewO11yAWSMetricsCollectionStrategy() *O11yAWSMetricsCollectionStrategy`

NewO11yAWSMetricsCollectionStrategy instantiates a new O11yAWSMetricsCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAWSMetricsCollectionStrategyWithDefaults

`func NewO11yAWSMetricsCollectionStrategyWithDefaults() *O11yAWSMetricsCollectionStrategy`

NewO11yAWSMetricsCollectionStrategyWithDefaults instantiates a new O11yAWSMetricsCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamFilters

`func (o *O11yAWSMetricsCollectionStrategy) GetStreamFilters() []O11yAWSCloudWatchMetricStreamFilter`

GetStreamFilters returns the StreamFilters field if non-nil, zero value otherwise.

### GetStreamFiltersOk

`func (o *O11yAWSMetricsCollectionStrategy) GetStreamFiltersOk() (*[]O11yAWSCloudWatchMetricStreamFilter, bool)`

GetStreamFiltersOk returns a tuple with the StreamFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamFilters

`func (o *O11yAWSMetricsCollectionStrategy) SetStreamFilters(v []O11yAWSCloudWatchMetricStreamFilter)`

SetStreamFilters sets StreamFilters field to given value.

### HasStreamFilters

`func (o *O11yAWSMetricsCollectionStrategy) HasStreamFilters() bool`

HasStreamFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


