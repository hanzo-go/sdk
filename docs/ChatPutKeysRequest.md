# ChatPutKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewChatPutKeysRequest

`func NewChatPutKeysRequest(name string, value string, expiresAt time.Time, ) *ChatPutKeysRequest`

NewChatPutKeysRequest instantiates a new ChatPutKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPutKeysRequestWithDefaults

`func NewChatPutKeysRequestWithDefaults() *ChatPutKeysRequest`

NewChatPutKeysRequestWithDefaults instantiates a new ChatPutKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChatPutKeysRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatPutKeysRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatPutKeysRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ChatPutKeysRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ChatPutKeysRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ChatPutKeysRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetExpiresAt

`func (o *ChatPutKeysRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChatPutKeysRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChatPutKeysRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


