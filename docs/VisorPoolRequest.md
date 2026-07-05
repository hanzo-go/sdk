# VisorPoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**MinNodes** | Pointer to **int32** |  | [optional] 
**MaxNodes** | Pointer to **int32** |  | [optional] 
**AutoScale** | Pointer to **bool** |  | [optional] 

## Methods

### NewVisorPoolRequest

`func NewVisorPoolRequest(provider string, ) *VisorPoolRequest`

NewVisorPoolRequest instantiates a new VisorPoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisorPoolRequestWithDefaults

`func NewVisorPoolRequestWithDefaults() *VisorPoolRequest`

NewVisorPoolRequestWithDefaults instantiates a new VisorPoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VisorPoolRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VisorPoolRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VisorPoolRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetName

`func (o *VisorPoolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisorPoolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisorPoolRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VisorPoolRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *VisorPoolRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VisorPoolRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VisorPoolRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VisorPoolRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCount

`func (o *VisorPoolRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VisorPoolRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VisorPoolRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VisorPoolRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMinNodes

`func (o *VisorPoolRequest) GetMinNodes() int32`

GetMinNodes returns the MinNodes field if non-nil, zero value otherwise.

### GetMinNodesOk

`func (o *VisorPoolRequest) GetMinNodesOk() (*int32, bool)`

GetMinNodesOk returns a tuple with the MinNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodes

`func (o *VisorPoolRequest) SetMinNodes(v int32)`

SetMinNodes sets MinNodes field to given value.

### HasMinNodes

`func (o *VisorPoolRequest) HasMinNodes() bool`

HasMinNodes returns a boolean if a field has been set.

### GetMaxNodes

`func (o *VisorPoolRequest) GetMaxNodes() int32`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *VisorPoolRequest) GetMaxNodesOk() (*int32, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *VisorPoolRequest) SetMaxNodes(v int32)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *VisorPoolRequest) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetAutoScale

`func (o *VisorPoolRequest) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *VisorPoolRequest) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *VisorPoolRequest) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *VisorPoolRequest) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


