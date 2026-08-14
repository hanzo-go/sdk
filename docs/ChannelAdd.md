# ChannelAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider account this channel runs under: an ad-account, a page, or a mailing-list id. | [optional] 
**Id** | Pointer to **string** | ID is the campaign to add the channel to, from the path. | [optional] 
**Kind** | Pointer to **string** | Kind is the channel kind and the identity a campaign holds at most one of: paid, organic or email. | [optional] 
**Platform** | Pointer to **string** | Platform is the provider within the kind — meta, google, x, instagram, or the email provider. | [optional] 

## Methods

### NewChannelAdd

`func NewChannelAdd() *ChannelAdd`

NewChannelAdd instantiates a new ChannelAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelAddWithDefaults

`func NewChannelAddWithDefaults() *ChannelAdd`

NewChannelAddWithDefaults instantiates a new ChannelAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ChannelAdd) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ChannelAdd) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ChannelAdd) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ChannelAdd) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetId

`func (o *ChannelAdd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelAdd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelAdd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelAdd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ChannelAdd) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ChannelAdd) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ChannelAdd) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ChannelAdd) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPlatform

`func (o *ChannelAdd) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ChannelAdd) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ChannelAdd) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ChannelAdd) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


