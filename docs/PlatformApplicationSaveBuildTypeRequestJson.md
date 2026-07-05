# PlatformApplicationSaveBuildTypeRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** |  | 
**BuildType** | **string** |  | 
**Dockerfile** | Pointer to **string** |  | [optional] 
**PublishDirectory** | Pointer to **string** |  | [optional] 
**DockerContextPath** | Pointer to **string** |  | [optional] 
**DockerBuildStage** | Pointer to **string** |  | [optional] 
**HerokuVersion** | Pointer to **string** |  | [optional] 
**IsStaticSpa** | Pointer to **bool** |  | [optional] 
**RailpackVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformApplicationSaveBuildTypeRequestJson

`func NewPlatformApplicationSaveBuildTypeRequestJson(applicationId string, buildType string, ) *PlatformApplicationSaveBuildTypeRequestJson`

NewPlatformApplicationSaveBuildTypeRequestJson instantiates a new PlatformApplicationSaveBuildTypeRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformApplicationSaveBuildTypeRequestJsonWithDefaults

`func NewPlatformApplicationSaveBuildTypeRequestJsonWithDefaults() *PlatformApplicationSaveBuildTypeRequestJson`

NewPlatformApplicationSaveBuildTypeRequestJsonWithDefaults instantiates a new PlatformApplicationSaveBuildTypeRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetBuildType

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.


### GetDockerfile

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *PlatformApplicationSaveBuildTypeRequestJson) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetPublishDirectory

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetPublishDirectory() string`

GetPublishDirectory returns the PublishDirectory field if non-nil, zero value otherwise.

### GetPublishDirectoryOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetPublishDirectoryOk() (*string, bool)`

GetPublishDirectoryOk returns a tuple with the PublishDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDirectory

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetPublishDirectory(v string)`

SetPublishDirectory sets PublishDirectory field to given value.

### HasPublishDirectory

`func (o *PlatformApplicationSaveBuildTypeRequestJson) HasPublishDirectory() bool`

HasPublishDirectory returns a boolean if a field has been set.

### GetDockerContextPath

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetDockerContextPath() string`

GetDockerContextPath returns the DockerContextPath field if non-nil, zero value otherwise.

### GetDockerContextPathOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetDockerContextPathOk() (*string, bool)`

GetDockerContextPathOk returns a tuple with the DockerContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContextPath

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetDockerContextPath(v string)`

SetDockerContextPath sets DockerContextPath field to given value.

### HasDockerContextPath

`func (o *PlatformApplicationSaveBuildTypeRequestJson) HasDockerContextPath() bool`

HasDockerContextPath returns a boolean if a field has been set.

### GetDockerBuildStage

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetDockerBuildStage() string`

GetDockerBuildStage returns the DockerBuildStage field if non-nil, zero value otherwise.

### GetDockerBuildStageOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetDockerBuildStageOk() (*string, bool)`

GetDockerBuildStageOk returns a tuple with the DockerBuildStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerBuildStage

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetDockerBuildStage(v string)`

SetDockerBuildStage sets DockerBuildStage field to given value.

### HasDockerBuildStage

`func (o *PlatformApplicationSaveBuildTypeRequestJson) HasDockerBuildStage() bool`

HasDockerBuildStage returns a boolean if a field has been set.

### GetHerokuVersion

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetHerokuVersion() string`

GetHerokuVersion returns the HerokuVersion field if non-nil, zero value otherwise.

### GetHerokuVersionOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetHerokuVersionOk() (*string, bool)`

GetHerokuVersionOk returns a tuple with the HerokuVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHerokuVersion

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetHerokuVersion(v string)`

SetHerokuVersion sets HerokuVersion field to given value.

### HasHerokuVersion

`func (o *PlatformApplicationSaveBuildTypeRequestJson) HasHerokuVersion() bool`

HasHerokuVersion returns a boolean if a field has been set.

### GetIsStaticSpa

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetIsStaticSpa() bool`

GetIsStaticSpa returns the IsStaticSpa field if non-nil, zero value otherwise.

### GetIsStaticSpaOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetIsStaticSpaOk() (*bool, bool)`

GetIsStaticSpaOk returns a tuple with the IsStaticSpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaticSpa

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetIsStaticSpa(v bool)`

SetIsStaticSpa sets IsStaticSpa field to given value.

### HasIsStaticSpa

`func (o *PlatformApplicationSaveBuildTypeRequestJson) HasIsStaticSpa() bool`

HasIsStaticSpa returns a boolean if a field has been set.

### GetRailpackVersion

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetRailpackVersion() string`

GetRailpackVersion returns the RailpackVersion field if non-nil, zero value otherwise.

### GetRailpackVersionOk

`func (o *PlatformApplicationSaveBuildTypeRequestJson) GetRailpackVersionOk() (*string, bool)`

GetRailpackVersionOk returns a tuple with the RailpackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailpackVersion

`func (o *PlatformApplicationSaveBuildTypeRequestJson) SetRailpackVersion(v string)`

SetRailpackVersion sets RailpackVersion field to given value.

### HasRailpackVersion

`func (o *PlatformApplicationSaveBuildTypeRequestJson) HasRailpackVersion() bool`

HasRailpackVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


