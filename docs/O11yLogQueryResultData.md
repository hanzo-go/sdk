# O11yLogQueryResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]O11yLogQueryResultDataResultInner**](O11yLogQueryResultDataResultInner.md) |  | [optional] 
**Stats** | Pointer to [**O11yLogQueryResultDataStats**](O11yLogQueryResultDataStats.md) |  | [optional] 

## Methods

### NewO11yLogQueryResultData

`func NewO11yLogQueryResultData() *O11yLogQueryResultData`

NewO11yLogQueryResultData instantiates a new O11yLogQueryResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yLogQueryResultDataWithDefaults

`func NewO11yLogQueryResultDataWithDefaults() *O11yLogQueryResultData`

NewO11yLogQueryResultDataWithDefaults instantiates a new O11yLogQueryResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *O11yLogQueryResultData) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *O11yLogQueryResultData) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *O11yLogQueryResultData) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *O11yLogQueryResultData) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetResult

`func (o *O11yLogQueryResultData) GetResult() []O11yLogQueryResultDataResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *O11yLogQueryResultData) GetResultOk() (*[]O11yLogQueryResultDataResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *O11yLogQueryResultData) SetResult(v []O11yLogQueryResultDataResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *O11yLogQueryResultData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStats

`func (o *O11yLogQueryResultData) GetStats() O11yLogQueryResultDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *O11yLogQueryResultData) GetStatsOk() (*O11yLogQueryResultDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *O11yLogQueryResultData) SetStats(v O11yLogQueryResultDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *O11yLogQueryResultData) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


