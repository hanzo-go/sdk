# CampaignChannelSpec

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

### NewCampaignChannelSpec

`func NewCampaignChannelSpec() *CampaignChannelSpec`

NewCampaignChannelSpec instantiates a new CampaignChannelSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignChannelSpecWithDefaults

`func NewCampaignChannelSpecWithDefaults() *CampaignChannelSpec`

NewCampaignChannelSpecWithDefaults instantiates a new CampaignChannelSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CampaignChannelSpec) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CampaignChannelSpec) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CampaignChannelSpec) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CampaignChannelSpec) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDetail

`func (o *CampaignChannelSpec) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CampaignChannelSpec) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CampaignChannelSpec) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CampaignChannelSpec) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetExternalId

`func (o *CampaignChannelSpec) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CampaignChannelSpec) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CampaignChannelSpec) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CampaignChannelSpec) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetKind

`func (o *CampaignChannelSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CampaignChannelSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CampaignChannelSpec) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CampaignChannelSpec) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPlatform

`func (o *CampaignChannelSpec) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CampaignChannelSpec) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CampaignChannelSpec) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CampaignChannelSpec) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *CampaignChannelSpec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignChannelSpec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignChannelSpec) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CampaignChannelSpec) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


