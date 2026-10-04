# DeployDeployHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crd** | Pointer to **bool** | CRD reports whether the App custom resource is served and listable. Absent when the apiserver was unreachable, because then it is unknown rather than false. | [optional] 
**K8s** | Pointer to **bool** | K8s reports whether the Kubernetes API is reachable. Absent when the probe did not get far enough to find out. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering, so a probe response is attributable when several are collected together. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;ok&#x60; when this deployment can serve the delivery plane, and &#x60;degraded&#x60; otherwise. It agrees with the HTTP status by construction — see StatusCode. | [optional] 

## Methods

### NewDeployDeployHealth

`func NewDeployDeployHealth() *DeployDeployHealth`

NewDeployDeployHealth instantiates a new DeployDeployHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployDeployHealthWithDefaults

`func NewDeployDeployHealthWithDefaults() *DeployDeployHealth`

NewDeployDeployHealthWithDefaults instantiates a new DeployDeployHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrd

`func (o *DeployDeployHealth) GetCrd() bool`

GetCrd returns the Crd field if non-nil, zero value otherwise.

### GetCrdOk

`func (o *DeployDeployHealth) GetCrdOk() (*bool, bool)`

GetCrdOk returns a tuple with the Crd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrd

`func (o *DeployDeployHealth) SetCrd(v bool)`

SetCrd sets Crd field to given value.

### HasCrd

`func (o *DeployDeployHealth) HasCrd() bool`

HasCrd returns a boolean if a field has been set.

### GetK8s

`func (o *DeployDeployHealth) GetK8s() bool`

GetK8s returns the K8s field if non-nil, zero value otherwise.

### GetK8sOk

`func (o *DeployDeployHealth) GetK8sOk() (*bool, bool)`

GetK8sOk returns a tuple with the K8s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8s

`func (o *DeployDeployHealth) SetK8s(v bool)`

SetK8s sets K8s field to given value.

### HasK8s

`func (o *DeployDeployHealth) HasK8s() bool`

HasK8s returns a boolean if a field has been set.

### GetService

`func (o *DeployDeployHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DeployDeployHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DeployDeployHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DeployDeployHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *DeployDeployHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeployDeployHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeployDeployHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeployDeployHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


