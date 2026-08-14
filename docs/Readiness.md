# Readiness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crd** | Pointer to **bool** | CRD is whether the operator App CRD was found, and is absent when no cluster client resolved and the question could not be asked. | [optional] 
**Error** | Pointer to **string** | Error is the real reason the plane is degraded; absent when it is not. | [optional] 
**K8s** | Pointer to **bool** | K8s is whether a cluster client resolved at all. False means no kubeconfig. | [optional] 
**Service** | Pointer to **string** | Service is always \&quot;platform\&quot; — which control plane answered. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when this plane can deploy, \&quot;degraded\&quot; when it cannot. | [optional] 

## Methods

### NewReadiness

`func NewReadiness() *Readiness`

NewReadiness instantiates a new Readiness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadinessWithDefaults

`func NewReadinessWithDefaults() *Readiness`

NewReadinessWithDefaults instantiates a new Readiness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrd

`func (o *Readiness) GetCrd() bool`

GetCrd returns the Crd field if non-nil, zero value otherwise.

### GetCrdOk

`func (o *Readiness) GetCrdOk() (*bool, bool)`

GetCrdOk returns a tuple with the Crd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrd

`func (o *Readiness) SetCrd(v bool)`

SetCrd sets Crd field to given value.

### HasCrd

`func (o *Readiness) HasCrd() bool`

HasCrd returns a boolean if a field has been set.

### GetError

`func (o *Readiness) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Readiness) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Readiness) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Readiness) HasError() bool`

HasError returns a boolean if a field has been set.

### GetK8s

`func (o *Readiness) GetK8s() bool`

GetK8s returns the K8s field if non-nil, zero value otherwise.

### GetK8sOk

`func (o *Readiness) GetK8sOk() (*bool, bool)`

GetK8sOk returns a tuple with the K8s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8s

`func (o *Readiness) SetK8s(v bool)`

SetK8s sets K8s field to given value.

### HasK8s

`func (o *Readiness) HasK8s() bool`

HasK8s returns a boolean if a field has been set.

### GetService

`func (o *Readiness) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Readiness) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Readiness) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *Readiness) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *Readiness) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Readiness) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Readiness) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Readiness) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


