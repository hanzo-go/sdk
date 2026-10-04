# ToolTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **bool** | Activated is filled by the registry from the activation store for the requesting (org,project); providers leave it zero. An unactivated tool is discoverable but refused 403 at dispatch. | [optional] 
**Description** | Pointer to **string** | Description is the prose a model reads to decide whether to call the tool. | [optional] 
**Dispatchable** | Pointer to **bool** | Dispatchable is whether the tool can be CALLED. False for a listing-only entry: a skill is activated and attached to an agent, never called. | [optional] 
**InputSchema** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** | Name is the tool&#39;s id in the flat, fleet-wide tool namespace — the value a tools/call passes. Unique across sources: a collision is resolved by source precedence before the caller ever sees it. | [optional] 
**Price** | Pointer to [**ToolPrice**](ToolPrice.md) | Price is what a call costs and who is paid, absent for a free tool. Enforcement is the x402 settlement client; this is the declaration. | [optional] 
**Source** | Pointer to **string** | Source is where the tool comes from: connector, function, zap-service, agent, skill or mcp. | [optional] 

## Methods

### NewToolTool

`func NewToolTool() *ToolTool`

NewToolTool instantiates a new ToolTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolToolWithDefaults

`func NewToolToolWithDefaults() *ToolTool`

NewToolToolWithDefaults instantiates a new ToolTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *ToolTool) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *ToolTool) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *ToolTool) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *ToolTool) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetDescription

`func (o *ToolTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDispatchable

`func (o *ToolTool) GetDispatchable() bool`

GetDispatchable returns the Dispatchable field if non-nil, zero value otherwise.

### GetDispatchableOk

`func (o *ToolTool) GetDispatchableOk() (*bool, bool)`

GetDispatchableOk returns a tuple with the Dispatchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchable

`func (o *ToolTool) SetDispatchable(v bool)`

SetDispatchable sets Dispatchable field to given value.

### HasDispatchable

`func (o *ToolTool) HasDispatchable() bool`

HasDispatchable returns a boolean if a field has been set.

### GetInputSchema

`func (o *ToolTool) GetInputSchema() interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *ToolTool) GetInputSchemaOk() (*interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *ToolTool) SetInputSchema(v interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *ToolTool) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### SetInputSchemaNil

`func (o *ToolTool) SetInputSchemaNil(b bool)`

 SetInputSchemaNil sets the value for InputSchema to be an explicit nil

### UnsetInputSchema
`func (o *ToolTool) UnsetInputSchema()`

UnsetInputSchema ensures that no value is present for InputSchema, not even an explicit nil
### GetName

`func (o *ToolTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolTool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolTool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *ToolTool) GetPrice() ToolPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ToolTool) GetPriceOk() (*ToolPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ToolTool) SetPrice(v ToolPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ToolTool) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSource

`func (o *ToolTool) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ToolTool) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ToolTool) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ToolTool) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


