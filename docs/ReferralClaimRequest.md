# ReferralClaimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the referrer&#39;s referral code, as it appeared in their ?ref&#x3D; link. Case and surrounding whitespace do not matter. | [optional] 
**Invite** | Pointer to **string** | Invite is the beta invite the referee&#39;s link carried (its utm_content), when it carried one. Optional; it only lets the referrer&#39;s invite be marked joined, and never changes who the referral credits. | [optional] 

## Methods

### NewReferralClaimRequest

`func NewReferralClaimRequest() *ReferralClaimRequest`

NewReferralClaimRequest instantiates a new ReferralClaimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralClaimRequestWithDefaults

`func NewReferralClaimRequestWithDefaults() *ReferralClaimRequest`

NewReferralClaimRequestWithDefaults instantiates a new ReferralClaimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ReferralClaimRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReferralClaimRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReferralClaimRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReferralClaimRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInvite

`func (o *ReferralClaimRequest) GetInvite() string`

GetInvite returns the Invite field if non-nil, zero value otherwise.

### GetInviteOk

`func (o *ReferralClaimRequest) GetInviteOk() (*string, bool)`

GetInviteOk returns a tuple with the Invite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvite

`func (o *ReferralClaimRequest) SetInvite(v string)`

SetInvite sets Invite field to given value.

### HasInvite

`func (o *ReferralClaimRequest) HasInvite() bool`

HasInvite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


