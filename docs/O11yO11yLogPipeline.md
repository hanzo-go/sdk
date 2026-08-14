# O11yO11yLogPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias is the pipeline&#39;s short name. | [optional] 
**Config** | Pointer to [**[]O11yO11yLogPipelineOperator**](O11yO11yLogPipelineOperator.md) | Config is the pipeline&#39;s processors, in order. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the pipeline was created. | [optional] 
**CreatedBy** | Pointer to **string** | CreatedBy is the id of who created the pipeline. | [optional] 
**Description** | Pointer to **string** | Description says what the pipeline is for. | [optional] 
**Enabled** | Pointer to **bool** | Enabled says whether the pipeline is on. | [optional] 
**Filter** | Pointer to [**O11yO11yLogFilter**](O11yO11yLogFilter.md) | Filter selects which records the pipeline processes. | [optional] 
**Id** | Pointer to **string** | ID is the pipeline&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the pipeline&#39;s display name. | [optional] 
**OrderId** | Pointer to **int32** | OrderID is the pipeline&#39;s 1-based position in the set. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the pipeline last changed. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is the id of who last changed it. | [optional] 

## Methods

### NewO11yO11yLogPipeline

`func NewO11yO11yLogPipeline() *O11yO11yLogPipeline`

NewO11yO11yLogPipeline instantiates a new O11yO11yLogPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogPipelineWithDefaults

`func NewO11yO11yLogPipelineWithDefaults() *O11yO11yLogPipeline`

NewO11yO11yLogPipelineWithDefaults instantiates a new O11yO11yLogPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *O11yO11yLogPipeline) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *O11yO11yLogPipeline) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *O11yO11yLogPipeline) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *O11yO11yLogPipeline) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetConfig

`func (o *O11yO11yLogPipeline) GetConfig() []O11yO11yLogPipelineOperator`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yO11yLogPipeline) GetConfigOk() (*[]O11yO11yLogPipelineOperator, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yO11yLogPipeline) SetConfig(v []O11yO11yLogPipelineOperator)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yO11yLogPipeline) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yLogPipeline) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yLogPipeline) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yLogPipeline) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yLogPipeline) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yLogPipeline) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yLogPipeline) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yLogPipeline) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yLogPipeline) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *O11yO11yLogPipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yLogPipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yLogPipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yLogPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *O11yO11yLogPipeline) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yO11yLogPipeline) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yO11yLogPipeline) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yO11yLogPipeline) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFilter

`func (o *O11yO11yLogPipeline) GetFilter() O11yO11yLogFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *O11yO11yLogPipeline) GetFilterOk() (*O11yO11yLogFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *O11yO11yLogPipeline) SetFilter(v O11yO11yLogFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *O11yO11yLogPipeline) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLogPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLogPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLogPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLogPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yLogPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yLogPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yLogPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yLogPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderId

`func (o *O11yO11yLogPipeline) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *O11yO11yLogPipeline) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *O11yO11yLogPipeline) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *O11yO11yLogPipeline) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yLogPipeline) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yLogPipeline) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yLogPipeline) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yLogPipeline) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11yLogPipeline) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11yLogPipeline) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11yLogPipeline) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11yLogPipeline) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


