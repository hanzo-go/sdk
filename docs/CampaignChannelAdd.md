# CampaignChannelAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider account this channel runs under: an ad-account, a page, or a mailing-list id. | [optional] 
**Id** | Pointer to **string** | ID is the campaign to add the channel to, from the path. | [optional] 
**Kind** | Pointer to **string** | Kind is the channel kind and the identity a campaign holds at most one of: paid, organic or email. | [optional] 
**Platform** | Pointer to **string** | Platform is the provider within the kind — meta, google, x, instagram, or the email provider. | [optional] 

## Methods

### NewCampaignChannelAdd

`func NewCampaignChannelAdd() *CampaignChannelAdd`

NewCampaignChannelAdd instantiates a new CampaignChannelAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignChannelAddWithDefaults

`func NewCampaignChannelAddWithDefaults() *CampaignChannelAdd`

NewCampaignChannelAddWithDefaults instantiates a new CampaignChannelAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CampaignChannelAdd) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CampaignChannelAdd) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CampaignChannelAdd) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CampaignChannelAdd) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetId

`func (o *CampaignChannelAdd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignChannelAdd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignChannelAdd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignChannelAdd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CampaignChannelAdd) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CampaignChannelAdd) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CampaignChannelAdd) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CampaignChannelAdd) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPlatform

`func (o *CampaignChannelAdd) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CampaignChannelAdd) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CampaignChannelAdd) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CampaignChannelAdd) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


