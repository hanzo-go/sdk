# CloudChallengeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **int32** | ExpiresAt is when the nonce stops being redeemable, as a Unix timestamp. | [optional] 
**Message** | Pointer to **string** | Message is the EXACT text to personal_sign. It is reconstructed server-side from the validated org, the slot and the nonce at redemption, so signing anything else cannot claim the slot. | [optional] 
**Nonce** | Pointer to **string** | Nonce is the single-use, org-bound challenge value to send back with the signature. | [optional] 
**TokenId** | Pointer to **int32** | TokenID is the slot the challenge was issued for. | [optional] 
**TtlSeconds** | Pointer to **int32** | TTLSeconds is the challenge lifetime in seconds. | [optional] 

## Methods

### NewCloudChallengeView

`func NewCloudChallengeView() *CloudChallengeView`

NewCloudChallengeView instantiates a new CloudChallengeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudChallengeViewWithDefaults

`func NewCloudChallengeViewWithDefaults() *CloudChallengeView`

NewCloudChallengeViewWithDefaults instantiates a new CloudChallengeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CloudChallengeView) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CloudChallengeView) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CloudChallengeView) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CloudChallengeView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetMessage

`func (o *CloudChallengeView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudChallengeView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudChallengeView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudChallengeView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNonce

`func (o *CloudChallengeView) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CloudChallengeView) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CloudChallengeView) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *CloudChallengeView) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetTokenId

`func (o *CloudChallengeView) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CloudChallengeView) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CloudChallengeView) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *CloudChallengeView) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTtlSeconds

`func (o *CloudChallengeView) GetTtlSeconds() int32`

GetTtlSeconds returns the TtlSeconds field if non-nil, zero value otherwise.

### GetTtlSecondsOk

`func (o *CloudChallengeView) GetTtlSecondsOk() (*int32, bool)`

GetTtlSecondsOk returns a tuple with the TtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSeconds

`func (o *CloudChallengeView) SetTtlSeconds(v int32)`

SetTtlSeconds sets TtlSeconds field to given value.

### HasTtlSeconds

`func (o *CloudChallengeView) HasTtlSeconds() bool`

HasTtlSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


