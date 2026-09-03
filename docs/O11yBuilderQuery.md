# O11yBuilderQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAnomaly** | Pointer to **bool** |  | [optional] 
**QueriesUsedInFormula** | Pointer to **[]string** |  | [optional] 
**ShiftBy** | Pointer to **int64** |  | [optional] 
**AggregateAttribute** | Pointer to [**O11yAttributeKey**](O11yAttributeKey.md) |  | [optional] 
**AggregateOperator** | Pointer to **string** |  | [optional] 
**DataSource** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**O11yFilterSet**](O11yFilterSet.md) |  | [optional] 
**Functions** | Pointer to [**[]O11yFunction**](O11yFunction.md) |  | [optional] 
**GroupBy** | Pointer to [**[]O11yAttributeKey**](O11yAttributeKey.md) |  | [optional] 
**Having** | Pointer to [**[]O11yHaving**](O11yHaving.md) |  | [optional] 
**Legend** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to [**[]O11yOrderBy**](O11yOrderBy.md) |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**QueryName** | Pointer to **string** |  | [optional] 
**ReduceTo** | Pointer to **string** |  | [optional] 
**SelectColumns** | Pointer to [**[]O11yAttributeKey**](O11yAttributeKey.md) |  | [optional] 
**SeriesAggregation** | Pointer to **string** |  | [optional] 
**SpaceAggregation** | Pointer to **string** |  | [optional] 
**StepInterval** | Pointer to **int64** |  | [optional] 
**Temporality** | Pointer to **string** |  | [optional] 
**TimeAggregation** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yBuilderQuery

`func NewO11yBuilderQuery() *O11yBuilderQuery`

NewO11yBuilderQuery instantiates a new O11yBuilderQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yBuilderQueryWithDefaults

`func NewO11yBuilderQueryWithDefaults() *O11yBuilderQuery`

NewO11yBuilderQueryWithDefaults instantiates a new O11yBuilderQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAnomaly

`func (o *O11yBuilderQuery) GetIsAnomaly() bool`

GetIsAnomaly returns the IsAnomaly field if non-nil, zero value otherwise.

### GetIsAnomalyOk

`func (o *O11yBuilderQuery) GetIsAnomalyOk() (*bool, bool)`

GetIsAnomalyOk returns a tuple with the IsAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnomaly

`func (o *O11yBuilderQuery) SetIsAnomaly(v bool)`

SetIsAnomaly sets IsAnomaly field to given value.

### HasIsAnomaly

`func (o *O11yBuilderQuery) HasIsAnomaly() bool`

HasIsAnomaly returns a boolean if a field has been set.

### GetQueriesUsedInFormula

`func (o *O11yBuilderQuery) GetQueriesUsedInFormula() []string`

GetQueriesUsedInFormula returns the QueriesUsedInFormula field if non-nil, zero value otherwise.

### GetQueriesUsedInFormulaOk

`func (o *O11yBuilderQuery) GetQueriesUsedInFormulaOk() (*[]string, bool)`

GetQueriesUsedInFormulaOk returns a tuple with the QueriesUsedInFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriesUsedInFormula

`func (o *O11yBuilderQuery) SetQueriesUsedInFormula(v []string)`

SetQueriesUsedInFormula sets QueriesUsedInFormula field to given value.

### HasQueriesUsedInFormula

`func (o *O11yBuilderQuery) HasQueriesUsedInFormula() bool`

HasQueriesUsedInFormula returns a boolean if a field has been set.

### GetShiftBy

`func (o *O11yBuilderQuery) GetShiftBy() int64`

GetShiftBy returns the ShiftBy field if non-nil, zero value otherwise.

### GetShiftByOk

`func (o *O11yBuilderQuery) GetShiftByOk() (*int64, bool)`

GetShiftByOk returns a tuple with the ShiftBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftBy

`func (o *O11yBuilderQuery) SetShiftBy(v int64)`

SetShiftBy sets ShiftBy field to given value.

### HasShiftBy

`func (o *O11yBuilderQuery) HasShiftBy() bool`

HasShiftBy returns a boolean if a field has been set.

### GetAggregateAttribute

`func (o *O11yBuilderQuery) GetAggregateAttribute() O11yAttributeKey`

