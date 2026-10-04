# ProviderSlackSendMessageOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the resolved channel ID where the message was posted. | [optional] 
**Ts** | Pointer to **string** | TS is the posted message&#39;s exact Slack timestamp. | [optional] 

## Methods

### NewProviderSlackSendMessageOut

`func NewProviderSlackSendMessageOut() *ProviderSlackSendMessageOut`

NewProviderSlackSendMessageOut instantiates a new ProviderSlackSendMessageOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackSendMessageOutWithDefaults

`func NewProviderSlackSendMessageOutWithDefaults() *ProviderSlackSendMessageOut`

NewProviderSlackSendMessageOutWithDefaults instantiates a new ProviderSlackSendMessageOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackSendMessageOut) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackSendMessageOut) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackSendMessageOut) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackSendMessageOut) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTs

`func (o *ProviderSlackSendMessageOut) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProviderSlackSendMessageOut) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProviderSlackSendMessageOut) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProviderSlackSendMessageOut) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


