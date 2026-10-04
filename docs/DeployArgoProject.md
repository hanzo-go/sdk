# DeployArgoProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion is the constant \&quot;argoproj.io/v1alpha1\&quot;. A project here is an IAM resource wearing that shape; no argoproj.io object is stored behind it. | [optional] 
**Kind** | Pointer to **string** | Kind is the constant \&quot;AppProject\&quot;. | [optional] 
**Metadata** | Pointer to [**DeployArgoMeta**](DeployArgoMeta.md) | Metadata is the project&#39;s identity: its name is the key an application&#39;s spec.project matches, and is the same string an App CR carries in its app.kubernetes.io/part-of label. | [optional] 
**Spec** | Pointer to [**DeployArgoProjectSpec**](DeployArgoProjectSpec.md) | Spec is the fence the SPA displays — repos, destinations, admitted kinds. | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeployArgoProject

`func NewDeployArgoProject() *DeployArgoProject`

NewDeployArgoProject instantiates a new DeployArgoProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoProjectWithDefaults

`func NewDeployArgoProjectWithDefaults() *DeployArgoProject`

NewDeployArgoProjectWithDefaults instantiates a new DeployArgoProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *DeployArgoProject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DeployArgoProject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DeployArgoProject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DeployArgoProject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *DeployArgoProject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeployArgoProject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeployArgoProject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DeployArgoProject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *DeployArgoProject) GetMetadata() DeployArgoMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeployArgoProject) GetMetadataOk() (*DeployArgoMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeployArgoProject) SetMetadata(v DeployArgoMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeployArgoProject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *DeployArgoProject) GetSpec() DeployArgoProjectSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DeployArgoProject) GetSpecOk() (*DeployArgoProjectSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DeployArgoProject) SetSpec(v DeployArgoProjectSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DeployArgoProject) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *DeployArgoProject) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeployArgoProject) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeployArgoProject) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeployArgoProject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


