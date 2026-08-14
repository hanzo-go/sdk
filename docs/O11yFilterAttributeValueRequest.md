# O11yFilterAttributeValueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateAttribute** | Pointer to **string** |  | [optional] 
**AggregateOperator** | Pointer to **string** |  | [optional] 
**DataSource** | Pointer to **string** |  | [optional] 
**EndTimeMillis** | Pointer to **int32** |  | [optional] 
**ExistingFilterItems** | Pointer to [**[]O11yFilterItem**](O11yFilterItem.md) |  | [optional] 
**FilterAttributeKey** | Pointer to **string** |  | [optional] 
**FilterAttributeKeyDataType** | Pointer to **string** |  | [optional] 
**IncludeRelated** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**MetricNames** | Pointer to **[]string** |  | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**StartTimeMillis** | Pointer to **int32** |  | [optional] 
**TagType** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yFilterAttributeValueRequest

`func NewO11yFilterAttributeValueRequest() *O11yFilterAttributeValueRequest`

NewO11yFilterAttributeValueRequest instantiates a new O11yFilterAttributeValueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yFilterAttributeValueRequestWithDefaults

`func NewO11yFilterAttributeValueRequestWithDefaults() *O11yFilterAttributeValueRequest`

NewO11yFilterAttributeValueRequestWithDefaults instantiates a new O11yFilterAttributeValueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregateAttribute

`func (o *O11yFilterAttributeValueRequest) GetAggregateAttribute() string`

GetAggregateAttribute returns the AggregateAttribute field if non-nil, zero value otherwise.

### GetAggregateAttributeOk

`func (o *O11yFilterAttributeValueRequest) GetAggregateAttributeOk() (*string, bool)`

GetAggregateAttributeOk returns a tuple with the AggregateAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateAttribute

`func (o *O11yFilterAttributeValueRequest) SetAggregateAttribute(v string)`

SetAggregateAttribute sets AggregateAttribute field to given value.

### HasAggregateAttribute

`func (o *O11yFilterAttributeValueRequest) HasAggregateAttribute() bool`

HasAggregateAttribute returns a boolean if a field has been set.

### GetAggregateOperator

`func (o *O11yFilterAttributeValueRequest) GetAggregateOperator() string`

GetAggregateOperator returns the AggregateOperator field if non-nil, zero value otherwise.

### GetAggregateOperatorOk

`func (o *O11yFilterAttributeValueRequest) GetAggregateOperatorOk() (*string, bool)`

GetAggregateOperatorOk returns a tuple with the AggregateOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateOperator

`func (o *O11yFilterAttributeValueRequest) SetAggregateOperator(v string)`

SetAggregateOperator sets AggregateOperator field to given value.

### HasAggregateOperator

`func (o *O11yFilterAttributeValueRequest) HasAggregateOperator() bool`

HasAggregateOperator returns a boolean if a field has been set.

### GetDataSource

