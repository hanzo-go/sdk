# ProductSearchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalDocuments** | **int64** | Total documents across all indexes. | 
**TotalSearches** | **int64** | Always 0 — Meilisearch keeps no query-history counters. | 
**TotalSessions** | **int64** | Always 0 — not derivable from the index. | 
**SearchesPerDay** | [**[]ProductDayCount**](ProductDayCount.md) | Always empty — no per-day query history is available. | 

## Methods

### NewProductSearchStats

`func NewProductSearchStats(totalDocuments int64, totalSearches int64, totalSessions int64, searchesPerDay []ProductDayCount, ) *ProductSearchStats`

NewProductSearchStats instantiates a new ProductSearchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSearchStatsWithDefaults

`func NewProductSearchStatsWithDefaults() *ProductSearchStats`

NewProductSearchStatsWithDefaults instantiates a new ProductSearchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalDocuments

`func (o *ProductSearchStats) GetTotalDocuments() int64`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *ProductSearchStats) GetTotalDocumentsOk() (*int64, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *ProductSearchStats) SetTotalDocuments(v int64)`

SetTotalDocuments sets TotalDocuments field to given value.


### GetTotalSearches

`func (o *ProductSearchStats) GetTotalSearches() int64`

GetTotalSearches returns the TotalSearches field if non-nil, zero value otherwise.

### GetTotalSearchesOk

`func (o *ProductSearchStats) GetTotalSearchesOk() (*int64, bool)`

GetTotalSearchesOk returns a tuple with the TotalSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSearches

`func (o *ProductSearchStats) SetTotalSearches(v int64)`

SetTotalSearches sets TotalSearches field to given value.


### GetTotalSessions

`func (o *ProductSearchStats) GetTotalSessions() int64`

GetTotalSessions returns the TotalSessions field if non-nil, zero value otherwise.

### GetTotalSessionsOk

`func (o *ProductSearchStats) GetTotalSessionsOk() (*int64, bool)`

GetTotalSessionsOk returns a tuple with the TotalSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSessions

`func (o *ProductSearchStats) SetTotalSessions(v int64)`

SetTotalSessions sets TotalSessions field to given value.


### GetSearchesPerDay

`func (o *ProductSearchStats) GetSearchesPerDay() []ProductDayCount`

GetSearchesPerDay returns the SearchesPerDay field if non-nil, zero value otherwise.

### GetSearchesPerDayOk

`func (o *ProductSearchStats) GetSearchesPerDayOk() (*[]ProductDayCount, bool)`

GetSearchesPerDayOk returns a tuple with the SearchesPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchesPerDay

`func (o *ProductSearchStats) SetSearchesPerDay(v []ProductDayCount)`

SetSearchesPerDay sets SearchesPerDay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


