# EvalDatasetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is free text about what this set measures; over 64 KiB is refused. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata is a free-form object stored with the set and echoed back verbatim. | [optional] 
**Name** | **string** | Name is the dataset&#39;s org-unique handle and the segment that will address it, so it must match ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. | 

## Methods

### NewEvalDatasetReq

`func NewEvalDatasetReq(name string, ) *EvalDatasetReq`

NewEvalDatasetReq instantiates a new EvalDatasetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalDatasetReqWithDefaults

`func NewEvalDatasetReqWithDefaults() *EvalDatasetReq`

NewEvalDatasetReqWithDefaults instantiates a new EvalDatasetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EvalDatasetReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EvalDatasetReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EvalDatasetReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EvalDatasetReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *EvalDatasetReq) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvalDatasetReq) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvalDatasetReq) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EvalDatasetReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *EvalDatasetReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalDatasetReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalDatasetReq) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


