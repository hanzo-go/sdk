# MarketingBetaInviteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counts** | Pointer to [**MarketingBetaInviteCounts**](MarketingBetaInviteCounts.md) | Counts are the caller&#39;s invites, counted. | [optional] 
**Invites** | Pointer to [**[]MarketingBetaInviteRow**](MarketingBetaInviteRow.md) | Invites are the caller&#39;s invites, newest first; a reservation that lapsed unsent is not among them. | [optional] 

## Methods

### NewMarketingBetaInviteList

`func NewMarketingBetaInviteList() *MarketingBetaInviteList`

NewMarketingBetaInviteList instantiates a new MarketingBetaInviteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaInviteListWithDefaults

`func NewMarketingBetaInviteListWithDefaults() *MarketingBetaInviteList`

NewMarketingBetaInviteListWithDefaults instantiates a new MarketingBetaInviteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounts

`func (o *MarketingBetaInviteList) GetCounts() MarketingBetaInviteCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *MarketingBetaInviteList) GetCountsOk() (*MarketingBetaInviteCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *MarketingBetaInviteList) SetCounts(v MarketingBetaInviteCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *MarketingBetaInviteList) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetInvites

`func (o *MarketingBetaInviteList) GetInvites() []MarketingBetaInviteRow`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *MarketingBetaInviteList) GetInvitesOk() (*[]MarketingBetaInviteRow, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *MarketingBetaInviteList) SetInvites(v []MarketingBetaInviteRow)`

SetInvites sets Invites field to given value.

### HasInvites

`func (o *MarketingBetaInviteList) HasInvites() bool`

HasInvites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


