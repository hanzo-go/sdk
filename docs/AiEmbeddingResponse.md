# AiEmbeddingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**AiUsage**](AiUsage.md) |  | [optional] 

## Methods

### NewAiEmbeddingResponse

`func NewAiEmbeddingResponse() *AiEmbeddingResponse`

NewAiEmbeddingResponse instantiates a new AiEmbeddingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEmbeddingResponseWithDefaults

`func NewAiEmbeddingResponseWithDefaults() *AiEmbeddingResponse`

NewAiEmbeddingResponseWithDefaults instantiates a new AiEmbeddingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *AiEmbeddingResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiEmbeddingResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiEmbeddingResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiEmbeddingResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *AiEmbeddingResponse) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiEmbeddingResponse) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiEmbeddingResponse) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AiEmbeddingResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModel

`func (o *AiEmbeddingResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiEmbeddingResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiEmbeddingResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiEmbeddingResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUsage

`func (o *AiEmbeddingResponse) GetUsage() AiUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiEmbeddingResponse) GetUsageOk() (*AiUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiEmbeddingResponse) SetUsage(v AiUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiEmbeddingResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


