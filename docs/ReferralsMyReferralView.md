# ReferralsMyReferralView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Referee** | Pointer to **string** | The org that signed up via my code. | [optional] 
**Status** | Pointer to [**ReferralsReferralStatus**](ReferralsReferralStatus.md) |  | [optional] 
**CreditsCents** | Pointer to **int64** | What I earned from this referral, in USD minor units (cents). | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp. | [optional] 
**QualifiedAt** | Pointer to **int64** | Unix timestamp; 0 if not yet qualified. | [optional] 
**CreditedAt** | Pointer to **int64** | Unix timestamp; 0 if not yet credited. | [optional] 

## Methods

### NewReferralsMyReferralView

`func NewReferralsMyReferralView() *ReferralsMyReferralView`

NewReferralsMyReferralView instantiates a new ReferralsMyReferralView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralsMyReferralViewWithDefaults

`func NewReferralsMyReferralViewWithDefaults() *ReferralsMyReferralView`

NewReferralsMyReferralViewWithDefaults instantiates a new ReferralsMyReferralView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferralsMyReferralView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferralsMyReferralView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferralsMyReferralView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReferralsMyReferralView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferee

`func (o *ReferralsMyReferralView) GetReferee() string`

GetReferee returns the Referee field if non-nil, zero value otherwise.

### GetRefereeOk

`func (o *ReferralsMyReferralView) GetRefereeOk() (*string, bool)`

GetRefereeOk returns a tuple with the Referee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferee

`func (o *ReferralsMyReferralView) SetReferee(v string)`

SetReferee sets Referee field to given value.

### HasReferee

`func (o *ReferralsMyReferralView) HasReferee() bool`

HasReferee returns a boolean if a field has been set.

### GetStatus

`func (o *ReferralsMyReferralView) GetStatus() ReferralsReferralStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReferralsMyReferralView) GetStatusOk() (*ReferralsReferralStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReferralsMyReferralView) SetStatus(v ReferralsReferralStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReferralsMyReferralView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreditsCents

`func (o *ReferralsMyReferralView) GetCreditsCents() int64`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *ReferralsMyReferralView) GetCreditsCentsOk() (*int64, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *ReferralsMyReferralView) SetCreditsCents(v int64)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *ReferralsMyReferralView) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReferralsMyReferralView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReferralsMyReferralView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReferralsMyReferralView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReferralsMyReferralView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetQualifiedAt

`func (o *ReferralsMyReferralView) GetQualifiedAt() int64`

GetQualifiedAt returns the QualifiedAt field if non-nil, zero value otherwise.

### GetQualifiedAtOk

`func (o *ReferralsMyReferralView) GetQualifiedAtOk() (*int64, bool)`

GetQualifiedAtOk returns a tuple with the QualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedAt

`func (o *ReferralsMyReferralView) SetQualifiedAt(v int64)`

SetQualifiedAt sets QualifiedAt field to given value.

### HasQualifiedAt

`func (o *ReferralsMyReferralView) HasQualifiedAt() bool`

HasQualifiedAt returns a boolean if a field has been set.

### GetCreditedAt

`func (o *ReferralsMyReferralView) GetCreditedAt() int64`

GetCreditedAt returns the CreditedAt field if non-nil, zero value otherwise.

### GetCreditedAtOk

`func (o *ReferralsMyReferralView) GetCreditedAtOk() (*int64, bool)`

GetCreditedAtOk returns a tuple with the CreditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedAt

`func (o *ReferralsMyReferralView) SetCreditedAt(v int64)`

SetCreditedAt sets CreditedAt field to given value.

### HasCreditedAt

`func (o *ReferralsMyReferralView) HasCreditedAt() bool`

HasCreditedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


