# CloudArgoProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**CloudArgoMeta**](CloudArgoMeta.md) |  | [optional] 
**Spec** | Pointer to [**CloudArgoProjectSpec**](CloudArgoProjectSpec.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCloudArgoProject

`func NewCloudArgoProject() *CloudArgoProject`

NewCloudArgoProject instantiates a new CloudArgoProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoProjectWithDefaults

`func NewCloudArgoProjectWithDefaults() *CloudArgoProject`

NewCloudArgoProjectWithDefaults instantiates a new CloudArgoProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CloudArgoProject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudArgoProject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudArgoProject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudArgoProject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CloudArgoProject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudArgoProject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudArgoProject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudArgoProject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudArgoProject) GetMetadata() CloudArgoMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudArgoProject) GetMetadataOk() (*CloudArgoMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudArgoProject) SetMetadata(v CloudArgoMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CloudArgoProject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CloudArgoProject) GetSpec() CloudArgoProjectSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudArgoProject) GetSpecOk() (*CloudArgoProjectSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudArgoProject) SetSpec(v CloudArgoProjectSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudArgoProject) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CloudArgoProject) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudArgoProject) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudArgoProject) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudArgoProject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


