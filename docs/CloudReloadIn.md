# CloudReloadIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the app, from the path. It must be one the manifest declares. | [optional] 
**Scope** | Pointer to **string** | Scope \&quot;host\&quot; applies here only. Default \&quot;fleet\&quot; rolls it out one host at a time, halting on the first host that fails to come up. | [optional] 
**Sum** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL is the artifact directly, for an origin with no index. Sum is its hex SHA-256 and is REQUIRED with it: zip refuses an unverified download, and so does this. | [optional] 
**Version** | Pointer to **string** | Version is a release tag, resolved to a URL and digest through the origin&#39;s binaries.json index — the same index CI publishes, so there is no second table mapping versions to digests. | [optional] 

## Methods

### NewCloudReloadIn

`func NewCloudReloadIn() *CloudReloadIn`

NewCloudReloadIn instantiates a new CloudReloadIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudReloadInWithDefaults

`func NewCloudReloadInWithDefaults() *CloudReloadIn`

NewCloudReloadInWithDefaults instantiates a new CloudReloadIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudReloadIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudReloadIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudReloadIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudReloadIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *CloudReloadIn) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudReloadIn) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudReloadIn) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudReloadIn) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSum

`func (o *CloudReloadIn) GetSum() string`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *CloudReloadIn) GetSumOk() (*string, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *CloudReloadIn) SetSum(v string)`

SetSum sets Sum field to given value.

### HasSum

`func (o *CloudReloadIn) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetUrl

`func (o *CloudReloadIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudReloadIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudReloadIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudReloadIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersion

`func (o *CloudReloadIn) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudReloadIn) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudReloadIn) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudReloadIn) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


