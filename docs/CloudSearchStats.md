# CloudSearchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchesPerDay** | Pointer to [**[]CloudDayCount**](CloudDayCount.md) | SearchesPerDay is always empty, for the same reason as totalSearches. | [optional] 
**TotalDocuments** | Pointer to **int32** | TotalDocuments is the sum of every index&#39;s document count. | [optional] 
**TotalSearches** | Pointer to **int32** | TotalSearches is always 0: Meilisearch keeps no query-history counter, so this surface reports the honest zero rather than an estimate. | [optional] 
**TotalSessions** | Pointer to **int32** | TotalSessions is always 0, for the same reason as totalSearches. | [optional] 

## Methods

### NewCloudSearchStats

`func NewCloudSearchStats() *CloudSearchStats`

NewCloudSearchStats instantiates a new CloudSearchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSearchStatsWithDefaults

`func NewCloudSearchStatsWithDefaults() *CloudSearchStats`

NewCloudSearchStatsWithDefaults instantiates a new CloudSearchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchesPerDay

`func (o *CloudSearchStats) GetSearchesPerDay() []CloudDayCount`

GetSearchesPerDay returns the SearchesPerDay field if non-nil, zero value otherwise.

### GetSearchesPerDayOk

`func (o *CloudSearchStats) GetSearchesPerDayOk() (*[]CloudDayCount, bool)`

GetSearchesPerDayOk returns a tuple with the SearchesPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchesPerDay

`func (o *CloudSearchStats) SetSearchesPerDay(v []CloudDayCount)`

SetSearchesPerDay sets SearchesPerDay field to given value.

### HasSearchesPerDay

`func (o *CloudSearchStats) HasSearchesPerDay() bool`

HasSearchesPerDay returns a boolean if a field has been set.

### GetTotalDocuments

`func (o *CloudSearchStats) GetTotalDocuments() int32`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *CloudSearchStats) GetTotalDocumentsOk() (*int32, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *CloudSearchStats) SetTotalDocuments(v int32)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *CloudSearchStats) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.

### GetTotalSearches

`func (o *CloudSearchStats) GetTotalSearches() int32`

GetTotalSearches returns the TotalSearches field if non-nil, zero value otherwise.

### GetTotalSearchesOk

`func (o *CloudSearchStats) GetTotalSearchesOk() (*int32, bool)`

GetTotalSearchesOk returns a tuple with the TotalSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSearches

`func (o *CloudSearchStats) SetTotalSearches(v int32)`

SetTotalSearches sets TotalSearches field to given value.

### HasTotalSearches

`func (o *CloudSearchStats) HasTotalSearches() bool`

HasTotalSearches returns a boolean if a field has been set.

### GetTotalSessions

`func (o *CloudSearchStats) GetTotalSessions() int32`

GetTotalSessions returns the TotalSessions field if non-nil, zero value otherwise.

### GetTotalSessionsOk

`func (o *CloudSearchStats) GetTotalSessionsOk() (*int32, bool)`

GetTotalSessionsOk returns a tuple with the TotalSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSessions

`func (o *CloudSearchStats) SetTotalSessions(v int32)`

SetTotalSessions sets TotalSessions field to given value.

### HasTotalSessions

`func (o *CloudSearchStats) HasTotalSessions() bool`

HasTotalSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


