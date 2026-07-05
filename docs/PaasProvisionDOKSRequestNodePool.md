# PaasProvisionDOKSRequestNodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Size** | **string** |  | 
**Count** | **int32** |  | 
**AutoScale** | Pointer to **bool** |  | [optional] 
**MinNodes** | Pointer to **int32** |  | [optional] 
**MaxNodes** | Pointer to **int32** |  | [optional] 

## Methods

### NewPaasProvisionDOKSRequestNodePool

`func NewPaasProvisionDOKSRequestNodePool(name string, size string, count int32, ) *PaasProvisionDOKSRequestNodePool`

NewPaasProvisionDOKSRequestNodePool instantiates a new PaasProvisionDOKSRequestNodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasProvisionDOKSRequestNodePoolWithDefaults

`func NewPaasProvisionDOKSRequestNodePoolWithDefaults() *PaasProvisionDOKSRequestNodePool`

NewPaasProvisionDOKSRequestNodePoolWithDefaults instantiates a new PaasProvisionDOKSRequestNodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaasProvisionDOKSRequestNodePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaasProvisionDOKSRequestNodePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaasProvisionDOKSRequestNodePool) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *PaasProvisionDOKSRequestNodePool) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PaasProvisionDOKSRequestNodePool) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PaasProvisionDOKSRequestNodePool) SetSize(v string)`

SetSize sets Size field to given value.


### GetCount

`func (o *PaasProvisionDOKSRequestNodePool) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaasProvisionDOKSRequestNodePool) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaasProvisionDOKSRequestNodePool) SetCount(v int32)`

SetCount sets Count field to given value.


### GetAutoScale

`func (o *PaasProvisionDOKSRequestNodePool) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *PaasProvisionDOKSRequestNodePool) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *PaasProvisionDOKSRequestNodePool) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *PaasProvisionDOKSRequestNodePool) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetMinNodes

`func (o *PaasProvisionDOKSRequestNodePool) GetMinNodes() int32`

GetMinNodes returns the MinNodes field if non-nil, zero value otherwise.

### GetMinNodesOk

`func (o *PaasProvisionDOKSRequestNodePool) GetMinNodesOk() (*int32, bool)`

GetMinNodesOk returns a tuple with the MinNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodes

`func (o *PaasProvisionDOKSRequestNodePool) SetMinNodes(v int32)`

SetMinNodes sets MinNodes field to given value.

### HasMinNodes

`func (o *PaasProvisionDOKSRequestNodePool) HasMinNodes() bool`

HasMinNodes returns a boolean if a field has been set.

### GetMaxNodes

`func (o *PaasProvisionDOKSRequestNodePool) GetMaxNodes() int32`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *PaasProvisionDOKSRequestNodePool) GetMaxNodesOk() (*int32, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *PaasProvisionDOKSRequestNodePool) SetMaxNodes(v int32)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *PaasProvisionDOKSRequestNodePool) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


