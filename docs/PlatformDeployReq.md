# PlatformDeployReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the application&#39;s slug, from the path. | [optional] 
**Commit** | Pointer to **string** | Commit is the git commit or ref to build, for a git-source app. Defaults to the app&#39;s branch. | [optional] 
**Project** | Pointer to **string** | Project is the project the application lives under, from the path. | [optional] 
**Tag** | Pointer to **string** | Tag is the image tag to deploy, for an image-source app. Defaults to the app&#39;s tag, then &#x60;latest&#x60;. | [optional] 

## Methods

### NewPlatformDeployReq

`func NewPlatformDeployReq() *PlatformDeployReq`

NewPlatformDeployReq instantiates a new PlatformDeployReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDeployReqWithDefaults

`func NewPlatformDeployReqWithDefaults() *PlatformDeployReq`

NewPlatformDeployReqWithDefaults instantiates a new PlatformDeployReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PlatformDeployReq) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PlatformDeployReq) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PlatformDeployReq) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PlatformDeployReq) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCommit

`func (o *PlatformDeployReq) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PlatformDeployReq) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PlatformDeployReq) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PlatformDeployReq) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetProject

`func (o *PlatformDeployReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformDeployReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformDeployReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformDeployReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTag

`func (o *PlatformDeployReq) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PlatformDeployReq) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PlatformDeployReq) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PlatformDeployReq) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


