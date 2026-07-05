# VectorCollectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | Pointer to [**VectorCollectionConfigParams**](VectorCollectionConfigParams.md) |  | [optional] 
**HnswConfig** | Pointer to [**VectorHnswConfig**](VectorHnswConfig.md) |  | [optional] 
**OptimizerConfig** | Pointer to [**VectorOptimizerConfig**](VectorOptimizerConfig.md) |  | [optional] 
**WalConfig** | Pointer to [**VectorWalConfig**](VectorWalConfig.md) |  | [optional] 

## Methods

### NewVectorCollectionConfig

`func NewVectorCollectionConfig() *VectorCollectionConfig`

NewVectorCollectionConfig instantiates a new VectorCollectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorCollectionConfigWithDefaults

`func NewVectorCollectionConfigWithDefaults() *VectorCollectionConfig`

NewVectorCollectionConfigWithDefaults instantiates a new VectorCollectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *VectorCollectionConfig) GetParams() VectorCollectionConfigParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *VectorCollectionConfig) GetParamsOk() (*VectorCollectionConfigParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *VectorCollectionConfig) SetParams(v VectorCollectionConfigParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *VectorCollectionConfig) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetHnswConfig

`func (o *VectorCollectionConfig) GetHnswConfig() VectorHnswConfig`

GetHnswConfig returns the HnswConfig field if non-nil, zero value otherwise.

### GetHnswConfigOk

`func (o *VectorCollectionConfig) GetHnswConfigOk() (*VectorHnswConfig, bool)`

GetHnswConfigOk returns a tuple with the HnswConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHnswConfig

`func (o *VectorCollectionConfig) SetHnswConfig(v VectorHnswConfig)`

SetHnswConfig sets HnswConfig field to given value.

### HasHnswConfig

`func (o *VectorCollectionConfig) HasHnswConfig() bool`

HasHnswConfig returns a boolean if a field has been set.

### GetOptimizerConfig

`func (o *VectorCollectionConfig) GetOptimizerConfig() VectorOptimizerConfig`

GetOptimizerConfig returns the OptimizerConfig field if non-nil, zero value otherwise.

### GetOptimizerConfigOk

`func (o *VectorCollectionConfig) GetOptimizerConfigOk() (*VectorOptimizerConfig, bool)`

GetOptimizerConfigOk returns a tuple with the OptimizerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizerConfig

`func (o *VectorCollectionConfig) SetOptimizerConfig(v VectorOptimizerConfig)`

SetOptimizerConfig sets OptimizerConfig field to given value.

### HasOptimizerConfig

`func (o *VectorCollectionConfig) HasOptimizerConfig() bool`

HasOptimizerConfig returns a boolean if a field has been set.

### GetWalConfig

`func (o *VectorCollectionConfig) GetWalConfig() VectorWalConfig`

GetWalConfig returns the WalConfig field if non-nil, zero value otherwise.

### GetWalConfigOk

`func (o *VectorCollectionConfig) GetWalConfigOk() (*VectorWalConfig, bool)`

GetWalConfigOk returns a tuple with the WalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalConfig

`func (o *VectorCollectionConfig) SetWalConfig(v VectorWalConfig)`

SetWalConfig sets WalConfig field to given value.

### HasWalConfig

`func (o *VectorCollectionConfig) HasWalConfig() bool`

HasWalConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


