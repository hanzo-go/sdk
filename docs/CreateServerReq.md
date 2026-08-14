# CreateServerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeader** | Pointer to **string** | AuthHeader is the request header the credential is injected into, e.g. \&quot;Authorization\&quot;. Empty means the server needs no credential. | [optional] 
**Listing** | Pointer to **string** | Listing enables a CATALOG entry instead — the id from GET /v1/tools/catalog. The endpoint is the listing&#39;s own streamable-http remote, so a listing that only ships a stdio package is refused: there is nothing to reach yet. | [optional] 
**Name** | Pointer to **string** | Name labels the server for the org. Required with URL; with Listing it defaults to the listing&#39;s own title. | [optional] 
**Secret** | Pointer to **string** | Secret is the credential VALUE. It is sealed into KMS under a per-org ref and never stored in SQLite, never listed, and never returned. | [optional] 
**Url** | Pointer to **string** | URL is the server&#39;s JSON-RPC endpoint. It must be an http(s) URL naming a PUBLIC host: loopback, link-local, private and cloud-metadata addresses are refused here and again when the dialer connects. | [optional] 

## Methods

### NewCreateServerReq

`func NewCreateServerReq() *CreateServerReq`

NewCreateServerReq instantiates a new CreateServerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerReqWithDefaults

`func NewCreateServerReqWithDefaults() *CreateServerReq`

NewCreateServerReqWithDefaults instantiates a new CreateServerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeader

`func (o *CreateServerReq) GetAuthHeader() string`

GetAuthHeader returns the AuthHeader field if non-nil, zero value otherwise.

### GetAuthHeaderOk

`func (o *CreateServerReq) GetAuthHeaderOk() (*string, bool)`

GetAuthHeaderOk returns a tuple with the AuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeader

`func (o *CreateServerReq) SetAuthHeader(v string)`

SetAuthHeader sets AuthHeader field to given value.

### HasAuthHeader

`func (o *CreateServerReq) HasAuthHeader() bool`

HasAuthHeader returns a boolean if a field has been set.

### GetListing

`func (o *CreateServerReq) GetListing() string`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *CreateServerReq) GetListingOk() (*string, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *CreateServerReq) SetListing(v string)`

SetListing sets Listing field to given value.

### HasListing

`func (o *CreateServerReq) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetName

`func (o *CreateServerReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServerReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServerReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateServerReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecret

`func (o *CreateServerReq) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateServerReq) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateServerReq) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CreateServerReq) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUrl

`func (o *CreateServerReq) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateServerReq) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateServerReq) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateServerReq) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


