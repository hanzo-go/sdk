# ProviderSlackReactOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the resolved channel ID the message sits in. | [optional] 
**Name** | Pointer to **string** | Name is the emoji that was added, without colons. | [optional] 
**Ts** | Pointer to **string** | TS is the reacted-to message&#39;s exact Slack timestamp. | [optional] 

## Methods

### NewProviderSlackReactOut

`func NewProviderSlackReactOut() *ProviderSlackReactOut`

NewProviderSlackReactOut instantiates a new ProviderSlackReactOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackReactOutWithDefaults

`func NewProviderSlackReactOutWithDefaults() *ProviderSlackReactOut`

NewProviderSlackReactOutWithDefaults instantiates a new ProviderSlackReactOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackReactOut) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackReactOut) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackReactOut) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackReactOut) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetName

`func (o *ProviderSlackReactOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSlackReactOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSlackReactOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderSlackReactOut) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTs

`func (o *ProviderSlackReactOut) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProviderSlackReactOut) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProviderSlackReactOut) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProviderSlackReactOut) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


