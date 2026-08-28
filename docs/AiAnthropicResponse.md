# AiAnthropicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]AiAnthropicContentBlock**](AiAnthropicContentBlock.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**StopReason** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**AiAnthropicUsage**](AiAnthropicUsage.md) |  | [optional] 

## Methods

### NewAiAnthropicResponse

`func NewAiAnthropicResponse() *AiAnthropicResponse`

NewAiAnthropicResponse instantiates a new AiAnthropicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAnthropicResponseWithDefaults

`func NewAiAnthropicResponseWithDefaults() *AiAnthropicResponse`

NewAiAnthropicResponseWithDefaults instantiates a new AiAnthropicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *AiAnthropicResponse) GetContent() []AiAnthropicContentBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiAnthropicResponse) GetContentOk() (*[]AiAnthropicContentBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiAnthropicResponse) SetContent(v []AiAnthropicContentBlock)`

SetContent sets Content field to given value.

### HasContent

`func (o *AiAnthropicResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetId

`func (o *AiAnthropicResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiAnthropicResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiAnthropicResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiAnthropicResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *AiAnthropicResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiAnthropicResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiAnthropicResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiAnthropicResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRole

`func (o *AiAnthropicResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AiAnthropicResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AiAnthropicResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AiAnthropicResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStopReason

`func (o *AiAnthropicResponse) GetStopReason() string`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *AiAnthropicResponse) GetStopReasonOk() (*string, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *AiAnthropicResponse) SetStopReason(v string)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *AiAnthropicResponse) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### GetType

`func (o *AiAnthropicResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiAnthropicResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiAnthropicResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AiAnthropicResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsage

`func (o *AiAnthropicResponse) GetUsage() AiAnthropicUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiAnthropicResponse) GetUsageOk() (*AiAnthropicUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiAnthropicResponse) SetUsage(v AiAnthropicUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiAnthropicResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


