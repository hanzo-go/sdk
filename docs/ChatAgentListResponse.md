# ChatAgentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ChatAgent**](ChatAgent.md) |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**After** | Pointer to **string** |  | [optional] 

## Methods

### NewChatAgentListResponse

`func NewChatAgentListResponse() *ChatAgentListResponse`

NewChatAgentListResponse instantiates a new ChatAgentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatAgentListResponseWithDefaults

`func NewChatAgentListResponseWithDefaults() *ChatAgentListResponse`

NewChatAgentListResponseWithDefaults instantiates a new ChatAgentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChatAgentListResponse) GetData() []ChatAgent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChatAgentListResponse) GetDataOk() (*[]ChatAgent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChatAgentListResponse) SetData(v []ChatAgent)`

SetData sets Data field to given value.

### HasData

`func (o *ChatAgentListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasMore

`func (o *ChatAgentListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ChatAgentListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ChatAgentListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ChatAgentListResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetAfter

`func (o *ChatAgentListResponse) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ChatAgentListResponse) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ChatAgentListResponse) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ChatAgentListResponse) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


