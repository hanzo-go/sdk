# O11yO11yDomainsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Results is one entry per query. Each entry&#39;s shape follows the answer&#39;s type — time-series data, scalar data or raw rows — so the bytes pass through verbatim. | [optional] 

## Methods

### NewO11yO11yDomainsData

`func NewO11yO11yDomainsData() *O11yO11yDomainsData`

NewO11yO11yDomainsData instantiates a new O11yO11yDomainsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDomainsDataWithDefaults

`func NewO11yO11yDomainsDataWithDefaults() *O11yO11yDomainsData`

NewO11yO11yDomainsDataWithDefaults instantiates a new O11yO11yDomainsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *O11yO11yDomainsData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *O11yO11yDomainsData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *O11yO11yDomainsData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *O11yO11yDomainsData) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


