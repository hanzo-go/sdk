# CloudCampaignSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to **int32** | Budget is the sum of every campaign&#39;s budget, in CENTS. | [optional] 
**Campaigns** | Pointer to **int32** | Campaigns is how many campaigns the org has, in any state. | [optional] 
**Channels** | Pointer to **[]string** | Channels are the channel kinds this deployment has an executor wired for. A kind absent here is a kind a launch will honestly record as unavailable. | [optional] 
**Live** | Pointer to **int32** | Live is how many of them are currently live. | [optional] 

## Methods

### NewCloudCampaignSummary

`func NewCloudCampaignSummary() *CloudCampaignSummary`

NewCloudCampaignSummary instantiates a new CloudCampaignSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCampaignSummaryWithDefaults

`func NewCloudCampaignSummaryWithDefaults() *CloudCampaignSummary`

NewCloudCampaignSummaryWithDefaults instantiates a new CloudCampaignSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *CloudCampaignSummary) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CloudCampaignSummary) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CloudCampaignSummary) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CloudCampaignSummary) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCampaigns

`func (o *CloudCampaignSummary) GetCampaigns() int32`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *CloudCampaignSummary) GetCampaignsOk() (*int32, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *CloudCampaignSummary) SetCampaigns(v int32)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *CloudCampaignSummary) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetChannels

`func (o *CloudCampaignSummary) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CloudCampaignSummary) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CloudCampaignSummary) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CloudCampaignSummary) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetLive

`func (o *CloudCampaignSummary) GetLive() int32`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *CloudCampaignSummary) GetLiveOk() (*int32, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *CloudCampaignSummary) SetLive(v int32)`

SetLive sets Live field to given value.

### HasLive

`func (o *CloudCampaignSummary) HasLive() bool`

HasLive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


