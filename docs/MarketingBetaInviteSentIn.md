# MarketingBetaInviteSentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | Pointer to [**[]MarketingBetaInviteFailure**](MarketingBetaInviteFailure.md) | Failed are the invites that did not go out. | [optional] 
**Ids** | Pointer to **[]string** | IDs are the invites sent. | [optional] 

## Methods

### NewMarketingBetaInviteSentIn

`func NewMarketingBetaInviteSentIn() *MarketingBetaInviteSentIn`

NewMarketingBetaInviteSentIn instantiates a new MarketingBetaInviteSentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaInviteSentInWithDefaults

`func NewMarketingBetaInviteSentInWithDefaults() *MarketingBetaInviteSentIn`

NewMarketingBetaInviteSentInWithDefaults instantiates a new MarketingBetaInviteSentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *MarketingBetaInviteSentIn) GetFailed() []MarketingBetaInviteFailure`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *MarketingBetaInviteSentIn) GetFailedOk() (*[]MarketingBetaInviteFailure, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *MarketingBetaInviteSentIn) SetFailed(v []MarketingBetaInviteFailure)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *MarketingBetaInviteSentIn) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetIds

`func (o *MarketingBetaInviteSentIn) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MarketingBetaInviteSentIn) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MarketingBetaInviteSentIn) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MarketingBetaInviteSentIn) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


