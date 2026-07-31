# CloudSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **bool** | Degraded is true when retrieval failed and the empty result set is an outage rather than a real absence of matches. Absent on a healthy answer. | [optional] 
**Query** | Pointer to **string** | Query echoes the query that was run. | [optional] 
**Results** | Pointer to [**[]CloudSpan**](CloudSpan.md) | Results are the matching spans, best first. Never null — an empty search is an empty array. | [optional] 
**Type** | Pointer to **string** | Type echoes the retrieval tier that ran, after defaulting. | [optional] 

## Methods

### NewCloudSearchResults

`func NewCloudSearchResults() *CloudSearchResults`

NewCloudSearchResults instantiates a new CloudSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSearchResultsWithDefaults

`func NewCloudSearchResultsWithDefaults() *CloudSearchResults`

NewCloudSearchResultsWithDefaults instantiates a new CloudSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *CloudSearchResults) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *CloudSearchResults) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *CloudSearchResults) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *CloudSearchResults) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetQuery

`func (o *CloudSearchResults) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CloudSearchResults) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CloudSearchResults) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CloudSearchResults) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResults

`func (o *CloudSearchResults) GetResults() []CloudSpan`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CloudSearchResults) GetResultsOk() (*[]CloudSpan, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CloudSearchResults) SetResults(v []CloudSpan)`

SetResults sets Results field to given value.

### HasResults

`func (o *CloudSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetType

`func (o *CloudSearchResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudSearchResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudSearchResults) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudSearchResults) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


