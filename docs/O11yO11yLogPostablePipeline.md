# O11yO11yLogPostablePipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias is the pipeline&#39;s short name. | [optional] 
**Config** | Pointer to [**[]O11yO11yLogPipelineOperator**](O11yO11yLogPipelineOperator.md) | Config is the pipeline&#39;s processors, in order. | [optional] 
**Description** | Pointer to **string** | Description says what the pipeline is for. | [optional] 
**Enabled** | Pointer to **bool** | Enabled turns the pipeline on. | [optional] 
**Filter** | Pointer to [**O11yO11yLogFilter**](O11yO11yLogFilter.md) | Filter selects which records the pipeline processes. | [optional] 
**Id** | Pointer to **string** | ID is the pipeline&#39;s id. Empty on a new pipeline; the id it was listed with to keep an existing one. | [optional] 
**Name** | Pointer to **string** | Name is the pipeline&#39;s display name. | [optional] 
**OrderId** | Pointer to **int64** | OrderID is the pipeline&#39;s 1-based position in the set. | [optional] 

## Methods

### NewO11yO11yLogPostablePipeline

`func NewO11yO11yLogPostablePipeline() *O11yO11yLogPostablePipeline`

NewO11yO11yLogPostablePipeline instantiates a new O11yO11yLogPostablePipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogPostablePipelineWithDefaults

`func NewO11yO11yLogPostablePipelineWithDefaults() *O11yO11yLogPostablePipeline`

NewO11yO11yLogPostablePipelineWithDefaults instantiates a new O11yO11yLogPostablePipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *O11yO11yLogPostablePipeline) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *O11yO11yLogPostablePipeline) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *O11yO11yLogPostablePipeline) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *O11yO11yLogPostablePipeline) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetConfig

`func (o *O11yO11yLogPostablePipeline) GetConfig() []O11yO11yLogPipelineOperator`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yO11yLogPostablePipeline) GetConfigOk() (*[]O11yO11yLogPipelineOperator, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yO11yLogPostablePipeline) SetConfig(v []O11yO11yLogPipelineOperator)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yO11yLogPostablePipeline) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *O11yO11yLogPostablePipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yLogPostablePipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yLogPostablePipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yLogPostablePipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *O11yO11yLogPostablePipeline) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yO11yLogPostablePipeline) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yO11yLogPostablePipeline) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yO11yLogPostablePipeline) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFilter

`func (o *O11yO11yLogPostablePipeline) GetFilter() O11yO11yLogFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *O11yO11yLogPostablePipeline) GetFilterOk() (*O11yO11yLogFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *O11yO11yLogPostablePipeline) SetFilter(v O11yO11yLogFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *O11yO11yLogPostablePipeline) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLogPostablePipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLogPostablePipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLogPostablePipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLogPostablePipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yLogPostablePipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yLogPostablePipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yLogPostablePipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yLogPostablePipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderId

`func (o *O11yO11yLogPostablePipeline) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *O11yO11yLogPostablePipeline) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *O11yO11yLogPostablePipeline) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *O11yO11yLogPostablePipeline) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


