# AiModels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to [**[]AiModel**](AiModel.md) | Models are the matches, in the catalog&#39;s order. | [optional] 
**Total** | Pointer to **int64** | Total is how many models matched, before Limit. | [optional] 

## Methods

### NewAiModels

`func NewAiModels() *AiModels`

NewAiModels instantiates a new AiModels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelsWithDefaults

`func NewAiModelsWithDefaults() *AiModels`

NewAiModelsWithDefaults instantiates a new AiModels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *AiModels) GetModels() []AiModel`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AiModels) GetModelsOk() (*[]AiModel, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AiModels) SetModels(v []AiModel)`

SetModels sets Models field to given value.

### HasModels

`func (o *AiModels) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetTotal

`func (o *AiModels) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AiModels) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AiModels) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AiModels) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


