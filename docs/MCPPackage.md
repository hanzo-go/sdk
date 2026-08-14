# MCPPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | Identifier is the package name or download URL. | [optional] 
**Registry** | Pointer to **string** | Registry is where the package is fetched from: npm, pypi, oci, nuget, mcpb. | [optional] 
**Runtime** | Pointer to **string** | Runtime is the publisher&#39;s hint for what launches it: npx, uvx, docker. | [optional] 
**Transport** | Pointer to **string** | Transport is what the launched process speaks: usually \&quot;stdio\&quot;. | [optional] 
**Version** | Pointer to **string** | Version is the exact published package version. | [optional] 

## Methods

### NewMCPPackage

`func NewMCPPackage() *MCPPackage`

NewMCPPackage instantiates a new MCPPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPPackageWithDefaults

`func NewMCPPackageWithDefaults() *MCPPackage`

NewMCPPackageWithDefaults instantiates a new MCPPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *MCPPackage) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MCPPackage) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MCPPackage) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MCPPackage) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRegistry

`func (o *MCPPackage) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *MCPPackage) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *MCPPackage) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *MCPPackage) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRuntime

`func (o *MCPPackage) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *MCPPackage) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *MCPPackage) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *MCPPackage) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTransport

`func (o *MCPPackage) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *MCPPackage) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *MCPPackage) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *MCPPackage) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetVersion

`func (o *MCPPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MCPPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MCPPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MCPPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


