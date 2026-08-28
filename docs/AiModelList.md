# AiModelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AiModelInfo**](AiModelInfo.md) |  | [optional] 
**Models** | Pointer to [**[]AiModelInfo**](AiModelInfo.md) |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 

## Methods

### NewAiModelList

`func NewAiModelList() *AiModelList`

NewAiModelList instantiates a new AiModelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelListWithDefaults

`func NewAiModelListWithDefaults() *AiModelList`

NewAiModelListWithDefaults instantiates a new AiModelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AiModelList) GetData() []AiModelInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiModelList) GetDataOk() (*[]AiModelInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiModelList) SetData(v []AiModelInfo)`

SetData sets Data field to given value.

### HasData

`func (o *AiModelList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModels

`func (o *AiModelList) GetModels() []AiModelInfo`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AiModelList) GetModelsOk() (*[]AiModelInfo, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AiModelList) SetModels(v []AiModelInfo)`

SetModels sets Models field to given value.

### HasModels

`func (o *AiModelList) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetObject

`func (o *AiModelList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiModelList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiModelList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiModelList) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


