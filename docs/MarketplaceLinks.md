# MarketplaceLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cli** | Pointer to **string** | CLI is the hanzo command that buys it. | [optional] 
**Docs** | Pointer to **string** | Docs is the listing&#39;s documentation, when the seller named one. | [optional] 
**Mcp** | Pointer to [**MarketplaceMCPRef**](MarketplaceMCPRef.md) | MCP is the MCP tool and operation that buys it. | [optional] 

## Methods

### NewMarketplaceLinks

`func NewMarketplaceLinks() *MarketplaceLinks`

NewMarketplaceLinks instantiates a new MarketplaceLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceLinksWithDefaults

`func NewMarketplaceLinksWithDefaults() *MarketplaceLinks`

NewMarketplaceLinksWithDefaults instantiates a new MarketplaceLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCli

`func (o *MarketplaceLinks) GetCli() string`

GetCli returns the Cli field if non-nil, zero value otherwise.

### GetCliOk

`func (o *MarketplaceLinks) GetCliOk() (*string, bool)`

GetCliOk returns a tuple with the Cli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCli

`func (o *MarketplaceLinks) SetCli(v string)`

SetCli sets Cli field to given value.

### HasCli

`func (o *MarketplaceLinks) HasCli() bool`

HasCli returns a boolean if a field has been set.

### GetDocs

`func (o *MarketplaceLinks) GetDocs() string`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *MarketplaceLinks) GetDocsOk() (*string, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *MarketplaceLinks) SetDocs(v string)`

SetDocs sets Docs field to given value.

### HasDocs

`func (o *MarketplaceLinks) HasDocs() bool`

HasDocs returns a boolean if a field has been set.

### GetMcp

`func (o *MarketplaceLinks) GetMcp() MarketplaceMCPRef`

GetMcp returns the Mcp field if non-nil, zero value otherwise.

### GetMcpOk

`func (o *MarketplaceLinks) GetMcpOk() (*MarketplaceMCPRef, bool)`

GetMcpOk returns a tuple with the Mcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcp

`func (o *MarketplaceLinks) SetMcp(v MarketplaceMCPRef)`

SetMcp sets Mcp field to given value.

### HasMcp

`func (o *MarketplaceLinks) HasMcp() bool`

HasMcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


