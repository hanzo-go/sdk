# CloudChannelAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider account this channel runs under: an ad-account, a page, or a mailing-list id. | [optional] 
**Id** | Pointer to **string** | ID is the campaign to add the channel to, from the path. | [optional] 
**Kind** | Pointer to **string** | Kind is the channel kind and the identity a campaign holds at most one of: paid, organic or email. | [optional] 
**Platform** | Pointer to **string** | Platform is the provider within the kind — meta, google, x, instagram, or the email provider. | [optional] 

## Methods

### NewCloudChannelAdd

`func NewCloudChannelAdd() *CloudChannelAdd`

NewCloudChannelAdd instantiates a new CloudChannelAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudChannelAddWithDefaults

`func NewCloudChannelAddWithDefaults() *CloudChannelAdd`

NewCloudChannelAddWithDefaults instantiates a new CloudChannelAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudChannelAdd) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudChannelAdd) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudChannelAdd) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudChannelAdd) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetId

`func (o *CloudChannelAdd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudChannelAdd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudChannelAdd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudChannelAdd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudChannelAdd) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudChannelAdd) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudChannelAdd) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudChannelAdd) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudChannelAdd) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudChannelAdd) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudChannelAdd) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudChannelAdd) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


