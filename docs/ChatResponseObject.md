# ChatResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Usage** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewChatResponseObject

`func NewChatResponseObject() *ChatResponseObject`

NewChatResponseObject instantiates a new ChatResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatResponseObjectWithDefaults

`func NewChatResponseObjectWithDefaults() *ChatResponseObject`

NewChatResponseObjectWithDefaults instantiates a new ChatResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatResponseObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatResponseObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatResponseObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatResponseObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ChatResponseObject) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChatResponseObject) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChatResponseObject) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChatResponseObject) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChatResponseObject) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatResponseObject) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatResponseObject) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChatResponseObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ChatResponseObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatResponseObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatResponseObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChatResponseObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetModel

`func (o *ChatResponseObject) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatResponseObject) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatResponseObject) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatResponseObject) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOutput

`func (o *ChatResponseObject) GetOutput() []map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ChatResponseObject) GetOutputOk() (*[]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ChatResponseObject) SetOutput(v []map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ChatResponseObject) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetUsage

`func (o *ChatResponseObject) GetUsage() map[string]interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChatResponseObject) GetUsageOk() (*map[string]interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChatResponseObject) SetUsage(v map[string]interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChatResponseObject) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