GetAggregateAttribute returns the AggregateAttribute field if non-nil, zero value otherwise.

### GetAggregateAttributeOk

`func (o *O11yBuilderQuery) GetAggregateAttributeOk() (*O11yAttributeKey, bool)`

GetAggregateAttributeOk returns a tuple with the AggregateAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateAttribute

`func (o *O11yBuilderQuery) SetAggregateAttribute(v O11yAttributeKey)`

SetAggregateAttribute sets AggregateAttribute field to given value.

### HasAggregateAttribute

`func (o *O11yBuilderQuery) HasAggregateAttribute() bool`

HasAggregateAttribute returns a boolean if a field has been set.

### GetAggregateOperator

`func (o *O11yBuilderQuery) GetAggregateOperator() string`

GetAggregateOperator returns the AggregateOperator field if non-nil, zero value otherwise.

### GetAggregateOperatorOk

`func (o *O11yBuilderQuery) GetAggregateOperatorOk() (*string, bool)`

GetAggregateOperatorOk returns a tuple with the AggregateOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateOperator

`func (o *O11yBuilderQuery) SetAggregateOperator(v string)`

SetAggregateOperator sets AggregateOperator field to given value.

### HasAggregateOperator

`func (o *O11yBuilderQuery) HasAggregateOperator() bool`

HasAggregateOperator returns a boolean if a field has been set.

### GetDataSource

`func (o *O11yBuilderQuery) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *O11yBuilderQuery) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *O11yBuilderQuery) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *O11yBuilderQuery) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDisabled

`func (o *O11yBuilderQuery) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *O11yBuilderQuery) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *O11yBuilderQuery) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *O11yBuilderQuery) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExpression

`func (o *O11yBuilderQuery) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *O11yBuilderQuery) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *O11yBuilderQuery) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *O11yBuilderQuery) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetFilters

`func (o *O11yBuilderQuery) GetFilters() O11yFilterSet`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *O11yBuilderQuery) GetFiltersOk() (*O11yFilterSet, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *O11yBuilderQuery) SetFilters(v O11yFilterSet)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *O11yBuilderQuery) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFunctions

`func (o *O11yBuilderQuery) GetFunctions() []O11yFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *O11yBuilderQuery) GetFunctionsOk() (*[]O11yFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *O11yBuilderQuery) SetFunctions(v []O11yFunction)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *O11yBuilderQuery) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetGroupBy

`func (o *O11yBuilderQuery) GetGroupBy() []O11yAttributeKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *O11yBuilderQuery) GetGroupByOk() (*[]O11yAttributeKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *O11yBuilderQuery) SetGroupBy(v []O11yAttributeKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *O11yBuilderQuery) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetHaving

`func (o *O11yBuilderQuery) GetHaving() []O11yHaving`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *O11yBuilderQuery) GetHavingOk() (*[]O11yHaving, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *O11yBuilderQuery) SetHaving(v []O11yHaving)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *O11yBuilderQuery) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### GetLegend

`func (o *O11yBuilderQuery) GetLegend() string`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *O11yBuilderQuery) GetLegendOk() (*string, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *O11yBuilderQuery) SetLegend(v string)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *O11yBuilderQuery) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### GetLimit

`func (o *O11yBuilderQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yBuilderQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yBuilderQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yBuilderQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yBuilderQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yBuilderQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yBuilderQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yBuilderQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *O11yBuilderQuery) GetOrderBy() []O11yOrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *O11yBuilderQuery) GetOrderByOk() (*[]O11yOrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *O11yBuilderQuery) SetOrderBy(v []O11yOrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *O11yBuilderQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPageSize

`func (o *O11yBuilderQuery) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *O11yBuilderQuery) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *O11yBuilderQuery) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *O11yBuilderQuery) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetQueryName

`func (o *O11yBuilderQuery) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *O11yBuilderQuery) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *O11yBuilderQuery) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.

### HasQueryName

`func (o *O11yBuilderQuery) HasQueryName() bool`

HasQueryName returns a boolean if a field has been set.

### GetReduceTo

`func (o *O11yBuilderQuery) GetReduceTo() string`

GetReduceTo returns the ReduceTo field if non-nil, zero value otherwise.

### GetReduceToOk

