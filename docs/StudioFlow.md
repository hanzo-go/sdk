# StudioFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Graph** | Pointer to [**StudioGraph**](StudioGraph.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to **map[string]interface{}** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewStudioFlow

`func NewStudioFlow() *StudioFlow`

NewStudioFlow instantiates a new StudioFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioFlowWithDefaults

`func NewStudioFlowWithDefaults() *StudioFlow`

NewStudioFlowWithDefaults instantiates a new StudioFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *StudioFlow) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StudioFlow) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StudioFlow) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StudioFlow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *StudioFlow) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StudioFlow) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StudioFlow) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StudioFlow) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *StudioFlow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StudioFlow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StudioFlow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StudioFlow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGraph

`func (o *StudioFlow) GetGraph() StudioGraph`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *StudioFlow) GetGraphOk() (*StudioGraph, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *StudioFlow) SetGraph(v StudioGraph)`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *StudioFlow) HasGraph() bool`

HasGraph returns a boolean if a field has been set.

### GetId

`func (o *StudioFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudioFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudioFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StudioFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputs

`func (o *StudioFlow) GetInputs() map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *StudioFlow) GetInputsOk() (*map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *StudioFlow) SetInputs(v map[string]interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *StudioFlow) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetKind

`func (o *StudioFlow) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StudioFlow) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StudioFlow) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StudioFlow) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *StudioFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudioFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudioFlow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StudioFlow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *StudioFlow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *StudioFlow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *StudioFlow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *StudioFlow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetShared

`func (o *StudioFlow) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StudioFlow) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StudioFlow) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StudioFlow) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetUpdated

`func (o *StudioFlow) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *StudioFlow) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *StudioFlow) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *StudioFlow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


