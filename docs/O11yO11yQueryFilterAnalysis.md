# O11yO11yQueryFilterAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]O11yO11yColumnInfo**](O11yO11yColumnInfo.md) | Groups are the columns the query groups by. | [optional] 
**MetricNames** | Pointer to **[]string** | MetricNames are the metrics the query reads. | [optional] 

## Methods

### NewO11yO11yQueryFilterAnalysis

`func NewO11yO11yQueryFilterAnalysis() *O11yO11yQueryFilterAnalysis`

NewO11yO11yQueryFilterAnalysis instantiates a new O11yO11yQueryFilterAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueryFilterAnalysisWithDefaults

`func NewO11yO11yQueryFilterAnalysisWithDefaults() *O11yO11yQueryFilterAnalysis`

NewO11yO11yQueryFilterAnalysisWithDefaults instantiates a new O11yO11yQueryFilterAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *O11yO11yQueryFilterAnalysis) GetGroups() []O11yO11yColumnInfo`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *O11yO11yQueryFilterAnalysis) GetGroupsOk() (*[]O11yO11yColumnInfo, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *O11yO11yQueryFilterAnalysis) SetGroups(v []O11yO11yColumnInfo)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *O11yO11yQueryFilterAnalysis) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetMetricNames

`func (o *O11yO11yQueryFilterAnalysis) GetMetricNames() []string`

GetMetricNames returns the MetricNames field if non-nil, zero value otherwise.

### GetMetricNamesOk

`func (o *O11yO11yQueryFilterAnalysis) GetMetricNamesOk() (*[]string, bool)`

GetMetricNamesOk returns a tuple with the MetricNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricNames

`func (o *O11yO11yQueryFilterAnalysis) SetMetricNames(v []string)`

SetMetricNames sets MetricNames field to given value.

### HasMetricNames

`func (o *O11yO11yQueryFilterAnalysis) HasMetricNames() bool`

HasMetricNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


