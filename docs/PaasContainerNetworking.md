# PaasContainerNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerPort** | Pointer to **int32** |  | [optional] 
**Ingress** | Pointer to [**PaasContainerNetworkingIngress**](PaasContainerNetworkingIngress.md) |  | [optional] 

## Methods

### NewPaasContainerNetworking

`func NewPaasContainerNetworking() *PaasContainerNetworking`

NewPaasContainerNetworking instantiates a new PaasContainerNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasContainerNetworkingWithDefaults

`func NewPaasContainerNetworkingWithDefaults() *PaasContainerNetworking`

NewPaasContainerNetworkingWithDefaults instantiates a new PaasContainerNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerPort

`func (o *PaasContainerNetworking) GetContainerPort() int32`

GetContainerPort returns the ContainerPort field if non-nil, zero value otherwise.

### GetContainerPortOk

`func (o *PaasContainerNetworking) GetContainerPortOk() (*int32, bool)`

GetContainerPortOk returns a tuple with the ContainerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPort

`func (o *PaasContainerNetworking) SetContainerPort(v int32)`

SetContainerPort sets ContainerPort field to given value.

### HasContainerPort

`func (o *PaasContainerNetworking) HasContainerPort() bool`

HasContainerPort returns a boolean if a field has been set.

### GetIngress

`func (o *PaasContainerNetworking) GetIngress() PaasContainerNetworkingIngress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *PaasContainerNetworking) GetIngressOk() (*PaasContainerNetworkingIngress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *PaasContainerNetworking) SetIngress(v PaasContainerNetworkingIngress)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *PaasContainerNetworking) HasIngress() bool`

HasIngress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


