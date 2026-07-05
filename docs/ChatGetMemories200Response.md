# ChatGetMemories200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memories** | Pointer to [**[]ChatMemory**](ChatMemory.md) |  | [optional] 
**TotalTokens** | Pointer to **int32** |  | [optional] 
**TokenLimit** | Pointer to **int32** |  | [optional] 
**CharLimit** | Pointer to **int32** |  | [optional] 
**UsagePercentage** | Pointer to **int32** |  | [optional] 

## Methods

### NewChatGetMemories200Response

`func NewChatGetMemories200Response() *ChatGetMemories200Response`

NewChatGetMemories200Response instantiates a new ChatGetMemories200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatGetMemories200ResponseWithDefaults

`func NewChatGetMemories200ResponseWithDefaults() *ChatGetMemories200Response`

NewChatGetMemories200ResponseWithDefaults instantiates a new ChatGetMemories200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemories

`func (o *ChatGetMemories200Response) GetMemories() []ChatMemory`

GetMemories returns the Memories field if non-nil, zero value otherwise.

### GetMemoriesOk

`func (o *ChatGetMemories200Response) GetMemoriesOk() (*[]ChatMemory, bool)`

GetMemoriesOk returns a tuple with the Memories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemories

`func (o *ChatGetMemories200Response) SetMemories(v []ChatMemory)`

SetMemories sets Memories field to given value.

### HasMemories

`func (o *ChatGetMemories200Response) HasMemories() bool`

HasMemories returns a boolean if a field has been set.

### GetTotalTokens

`func (o *ChatGetMemories200Response) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *ChatGetMemories200Response) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *ChatGetMemories200Response) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *ChatGetMemories200Response) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetTokenLimit

`func (o *ChatGetMemories200Response) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *ChatGetMemories200Response) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *ChatGetMemories200Response) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *ChatGetMemories200Response) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### GetCharLimit

`func (o *ChatGetMemories200Response) GetCharLimit() int32`

GetCharLimit returns the CharLimit field if non-nil, zero value otherwise.

### GetCharLimitOk

`func (o *ChatGetMemories200Response) GetCharLimitOk() (*int32, bool)`

GetCharLimitOk returns a tuple with the CharLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharLimit

`func (o *ChatGetMemories200Response) SetCharLimit(v int32)`

SetCharLimit sets CharLimit field to given value.

### HasCharLimit

`func (o *ChatGetMemories200Response) HasCharLimit() bool`

HasCharLimit returns a boolean if a field has been set.

### GetUsagePercentage

`func (o *ChatGetMemories200Response) GetUsagePercentage() int32`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *ChatGetMemories200Response) GetUsagePercentageOk() (*int32, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *ChatGetMemories200Response) SetUsagePercentage(v int32)`

SetUsagePercentage sets UsagePercentage field to given value.

### HasUsagePercentage

`func (o *ChatGetMemories200Response) HasUsagePercentage() bool`

HasUsagePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


