# PlatformApplicationSaveGithubProviderRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** |  | 
**Repository** | **string** |  | 
**Branch** | **string** |  | 
**Owner** | **string** |  | 
**BuildPath** | Pointer to **string** |  | [optional] 
**GithubId** | Pointer to **string** |  | [optional] 
**WatchPaths** | Pointer to **string** |  | [optional] 
**TriggerType** | Pointer to **string** |  | [optional] 
**EnableSubmodules** | Pointer to **bool** |  | [optional] 

## Methods

### NewPlatformApplicationSaveGithubProviderRequestJson

`func NewPlatformApplicationSaveGithubProviderRequestJson(applicationId string, repository string, branch string, owner string, ) *PlatformApplicationSaveGithubProviderRequestJson`

NewPlatformApplicationSaveGithubProviderRequestJson instantiates a new PlatformApplicationSaveGithubProviderRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformApplicationSaveGithubProviderRequestJsonWithDefaults

`func NewPlatformApplicationSaveGithubProviderRequestJsonWithDefaults() *PlatformApplicationSaveGithubProviderRequestJson`

NewPlatformApplicationSaveGithubProviderRequestJsonWithDefaults instantiates a new PlatformApplicationSaveGithubProviderRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetRepository

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetBranch

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetOwner

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetBuildPath

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetBuildPath() string`

GetBuildPath returns the BuildPath field if non-nil, zero value otherwise.

### GetBuildPathOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetBuildPathOk() (*string, bool)`

GetBuildPathOk returns a tuple with the BuildPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildPath

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetBuildPath(v string)`

SetBuildPath sets BuildPath field to given value.

### HasBuildPath

`func (o *PlatformApplicationSaveGithubProviderRequestJson) HasBuildPath() bool`

HasBuildPath returns a boolean if a field has been set.

### GetGithubId

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetGithubId() string`

GetGithubId returns the GithubId field if non-nil, zero value otherwise.

### GetGithubIdOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetGithubIdOk() (*string, bool)`

GetGithubIdOk returns a tuple with the GithubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubId

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetGithubId(v string)`

SetGithubId sets GithubId field to given value.

### HasGithubId

`func (o *PlatformApplicationSaveGithubProviderRequestJson) HasGithubId() bool`

HasGithubId returns a boolean if a field has been set.

### GetWatchPaths

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetWatchPaths() string`

GetWatchPaths returns the WatchPaths field if non-nil, zero value otherwise.

### GetWatchPathsOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetWatchPathsOk() (*string, bool)`

GetWatchPathsOk returns a tuple with the WatchPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchPaths

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetWatchPaths(v string)`

SetWatchPaths sets WatchPaths field to given value.

### HasWatchPaths

`func (o *PlatformApplicationSaveGithubProviderRequestJson) HasWatchPaths() bool`

HasWatchPaths returns a boolean if a field has been set.

### GetTriggerType

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *PlatformApplicationSaveGithubProviderRequestJson) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetEnableSubmodules

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetEnableSubmodules() bool`

GetEnableSubmodules returns the EnableSubmodules field if non-nil, zero value otherwise.

### GetEnableSubmodulesOk

`func (o *PlatformApplicationSaveGithubProviderRequestJson) GetEnableSubmodulesOk() (*bool, bool)`

GetEnableSubmodulesOk returns a tuple with the EnableSubmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubmodules

`func (o *PlatformApplicationSaveGithubProviderRequestJson) SetEnableSubmodules(v bool)`

SetEnableSubmodules sets EnableSubmodules field to given value.

### HasEnableSubmodules

`func (o *PlatformApplicationSaveGithubProviderRequestJson) HasEnableSubmodules() bool`

HasEnableSubmodules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


