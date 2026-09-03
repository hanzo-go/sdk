# ProjectsUploadGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **int64** | ExpiresAt is when the grant stops being accepted, as Unix seconds. It is short-lived by design and is handed out ONCE, on the response that queues the deployment — a later read of that deployment does not carry it, so a grant cannot be fetched again after the build it was minted for. | [optional] 
**Fields** | Pointer to **map[string]string** | Fields are form values every POST must carry VERBATIM, alongside &#x60;key&#x60; and &#x60;file&#x60;. The signature covers them, so altering any one of them — including widening the key to reach outside the prefix — invalidates the grant rather than extending it. | [optional] 
**MaxBytes** | Pointer to **int64** | MaxBytes bounds ONE object, not the upload as a whole. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the only place this grant can write: the deployment&#39;s own key prefix. It authorizes WRITES ONLY, which is why completing a deployment reconciles the prefix against a manifest instead of letting CI delete. | [optional] 
**Url** | Pointer to **string** | URL is the address to POST each object to. It is signed for the PUBLIC endpoint, because the signature covers the host and CI posts from outside the cluster. | [optional] 

## Methods

### NewProjectsUploadGrant

`func NewProjectsUploadGrant() *ProjectsUploadGrant`

NewProjectsUploadGrant instantiates a new ProjectsUploadGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsUploadGrantWithDefaults

`func NewProjectsUploadGrantWithDefaults() *ProjectsUploadGrant`

NewProjectsUploadGrantWithDefaults instantiates a new ProjectsUploadGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *ProjectsUploadGrant) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ProjectsUploadGrant) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ProjectsUploadGrant) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ProjectsUploadGrant) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFields

`func (o *ProjectsUploadGrant) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ProjectsUploadGrant) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ProjectsUploadGrant) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ProjectsUploadGrant) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMaxBytes

`func (o *ProjectsUploadGrant) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *ProjectsUploadGrant) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *ProjectsUploadGrant) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *ProjectsUploadGrant) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetPrefix

`func (o *ProjectsUploadGrant) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ProjectsUploadGrant) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ProjectsUploadGrant) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ProjectsUploadGrant) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectsUploadGrant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectsUploadGrant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectsUploadGrant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectsUploadGrant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


