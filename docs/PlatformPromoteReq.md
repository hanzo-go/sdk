# PlatformPromoteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the application&#39;s slug, from the path. | [optional] 
**DeploymentId** | Pointer to **string** | DeploymentID promotes that deployment&#39;s exact built image. One of this and Tag is required. | [optional] 
**Project** | Pointer to **string** | Project is the project the application lives under, from the path. | [optional] 
**Tag** | Pointer to **string** | Tag promotes an image tag, resolved the same way a deploy resolves one. | [optional] 

## Methods

### NewPlatformPromoteReq

`func NewPlatformPromoteReq() *PlatformPromoteReq`

NewPlatformPromoteReq instantiates a new PlatformPromoteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformPromoteReqWithDefaults

`func NewPlatformPromoteReqWithDefaults() *PlatformPromoteReq`

NewPlatformPromoteReqWithDefaults instantiates a new PlatformPromoteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PlatformPromoteReq) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PlatformPromoteReq) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PlatformPromoteReq) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PlatformPromoteReq) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetDeploymentId

`func (o *PlatformPromoteReq) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *PlatformPromoteReq) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *PlatformPromoteReq) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *PlatformPromoteReq) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetProject

`func (o *PlatformPromoteReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformPromoteReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformPromoteReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformPromoteReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTag

`func (o *PlatformPromoteReq) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PlatformPromoteReq) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PlatformPromoteReq) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PlatformPromoteReq) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


