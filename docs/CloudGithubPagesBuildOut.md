# CloudGithubPagesBuildOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repo** | Pointer to **string** | Repo is the repository the build was queued for. | [optional] 
**Status** | Pointer to **string** | Status is GitHub&#39;s build state at the moment it was queued (\&quot;queued\&quot;). | [optional] 
**Url** | Pointer to **string** | URL is GitHub&#39;s API URL for the build, for polling it there. | [optional] 

## Methods

### NewCloudGithubPagesBuildOut

`func NewCloudGithubPagesBuildOut() *CloudGithubPagesBuildOut`

NewCloudGithubPagesBuildOut instantiates a new CloudGithubPagesBuildOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubPagesBuildOutWithDefaults

`func NewCloudGithubPagesBuildOutWithDefaults() *CloudGithubPagesBuildOut`

NewCloudGithubPagesBuildOutWithDefaults instantiates a new CloudGithubPagesBuildOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepo

`func (o *CloudGithubPagesBuildOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudGithubPagesBuildOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudGithubPagesBuildOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudGithubPagesBuildOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetStatus

`func (o *CloudGithubPagesBuildOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudGithubPagesBuildOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudGithubPagesBuildOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudGithubPagesBuildOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *CloudGithubPagesBuildOut) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudGithubPagesBuildOut) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudGithubPagesBuildOut) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudGithubPagesBuildOut) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


