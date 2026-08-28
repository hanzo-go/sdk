# OpenaiEmbeddingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]OpenaiEmbedding**](OpenaiEmbedding.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**OpenaiUsage**](OpenaiUsage.md) |  | [optional] 

## Methods

### NewOpenaiEmbeddingResponse

`func NewOpenaiEmbeddingResponse() *OpenaiEmbeddingResponse`

NewOpenaiEmbeddingResponse instantiates a new OpenaiEmbeddingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiEmbeddingResponseWithDefaults

`func NewOpenaiEmbeddingResponseWithDefaults() *OpenaiEmbeddingResponse`

NewOpenaiEmbeddingResponseWithDefaults instantiates a new OpenaiEmbeddingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OpenaiEmbeddingResponse) GetData() []OpenaiEmbedding`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpenaiEmbeddingResponse) GetDataOk() (*[]OpenaiEmbedding, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpenaiEmbeddingResponse) SetData(v []OpenaiEmbedding)`

SetData sets Data field to given value.

### HasData

`func (o *OpenaiEmbeddingResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModel

`func (o *OpenaiEmbeddingResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OpenaiEmbeddingResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OpenaiEmbeddingResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OpenaiEmbeddingResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObject

`func (o *OpenaiEmbeddingResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OpenaiEmbeddingResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OpenaiEmbeddingResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OpenaiEmbeddingResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUsage

`func (o *OpenaiEmbeddingResponse) GetUsage() OpenaiUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *OpenaiEmbeddingResponse) GetUsageOk() (*OpenaiUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *OpenaiEmbeddingResponse) SetUsage(v OpenaiUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *OpenaiEmbeddingResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


