# MarketingBetaInviteCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Joined** | Pointer to **int64** | Joined counts invites whose invitee signed up with the member&#39;s code. | [optional] 
**Left** | Pointer to **int64** | Left is how many more the member may invite now. | [optional] 
**Sent** | Pointer to **int64** | Sent counts invites sent, joined ones included. | [optional] 

## Methods

### NewMarketingBetaInviteCounts

`func NewMarketingBetaInviteCounts() *MarketingBetaInviteCounts`

NewMarketingBetaInviteCounts instantiates a new MarketingBetaInviteCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaInviteCountsWithDefaults

`func NewMarketingBetaInviteCountsWithDefaults() *MarketingBetaInviteCounts`

NewMarketingBetaInviteCountsWithDefaults instantiates a new MarketingBetaInviteCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoined

`func (o *MarketingBetaInviteCounts) GetJoined() int64`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *MarketingBetaInviteCounts) GetJoinedOk() (*int64, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *MarketingBetaInviteCounts) SetJoined(v int64)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *MarketingBetaInviteCounts) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### GetLeft

`func (o *MarketingBetaInviteCounts) GetLeft() int64`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *MarketingBetaInviteCounts) GetLeftOk() (*int64, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *MarketingBetaInviteCounts) SetLeft(v int64)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *MarketingBetaInviteCounts) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### GetSent

`func (o *MarketingBetaInviteCounts) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MarketingBetaInviteCounts) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MarketingBetaInviteCounts) SetSent(v int64)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MarketingBetaInviteCounts) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


