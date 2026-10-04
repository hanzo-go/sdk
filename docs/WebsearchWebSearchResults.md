# WebsearchWebSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engines** | Pointer to [**[]WebsearchWebEngine**](WebsearchWebEngine.md) | Engines is one entry per engine asked, in the order they were asked. It is ADDITIVE to the SearXNG contract, which the LibreChat client ignores as an unknown field exactly as it ignores &#x60;engine&#x60; on a result. | [optional] 
**NumberOfResults** | Pointer to **int64** | NumberOfResults is len(results) — what this answer carries, never an estimate of what the web holds. | [optional] 
**Query** | Pointer to **string** | Query is the query that ran, echoed back. | [optional] 
**Results** | Pointer to [**[]WebsearchWebResult**](WebsearchWebResult.md) | Results are the merged hits, deduplicated by normalised URL and capped at 30. Always an array and never null: no hits is an ANSWER, not a fault. | [optional] 

## Methods

### NewWebsearchWebSearchResults

`func NewWebsearchWebSearchResults() *WebsearchWebSearchResults`

NewWebsearchWebSearchResults instantiates a new WebsearchWebSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchWebSearchResultsWithDefaults

`func NewWebsearchWebSearchResultsWithDefaults() *WebsearchWebSearchResults`

NewWebsearchWebSearchResultsWithDefaults instantiates a new WebsearchWebSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngines

`func (o *WebsearchWebSearchResults) GetEngines() []WebsearchWebEngine`

GetEngines returns the Engines field if non-nil, zero value otherwise.

### GetEnginesOk

`func (o *WebsearchWebSearchResults) GetEnginesOk() (*[]WebsearchWebEngine, bool)`

GetEnginesOk returns a tuple with the Engines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngines

`func (o *WebsearchWebSearchResults) SetEngines(v []WebsearchWebEngine)`

SetEngines sets Engines field to given value.

### HasEngines

`func (o *WebsearchWebSearchResults) HasEngines() bool`

HasEngines returns a boolean if a field has been set.

### GetNumberOfResults

`func (o *WebsearchWebSearchResults) GetNumberOfResults() int64`

GetNumberOfResults returns the NumberOfResults field if non-nil, zero value otherwise.

### GetNumberOfResultsOk

`func (o *WebsearchWebSearchResults) GetNumberOfResultsOk() (*int64, bool)`

GetNumberOfResultsOk returns a tuple with the NumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfResults

`func (o *WebsearchWebSearchResults) SetNumberOfResults(v int64)`

SetNumberOfResults sets NumberOfResults field to given value.

### HasNumberOfResults

`func (o *WebsearchWebSearchResults) HasNumberOfResults() bool`

HasNumberOfResults returns a boolean if a field has been set.

### GetQuery

`func (o *WebsearchWebSearchResults) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *WebsearchWebSearchResults) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *WebsearchWebSearchResults) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *WebsearchWebSearchResults) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResults

`func (o *WebsearchWebSearchResults) GetResults() []WebsearchWebResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WebsearchWebSearchResults) GetResultsOk() (*[]WebsearchWebResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WebsearchWebSearchResults) SetResults(v []WebsearchWebResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *WebsearchWebSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


