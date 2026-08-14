# O11yO11yQueryRangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]O11yO11yReductionSeriesResult**](O11yO11yReductionSeriesResult.md) | Results are the per-query results, each a set of aggregated series. | [optional] 

## Methods

### NewO11yO11yQueryRangeData

`func NewO11yO11yQueryRangeData() *O11yO11yQueryRangeData`

NewO11yO11yQueryRangeData instantiates a new O11yO11yQueryRangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueryRangeDataWithDefaults

`func NewO11yO11yQueryRangeDataWithDefaults() *O11yO11yQueryRangeData`

NewO11yO11yQueryRangeDataWithDefaults instantiates a new O11yO11yQueryRangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *O11yO11yQueryRangeData) GetResults() []O11yO11yReductionSeriesResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *O11yO11yQueryRangeData) GetResultsOk() (*[]O11yO11yReductionSeriesResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *O11yO11yQueryRangeData) SetResults(v []O11yO11yReductionSeriesResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *O11yO11yQueryRangeData) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


