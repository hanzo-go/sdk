# ProviderSlackUpdateMessageIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is a Slack channel ID, name (#hanzo-gtm), or channel mention. | [optional] 
**Text** | Pointer to **string** | Text is the replacement, up to 40000 characters. | [optional] 
**Ts** | Pointer to **string** | TS is the exact timestamp of the message to rewrite. | [optional] 

## Methods

### NewProviderSlackUpdateMessageIn

`func NewProviderSlackUpdateMessageIn() *ProviderSlackUpdateMessageIn`

NewProviderSlackUpdateMessageIn instantiates a new ProviderSlackUpdateMessageIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackUpdateMessageInWithDefaults

`func NewProviderSlackUpdateMessageInWithDefaults() *ProviderSlackUpdateMessageIn`

NewProviderSlackUpdateMessageInWithDefaults instantiates a new ProviderSlackUpdateMessageIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackUpdateMessageIn) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackUpdateMessageIn) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackUpdateMessageIn) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackUpdateMessageIn) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetText

`func (o *ProviderSlackUpdateMessageIn) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProviderSlackUpdateMessageIn) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProviderSlackUpdateMessageIn) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ProviderSlackUpdateMessageIn) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTs

`func (o *ProviderSlackUpdateMessageIn) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProviderSlackUpdateMessageIn) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProviderSlackUpdateMessageIn) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProviderSlackUpdateMessageIn) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


