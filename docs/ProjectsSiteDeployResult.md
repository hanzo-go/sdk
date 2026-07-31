# ProjectsSiteDeployResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Canonical live URL, https://&lt;slug&gt;.&lt;apex&gt;. | 
**Slug** | **string** |  | 
**Name** | **string** |  | 
**DeploymentId** | **string** | The recorded deployment id (dep_...). | 
**Files** | **[]string** | Deployed file paths, sorted. | 
**Status** | **string** | Deployment status (live on success). | 

## Methods

### NewProjectsSiteDeployResult

`func NewProjectsSiteDeployResult(url string, slug string, name string, deploymentId string, files []string, status string, ) *ProjectsSiteDeployResult`

NewProjectsSiteDeployResult instantiates a new ProjectsSiteDeployResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsSiteDeployResultWithDefaults

`func NewProjectsSiteDeployResultWithDefaults() *ProjectsSiteDeployResult`

NewProjectsSiteDeployResultWithDefaults instantiates a new ProjectsSiteDeployResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProjectsSiteDeployResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectsSiteDeployResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectsSiteDeployResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSlug

`func (o *ProjectsSiteDeployResult) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsSiteDeployResult) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsSiteDeployResult) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *ProjectsSiteDeployResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsSiteDeployResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsSiteDeployResult) SetName(v string)`

SetName sets Name field to given value.


### GetDeploymentId

`func (o *ProjectsSiteDeployResult) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ProjectsSiteDeployResult) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ProjectsSiteDeployResult) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetFiles

`func (o *ProjectsSiteDeployResult) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectsSiteDeployResult) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectsSiteDeployResult) SetFiles(v []string)`

SetFiles sets Files field to given value.


### GetStatus

`func (o *ProjectsSiteDeployResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsSiteDeployResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsSiteDeployResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


