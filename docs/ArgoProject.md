# ArgoProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ArgoMeta**](ArgoMeta.md) |  | [optional] 
**Spec** | Pointer to [**ArgoProjectSpec**](ArgoProjectSpec.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewArgoProject

`func NewArgoProject() *ArgoProject`

NewArgoProject instantiates a new ArgoProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoProjectWithDefaults

`func NewArgoProjectWithDefaults() *ArgoProject`

NewArgoProjectWithDefaults instantiates a new ArgoProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArgoProject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArgoProject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArgoProject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArgoProject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArgoProject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArgoProject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArgoProject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArgoProject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ArgoProject) GetMetadata() ArgoMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArgoProject) GetMetadataOk() (*ArgoMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArgoProject) SetMetadata(v ArgoMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArgoProject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *ArgoProject) GetSpec() ArgoProjectSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ArgoProject) GetSpecOk() (*ArgoProjectSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ArgoProject) SetSpec(v ArgoProjectSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ArgoProject) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ArgoProject) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArgoProject) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArgoProject) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArgoProject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


