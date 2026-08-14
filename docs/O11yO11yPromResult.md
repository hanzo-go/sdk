# O11yO11yPromResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **interface{}** |  | [optional] 
**ResultType** | Pointer to **string** | ResultType discriminates Result: matrix, vector, scalar or string. | [optional] 
**Stats** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yO11yPromResult

`func NewO11yO11yPromResult() *O11yO11yPromResult`

NewO11yO11yPromResult instantiates a new O11yO11yPromResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yPromResultWithDefaults

`func NewO11yO11yPromResultWithDefaults() *O11yO11yPromResult`

NewO11yO11yPromResultWithDefaults instantiates a new O11yO11yPromResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *O11yO11yPromResult) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *O11yO11yPromResult) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *O11yO11yPromResult) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *O11yO11yPromResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *O11yO11yPromResult) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *O11yO11yPromResult) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetResultType

`func (o *O11yO11yPromResult) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *O11yO11yPromResult) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *O11yO11yPromResult) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *O11yO11yPromResult) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetStats

`func (o *O11yO11yPromResult) GetStats() interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *O11yO11yPromResult) GetStatsOk() (*interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *O11yO11yPromResult) SetStats(v interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *O11yO11yPromResult) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *O11yO11yPromResult) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *O11yO11yPromResult) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


