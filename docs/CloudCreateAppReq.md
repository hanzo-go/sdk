# CloudCreateAppReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Dockerfile** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to [**[]CloudEnvVarJSON**](CloudEnvVarJSON.md) |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**CloudCreateAppReqImage**](CloudCreateAppReqImage.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Repo** | Pointer to [**CloudCreateAppReqRepo**](CloudCreateAppReqRepo.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**StorageGb** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudCreateAppReq

`func NewCloudCreateAppReq() *CloudCreateAppReq`

NewCloudCreateAppReq instantiates a new CloudCreateAppReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateAppReqWithDefaults

`func NewCloudCreateAppReqWithDefaults() *CloudCreateAppReq`

NewCloudCreateAppReqWithDefaults instantiates a new CloudCreateAppReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *CloudCreateAppReq) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *CloudCreateAppReq) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *CloudCreateAppReq) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *CloudCreateAppReq) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetDescription

`func (o *CloudCreateAppReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudCreateAppReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudCreateAppReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudCreateAppReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDockerfile

`func (o *CloudCreateAppReq) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *CloudCreateAppReq) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *CloudCreateAppReq) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *CloudCreateAppReq) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetDomains

`func (o *CloudCreateAppReq) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CloudCreateAppReq) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CloudCreateAppReq) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CloudCreateAppReq) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEnv

`func (o *CloudCreateAppReq) GetEnv() []CloudEnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CloudCreateAppReq) GetEnvOk() (*[]CloudEnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CloudCreateAppReq) SetEnv(v []CloudEnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CloudCreateAppReq) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvironment

`func (o *CloudCreateAppReq) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CloudCreateAppReq) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CloudCreateAppReq) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CloudCreateAppReq) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetImage

`func (o *CloudCreateAppReq) GetImage() CloudCreateAppReqImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudCreateAppReq) GetImageOk() (*CloudCreateAppReqImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudCreateAppReq) SetImage(v CloudCreateAppReqImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudCreateAppReq) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *CloudCreateAppReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCreateAppReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCreateAppReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCreateAppReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *CloudCreateAppReq) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudCreateAppReq) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudCreateAppReq) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CloudCreateAppReq) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetReplicas

`func (o *CloudCreateAppReq) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CloudCreateAppReq) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CloudCreateAppReq) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *CloudCreateAppReq) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepo

`func (o *CloudCreateAppReq) GetRepo() CloudCreateAppReqRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudCreateAppReq) GetRepoOk() (*CloudCreateAppReqRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudCreateAppReq) SetRepo(v CloudCreateAppReqRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudCreateAppReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSlug

`func (o *CloudCreateAppReq) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CloudCreateAppReq) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CloudCreateAppReq) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CloudCreateAppReq) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *CloudCreateAppReq) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudCreateAppReq) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudCreateAppReq) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudCreateAppReq) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStorageGb

`func (o *CloudCreateAppReq) GetStorageGb() int32`

GetStorageGb returns the StorageGb field if non-nil, zero value otherwise.

### GetStorageGbOk

`func (o *CloudCreateAppReq) GetStorageGbOk() (*int32, bool)`

GetStorageGbOk returns a tuple with the StorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGb

`func (o *CloudCreateAppReq) SetStorageGb(v int32)`

SetStorageGb sets StorageGb field to given value.

### HasStorageGb

`func (o *CloudCreateAppReq) HasStorageGb() bool`

HasStorageGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


