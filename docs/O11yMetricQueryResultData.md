# O11yMetricQueryResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]O11yMetricQueryResultDataResultInner**](O11yMetricQueryResultDataResultInner.md) |  | [optional] 

## Methods

### NewO11yMetricQueryResultData

`func NewO11yMetricQueryResultData() *O11yMetricQueryResultData`

NewO11yMetricQueryResultData instantiates a new O11yMetricQueryResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMetricQueryResultDataWithDefaults

`func NewO11yMetricQueryResultDataWithDefaults() *O11yMetricQueryResultData`

NewO11yMetricQueryResultDataWithDefaults instantiates a new O11yMetricQueryResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *O11yMetricQueryResultData) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *O11yMetricQueryResultData) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *O11yMetricQueryResultData) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *O11yMetricQueryResultData) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetResult

`func (o *O11yMetricQueryResultData) GetResult() []O11yMetricQueryResultDataResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *O11yMetricQueryResultData) GetResultOk() (*[]O11yMetricQueryResultDataResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *O11yMetricQueryResultData) SetResult(v []O11yMetricQueryResultDataResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *O11yMetricQueryResultData) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


