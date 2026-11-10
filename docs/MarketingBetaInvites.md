# MarketingBetaInvites

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invites** | Pointer to [**[]MarketingBetaInvite**](MarketingBetaInvite.md) | Invites are the reserved invites, each with its message. | [optional] 
**Left** | Pointer to **int64** | Left is how many more the member may invite now: the smaller of what is left today and what is left this month. | [optional] 
**Skipped** | Pointer to [**[]MarketingBetaInviteSkip**](MarketingBetaInviteSkip.md) | Skipped are the addresses not reserved. | [optional] 

## Methods

### NewMarketingBetaInvites

`func NewMarketingBetaInvites() *MarketingBetaInvites`

NewMarketingBetaInvites instantiates a new MarketingBetaInvites object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaInvitesWithDefaults

`func NewMarketingBetaInvitesWithDefaults() *MarketingBetaInvites`

NewMarketingBetaInvitesWithDefaults instantiates a new MarketingBetaInvites object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvites

`func (o *MarketingBetaInvites) GetInvites() []MarketingBetaInvite`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *MarketingBetaInvites) GetInvitesOk() (*[]MarketingBetaInvite, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *MarketingBetaInvites) SetInvites(v []MarketingBetaInvite)`

SetInvites sets Invites field to given value.

### HasInvites

`func (o *MarketingBetaInvites) HasInvites() bool`

HasInvites returns a boolean if a field has been set.

### GetLeft

`func (o *MarketingBetaInvites) GetLeft() int64`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *MarketingBetaInvites) GetLeftOk() (*int64, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *MarketingBetaInvites) SetLeft(v int64)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *MarketingBetaInvites) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### GetSkipped

`func (o *MarketingBetaInvites) GetSkipped() []MarketingBetaInviteSkip`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *MarketingBetaInvites) GetSkippedOk() (*[]MarketingBetaInviteSkip, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *MarketingBetaInvites) SetSkipped(v []MarketingBetaInviteSkip)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *MarketingBetaInvites) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


