# CloudMyReferralView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the referral was recorded, as a Unix timestamp. | [optional] 
**CreditedAt** | Pointer to **int32** | CreditedAt is when the bonuses were latched and granted, as a Unix timestamp; 0 until they are. It is the at-most-once latch. | [optional] 
**CreditsCents** | Pointer to **int32** | CreditsCents is what I earned from this referral, in USD cents. It is 0 until the referee qualifies. | [optional] 
**Id** | Pointer to **string** | ID is the referral&#39;s handle. | [optional] 
**QualifiedAt** | Pointer to **int32** | QualifiedAt is when the referee first made metered spend, as a Unix timestamp; 0 while the referral is still pending. | [optional] 
**Referee** | Pointer to **string** | Referee is the org that signed up with my code. | [optional] 
**Status** | Pointer to **string** | Status is the referral&#39;s lifecycle state: \&quot;signup\&quot; until the referee makes metered spend, then \&quot;qualified\&quot;, then \&quot;credited\&quot;. | [optional] 

## Methods

### NewCloudMyReferralView

`func NewCloudMyReferralView() *CloudMyReferralView`

NewCloudMyReferralView instantiates a new CloudMyReferralView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMyReferralViewWithDefaults

`func NewCloudMyReferralViewWithDefaults() *CloudMyReferralView`

NewCloudMyReferralViewWithDefaults instantiates a new CloudMyReferralView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudMyReferralView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudMyReferralView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudMyReferralView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudMyReferralView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditedAt

`func (o *CloudMyReferralView) GetCreditedAt() int32`

GetCreditedAt returns the CreditedAt field if non-nil, zero value otherwise.

### GetCreditedAtOk

`func (o *CloudMyReferralView) GetCreditedAtOk() (*int32, bool)`

GetCreditedAtOk returns a tuple with the CreditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedAt

`func (o *CloudMyReferralView) SetCreditedAt(v int32)`

SetCreditedAt sets CreditedAt field to given value.

### HasCreditedAt

`func (o *CloudMyReferralView) HasCreditedAt() bool`

HasCreditedAt returns a boolean if a field has been set.

### GetCreditsCents

`func (o *CloudMyReferralView) GetCreditsCents() int32`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *CloudMyReferralView) GetCreditsCentsOk() (*int32, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *CloudMyReferralView) SetCreditsCents(v int32)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *CloudMyReferralView) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetId

`func (o *CloudMyReferralView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudMyReferralView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudMyReferralView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudMyReferralView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQualifiedAt

`func (o *CloudMyReferralView) GetQualifiedAt() int32`

GetQualifiedAt returns the QualifiedAt field if non-nil, zero value otherwise.

### GetQualifiedAtOk

`func (o *CloudMyReferralView) GetQualifiedAtOk() (*int32, bool)`

GetQualifiedAtOk returns a tuple with the QualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedAt

`func (o *CloudMyReferralView) SetQualifiedAt(v int32)`

SetQualifiedAt sets QualifiedAt field to given value.

### HasQualifiedAt

`func (o *CloudMyReferralView) HasQualifiedAt() bool`

HasQualifiedAt returns a boolean if a field has been set.

### GetReferee

`func (o *CloudMyReferralView) GetReferee() string`

GetReferee returns the Referee field if non-nil, zero value otherwise.

### GetRefereeOk

`func (o *CloudMyReferralView) GetRefereeOk() (*string, bool)`

GetRefereeOk returns a tuple with the Referee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferee

`func (o *CloudMyReferralView) SetReferee(v string)`

SetReferee sets Referee field to given value.

### HasReferee

`func (o *CloudMyReferralView) HasReferee() bool`

HasReferee returns a boolean if a field has been set.

### GetStatus

`func (o *CloudMyReferralView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudMyReferralView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudMyReferralView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudMyReferralView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


