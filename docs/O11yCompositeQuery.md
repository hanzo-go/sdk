# O11yCompositeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | Pointer to [**[]O11yQueryEnvelope**](O11yQueryEnvelope.md) | Queries is the queries to use for the request. | [optional] 

## Methods

### NewO11yCompositeQuery

`func NewO11yCompositeQuery() *O11yCompositeQuery`

NewO11yCompositeQuery instantiates a new O11yCompositeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yCompositeQueryWithDefaults

`func NewO11yCompositeQueryWithDefaults() *O11yCompositeQuery`

NewO11yCompositeQueryWithDefaults instantiates a new O11yCompositeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *O11yCompositeQuery) GetQueries() []O11yQueryEnvelope`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *O11yCompositeQuery) GetQueriesOk() (*[]O11yQueryEnvelope, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *O11yCompositeQuery) SetQueries(v []O11yQueryEnvelope)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *O11yCompositeQuery) HasQueries() bool`

HasQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