`func (o *O11yBuilderQuery) GetReduceToOk() (*string, bool)`

GetReduceToOk returns a tuple with the ReduceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceTo

`func (o *O11yBuilderQuery) SetReduceTo(v string)`

SetReduceTo sets ReduceTo field to given value.

### HasReduceTo

`func (o *O11yBuilderQuery) HasReduceTo() bool`

HasReduceTo returns a boolean if a field has been set.

### GetSelectColumns

`func (o *O11yBuilderQuery) GetSelectColumns() []O11yAttributeKey`

GetSelectColumns returns the SelectColumns field if non-nil, zero value otherwise.

### GetSelectColumnsOk

`func (o *O11yBuilderQuery) GetSelectColumnsOk() (*[]O11yAttributeKey, bool)`

GetSelectColumnsOk returns a tuple with the SelectColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectColumns

`func (o *O11yBuilderQuery) SetSelectColumns(v []O11yAttributeKey)`

SetSelectColumns sets SelectColumns field to given value.

### HasSelectColumns

`func (o *O11yBuilderQuery) HasSelectColumns() bool`

HasSelectColumns returns a boolean if a field has been set.

### GetSeriesAggregation

`func (o *O11yBuilderQuery) GetSeriesAggregation() string`

GetSeriesAggregation returns the SeriesAggregation field if non-nil, zero value otherwise.

### GetSeriesAggregationOk

`func (o *O11yBuilderQuery) GetSeriesAggregationOk() (*string, bool)`

GetSeriesAggregationOk returns a tuple with the SeriesAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesAggregation

`func (o *O11yBuilderQuery) SetSeriesAggregation(v string)`

SetSeriesAggregation sets SeriesAggregation field to given value.

### HasSeriesAggregation

`func (o *O11yBuilderQuery) HasSeriesAggregation() bool`

HasSeriesAggregation returns a boolean if a field has been set.

### GetSpaceAggregation

`func (o *O11yBuilderQuery) GetSpaceAggregation() string`

GetSpaceAggregation returns the SpaceAggregation field if non-nil, zero value otherwise.

### GetSpaceAggregationOk

`func (o *O11yBuilderQuery) GetSpaceAggregationOk() (*string, bool)`

GetSpaceAggregationOk returns a tuple with the SpaceAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAggregation

`func (o *O11yBuilderQuery) SetSpaceAggregation(v string)`

SetSpaceAggregation sets SpaceAggregation field to given value.

### HasSpaceAggregation

`func (o *O11yBuilderQuery) HasSpaceAggregation() bool`

HasSpaceAggregation returns a boolean if a field has been set.

### GetStepInterval

`func (o *O11yBuilderQuery) GetStepInterval() int64`

GetStepInterval returns the StepInterval field if non-nil, zero value otherwise.

### GetStepIntervalOk

`func (o *O11yBuilderQuery) GetStepIntervalOk() (*int64, bool)`

GetStepIntervalOk returns a tuple with the StepInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepInterval

`func (o *O11yBuilderQuery) SetStepInterval(v int64)`

SetStepInterval sets StepInterval field to given value.

### HasStepInterval

`func (o *O11yBuilderQuery) HasStepInterval() bool`

HasStepInterval returns a boolean if a field has been set.

### GetTemporality

`func (o *O11yBuilderQuery) GetTemporality() string`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *O11yBuilderQuery) GetTemporalityOk() (*string, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *O11yBuilderQuery) SetTemporality(v string)`

SetTemporality sets Temporality field to given value.

### HasTemporality

`func (o *O11yBuilderQuery) HasTemporality() bool`

HasTemporality returns a boolean if a field has been set.

### GetTimeAggregation

`func (o *O11yBuilderQuery) GetTimeAggregation() string`

GetTimeAggregation returns the TimeAggregation field if non-nil, zero value otherwise.

### GetTimeAggregationOk

`func (o *O11yBuilderQuery) GetTimeAggregationOk() (*string, bool)`

GetTimeAggregationOk returns a tuple with the TimeAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAggregation

`func (o *O11yBuilderQuery) SetTimeAggregation(v string)`

SetTimeAggregation sets TimeAggregation field to given value.

### HasTimeAggregation

`func (o *O11yBuilderQuery) HasTimeAggregation() bool`

HasTimeAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


