# OpenaiChatMessagePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageUrl** | Pointer to [**OpenaiChatMessageImageURL**](OpenaiChatMessageImageURL.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenaiChatMessagePart

`func NewOpenaiChatMessagePart() *OpenaiChatMessagePart`

NewOpenaiChatMessagePart instantiates a new OpenaiChatMessagePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiChatMessagePartWithDefaults

`func NewOpenaiChatMessagePartWithDefaults() *OpenaiChatMessagePart`

NewOpenaiChatMessagePartWithDefaults instantiates a new OpenaiChatMessagePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageUrl

`func (o *OpenaiChatMessagePart) GetImageUrl() OpenaiChatMessageImageURL`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *OpenaiChatMessagePart) GetImageUrlOk() (*OpenaiChatMessageImageURL, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *OpenaiChatMessagePart) SetImageUrl(v OpenaiChatMessageImageURL)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *OpenaiChatMessagePart) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetText

`func (o *OpenaiChatMessagePart) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OpenaiChatMessagePart) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OpenaiChatMessagePart) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *OpenaiChatMessagePart) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *OpenaiChatMessagePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenaiChatMessagePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenaiChatMessagePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenaiChatMessagePart) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


