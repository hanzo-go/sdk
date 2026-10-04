# MarketingSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the recipient, normalized (lower-cased, trimmed) so an opt-out cannot be slipped past on a case or whitespace difference. Required. | [optional] 
**Channel** | Pointer to **string** | Channel is the surface opted out of: email, sms, social, meta, google or tiktok. Empty means email. Opting out of one leaves the others reachable. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is unix seconds, server-assigned. | [optional] 
**Reason** | Pointer to **string** | Reason is a free-text note, capped at 1024 bytes. The public one-click endpoint records \&quot;one-click unsubscribe\&quot;. | [optional] 

## Methods

### NewMarketingSuppression

`func NewMarketingSuppression() *MarketingSuppression`

NewMarketingSuppression instantiates a new MarketingSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSuppressionWithDefaults

`func NewMarketingSuppressionWithDefaults() *MarketingSuppression`

NewMarketingSuppressionWithDefaults instantiates a new MarketingSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MarketingSuppression) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MarketingSuppression) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MarketingSuppression) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MarketingSuppression) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetChannel

`func (o *MarketingSuppression) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketingSuppression) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketingSuppression) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *MarketingSuppression) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingSuppression) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingSuppression) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingSuppression) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketingSuppression) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetReason

`func (o *MarketingSuppression) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MarketingSuppression) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MarketingSuppression) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MarketingSuppression) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


