# CloudTargetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**CloudMetrics**](CloudMetrics.md) |  | [optional] 
**Spec** | Pointer to [**CloudSpec**](CloudSpec.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudTargetReq

`func NewCloudTargetReq() *CloudTargetReq`

NewCloudTargetReq instantiates a new CloudTargetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTargetReqWithDefaults

`func NewCloudTargetReqWithDefaults() *CloudTargetReq`

NewCloudTargetReqWithDefaults instantiates a new CloudTargetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *CloudTargetReq) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CloudTargetReq) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CloudTargetReq) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CloudTargetReq) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHost

`func (o *CloudTargetReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudTargetReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudTargetReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudTargetReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *CloudTargetReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudTargetReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudTargetReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudTargetReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *CloudTargetReq) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudTargetReq) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudTargetReq) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudTargetReq) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *CloudTargetReq) GetMetrics() CloudMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CloudTargetReq) GetMetricsOk() (*CloudMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CloudTargetReq) SetMetrics(v CloudMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CloudTargetReq) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetSpec

`func (o *CloudTargetReq) GetSpec() CloudSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudTargetReq) GetSpecOk() (*CloudSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudTargetReq) SetSpec(v CloudSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudTargetReq) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CloudTargetReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudTargetReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudTargetReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudTargetReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


