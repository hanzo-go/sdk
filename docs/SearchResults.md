# SearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **bool** | Degraded is true when retrieval failed and the empty result set is an outage rather than a real absence of matches. Absent on a healthy answer. | [optional] 
**Query** | Pointer to **string** | Query echoes the query that was run. | [optional] 
**Results** | Pointer to [**[]Span**](Span.md) | Results are the matching spans, best first. Never null — an empty search is an empty array. | [optional] 
**Type** | Pointer to **string** | Type echoes the retrieval tier that ran, after defaulting. | [optional] 

## Methods

### NewSearchResults

`func NewSearchResults() *SearchResults`

NewSearchResults instantiates a new SearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsWithDefaults

`func NewSearchResultsWithDefaults() *SearchResults`

NewSearchResultsWithDefaults instantiates a new SearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *SearchResults) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *SearchResults) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *SearchResults) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *SearchResults) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetQuery

`func (o *SearchResults) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchResults) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchResults) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchResults) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResults

`func (o *SearchResults) GetResults() []Span`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchResults) GetResultsOk() (*[]Span, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchResults) SetResults(v []Span)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetType

`func (o *SearchResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchResults) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchResults) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


