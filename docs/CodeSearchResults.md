# CodeSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capped** | Pointer to **bool** | Capped is true when the semantic tier stopped at its scan cap — it compares a query with at most 50000 stored vectors — with more left unread, so a chunk past the cap could not be found by meaning. The text, regex and symbol tiers read the whole index either way. Absent otherwise. | [optional] 
**Degraded** | Pointer to **bool** | Degraded is true when retrieval failed and the empty result set is an outage rather than a real absence of matches. Absent on a healthy answer. | [optional] 
**Query** | Pointer to **string** | Query echoes the query that was run. | [optional] 
**Results** | Pointer to [**[]CodeSpan**](CodeSpan.md) | Results are the matching spans, best first. Never null — an empty search is an empty array. | [optional] 
**Type** | Pointer to **string** | Type echoes the retrieval tier that ran, after defaulting. | [optional] 

## Methods

### NewCodeSearchResults

`func NewCodeSearchResults() *CodeSearchResults`

NewCodeSearchResults instantiates a new CodeSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeSearchResultsWithDefaults

`func NewCodeSearchResultsWithDefaults() *CodeSearchResults`

NewCodeSearchResultsWithDefaults instantiates a new CodeSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapped

`func (o *CodeSearchResults) GetCapped() bool`

GetCapped returns the Capped field if non-nil, zero value otherwise.

### GetCappedOk

`func (o *CodeSearchResults) GetCappedOk() (*bool, bool)`

GetCappedOk returns a tuple with the Capped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapped

`func (o *CodeSearchResults) SetCapped(v bool)`

SetCapped sets Capped field to given value.

### HasCapped

`func (o *CodeSearchResults) HasCapped() bool`

HasCapped returns a boolean if a field has been set.

### GetDegraded

`func (o *CodeSearchResults) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *CodeSearchResults) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *CodeSearchResults) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *CodeSearchResults) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetQuery

`func (o *CodeSearchResults) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CodeSearchResults) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CodeSearchResults) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CodeSearchResults) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResults

`func (o *CodeSearchResults) GetResults() []CodeSpan`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CodeSearchResults) GetResultsOk() (*[]CodeSpan, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CodeSearchResults) SetResults(v []CodeSpan)`

SetResults sets Results field to given value.

### HasResults

`func (o *CodeSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetType

`func (o *CodeSearchResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CodeSearchResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CodeSearchResults) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CodeSearchResults) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


