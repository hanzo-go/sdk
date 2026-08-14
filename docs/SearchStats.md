# SearchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchesPerDay** | Pointer to [**[]DayCount**](DayCount.md) | SearchesPerDay is always empty, for the same reason as totalSearches. | [optional] 
**TotalDocuments** | Pointer to **int32** | TotalDocuments is the sum of every index&#39;s document count. | [optional] 
**TotalSearches** | Pointer to **int32** | TotalSearches is always 0: Meilisearch keeps no query-history counter, so this surface reports the honest zero rather than an estimate. | [optional] 
**TotalSessions** | Pointer to **int32** | TotalSessions is always 0, for the same reason as totalSearches. | [optional] 

## Methods

### NewSearchStats

`func NewSearchStats() *SearchStats`

NewSearchStats instantiates a new SearchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStatsWithDefaults

`func NewSearchStatsWithDefaults() *SearchStats`

NewSearchStatsWithDefaults instantiates a new SearchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchesPerDay

`func (o *SearchStats) GetSearchesPerDay() []DayCount`

GetSearchesPerDay returns the SearchesPerDay field if non-nil, zero value otherwise.

### GetSearchesPerDayOk

`func (o *SearchStats) GetSearchesPerDayOk() (*[]DayCount, bool)`

GetSearchesPerDayOk returns a tuple with the SearchesPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchesPerDay

`func (o *SearchStats) SetSearchesPerDay(v []DayCount)`

SetSearchesPerDay sets SearchesPerDay field to given value.

### HasSearchesPerDay

`func (o *SearchStats) HasSearchesPerDay() bool`

HasSearchesPerDay returns a boolean if a field has been set.

### GetTotalDocuments

`func (o *SearchStats) GetTotalDocuments() int32`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *SearchStats) GetTotalDocumentsOk() (*int32, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *SearchStats) SetTotalDocuments(v int32)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *SearchStats) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.

### GetTotalSearches

`func (o *SearchStats) GetTotalSearches() int32`

GetTotalSearches returns the TotalSearches field if non-nil, zero value otherwise.

### GetTotalSearchesOk

`func (o *SearchStats) GetTotalSearchesOk() (*int32, bool)`

GetTotalSearchesOk returns a tuple with the TotalSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSearches

`func (o *SearchStats) SetTotalSearches(v int32)`

SetTotalSearches sets TotalSearches field to given value.

### HasTotalSearches

`func (o *SearchStats) HasTotalSearches() bool`

HasTotalSearches returns a boolean if a field has been set.

### GetTotalSessions

`func (o *SearchStats) GetTotalSessions() int32`

GetTotalSessions returns the TotalSessions field if non-nil, zero value otherwise.

### GetTotalSessionsOk

`func (o *SearchStats) GetTotalSessionsOk() (*int32, bool)`

GetTotalSessionsOk returns a tuple with the TotalSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSessions

`func (o *SearchStats) SetTotalSessions(v int32)`

SetTotalSessions sets TotalSessions field to given value.

### HasTotalSessions

`func (o *SearchStats) HasTotalSessions() bool`

HasTotalSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


