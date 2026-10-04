# CampaignCampaignSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to **int64** | Budget is the sum of every campaign&#39;s budget, in CENTS. | [optional] 
**Campaigns** | Pointer to **int64** | Campaigns is how many campaigns the org has, in any state. | [optional] 
**Channels** | Pointer to **[]string** | Channels are the channel kinds this deployment has an executor wired for. A kind absent here is a kind a launch will honestly record as unavailable. | [optional] 
**Live** | Pointer to **int64** | Live is how many of them are currently live. | [optional] 

## Methods

### NewCampaignCampaignSummary

`func NewCampaignCampaignSummary() *CampaignCampaignSummary`

NewCampaignCampaignSummary instantiates a new CampaignCampaignSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCampaignSummaryWithDefaults

`func NewCampaignCampaignSummaryWithDefaults() *CampaignCampaignSummary`

NewCampaignCampaignSummaryWithDefaults instantiates a new CampaignCampaignSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *CampaignCampaignSummary) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CampaignCampaignSummary) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CampaignCampaignSummary) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CampaignCampaignSummary) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCampaigns

`func (o *CampaignCampaignSummary) GetCampaigns() int64`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *CampaignCampaignSummary) GetCampaignsOk() (*int64, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *CampaignCampaignSummary) SetCampaigns(v int64)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *CampaignCampaignSummary) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetChannels

`func (o *CampaignCampaignSummary) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CampaignCampaignSummary) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CampaignCampaignSummary) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CampaignCampaignSummary) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetLive

`func (o *CampaignCampaignSummary) GetLive() int64`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *CampaignCampaignSummary) GetLiveOk() (*int64, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *CampaignCampaignSummary) SetLive(v int64)`

SetLive sets Live field to given value.

### HasLive

`func (o *CampaignCampaignSummary) HasLive() bool`

HasLive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


