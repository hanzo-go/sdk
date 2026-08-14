# O11yO11yAnalyzeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query is the query text. Required. | 
**QueryType** | **string** | QueryType says which language the query is in — promql or the datastore&#39;s SQL dialect. Required. | 

## Methods

### NewO11yO11yAnalyzeIn

`func NewO11yO11yAnalyzeIn(query string, queryType string, ) *O11yO11yAnalyzeIn`

NewO11yO11yAnalyzeIn instantiates a new O11yO11yAnalyzeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAnalyzeInWithDefaults

`func NewO11yO11yAnalyzeInWithDefaults() *O11yO11yAnalyzeIn`

NewO11yO11yAnalyzeInWithDefaults instantiates a new O11yO11yAnalyzeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *O11yO11yAnalyzeIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *O11yO11yAnalyzeIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *O11yO11yAnalyzeIn) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetQueryType

`func (o *O11yO11yAnalyzeIn) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *O11yO11yAnalyzeIn) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *O11yO11yAnalyzeIn) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


