# CloudProductControllerChatDocsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]CloudProductControllerChatDocsRequestMessagesInner**](CloudProductControllerChatDocsRequestMessagesInner.md) |  | 
**Index** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudProductControllerChatDocsRequest

`func NewCloudProductControllerChatDocsRequest(messages []CloudProductControllerChatDocsRequestMessagesInner, ) *CloudProductControllerChatDocsRequest`

NewCloudProductControllerChatDocsRequest instantiates a new CloudProductControllerChatDocsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProductControllerChatDocsRequestWithDefaults

`func NewCloudProductControllerChatDocsRequestWithDefaults() *CloudProductControllerChatDocsRequest`

NewCloudProductControllerChatDocsRequestWithDefaults instantiates a new CloudProductControllerChatDocsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *CloudProductControllerChatDocsRequest) GetMessages() []CloudProductControllerChatDocsRequestMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CloudProductControllerChatDocsRequest) GetMessagesOk() (*[]CloudProductControllerChatDocsRequestMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CloudProductControllerChatDocsRequest) SetMessages(v []CloudProductControllerChatDocsRequestMessagesInner)`

SetMessages sets Messages field to given value.


### GetIndex

`func (o *CloudProductControllerChatDocsRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CloudProductControllerChatDocsRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CloudProductControllerChatDocsRequest) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CloudProductControllerChatDocsRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetStream

`func (o *CloudProductControllerChatDocsRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CloudProductControllerChatDocsRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CloudProductControllerChatDocsRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CloudProductControllerChatDocsRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


