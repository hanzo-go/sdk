# ProjectsSiteDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** | DeploymentID is the deployment this publish recorded, for the history. | [optional] 
**Files** | Pointer to **[]string** | Files are the site-relative paths that were uploaded, sorted. | [optional] 
**Name** | Pointer to **string** | Name is the project&#39;s display name. | [optional] 
**Slug** | Pointer to **string** | Slug is the project the site was published into, created on the fly when the slug was free. | [optional] 
**Status** | Pointer to **string** | Status is the deployment status, \&quot;live\&quot; on success. | [optional] 
**Url** | Pointer to **string** | URL is the canonical live URL, https://&lt;slug&gt;.&lt;apex&gt; — empty when the subdomain belongs to another tenant and this site has none. | [optional] 

## Methods

### NewProjectsSiteDeploy

`func NewProjectsSiteDeploy() *ProjectsSiteDeploy`

NewProjectsSiteDeploy instantiates a new ProjectsSiteDeploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsSiteDeployWithDefaults

`func NewProjectsSiteDeployWithDefaults() *ProjectsSiteDeploy`

NewProjectsSiteDeployWithDefaults instantiates a new ProjectsSiteDeploy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *ProjectsSiteDeploy) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ProjectsSiteDeploy) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ProjectsSiteDeploy) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ProjectsSiteDeploy) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetFiles

`func (o *ProjectsSiteDeploy) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectsSiteDeploy) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectsSiteDeploy) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProjectsSiteDeploy) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetName

`func (o *ProjectsSiteDeploy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsSiteDeploy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsSiteDeploy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsSiteDeploy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsSiteDeploy) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsSiteDeploy) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsSiteDeploy) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsSiteDeploy) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectsSiteDeploy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsSiteDeploy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsSiteDeploy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectsSiteDeploy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectsSiteDeploy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectsSiteDeploy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectsSiteDeploy) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectsSiteDeploy) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


