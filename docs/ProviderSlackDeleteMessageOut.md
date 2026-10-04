# ProviderSlackDeleteMessageOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the resolved channel ID the message was in. | [optional] 
**Ts** | Pointer to **string** | TS is the deleted message&#39;s timestamp, which now addresses nothing. | [optional] 

## Methods

### NewProviderSlackDeleteMessageOut

`func NewProviderSlackDeleteMessageOut() *ProviderSlackDeleteMessageOut`

NewProviderSlackDeleteMessageOut instantiates a new ProviderSlackDeleteMessageOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackDeleteMessageOutWithDefaults

`func NewProviderSlackDeleteMessageOutWithDefaults() *ProviderSlackDeleteMessageOut`

NewProviderSlackDeleteMessageOutWithDefaults instantiates a new ProviderSlackDeleteMessageOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackDeleteMessageOut) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackDeleteMessageOut) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackDeleteMessageOut) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackDeleteMessageOut) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTs

`func (o *ProviderSlackDeleteMessageOut) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProviderSlackDeleteMessageOut) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProviderSlackDeleteMessageOut) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProviderSlackDeleteMessageOut) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


