# FlowTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Pieces** | Pointer to **[]string** |  | [optional] 
**Template** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlowTemplate

`func NewFlowTemplate() *FlowTemplate`

NewFlowTemplate instantiates a new FlowTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowTemplateWithDefaults

`func NewFlowTemplateWithDefaults() *FlowTemplate`

NewFlowTemplateWithDefaults instantiates a new FlowTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlowTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FlowTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlowTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlowTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlowTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPieces

`func (o *FlowTemplate) GetPieces() []string`

GetPieces returns the Pieces field if non-nil, zero value otherwise.

### GetPiecesOk

`func (o *FlowTemplate) GetPiecesOk() (*[]string, bool)`

GetPiecesOk returns a tuple with the Pieces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieces

`func (o *FlowTemplate) SetPieces(v []string)`

SetPieces sets Pieces field to given value.

### HasPieces

`func (o *FlowTemplate) HasPieces() bool`

HasPieces returns a boolean if a field has been set.

### GetTemplate

`func (o *FlowTemplate) GetTemplate() map[string]interface{}`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *FlowTemplate) GetTemplateOk() (*map[string]interface{}, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *FlowTemplate) SetTemplate(v map[string]interface{})`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *FlowTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetCreated

`func (o *FlowTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *FlowTemplate) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FlowTemplate) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FlowTemplate) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FlowTemplate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


