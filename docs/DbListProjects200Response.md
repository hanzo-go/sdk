# DbListProjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]DbProject**](DbProject.md) |  | [optional] 
**Pagination** | Pointer to [**DbListProjects200ResponsePagination**](DbListProjects200ResponsePagination.md) |  | [optional] 

## Methods

### NewDbListProjects200Response

`func NewDbListProjects200Response() *DbListProjects200Response`

NewDbListProjects200Response instantiates a new DbListProjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbListProjects200ResponseWithDefaults

`func NewDbListProjects200ResponseWithDefaults() *DbListProjects200Response`

NewDbListProjects200ResponseWithDefaults instantiates a new DbListProjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *DbListProjects200Response) GetProjects() []DbProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *DbListProjects200Response) GetProjectsOk() (*[]DbProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *DbListProjects200Response) SetProjects(v []DbProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *DbListProjects200Response) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetPagination

`func (o *DbListProjects200Response) GetPagination() DbListProjects200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DbListProjects200Response) GetPaginationOk() (*DbListProjects200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DbListProjects200Response) SetPagination(v DbListProjects200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DbListProjects200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


