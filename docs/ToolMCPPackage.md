# ToolMCPPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | Identifier is the package name or download URL. | [optional] 
**Registry** | Pointer to **string** | Registry is where the package is fetched from: npm, pypi, oci, nuget, mcpb. | [optional] 
**Runtime** | Pointer to **string** | Runtime is the publisher&#39;s hint for what launches it: npx, uvx, docker. | [optional] 
**Transport** | Pointer to **string** | Transport is what the launched process speaks: usually \&quot;stdio\&quot;. | [optional] 
**Version** | Pointer to **string** | Version is the exact published package version. | [optional] 

## Methods

### NewToolMCPPackage

`func NewToolMCPPackage() *ToolMCPPackage`

NewToolMCPPackage instantiates a new ToolMCPPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolMCPPackageWithDefaults

`func NewToolMCPPackageWithDefaults() *ToolMCPPackage`

NewToolMCPPackageWithDefaults instantiates a new ToolMCPPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ToolMCPPackage) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ToolMCPPackage) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ToolMCPPackage) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ToolMCPPackage) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRegistry

`func (o *ToolMCPPackage) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ToolMCPPackage) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ToolMCPPackage) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *ToolMCPPackage) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRuntime

`func (o *ToolMCPPackage) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ToolMCPPackage) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ToolMCPPackage) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *ToolMCPPackage) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTransport

`func (o *ToolMCPPackage) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ToolMCPPackage) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ToolMCPPackage) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ToolMCPPackage) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetVersion

`func (o *ToolMCPPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ToolMCPPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ToolMCPPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ToolMCPPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


