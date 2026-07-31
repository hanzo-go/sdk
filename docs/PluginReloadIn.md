# PluginReloadIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the app, from the path. It must be one the manifest declares.  | [optional] 
**Version** | Pointer to **string** | Version is a release tag, resolved to a URL and digest through the origin&#39;s binaries.json index — the same index CI publishes, so there is no second table mapping versions to digests.  | [optional] 
**Url** | Pointer to **string** | URL is the artifact directly, for an origin with no index. Sum is its hex SHA-256 and is REQUIRED with it: zip refuses an unverified download, and so does this.  | [optional] 
**Sum** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | Scope \&quot;host\&quot; applies here only. Default \&quot;fleet\&quot; rolls it out one host at a time, halting on the first host that fails to come up.  | [optional] 

## Methods

### NewPluginReloadIn

`func NewPluginReloadIn() *PluginReloadIn`

NewPluginReloadIn instantiates a new PluginReloadIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginReloadInWithDefaults

`func NewPluginReloadInWithDefaults() *PluginReloadIn`

NewPluginReloadInWithDefaults instantiates a new PluginReloadIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginReloadIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginReloadIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginReloadIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginReloadIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *PluginReloadIn) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginReloadIn) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginReloadIn) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PluginReloadIn) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUrl

`func (o *PluginReloadIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PluginReloadIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PluginReloadIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PluginReloadIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSum

`func (o *PluginReloadIn) GetSum() string`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *PluginReloadIn) GetSumOk() (*string, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *PluginReloadIn) SetSum(v string)`

SetSum sets Sum field to given value.

### HasSum

`func (o *PluginReloadIn) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetScope

`func (o *PluginReloadIn) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PluginReloadIn) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PluginReloadIn) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PluginReloadIn) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


