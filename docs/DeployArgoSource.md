# DeployArgoSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path is the directory within RepoURL. Display-only alongside a display-only RepoURL; CD&#39;s own value for a CD row. | [optional] 
**RepoURL** | Pointer to **string** | RepoURL is the git repository the desired state comes from. For an application projected from an App CR it is the fleet manifest repo and is DISPLAY ONLY — an App CR pins an image, and nothing is rendered from this repo to produce it. For a CD row it is the repo CD actually polls. | [optional] 
**TargetRevision** | Pointer to **string** | TargetRevision is the git ref tracked there — a branch such as \&quot;main\&quot;. Display-only for a projected App CR; the ref CD tracks for a CD row. | [optional] 

## Methods

### NewDeployArgoSource

`func NewDeployArgoSource() *DeployArgoSource`

NewDeployArgoSource instantiates a new DeployArgoSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoSourceWithDefaults

`func NewDeployArgoSourceWithDefaults() *DeployArgoSource`

NewDeployArgoSourceWithDefaults instantiates a new DeployArgoSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *DeployArgoSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeployArgoSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeployArgoSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeployArgoSource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepoURL

`func (o *DeployArgoSource) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *DeployArgoSource) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *DeployArgoSource) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.

### HasRepoURL

`func (o *DeployArgoSource) HasRepoURL() bool`

HasRepoURL returns a boolean if a field has been set.

### GetTargetRevision

`func (o *DeployArgoSource) GetTargetRevision() string`

GetTargetRevision returns the TargetRevision field if non-nil, zero value otherwise.

### GetTargetRevisionOk

`func (o *DeployArgoSource) GetTargetRevisionOk() (*string, bool)`

GetTargetRevisionOk returns a tuple with the TargetRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRevision

`func (o *DeployArgoSource) SetTargetRevision(v string)`

SetTargetRevision sets TargetRevision field to given value.

### HasTargetRevision

`func (o *DeployArgoSource) HasTargetRevision() bool`

HasTargetRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


