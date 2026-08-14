# O11yGettableIngestionKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**O11yPagination**](O11yPagination.md) |  | [optional] 
**Keys** | Pointer to [**[]O11yIngestionKey**](O11yIngestionKey.md) |  | [optional] 

## Methods

### NewO11yGettableIngestionKeys

`func NewO11yGettableIngestionKeys() *O11yGettableIngestionKeys`

NewO11yGettableIngestionKeys instantiates a new O11yGettableIngestionKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableIngestionKeysWithDefaults

`func NewO11yGettableIngestionKeysWithDefaults() *O11yGettableIngestionKeys`

NewO11yGettableIngestionKeysWithDefaults instantiates a new O11yGettableIngestionKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *O11yGettableIngestionKeys) GetPagination() O11yPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *O11yGettableIngestionKeys) GetPaginationOk() (*O11yPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *O11yGettableIngestionKeys) SetPagination(v O11yPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *O11yGettableIngestionKeys) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetKeys

`func (o *O11yGettableIngestionKeys) GetKeys() []O11yIngestionKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *O11yGettableIngestionKeys) GetKeysOk() (*[]O11yIngestionKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *O11yGettableIngestionKeys) SetKeys(v []O11yIngestionKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *O11yGettableIngestionKeys) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


