# AiMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StopReason** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAiMessageResponse

`func NewAiMessageResponse() *AiMessageResponse`

NewAiMessageResponse instantiates a new AiMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMessageResponseWithDefaults

`func NewAiMessageResponseWithDefaults() *AiMessageResponse`

NewAiMessageResponseWithDefaults instantiates a new AiMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiMessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiMessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiMessageResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiMessageResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AiMessageResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiMessageResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiMessageResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AiMessageResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRole

`func (o *AiMessageResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AiMessageResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AiMessageResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AiMessageResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetModel

`func (o *AiMessageResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiMessageResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiMessageResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiMessageResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetContent

`func (o *AiMessageResponse) GetContent() []map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiMessageResponse) GetContentOk() (*[]map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiMessageResponse) SetContent(v []map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *AiMessageResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStopReason

`func (o *AiMessageResponse) GetStopReason() string`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *AiMessageResponse) GetStopReasonOk() (*string, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *AiMessageResponse) SetStopReason(v string)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *AiMessageResponse) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### GetUsage

`func (o *AiMessageResponse) GetUsage() map[string]interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiMessageResponse) GetUsageOk() (*map[string]interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiMessageResponse) SetUsage(v map[string]interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiMessageResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


