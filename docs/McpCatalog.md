# McpCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to [**[]MCPListing**](MCPListing.md) | Catalog is this page of listings, featured first, then by name. | [optional] 
**Limit** | Pointer to **int32** | Limit is the page size that was actually applied — the default or the clamp, when the request asked for neither or for too much. | [optional] 
**Offset** | Pointer to **int32** | Offset is where this page started, so a caller pages from what the server did rather than from what it asked for. | [optional] 
**Total** | Pointer to **int32** | Total is how many listings the filter matched, which is more than this page holds whenever there is a next one. | [optional] 

## Methods

### NewMcpCatalog

`func NewMcpCatalog() *McpCatalog`

NewMcpCatalog instantiates a new McpCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpCatalogWithDefaults

`func NewMcpCatalogWithDefaults() *McpCatalog`

NewMcpCatalogWithDefaults instantiates a new McpCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *McpCatalog) GetCatalog() []MCPListing`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *McpCatalog) GetCatalogOk() (*[]MCPListing, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *McpCatalog) SetCatalog(v []MCPListing)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *McpCatalog) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetLimit

`func (o *McpCatalog) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *McpCatalog) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *McpCatalog) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *McpCatalog) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *McpCatalog) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *McpCatalog) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *McpCatalog) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *McpCatalog) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *McpCatalog) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *McpCatalog) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *McpCatalog) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *McpCatalog) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


