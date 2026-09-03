# IamListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ItemsPerPage** | Pointer to **int64** |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 
**StartIndex** | Pointer to **int64** |  | [optional] 
**TotalResults** | Pointer to **int64** |  | [optional] 

## Methods

### NewIamListResponse

`func NewIamListResponse() *IamListResponse`

NewIamListResponse instantiates a new IamListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamListResponseWithDefaults

`func NewIamListResponseWithDefaults() *IamListResponse`

NewIamListResponseWithDefaults instantiates a new IamListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *IamListResponse) GetResources() []map[string]interface{}`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IamListResponse) GetResourcesOk() (*[]map[string]interface{}, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IamListResponse) SetResources(v []map[string]interface{})`

SetResources sets Resources field to given value.

### HasResources

`func (o *IamListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *IamListResponse) GetItemsPerPage() int64`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *IamListResponse) GetItemsPerPageOk() (*int64, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *IamListResponse) SetItemsPerPage(v int64)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *IamListResponse) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetSchemas

`func (o *IamListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *IamListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *IamListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *IamListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetStartIndex

`func (o *IamListResponse) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *IamListResponse) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *IamListResponse) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *IamListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetTotalResults

`func (o *IamListResponse) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *IamListResponse) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *IamListResponse) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *IamListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


