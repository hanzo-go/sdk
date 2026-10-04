# ToolMCPServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admitted** | Pointer to **bool** | Admitted is whether an admin of the org, or a SuperAdmin, registered the server. Only an admitted server is carried into the org&#39;s agent runs: a server registered before registering one took an admin is not, until an admin registers it again. | [optional] 
**AuthHeader** | Pointer to **string** | AuthHeader is the request header the KMS-held credential is injected into, e.g. \&quot;Authorization\&quot;. Absent when the server needs no credential. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the server was registered, Unix seconds. | [optional] 
**HasSecret** | Pointer to **bool** | HasSecret is whether a credential is sealed in KMS for this server. The VALUE is never returned by any route. | [optional] 
**Id** | Pointer to **string** | ID is the server&#39;s id within the org. It also PREFIXES every tool name the server contributes, which is what keeps two servers&#39; \&quot;search\&quot; apart. | [optional] 
**Listing** | Pointer to **string** | Listing is the catalog entry this server was enabled from, when it was. Empty means the org typed the URL in itself. | [optional] 
**Name** | Pointer to **string** | Name is the org&#39;s label for the server. | [optional] 
**Org** | Pointer to **string** | Org is the org that registered the server — the validated caller&#39;s. | [optional] 
**Reason** | Pointer to **string** | Reason says why a server that is not \&quot;ok\&quot; is not, in words any member may be shown: never the server&#39;s address. | [optional] 
**Source** | Pointer to **string** | Source is where the registration came from: \&quot;catalog\&quot; when it was enabled off the shelf, \&quot;org\&quot; when the org registered the URL itself. It is DERIVED from Listing rather than stored, because two columns for one fact is two chances to disagree. | [optional] 
**Status** | Pointer to **string** | Status is whether the server&#39;s tools could be listed when the org&#39;s servers were read: \&quot;ok\&quot;, \&quot;unreachable\&quot;, \&quot;refused\&quot; or \&quot;error\&quot;. Only the listing answers it; a registration does not. | [optional] 
**Tools** | Pointer to **int64** | Tools is how many tools the server listed, when it listed them. | [optional] 
**Url** | Pointer to **string** | URL is the server&#39;s JSON-RPC endpoint. Always a public http(s) host: the registration boundary and the dialer both refuse anything else. | [optional] 

## Methods

### NewToolMCPServer

`func NewToolMCPServer() *ToolMCPServer`

NewToolMCPServer instantiates a new ToolMCPServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolMCPServerWithDefaults

`func NewToolMCPServerWithDefaults() *ToolMCPServer`

NewToolMCPServerWithDefaults instantiates a new ToolMCPServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmitted

`func (o *ToolMCPServer) GetAdmitted() bool`

GetAdmitted returns the Admitted field if non-nil, zero value otherwise.

### GetAdmittedOk

`func (o *ToolMCPServer) GetAdmittedOk() (*bool, bool)`

GetAdmittedOk returns a tuple with the Admitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmitted

`func (o *ToolMCPServer) SetAdmitted(v bool)`

SetAdmitted sets Admitted field to given value.

### HasAdmitted

`func (o *ToolMCPServer) HasAdmitted() bool`

HasAdmitted returns a boolean if a field has been set.

### GetAuthHeader

`func (o *ToolMCPServer) GetAuthHeader() string`

GetAuthHeader returns the AuthHeader field if non-nil, zero value otherwise.

### GetAuthHeaderOk

`func (o *ToolMCPServer) GetAuthHeaderOk() (*string, bool)`

GetAuthHeaderOk returns a tuple with the AuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeader

`func (o *ToolMCPServer) SetAuthHeader(v string)`

SetAuthHeader sets AuthHeader field to given value.

### HasAuthHeader

`func (o *ToolMCPServer) HasAuthHeader() bool`

HasAuthHeader returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ToolMCPServer) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolMCPServer) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolMCPServer) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ToolMCPServer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHasSecret

`func (o *ToolMCPServer) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *ToolMCPServer) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *ToolMCPServer) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.

### HasHasSecret

`func (o *ToolMCPServer) HasHasSecret() bool`

HasHasSecret returns a boolean if a field has been set.

### GetId

`func (o *ToolMCPServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolMCPServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolMCPServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ToolMCPServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListing

`func (o *ToolMCPServer) GetListing() string`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *ToolMCPServer) GetListingOk() (*string, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *ToolMCPServer) SetListing(v string)`

SetListing sets Listing field to given value.

### HasListing

`func (o *ToolMCPServer) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetName

`func (o *ToolMCPServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolMCPServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolMCPServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolMCPServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *ToolMCPServer) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ToolMCPServer) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ToolMCPServer) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ToolMCPServer) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetReason

`func (o *ToolMCPServer) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ToolMCPServer) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ToolMCPServer) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ToolMCPServer) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSource

`func (o *ToolMCPServer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ToolMCPServer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ToolMCPServer) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ToolMCPServer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ToolMCPServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ToolMCPServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ToolMCPServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ToolMCPServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTools

`func (o *ToolMCPServer) GetTools() int64`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ToolMCPServer) GetToolsOk() (*int64, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ToolMCPServer) SetTools(v int64)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ToolMCPServer) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetUrl

`func (o *ToolMCPServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ToolMCPServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ToolMCPServer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ToolMCPServer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


