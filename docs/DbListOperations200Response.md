# DbListOperations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | Pointer to [**[]DbOperation**](DbOperation.md) |  | [optional] 
**Pagination** | Pointer to [**DbListProjects200ResponsePagination**](DbListProjects200ResponsePagination.md) |  | [optional] 

## Methods

### NewDbListOperations200Response

`func NewDbListOperations200Response() *DbListOperations200Response`

NewDbListOperations200Response instantiates a new DbListOperations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbListOperations200ResponseWithDefaults

`func NewDbListOperations200ResponseWithDefaults() *DbListOperations200Response`

NewDbListOperations200ResponseWithDefaults instantiates a new DbListOperations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *DbListOperations200Response) GetOperations() []DbOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *DbListOperations200Response) GetOperationsOk() (*[]DbOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *DbListOperations200Response) SetOperations(v []DbOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *DbListOperations200Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPagination

`func (o *DbListOperations200Response) GetPagination() DbListProjects200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DbListOperations200Response) GetPaginationOk() (*DbListProjects200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DbListOperations200Response) SetPagination(v DbListProjects200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DbListOperations200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


