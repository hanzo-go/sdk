# O11yO11yChannelUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | Pointer to [**O11yReceiver**](O11yReceiver.md) |  | [optional] 
**GooglechatConfigs** | Pointer to [**[]O11yGoogleChatReceiverConfig**](O11yGoogleChatReceiverConfig.md) |  | [optional] 

## Methods

### NewO11yO11yChannelUpdateIn

`func NewO11yO11yChannelUpdateIn() *O11yO11yChannelUpdateIn`

NewO11yO11yChannelUpdateIn instantiates a new O11yO11yChannelUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yChannelUpdateInWithDefaults

`func NewO11yO11yChannelUpdateInWithDefaults() *O11yO11yChannelUpdateIn`

NewO11yO11yChannelUpdateInWithDefaults instantiates a new O11yO11yChannelUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *O11yO11yChannelUpdateIn) GetReceiver() O11yReceiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *O11yO11yChannelUpdateIn) GetReceiverOk() (*O11yReceiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *O11yO11yChannelUpdateIn) SetReceiver(v O11yReceiver)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *O11yO11yChannelUpdateIn) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetGooglechatConfigs

`func (o *O11yO11yChannelUpdateIn) GetGooglechatConfigs() []O11yGoogleChatReceiverConfig`

GetGooglechatConfigs returns the GooglechatConfigs field if non-nil, zero value otherwise.

### GetGooglechatConfigsOk

`func (o *O11yO11yChannelUpdateIn) GetGooglechatConfigsOk() (*[]O11yGoogleChatReceiverConfig, bool)`

GetGooglechatConfigsOk returns a tuple with the GooglechatConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglechatConfigs

`func (o *O11yO11yChannelUpdateIn) SetGooglechatConfigs(v []O11yGoogleChatReceiverConfig)`

SetGooglechatConfigs sets GooglechatConfigs field to given value.

### HasGooglechatConfigs

`func (o *O11yO11yChannelUpdateIn) HasGooglechatConfigs() bool`

HasGooglechatConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


