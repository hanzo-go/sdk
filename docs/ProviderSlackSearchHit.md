# ProviderSlackSearchHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the conversation ID the match sits in. | [optional] 
**ChannelName** | Pointer to **string** | ChannelName is that conversation&#39;s name, when Slack reports one. | [optional] 
**Permalink** | Pointer to **string** | Permalink addresses the message in a browser. | [optional] 
**Text** | Pointer to **string** | Text is the matching message&#39;s text. | [optional] 
**Ts** | Pointer to **string** | TS is the message&#39;s exact Slack timestamp, the handle for replying to or reacting to it. | [optional] 
**User** | Pointer to **string** | User is the author&#39;s Slack user ID. | [optional] 
**Username** | Pointer to **string** | Username is the author&#39;s display name, when Slack reports one. | [optional] 

## Methods

### NewProviderSlackSearchHit

`func NewProviderSlackSearchHit() *ProviderSlackSearchHit`

NewProviderSlackSearchHit instantiates a new ProviderSlackSearchHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackSearchHitWithDefaults

`func NewProviderSlackSearchHitWithDefaults() *ProviderSlackSearchHit`

NewProviderSlackSearchHitWithDefaults instantiates a new ProviderSlackSearchHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackSearchHit) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackSearchHit) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackSearchHit) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackSearchHit) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelName

`func (o *ProviderSlackSearchHit) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ProviderSlackSearchHit) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ProviderSlackSearchHit) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ProviderSlackSearchHit) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetPermalink

`func (o *ProviderSlackSearchHit) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *ProviderSlackSearchHit) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *ProviderSlackSearchHit) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *ProviderSlackSearchHit) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetText

`func (o *ProviderSlackSearchHit) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProviderSlackSearchHit) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProviderSlackSearchHit) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ProviderSlackSearchHit) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTs

`func (o *ProviderSlackSearchHit) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProviderSlackSearchHit) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProviderSlackSearchHit) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProviderSlackSearchHit) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUser

`func (o *ProviderSlackSearchHit) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProviderSlackSearchHit) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProviderSlackSearchHit) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ProviderSlackSearchHit) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUsername

`func (o *ProviderSlackSearchHit) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProviderSlackSearchHit) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProviderSlackSearchHit) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProviderSlackSearchHit) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


