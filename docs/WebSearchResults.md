# WebSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engines** | Pointer to [**[]WebEngine**](WebEngine.md) | Engines is one entry per engine asked, in the order they were asked. It is ADDITIVE to the SearXNG contract, which the LibreChat client ignores as an unknown field exactly as it ignores &#x60;engine&#x60; on a result. | [optional] 
**NumberOfResults** | Pointer to **int64** | NumberOfResults is len(results) — what this answer carries, never an estimate of what the web holds. | [optional] 
**Query** | Pointer to **string** | Query is the query that ran, echoed back. | [optional] 
**Results** | Pointer to [**[]WebResult**](WebResult.md) | Results are the merged hits, deduplicated by normalised URL and capped at 30. Always an array and never null: no hits is an ANSWER, not a fault. | [optional] 

## Methods

### NewWebSearchResults

`func NewWebSearchResults() *WebSearchResults`

NewWebSearchResults instantiates a new WebSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSearchResultsWithDefaults

`func NewWebSearchResultsWithDefaults() *WebSearchResults`

NewWebSearchResultsWithDefaults instantiates a new WebSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngines

`func (o *WebSearchResults) GetEngines() []WebEngine`

GetEngines returns the Engines field if non-nil, zero value otherwise.

### GetEnginesOk

`func (o *WebSearchResults) GetEnginesOk() (*[]WebEngine, bool)`

GetEnginesOk returns a tuple with the Engines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngines

`func (o *WebSearchResults) SetEngines(v []WebEngine)`

SetEngines sets Engines field to given value.

### HasEngines

`func (o *WebSearchResults) HasEngines() bool`

HasEngines returns a boolean if a field has been set.

### GetNumberOfResults

`func (o *WebSearchResults) GetNumberOfResults() int64`

GetNumberOfResults returns the NumberOfResults field if non-nil, zero value otherwise.

### GetNumberOfResultsOk

`func (o *WebSearchResults) GetNumberOfResultsOk() (*int64, bool)`

GetNumberOfResultsOk returns a tuple with the NumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfResults

`func (o *WebSearchResults) SetNumberOfResults(v int64)`

SetNumberOfResults sets NumberOfResults field to given value.

### HasNumberOfResults

`func (o *WebSearchResults) HasNumberOfResults() bool`

HasNumberOfResults returns a boolean if a field has been set.

### GetQuery

`func (o *WebSearchResults) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *WebSearchResults) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *WebSearchResults) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *WebSearchResults) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResults

`func (o *WebSearchResults) GetResults() []WebResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WebSearchResults) GetResultsOk() (*[]WebResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WebSearchResults) SetResults(v []WebResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *WebSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


