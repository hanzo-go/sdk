# CloudSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the recipient, normalized (lower-cased, trimmed) so an opt-out cannot be slipped past on a case or whitespace difference. Required. | [optional] 
**Channel** | Pointer to **string** | Channel is the surface opted out of: email, sms, social, meta, google or tiktok. Empty means email. Opting out of one leaves the others reachable. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is unix seconds, server-assigned. | [optional] 
**Reason** | Pointer to **string** | Reason is a free-text note, capped at 1024 bytes. The public one-click endpoint records \&quot;one-click unsubscribe\&quot;. | [optional] 

## Methods

### NewCloudSuppression

`func NewCloudSuppression() *CloudSuppression`

NewCloudSuppression instantiates a new CloudSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSuppressionWithDefaults

`func NewCloudSuppressionWithDefaults() *CloudSuppression`

NewCloudSuppressionWithDefaults instantiates a new CloudSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CloudSuppression) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CloudSuppression) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CloudSuppression) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CloudSuppression) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetChannel

`func (o *CloudSuppression) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CloudSuppression) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CloudSuppression) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CloudSuppression) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudSuppression) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSuppression) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSuppression) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSuppression) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetReason

`func (o *CloudSuppression) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudSuppression) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudSuppression) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudSuppression) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


