# ProviderSlackSendMessageIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is a Slack channel ID, name (#hanzo-gtm), or Slack channel mention. | [optional] 
**Text** | Pointer to **string** | Text is the message to post as Hanzo, up to 40000 characters. | [optional] 
**ThreadTs** | Pointer to **string** | ThreadTS optionally replies under this parent message&#39;s exact ts. | [optional] 

## Methods

### NewProviderSlackSendMessageIn

`func NewProviderSlackSendMessageIn() *ProviderSlackSendMessageIn`

NewProviderSlackSendMessageIn instantiates a new ProviderSlackSendMessageIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackSendMessageInWithDefaults

`func NewProviderSlackSendMessageInWithDefaults() *ProviderSlackSendMessageIn`

NewProviderSlackSendMessageInWithDefaults instantiates a new ProviderSlackSendMessageIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackSendMessageIn) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackSendMessageIn) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackSendMessageIn) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackSendMessageIn) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetText

`func (o *ProviderSlackSendMessageIn) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProviderSlackSendMessageIn) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProviderSlackSendMessageIn) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ProviderSlackSendMessageIn) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThreadTs

`func (o *ProviderSlackSendMessageIn) GetThreadTs() string`

GetThreadTs returns the ThreadTs field if non-nil, zero value otherwise.

### GetThreadTsOk

`func (o *ProviderSlackSendMessageIn) GetThreadTsOk() (*string, bool)`

GetThreadTsOk returns a tuple with the ThreadTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadTs

`func (o *ProviderSlackSendMessageIn) SetThreadTs(v string)`

SetThreadTs sets ThreadTs field to given value.

### HasThreadTs

`func (o *ProviderSlackSendMessageIn) HasThreadTs() bool`

HasThreadTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


