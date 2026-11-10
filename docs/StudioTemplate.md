# StudioTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Graph** | Pointer to [**StudioGraph**](StudioGraph.md) |  | [optional] 
**Id** | Pointer to **string** | ID is a catalog slug (image, poster) or an org template&#39;s id. | [optional] 
**Inputs** | Pointer to **map[string]interface{}** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Preview** | Pointer to **string** | Preview is an asset id or https URL of a still that shows what it makes. | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewStudioTemplate

`func NewStudioTemplate() *StudioTemplate`

NewStudioTemplate instantiates a new StudioTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioTemplateWithDefaults

`func NewStudioTemplateWithDefaults() *StudioTemplate`

NewStudioTemplateWithDefaults instantiates a new StudioTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *StudioTemplate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StudioTemplate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StudioTemplate) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StudioTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *StudioTemplate) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StudioTemplate) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StudioTemplate) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StudioTemplate) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *StudioTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StudioTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StudioTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StudioTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGraph

`func (o *StudioTemplate) GetGraph() StudioGraph`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *StudioTemplate) GetGraphOk() (*StudioGraph, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *StudioTemplate) SetGraph(v StudioGraph)`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *StudioTemplate) HasGraph() bool`

HasGraph returns a boolean if a field has been set.

### GetId

`func (o *StudioTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudioTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudioTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StudioTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputs

`func (o *StudioTemplate) GetInputs() map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *StudioTemplate) GetInputsOk() (*map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *StudioTemplate) SetInputs(v map[string]interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *StudioTemplate) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetKind

`func (o *StudioTemplate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StudioTemplate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StudioTemplate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StudioTemplate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrg

`func (o *StudioTemplate) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *StudioTemplate) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *StudioTemplate) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *StudioTemplate) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPreview

`func (o *StudioTemplate) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *StudioTemplate) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *StudioTemplate) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *StudioTemplate) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetShared

`func (o *StudioTemplate) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StudioTemplate) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StudioTemplate) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StudioTemplate) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetTitle

`func (o *StudioTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StudioTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StudioTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StudioTemplate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *StudioTemplate) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *StudioTemplate) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *StudioTemplate) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *StudioTemplate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


