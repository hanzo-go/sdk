# ToolMCPRemote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transport** | Pointer to **string** | Transport is \&quot;streamable-http\&quot; or \&quot;sse\&quot;. | [optional] 
**Url** | Pointer to **string** | URL is the endpoint. | [optional] 

## Methods

### NewToolMCPRemote

`func NewToolMCPRemote() *ToolMCPRemote`

NewToolMCPRemote instantiates a new ToolMCPRemote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolMCPRemoteWithDefaults

`func NewToolMCPRemoteWithDefaults() *ToolMCPRemote`

NewToolMCPRemoteWithDefaults instantiates a new ToolMCPRemote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransport

`func (o *ToolMCPRemote) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ToolMCPRemote) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ToolMCPRemote) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ToolMCPRemote) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetUrl

`func (o *ToolMCPRemote) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ToolMCPRemote) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ToolMCPRemote) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ToolMCPRemote) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


