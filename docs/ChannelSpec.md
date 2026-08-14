# ChannelSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider account this channel runs under: an ad-account, a page or a mailing-list id. An executor may replace it at launch with the account it actually used. | [optional] 
**Detail** | Pointer to **string** | Detail is the last outcome in one secret-free line — the failure reason, or what the executor reported. Absent when there is nothing to explain. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the provider-side id of the running execution, recorded by the orchestrator at launch and handed back verbatim to read spend or to pause. Server-owned and absent until this channel has launched; anything a caller sends for it is dropped. | [optional] 
**Kind** | Pointer to **string** | Kind is the channel and the identity a campaign holds at most one of: paid, organic or email. It picks the executor the launch fans out to. | [optional] 
**Platform** | Pointer to **string** | Platform is the provider within the kind — meta, google, x, instagram, or the email provider. | [optional] 
**Status** | Pointer to **string** | Status is this channel&#39;s own launch outcome, not the campaign&#39;s: pending (added, never launched), live, paused, failed (Detail says why) or unavailable (no executor wired on this deployment). Server-owned — a caller can never assert it. | [optional] 

## Methods

### NewChannelSpec

`func NewChannelSpec() *ChannelSpec`

NewChannelSpec instantiates a new ChannelSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelSpecWithDefaults

`func NewChannelSpecWithDefaults() *ChannelSpec`

NewChannelSpecWithDefaults instantiates a new ChannelSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ChannelSpec) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ChannelSpec) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ChannelSpec) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ChannelSpec) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDetail

`func (o *ChannelSpec) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ChannelSpec) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ChannelSpec) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ChannelSpec) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetExternalId

`func (o *ChannelSpec) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ChannelSpec) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ChannelSpec) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ChannelSpec) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetKind

`func (o *ChannelSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ChannelSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ChannelSpec) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ChannelSpec) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPlatform

`func (o *ChannelSpec) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ChannelSpec) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ChannelSpec) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ChannelSpec) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *ChannelSpec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChannelSpec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChannelSpec) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChannelSpec) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


