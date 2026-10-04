# ProviderSlackReactIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is a Slack channel ID, name (#hanzo-gtm), or channel mention. | [optional] 
**Name** | Pointer to **string** | Name is the emoji&#39;s name WITHOUT colons — \&quot;eyes\&quot;, not \&quot;:eyes:\&quot;. | [optional] 
**Ts** | Pointer to **string** | TS is the exact timestamp of the message to react to. | [optional] 

## Methods

### NewProviderSlackReactIn

`func NewProviderSlackReactIn() *ProviderSlackReactIn`

NewProviderSlackReactIn instantiates a new ProviderSlackReactIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackReactInWithDefaults

`func NewProviderSlackReactInWithDefaults() *ProviderSlackReactIn`

NewProviderSlackReactInWithDefaults instantiates a new ProviderSlackReactIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackReactIn) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackReactIn) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackReactIn) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackReactIn) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetName

`func (o *ProviderSlackReactIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSlackReactIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSlackReactIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderSlackReactIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTs

`func (o *ProviderSlackReactIn) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProviderSlackReactIn) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProviderSlackReactIn) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProviderSlackReactIn) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


