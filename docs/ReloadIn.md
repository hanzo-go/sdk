# ReloadIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the app, from the path. It must be one the manifest declares. | [optional] 
**Scope** | Pointer to **string** | Scope \&quot;host\&quot; applies here only. Default \&quot;fleet\&quot; rolls it out one host at a time, halting on the first host that fails to come up. | [optional] 
**Sum** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL is the artifact directly, for an origin with no index. Sum is its hex SHA-256 and is REQUIRED with it: zip refuses an unverified download, and so does this. | [optional] 
**Version** | Pointer to **string** | Version is a release tag, resolved to a URL and digest through the origin&#39;s binaries.json index — the same index CI publishes, so there is no second table mapping versions to digests. | [optional] 

## Methods

### NewReloadIn

`func NewReloadIn() *ReloadIn`

NewReloadIn instantiates a new ReloadIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReloadInWithDefaults

`func NewReloadInWithDefaults() *ReloadIn`

NewReloadInWithDefaults instantiates a new ReloadIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReloadIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReloadIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReloadIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReloadIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *ReloadIn) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ReloadIn) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ReloadIn) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ReloadIn) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSum

`func (o *ReloadIn) GetSum() string`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *ReloadIn) GetSumOk() (*string, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *ReloadIn) SetSum(v string)`

SetSum sets Sum field to given value.

### HasSum

`func (o *ReloadIn) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetUrl

`func (o *ReloadIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReloadIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReloadIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ReloadIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersion

`func (o *ReloadIn) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReloadIn) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReloadIn) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReloadIn) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


