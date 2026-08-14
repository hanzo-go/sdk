# RollbackReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the application&#39;s slug, from the path. | [optional] 
**DeploymentId** | Pointer to **string** | DeploymentID is the deployment to redeploy. Omit it to return to the previous release. | [optional] 
**Project** | Pointer to **string** | Project is the project the application lives under, from the path. | [optional] 

## Methods

### NewRollbackReq

`func NewRollbackReq() *RollbackReq`

NewRollbackReq instantiates a new RollbackReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackReqWithDefaults

`func NewRollbackReqWithDefaults() *RollbackReq`

NewRollbackReqWithDefaults instantiates a new RollbackReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *RollbackReq) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *RollbackReq) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *RollbackReq) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *RollbackReq) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetDeploymentId

`func (o *RollbackReq) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *RollbackReq) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *RollbackReq) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *RollbackReq) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetProject

`func (o *RollbackReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RollbackReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RollbackReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *RollbackReq) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


