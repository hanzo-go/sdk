# O11yLogQueryResultDataStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingester** | Pointer to **map[string]interface{}** |  | [optional] 
**Store** | Pointer to **map[string]interface{}** |  | [optional] 
**Summary** | Pointer to [**O11yLogQueryResultDataStatsSummary**](O11yLogQueryResultDataStatsSummary.md) |  | [optional] 

## Methods

### NewO11yLogQueryResultDataStats

`func NewO11yLogQueryResultDataStats() *O11yLogQueryResultDataStats`

NewO11yLogQueryResultDataStats instantiates a new O11yLogQueryResultDataStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yLogQueryResultDataStatsWithDefaults

`func NewO11yLogQueryResultDataStatsWithDefaults() *O11yLogQueryResultDataStats`

NewO11yLogQueryResultDataStatsWithDefaults instantiates a new O11yLogQueryResultDataStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngester

`func (o *O11yLogQueryResultDataStats) GetIngester() map[string]interface{}`

GetIngester returns the Ingester field if non-nil, zero value otherwise.

### GetIngesterOk

`func (o *O11yLogQueryResultDataStats) GetIngesterOk() (*map[string]interface{}, bool)`

GetIngesterOk returns a tuple with the Ingester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngester

`func (o *O11yLogQueryResultDataStats) SetIngester(v map[string]interface{})`

SetIngester sets Ingester field to given value.

### HasIngester

`func (o *O11yLogQueryResultDataStats) HasIngester() bool`

HasIngester returns a boolean if a field has been set.

### GetStore

`func (o *O11yLogQueryResultDataStats) GetStore() map[string]interface{}`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *O11yLogQueryResultDataStats) GetStoreOk() (*map[string]interface{}, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *O11yLogQueryResultDataStats) SetStore(v map[string]interface{})`

SetStore sets Store field to given value.

### HasStore

`func (o *O11yLogQueryResultDataStats) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetSummary

`func (o *O11yLogQueryResultDataStats) GetSummary() O11yLogQueryResultDataStatsSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *O11yLogQueryResultDataStats) GetSummaryOk() (*O11yLogQueryResultDataStatsSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *O11yLogQueryResultDataStats) SetSummary(v O11yLogQueryResultDataStatsSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *O11yLogQueryResultDataStats) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


