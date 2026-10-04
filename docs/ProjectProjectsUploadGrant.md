# ProjectProjectsUploadGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Copy** | Pointer to **bool** | Copy is true when the completion of this deployment accepts copyFrom and copy. A client that does not see it uploads every object itself, since a server that ignores those fields would leave the copied keys missing. | [optional] 
**ExpiresAt** | Pointer to **int64** | ExpiresAt is when the grant stops being accepted, as Unix seconds. It is short-lived by design and is handed out ONCE, on the response that queues the deployment — a later read of that deployment does not carry it, so a grant cannot be fetched again after the build it was minted for. | [optional] 
**Fields** | Pointer to **map[string]string** | Fields are form values every POST must carry VERBATIM, alongside &#x60;key&#x60; and &#x60;file&#x60;. The signature covers them, so altering any one of them — including widening the key to reach outside the prefix — invalidates the grant rather than extending it. | [optional] 
**MaxBytes** | Pointer to **int64** | MaxBytes bounds ONE object, not the upload as a whole. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the only place this grant can write: the deployment&#39;s own key prefix. It authorizes WRITES ONLY, which is why completing a deployment reconciles the prefix against a manifest instead of letting CI delete. | [optional] 
**Url** | Pointer to **string** | URL is the address to POST each object to. It is signed for the PUBLIC endpoint, because the signature covers the host and CI posts from outside the cluster. | [optional] 

## Methods

### NewProjectProjectsUploadGrant

`func NewProjectProjectsUploadGrant() *ProjectProjectsUploadGrant`

NewProjectProjectsUploadGrant instantiates a new ProjectProjectsUploadGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsUploadGrantWithDefaults

`func NewProjectProjectsUploadGrantWithDefaults() *ProjectProjectsUploadGrant`

NewProjectProjectsUploadGrantWithDefaults instantiates a new ProjectProjectsUploadGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopy

`func (o *ProjectProjectsUploadGrant) GetCopy() bool`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *ProjectProjectsUploadGrant) GetCopyOk() (*bool, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *ProjectProjectsUploadGrant) SetCopy(v bool)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *ProjectProjectsUploadGrant) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ProjectProjectsUploadGrant) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ProjectProjectsUploadGrant) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ProjectProjectsUploadGrant) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ProjectProjectsUploadGrant) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFields

`func (o *ProjectProjectsUploadGrant) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ProjectProjectsUploadGrant) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ProjectProjectsUploadGrant) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ProjectProjectsUploadGrant) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMaxBytes

`func (o *ProjectProjectsUploadGrant) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *ProjectProjectsUploadGrant) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *ProjectProjectsUploadGrant) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *ProjectProjectsUploadGrant) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetPrefix

`func (o *ProjectProjectsUploadGrant) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ProjectProjectsUploadGrant) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ProjectProjectsUploadGrant) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ProjectProjectsUploadGrant) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectProjectsUploadGrant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectProjectsUploadGrant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectProjectsUploadGrant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectProjectsUploadGrant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


