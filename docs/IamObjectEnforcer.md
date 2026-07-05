# IamObjectEnforcer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adapter** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelCfg** | Pointer to  |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectEnforcer

`func NewIamObjectEnforcer() *IamObjectEnforcer`

NewIamObjectEnforcer instantiates a new IamObjectEnforcer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectEnforcerWithDefaults

`func NewIamObjectEnforcerWithDefaults() *IamObjectEnforcer`

NewIamObjectEnforcerWithDefaults instantiates a new IamObjectEnforcer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapter

`func (o *IamObjectEnforcer) GetAdapter() string`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *IamObjectEnforcer) GetAdapterOk() (*string, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *IamObjectEnforcer) SetAdapter(v string)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *IamObjectEnforcer) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectEnforcer) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectEnforcer) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectEnforcer) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectEnforcer) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *IamObjectEnforcer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamObjectEnforcer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamObjectEnforcer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamObjectEnforcer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectEnforcer) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectEnforcer) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectEnforcer) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectEnforcer) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetModel

`func (o *IamObjectEnforcer) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IamObjectEnforcer) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IamObjectEnforcer) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *IamObjectEnforcer) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelCfg

`func (o *IamObjectEnforcer) GetModelCfg() map[string]string`

GetModelCfg returns the ModelCfg field if non-nil, zero value otherwise.

### GetModelCfgOk

`func (o *IamObjectEnforcer) GetModelCfgOk() (*map[string]string, bool)`

GetModelCfgOk returns a tuple with the ModelCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCfg

`func (o *IamObjectEnforcer) SetModelCfg(v map[string]string)`

SetModelCfg sets ModelCfg field to given value.

### HasModelCfg

`func (o *IamObjectEnforcer) HasModelCfg() bool`

HasModelCfg returns a boolean if a field has been set.

### SetModelCfgNil

`func (o *IamObjectEnforcer) SetModelCfgNil(b bool)`

 SetModelCfgNil sets the value for ModelCfg to be an explicit nil

### UnsetModelCfg
`func (o *IamObjectEnforcer) UnsetModelCfg()`

UnsetModelCfg ensures that no value is present for ModelCfg, not even an explicit nil
### GetName

`func (o *IamObjectEnforcer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectEnforcer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectEnforcer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectEnforcer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectEnforcer) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectEnforcer) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectEnforcer) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectEnforcer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *IamObjectEnforcer) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *IamObjectEnforcer) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *IamObjectEnforcer) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *IamObjectEnforcer) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


