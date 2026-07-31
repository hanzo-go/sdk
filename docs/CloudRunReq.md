# CloudRunReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**[]CloudEnvVarJSON**](CloudEnvVarJSON.md) |  | [optional] 
**Gpu** | Pointer to **int32** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**MaxScale** | Pointer to **int32** |  | [optional] 
**MinScale** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Runtime** | Pointer to **string** |  | [optional] 
**Shape** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudRunReq

`func NewCloudRunReq() *CloudRunReq`

NewCloudRunReq instantiates a new CloudRunReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRunReqWithDefaults

`func NewCloudRunReqWithDefaults() *CloudRunReq`

NewCloudRunReqWithDefaults instantiates a new CloudRunReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *CloudRunReq) GetEnv() []CloudEnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CloudRunReq) GetEnvOk() (*[]CloudEnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CloudRunReq) SetEnv(v []CloudEnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CloudRunReq) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetGpu

`func (o *CloudRunReq) GetGpu() int32`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *CloudRunReq) GetGpuOk() (*int32, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *CloudRunReq) SetGpu(v int32)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *CloudRunReq) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetImage

`func (o *CloudRunReq) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudRunReq) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudRunReq) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudRunReq) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMaxScale

`func (o *CloudRunReq) GetMaxScale() int32`

GetMaxScale returns the MaxScale field if non-nil, zero value otherwise.

### GetMaxScaleOk

`func (o *CloudRunReq) GetMaxScaleOk() (*int32, bool)`

GetMaxScaleOk returns a tuple with the MaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScale

`func (o *CloudRunReq) SetMaxScale(v int32)`

SetMaxScale sets MaxScale field to given value.

### HasMaxScale

`func (o *CloudRunReq) HasMaxScale() bool`

HasMaxScale returns a boolean if a field has been set.

### GetMinScale

`func (o *CloudRunReq) GetMinScale() int32`

GetMinScale returns the MinScale field if non-nil, zero value otherwise.

### GetMinScaleOk

`func (o *CloudRunReq) GetMinScaleOk() (*int32, bool)`

GetMinScaleOk returns a tuple with the MinScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScale

`func (o *CloudRunReq) SetMinScale(v int32)`

SetMinScale sets MinScale field to given value.

### HasMinScale

`func (o *CloudRunReq) HasMinScale() bool`

HasMinScale returns a boolean if a field has been set.

### GetName

`func (o *CloudRunReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRunReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRunReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRunReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *CloudRunReq) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudRunReq) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudRunReq) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CloudRunReq) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRuntime

`func (o *CloudRunReq) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *CloudRunReq) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *CloudRunReq) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *CloudRunReq) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetShape

`func (o *CloudRunReq) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *CloudRunReq) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *CloudRunReq) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *CloudRunReq) HasShape() bool`

HasShape returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