`func (o *O11yFilterAttributeValueRequest) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *O11yFilterAttributeValueRequest) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *O11yFilterAttributeValueRequest) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *O11yFilterAttributeValueRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetEndTimeMillis

`func (o *O11yFilterAttributeValueRequest) GetEndTimeMillis() int32`

GetEndTimeMillis returns the EndTimeMillis field if non-nil, zero value otherwise.

### GetEndTimeMillisOk

`func (o *O11yFilterAttributeValueRequest) GetEndTimeMillisOk() (*int32, bool)`

GetEndTimeMillisOk returns a tuple with the EndTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeMillis

`func (o *O11yFilterAttributeValueRequest) SetEndTimeMillis(v int32)`

SetEndTimeMillis sets EndTimeMillis field to given value.

### HasEndTimeMillis

`func (o *O11yFilterAttributeValueRequest) HasEndTimeMillis() bool`

HasEndTimeMillis returns a boolean if a field has been set.

### GetExistingFilterItems

`func (o *O11yFilterAttributeValueRequest) GetExistingFilterItems() []O11yFilterItem`

GetExistingFilterItems returns the ExistingFilterItems field if non-nil, zero value otherwise.

### GetExistingFilterItemsOk

`func (o *O11yFilterAttributeValueRequest) GetExistingFilterItemsOk() (*[]O11yFilterItem, bool)`

GetExistingFilterItemsOk returns a tuple with the ExistingFilterItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingFilterItems

`func (o *O11yFilterAttributeValueRequest) SetExistingFilterItems(v []O11yFilterItem)`

SetExistingFilterItems sets ExistingFilterItems field to given value.

### HasExistingFilterItems

`func (o *O11yFilterAttributeValueRequest) HasExistingFilterItems() bool`

HasExistingFilterItems returns a boolean if a field has been set.

### GetFilterAttributeKey

`func (o *O11yFilterAttributeValueRequest) GetFilterAttributeKey() string`

GetFilterAttributeKey returns the FilterAttributeKey field if non-nil, zero value otherwise.

### GetFilterAttributeKeyOk

`func (o *O11yFilterAttributeValueRequest) GetFilterAttributeKeyOk() (*string, bool)`

GetFilterAttributeKeyOk returns a tuple with the FilterAttributeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAttributeKey

`func (o *O11yFilterAttributeValueRequest) SetFilterAttributeKey(v string)`

SetFilterAttributeKey sets FilterAttributeKey field to given value.

### HasFilterAttributeKey

`func (o *O11yFilterAttributeValueRequest) HasFilterAttributeKey() bool`

HasFilterAttributeKey returns a boolean if a field has been set.

### GetFilterAttributeKeyDataType

`func (o *O11yFilterAttributeValueRequest) GetFilterAttributeKeyDataType() string`

GetFilterAttributeKeyDataType returns the FilterAttributeKeyDataType field if non-nil, zero value otherwise.

### GetFilterAttributeKeyDataTypeOk

`func (o *O11yFilterAttributeValueRequest) GetFilterAttributeKeyDataTypeOk() (*string, bool)`

GetFilterAttributeKeyDataTypeOk returns a tuple with the FilterAttributeKeyDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAttributeKeyDataType

`func (o *O11yFilterAttributeValueRequest) SetFilterAttributeKeyDataType(v string)`

SetFilterAttributeKeyDataType sets FilterAttributeKeyDataType field to given value.

### HasFilterAttributeKeyDataType

`func (o *O11yFilterAttributeValueRequest) HasFilterAttributeKeyDataType() bool`

HasFilterAttributeKeyDataType returns a boolean if a field has been set.

### GetIncludeRelated

`func (o *O11yFilterAttributeValueRequest) GetIncludeRelated() bool`

GetIncludeRelated returns the IncludeRelated field if non-nil, zero value otherwise.

### GetIncludeRelatedOk

`func (o *O11yFilterAttributeValueRequest) GetIncludeRelatedOk() (*bool, bool)`

GetIncludeRelatedOk returns a tuple with the IncludeRelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRelated

`func (o *O11yFilterAttributeValueRequest) SetIncludeRelated(v bool)`

SetIncludeRelated sets IncludeRelated field to given value.

### HasIncludeRelated

`func (o *O11yFilterAttributeValueRequest) HasIncludeRelated() bool`

HasIncludeRelated returns a boolean if a field has been set.

### GetLimit

`func (o *O11yFilterAttributeValueRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yFilterAttributeValueRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yFilterAttributeValueRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yFilterAttributeValueRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetricNames

`func (o *O11yFilterAttributeValueRequest) GetMetricNames() []string`

GetMetricNames returns the MetricNames field if non-nil, zero value otherwise.

### GetMetricNamesOk

`func (o *O11yFilterAttributeValueRequest) GetMetricNamesOk() (*[]string, bool)`

GetMetricNamesOk returns a tuple with the MetricNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricNames

`func (o *O11yFilterAttributeValueRequest) SetMetricNames(v []string)`

SetMetricNames sets MetricNames field to given value.

### HasMetricNames

`func (o *O11yFilterAttributeValueRequest) HasMetricNames() bool`

HasMetricNames returns a boolean if a field has been set.

### GetSearchText

`func (o *O11yFilterAttributeValueRequest) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *O11yFilterAttributeValueRequest) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *O11yFilterAttributeValueRequest) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *O11yFilterAttributeValueRequest) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetStartTimeMillis

`func (o *O11yFilterAttributeValueRequest) GetStartTimeMillis() int32`

GetStartTimeMillis returns the StartTimeMillis field if non-nil, zero value otherwise.

### GetStartTimeMillisOk

`func (o *O11yFilterAttributeValueRequest) GetStartTimeMillisOk() (*int32, bool)`

GetStartTimeMillisOk returns a tuple with the StartTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeMillis

`func (o *O11yFilterAttributeValueRequest) SetStartTimeMillis(v int32)`

SetStartTimeMillis sets StartTimeMillis field to given value.

### HasStartTimeMillis

`func (o *O11yFilterAttributeValueRequest) HasStartTimeMillis() bool`

HasStartTimeMillis returns a boolean if a field has been set.

### GetTagType

`func (o *O11yFilterAttributeValueRequest) GetTagType() string`

GetTagType returns the TagType field if non-nil, zero value otherwise.

### GetTagTypeOk

`func (o *O11yFilterAttributeValueRequest) GetTagTypeOk() (*string, bool)`

GetTagTypeOk returns a tuple with the TagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagType

`func (o *O11yFilterAttributeValueRequest) SetTagType(v string)`

SetTagType sets TagType field to given value.

### HasTagType

`func (o *O11yFilterAttributeValueRequest) HasTagType() bool`

HasTagType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


