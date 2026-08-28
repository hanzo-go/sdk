# OpenaiUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int32** |  | [optional] 
**CompletionTokensDetails** | Pointer to [**OpenaiCompletionTokensDetails**](OpenaiCompletionTokensDetails.md) |  | [optional] 
**PromptTokens** | Pointer to **int32** |  | [optional] 
**PromptTokensDetails** | Pointer to [**OpenaiPromptTokensDetails**](OpenaiPromptTokensDetails.md) |  | [optional] 
**TotalTokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewOpenaiUsage

`func NewOpenaiUsage() *OpenaiUsage`

NewOpenaiUsage instantiates a new OpenaiUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiUsageWithDefaults

`func NewOpenaiUsageWithDefaults() *OpenaiUsage`

NewOpenaiUsageWithDefaults instantiates a new OpenaiUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *OpenaiUsage) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *OpenaiUsage) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *OpenaiUsage) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *OpenaiUsage) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCompletionTokensDetails

`func (o *OpenaiUsage) GetCompletionTokensDetails() OpenaiCompletionTokensDetails`

GetCompletionTokensDetails returns the CompletionTokensDetails field if non-nil, zero value otherwise.

### GetCompletionTokensDetailsOk

`func (o *OpenaiUsage) GetCompletionTokensDetailsOk() (*OpenaiCompletionTokensDetails, bool)`

GetCompletionTokensDetailsOk returns a tuple with the CompletionTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokensDetails

`func (o *OpenaiUsage) SetCompletionTokensDetails(v OpenaiCompletionTokensDetails)`

SetCompletionTokensDetails sets CompletionTokensDetails field to given value.

### HasCompletionTokensDetails

`func (o *OpenaiUsage) HasCompletionTokensDetails() bool`

HasCompletionTokensDetails returns a boolean if a field has been set.

### GetPromptTokens

`func (o *OpenaiUsage) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *OpenaiUsage) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *OpenaiUsage) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *OpenaiUsage) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetPromptTokensDetails

`func (o *OpenaiUsage) GetPromptTokensDetails() OpenaiPromptTokensDetails`

GetPromptTokensDetails returns the PromptTokensDetails field if non-nil, zero value otherwise.

### GetPromptTokensDetailsOk

`func (o *OpenaiUsage) GetPromptTokensDetailsOk() (*OpenaiPromptTokensDetails, bool)`

GetPromptTokensDetailsOk returns a tuple with the PromptTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokensDetails

`func (o *OpenaiUsage) SetPromptTokensDetails(v OpenaiPromptTokensDetails)`

SetPromptTokensDetails sets PromptTokensDetails field to given value.

### HasPromptTokensDetails

`func (o *OpenaiUsage) HasPromptTokensDetails() bool`

HasPromptTokensDetails returns a boolean if a field has been set.

### GetTotalTokens

`func (o *OpenaiUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *OpenaiUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *OpenaiUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *OpenaiUsage) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


