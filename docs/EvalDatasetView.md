# EvalDatasetView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the name was first written, kept across later edits. | [optional] 
**Description** | Pointer to **string** | Description is free text the org wrote about what this set measures. | [optional] 
**Items** | Pointer to **int64** | Items is how many examples the set holds. It is filled only by the single read — a listing does not count, so it is absent there rather than zero. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata is the free-form object stored with the set, echoed back verbatim. | [optional] 
**Name** | Pointer to **string** | Name is the dataset&#39;s org-unique handle and the segment that addresses it. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when the description or metadata last changed. | [optional] 

## Methods

### NewEvalDatasetView

`func NewEvalDatasetView() *EvalDatasetView`

NewEvalDatasetView instantiates a new EvalDatasetView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalDatasetViewWithDefaults

`func NewEvalDatasetViewWithDefaults() *EvalDatasetView`

NewEvalDatasetViewWithDefaults instantiates a new EvalDatasetView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EvalDatasetView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvalDatasetView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvalDatasetView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EvalDatasetView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *EvalDatasetView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EvalDatasetView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EvalDatasetView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EvalDatasetView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItems

`func (o *EvalDatasetView) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EvalDatasetView) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EvalDatasetView) SetItems(v int64)`

SetItems sets Items field to given value.

### HasItems

`func (o *EvalDatasetView) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *EvalDatasetView) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvalDatasetView) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvalDatasetView) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EvalDatasetView) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *EvalDatasetView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalDatasetView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalDatasetView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EvalDatasetView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EvalDatasetView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvalDatasetView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvalDatasetView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EvalDatasetView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


