# CloudMCPServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeader** | Pointer to **string** | AuthHeader is the request header the KMS-held credential is injected into, e.g. \&quot;Authorization\&quot;. Absent when the server needs no credential. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the server was registered, Unix seconds. | [optional] 
**HasSecret** | Pointer to **bool** | HasSecret is whether a credential is sealed in KMS for this server. The VALUE is never returned by any route. | [optional] 
**Id** | Pointer to **string** | ID is the server&#39;s id within the org. It also PREFIXES every tool name the server contributes, which is what keeps two servers&#39; \&quot;search\&quot; apart. | [optional] 
**Listing** | Pointer to **string** | Listing is the catalog entry this server was enabled from, when it was. Empty means the org typed the URL in itself. | [optional] 
**Name** | Pointer to **string** | Name is the org&#39;s label for the server. | [optional] 
**Org** | Pointer to **string** | Org is the org that registered the server — the validated caller&#39;s. | [optional] 
**Source** | Pointer to **string** | Source is where the registration came from: \&quot;catalog\&quot; when it was enabled off the shelf, \&quot;org\&quot; when the org registered the URL itself. It is DERIVED from Listing rather than stored, because two columns for one fact is two chances to disagree. | [optional] 
**Url** | Pointer to **string** | URL is the server&#39;s JSON-RPC endpoint. Always a public http(s) host: the registration boundary and the dialer both refuse anything else. | [optional] 

## Methods

### NewCloudMCPServer

`func NewCloudMCPServer() *CloudMCPServer`

NewCloudMCPServer instantiates a new CloudMCPServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMCPServerWithDefaults

`func NewCloudMCPServerWithDefaults() *CloudMCPServer`

NewCloudMCPServerWithDefaults instantiates a new CloudMCPServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeader

`func (o *CloudMCPServer) GetAuthHeader() string`

GetAuthHeader returns the AuthHeader field if non-nil, zero value otherwise.

### GetAuthHeaderOk

`func (o *CloudMCPServer) GetAuthHeaderOk() (*string, bool)`

GetAuthHeaderOk returns a tuple with the AuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeader

`func (o *CloudMCPServer) SetAuthHeader(v string)`

SetAuthHeader sets AuthHeader field to given value.

### HasAuthHeader

`func (o *CloudMCPServer) HasAuthHeader() bool`

HasAuthHeader returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudMCPServer) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudMCPServer) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudMCPServer) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudMCPServer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHasSecret

`func (o *CloudMCPServer) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *CloudMCPServer) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *CloudMCPServer) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.

### HasHasSecret

`func (o *CloudMCPServer) HasHasSecret() bool`

HasHasSecret returns a boolean if a field has been set.

### GetId

`func (o *CloudMCPServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudMCPServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudMCPServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudMCPServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListing

`func (o *CloudMCPServer) GetListing() string`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *CloudMCPServer) GetListingOk() (*string, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *CloudMCPServer) SetListing(v string)`

SetListing sets Listing field to given value.

### HasListing

`func (o *CloudMCPServer) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetName

`func (o *CloudMCPServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudMCPServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudMCPServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudMCPServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *CloudMCPServer) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudMCPServer) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudMCPServer) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudMCPServer) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSource

`func (o *CloudMCPServer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudMCPServer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudMCPServer) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudMCPServer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrl

`func (o *CloudMCPServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudMCPServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudMCPServer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudMCPServer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


