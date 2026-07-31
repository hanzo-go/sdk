# CloudMcpCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to [**[]CloudMCPListing**](CloudMCPListing.md) | Catalog is this page of listings, featured first, then by name. | [optional] 
**Limit** | Pointer to **int32** | Limit is the page size that was actually applied — the default or the clamp, when the request asked for neither or for too much. | [optional] 
**Offset** | Pointer to **int32** | Offset is where this page started, so a caller pages from what the server did rather than from what it asked for. | [optional] 
**Total** | Pointer to **int32** | Total is how many listings the filter matched, which is more than this page holds whenever there is a next one. | [optional] 

## Methods

### NewCloudMcpCatalog

`func NewCloudMcpCatalog() *CloudMcpCatalog`

NewCloudMcpCatalog instantiates a new CloudMcpCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMcpCatalogWithDefaults

`func NewCloudMcpCatalogWithDefaults() *CloudMcpCatalog`

NewCloudMcpCatalogWithDefaults instantiates a new CloudMcpCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *CloudMcpCatalog) GetCatalog() []CloudMCPListing`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *CloudMcpCatalog) GetCatalogOk() (*[]CloudMCPListing, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *CloudMcpCatalog) SetCatalog(v []CloudMCPListing)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *CloudMcpCatalog) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetLimit

`func (o *CloudMcpCatalog) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CloudMcpCatalog) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CloudMcpCatalog) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CloudMcpCatalog) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CloudMcpCatalog) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CloudMcpCatalog) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CloudMcpCatalog) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CloudMcpCatalog) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *CloudMcpCatalog) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudMcpCatalog) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudMcpCatalog) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudMcpCatalog) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


