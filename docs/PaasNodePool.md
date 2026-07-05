# PaasNodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** | Droplet size slug (e.g. s-4vcpu-8gb) | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**AutoScale** | Pointer to **bool** |  | [optional] 
**MinNodes** | Pointer to **int32** |  | [optional] 
**MaxNodes** | Pointer to **int32** |  | [optional] 
**Nodes** | Pointer to [**[]PaasNodePoolNodesInner**](PaasNodePoolNodesInner.md) |  | [optional] 

## Methods

### NewPaasNodePool

`func NewPaasNodePool() *PaasNodePool`

NewPaasNodePool instantiates a new PaasNodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasNodePoolWithDefaults

`func NewPaasNodePoolWithDefaults() *PaasNodePool`

NewPaasNodePoolWithDefaults instantiates a new PaasNodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaasNodePool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaasNodePool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaasNodePool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaasNodePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PaasNodePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaasNodePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaasNodePool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaasNodePool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *PaasNodePool) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PaasNodePool) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PaasNodePool) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *PaasNodePool) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCount

`func (o *PaasNodePool) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaasNodePool) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaasNodePool) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaasNodePool) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAutoScale

`func (o *PaasNodePool) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *PaasNodePool) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *PaasNodePool) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *PaasNodePool) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetMinNodes

`func (o *PaasNodePool) GetMinNodes() int32`

GetMinNodes returns the MinNodes field if non-nil, zero value otherwise.

### GetMinNodesOk

`func (o *PaasNodePool) GetMinNodesOk() (*int32, bool)`

GetMinNodesOk returns a tuple with the MinNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodes

`func (o *PaasNodePool) SetMinNodes(v int32)`

SetMinNodes sets MinNodes field to given value.

### HasMinNodes

`func (o *PaasNodePool) HasMinNodes() bool`

HasMinNodes returns a boolean if a field has been set.

### GetMaxNodes

`func (o *PaasNodePool) GetMaxNodes() int32`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *PaasNodePool) GetMaxNodesOk() (*int32, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *PaasNodePool) SetMaxNodes(v int32)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *PaasNodePool) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetNodes

`func (o *PaasNodePool) GetNodes() []PaasNodePoolNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *PaasNodePool) GetNodesOk() (*[]PaasNodePoolNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *PaasNodePool) SetNodes(v []PaasNodePoolNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *PaasNodePool) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


