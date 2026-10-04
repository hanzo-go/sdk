# DeclareReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to **bool** |  | [optional] 
**Dockerfile** | Pointer to **string** |  | [optional] 
**Env** | Pointer to [**[]DeclareEnv**](DeclareEnv.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**PartOf** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewDeclareReq

`func NewDeclareReq() *DeclareReq`

NewDeclareReq instantiates a new DeclareReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclareReqWithDefaults

`func NewDeclareReqWithDefaults() *DeclareReq`

NewDeclareReqWithDefaults instantiates a new DeclareReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *DeclareReq) GetBuild() bool`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *DeclareReq) GetBuildOk() (*bool, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *DeclareReq) SetBuild(v bool)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *DeclareReq) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetDockerfile

`func (o *DeclareReq) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *DeclareReq) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *DeclareReq) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *DeclareReq) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetEnv

`func (o *DeclareReq) GetEnv() []DeclareEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *DeclareReq) GetEnvOk() (*[]DeclareEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *DeclareReq) SetEnv(v []DeclareEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *DeclareReq) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHost

`func (o *DeclareReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DeclareReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DeclareReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DeclareReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMode

`func (o *DeclareReq) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DeclareReq) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DeclareReq) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DeclareReq) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *DeclareReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeclareReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeclareReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeclareReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *DeclareReq) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *DeclareReq) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *DeclareReq) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *DeclareReq) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPartOf

`func (o *DeclareReq) GetPartOf() string`

GetPartOf returns the PartOf field if non-nil, zero value otherwise.

### GetPartOfOk

`func (o *DeclareReq) GetPartOfOk() (*string, bool)`

GetPartOfOk returns a tuple with the PartOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOf

`func (o *DeclareReq) SetPartOf(v string)`

SetPartOf sets PartOf field to given value.

### HasPartOf

`func (o *DeclareReq) HasPartOf() bool`

HasPartOf returns a boolean if a field has been set.

### GetPort

`func (o *DeclareReq) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeclareReq) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeclareReq) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DeclareReq) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRef

`func (o *DeclareReq) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DeclareReq) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DeclareReq) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DeclareReq) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetReplicas

`func (o *DeclareReq) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *DeclareReq) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *DeclareReq) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *DeclareReq) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepo

`func (o *DeclareReq) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *DeclareReq) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *DeclareReq) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *DeclareReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetTag

`func (o *DeclareReq) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DeclareReq) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DeclareReq) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DeclareReq) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


