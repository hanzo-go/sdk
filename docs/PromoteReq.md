# PromoteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the application&#39;s slug, from the path. | [optional] 
**DeploymentId** | Pointer to **string** | DeploymentID promotes that deployment&#39;s exact built image. One of this and Tag is required. | [optional] 
**Project** | Pointer to **string** | Project is the project the application lives under, from the path. | [optional] 
**Tag** | Pointer to **string** | Tag promotes an image tag, resolved the same way a deploy resolves one. | [optional] 

## Methods

### NewPromoteReq

`func NewPromoteReq() *PromoteReq`

NewPromoteReq instantiates a new PromoteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoteReqWithDefaults

`func NewPromoteReqWithDefaults() *PromoteReq`

NewPromoteReqWithDefaults instantiates a new PromoteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PromoteReq) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PromoteReq) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PromoteReq) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PromoteReq) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetDeploymentId

`func (o *PromoteReq) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *PromoteReq) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *PromoteReq) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *PromoteReq) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetProject

`func (o *PromoteReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PromoteReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PromoteReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PromoteReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTag

`func (o *PromoteReq) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PromoteReq) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PromoteReq) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PromoteReq) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


