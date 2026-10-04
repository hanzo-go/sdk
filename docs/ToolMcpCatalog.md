# ToolMcpCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to [**[]ToolMCPListing**](ToolMCPListing.md) | Catalog is this page of listings, featured first, then by name. | [optional] 
**Limit** | Pointer to **int64** | Limit is the page size that was actually applied — the default or the clamp, when the request asked for neither or for too much. | [optional] 
**Offset** | Pointer to **int64** | Offset is where this page started, so a caller pages from what the server did rather than from what it asked for. | [optional] 
**Total** | Pointer to **int64** | Total is how many listings the filter matched, which is more than this page holds whenever there is a next one. | [optional] 

## Methods

### NewToolMcpCatalog

`func NewToolMcpCatalog() *ToolMcpCatalog`

NewToolMcpCatalog instantiates a new ToolMcpCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolMcpCatalogWithDefaults

`func NewToolMcpCatalogWithDefaults() *ToolMcpCatalog`

NewToolMcpCatalogWithDefaults instantiates a new ToolMcpCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *ToolMcpCatalog) GetCatalog() []ToolMCPListing`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ToolMcpCatalog) GetCatalogOk() (*[]ToolMCPListing, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ToolMcpCatalog) SetCatalog(v []ToolMCPListing)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *ToolMcpCatalog) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetLimit

`func (o *ToolMcpCatalog) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ToolMcpCatalog) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ToolMcpCatalog) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ToolMcpCatalog) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ToolMcpCatalog) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ToolMcpCatalog) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ToolMcpCatalog) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ToolMcpCatalog) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *ToolMcpCatalog) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ToolMcpCatalog) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ToolMcpCatalog) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ToolMcpCatalog) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


