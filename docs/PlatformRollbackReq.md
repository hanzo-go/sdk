# PlatformRollbackReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the application&#39;s slug, from the path. | [optional] 
**DeploymentId** | Pointer to **string** | DeploymentID is the deployment to redeploy. Omit it to return to the previous release. | [optional] 
**Project** | Pointer to **string** | Project is the project the application lives under, from the path. | [optional] 

## Methods

### NewPlatformRollbackReq

`func NewPlatformRollbackReq() *PlatformRollbackReq`

NewPlatformRollbackReq instantiates a new PlatformRollbackReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformRollbackReqWithDefaults

`func NewPlatformRollbackReqWithDefaults() *PlatformRollbackReq`

NewPlatformRollbackReqWithDefaults instantiates a new PlatformRollbackReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PlatformRollbackReq) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PlatformRollbackReq) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PlatformRollbackReq) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PlatformRollbackReq) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetDeploymentId

`func (o *PlatformRollbackReq) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *PlatformRollbackReq) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *PlatformRollbackReq) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *PlatformRollbackReq) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetProject

`func (o *PlatformRollbackReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformRollbackReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformRollbackReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformRollbackReq) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


