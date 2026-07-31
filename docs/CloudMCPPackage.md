# CloudMCPPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | Identifier is the package name or download URL. | [optional] 
**Registry** | Pointer to **string** | Registry is where the package is fetched from: npm, pypi, oci, nuget, mcpb. | [optional] 
**Runtime** | Pointer to **string** | Runtime is the publisher&#39;s hint for what launches it: npx, uvx, docker. | [optional] 
**Transport** | Pointer to **string** | Transport is what the launched process speaks: usually \&quot;stdio\&quot;. | [optional] 
**Version** | Pointer to **string** | Version is the exact published package version. | [optional] 

## Methods

### NewCloudMCPPackage

`func NewCloudMCPPackage() *CloudMCPPackage`

NewCloudMCPPackage instantiates a new CloudMCPPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMCPPackageWithDefaults

`func NewCloudMCPPackageWithDefaults() *CloudMCPPackage`

NewCloudMCPPackageWithDefaults instantiates a new CloudMCPPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CloudMCPPackage) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CloudMCPPackage) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CloudMCPPackage) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CloudMCPPackage) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRegistry

`func (o *CloudMCPPackage) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *CloudMCPPackage) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *CloudMCPPackage) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *CloudMCPPackage) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRuntime

`func (o *CloudMCPPackage) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *CloudMCPPackage) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *CloudMCPPackage) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *CloudMCPPackage) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTransport

`func (o *CloudMCPPackage) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *CloudMCPPackage) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *CloudMCPPackage) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *CloudMCPPackage) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetVersion

`func (o *CloudMCPPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudMCPPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudMCPPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudMCPPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


