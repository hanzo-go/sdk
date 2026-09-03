# RunReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**[]EnvVarJSON**](EnvVarJSON.md) | Env is the run&#39;s environment. Keys must match &#x60;^[A-Za-z_][A-Za-z0-9_]*$&#x60;; a variable marked &#x60;secret: true&#x60; is sealed into KMS. | [optional] 
**Gpu** | Pointer to **int64** | GPU is how many GPUs the run asks for; a negative value is 400. | [optional] 
**Image** | Pointer to **string** | Image is the container image to run. Required. | [optional] 
**MaxScale** | Pointer to **int64** | MaxScale above the floor declares an autoscaling ceiling; 0 means no autoscaler at all — a fixed run at the floor. | [optional] 
**MinScale** | Pointer to **int64** | MinScale is the replica floor, clamped to the deployment&#39;s limit. | [optional] 
**Name** | Pointer to **string** | Name is the run&#39;s name, and the slug is derived from it. Required, and it must resolve to &#x60;^[a-z0-9]([a-z0-9-]{0,38}[a-z0-9])?$&#x60;. Re-running the same name updates that run in place. | [optional] 
**Port** | Pointer to **int64** | Port is the container port the run listens on. | [optional] 
**Runtime** | Pointer to **string** | Runtime is accepted for the client contract and echoed nowhere: the image IS the runtime unit. | [optional] 
**Shape** | Pointer to **string** | Shape is a compute size label, echoed back; sizing is the operator&#39;s default. Defaults to \&quot;auto\&quot;. | [optional] 

## Methods

### NewRunReq

`func NewRunReq() *RunReq`

NewRunReq instantiates a new RunReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunReqWithDefaults

`func NewRunReqWithDefaults() *RunReq`

NewRunReqWithDefaults instantiates a new RunReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *RunReq) GetEnv() []EnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *RunReq) GetEnvOk() (*[]EnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *RunReq) SetEnv(v []EnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *RunReq) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetGpu

`func (o *RunReq) GetGpu() int64`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *RunReq) GetGpuOk() (*int64, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *RunReq) SetGpu(v int64)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *RunReq) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetImage

`func (o *RunReq) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RunReq) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RunReq) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *RunReq) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMaxScale

`func (o *RunReq) GetMaxScale() int64`

GetMaxScale returns the MaxScale field if non-nil, zero value otherwise.

### GetMaxScaleOk

`func (o *RunReq) GetMaxScaleOk() (*int64, bool)`

GetMaxScaleOk returns a tuple with the MaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScale

`func (o *RunReq) SetMaxScale(v int64)`

SetMaxScale sets MaxScale field to given value.

### HasMaxScale

`func (o *RunReq) HasMaxScale() bool`

HasMaxScale returns a boolean if a field has been set.

### GetMinScale

`func (o *RunReq) GetMinScale() int64`

GetMinScale returns the MinScale field if non-nil, zero value otherwise.

### GetMinScaleOk

`func (o *RunReq) GetMinScaleOk() (*int64, bool)`

GetMinScaleOk returns a tuple with the MinScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScale

`func (o *RunReq) SetMinScale(v int64)`

SetMinScale sets MinScale field to given value.

### HasMinScale

`func (o *RunReq) HasMinScale() bool`

HasMinScale returns a boolean if a field has been set.

### GetName

`func (o *RunReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RunReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RunReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RunReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *RunReq) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RunReq) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RunReq) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *RunReq) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRuntime

`func (o *RunReq) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *RunReq) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *RunReq) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *RunReq) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetShape

`func (o *RunReq) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *RunReq) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *RunReq) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *RunReq) HasShape() bool`

HasShape returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


