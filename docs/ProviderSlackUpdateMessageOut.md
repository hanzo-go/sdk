# ProviderSlackUpdateMessageOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the resolved channel ID the message sits in. | [optional] 
**Ts** | Pointer to **string** | TS is the rewritten message&#39;s exact Slack timestamp, unchanged by an edit. | [optional] 

## Methods

### NewProviderSlackUpdateMessageOut

`func NewProviderSlackUpdateMessageOut() *ProviderSlackUpdateMessageOut`

NewProviderSlackUpdateMessageOut instantiates a new ProviderSlackUpdateMessageOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackUpdateMessageOutWithDefaults

`func NewProviderSlackUpdateMessageOutWithDefaults() *ProviderSlackUpdateMessageOut`

NewProviderSlackUpdateMessageOutWithDefaults instantiates a new ProviderSlackUpdateMessageOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackUpdateMessageOut) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackUpdateMessageOut) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackUpdateMessageOut) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackUpdateMessageOut) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTs

`func (o *ProviderSlackUpdateMessageOut) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProviderSlackUpdateMessageOut) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProviderSlackUpdateMessageOut) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProviderSlackUpdateMessageOut) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


