# CloudPatchTargetIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | ID is the target to update, from the path. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**CloudMetrics**](CloudMetrics.md) | present &#x3D;&gt; a heartbeat; the server stamps its time | [optional] 
**Spec** | Pointer to [**CloudSpec**](CloudSpec.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudPatchTargetIn

`func NewCloudPatchTargetIn() *CloudPatchTargetIn`

NewCloudPatchTargetIn instantiates a new CloudPatchTargetIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPatchTargetInWithDefaults

`func NewCloudPatchTargetInWithDefaults() *CloudPatchTargetIn`

NewCloudPatchTargetInWithDefaults instantiates a new CloudPatchTargetIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *CloudPatchTargetIn) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CloudPatchTargetIn) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CloudPatchTargetIn) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CloudPatchTargetIn) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHost

`func (o *CloudPatchTargetIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudPatchTargetIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudPatchTargetIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudPatchTargetIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *CloudPatchTargetIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudPatchTargetIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudPatchTargetIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudPatchTargetIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudPatchTargetIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudPatchTargetIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudPatchTargetIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudPatchTargetIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *CloudPatchTargetIn) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudPatchTargetIn) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudPatchTargetIn) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudPatchTargetIn) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *CloudPatchTargetIn) GetMetrics() CloudMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CloudPatchTargetIn) GetMetricsOk() (*CloudMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CloudPatchTargetIn) SetMetrics(v CloudMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CloudPatchTargetIn) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetSpec

`func (o *CloudPatchTargetIn) GetSpec() CloudSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudPatchTargetIn) GetSpecOk() (*CloudSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudPatchTargetIn) SetSpec(v CloudSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudPatchTargetIn) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CloudPatchTargetIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudPatchTargetIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudPatchTargetIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudPatchTargetIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


