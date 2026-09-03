# MyReferralView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the referral was recorded, as a Unix timestamp. | [optional] 
**Id** | Pointer to **string** | ID is the referral&#39;s handle. | [optional] 
**QualifiedAt** | Pointer to **int64** | QualifiedAt is when the referee first made metered spend, as a Unix timestamp; 0 while the referral is still pending. | [optional] 
**Referee** | Pointer to **string** | Referee is the org that signed up with my code. | [optional] 
**Status** | Pointer to **string** | Status is the referral&#39;s lifecycle state: \&quot;signup\&quot; until the referee makes metered spend, then \&quot;qualified\&quot;. | [optional] 

## Methods

### NewMyReferralView

`func NewMyReferralView() *MyReferralView`

NewMyReferralView instantiates a new MyReferralView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyReferralViewWithDefaults

`func NewMyReferralViewWithDefaults() *MyReferralView`

NewMyReferralViewWithDefaults instantiates a new MyReferralView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MyReferralView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MyReferralView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MyReferralView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MyReferralView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *MyReferralView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyReferralView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyReferralView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MyReferralView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQualifiedAt

`func (o *MyReferralView) GetQualifiedAt() int64`

GetQualifiedAt returns the QualifiedAt field if non-nil, zero value otherwise.

### GetQualifiedAtOk

`func (o *MyReferralView) GetQualifiedAtOk() (*int64, bool)`

GetQualifiedAtOk returns a tuple with the QualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedAt

`func (o *MyReferralView) SetQualifiedAt(v int64)`

SetQualifiedAt sets QualifiedAt field to given value.

### HasQualifiedAt

`func (o *MyReferralView) HasQualifiedAt() bool`

HasQualifiedAt returns a boolean if a field has been set.

### GetReferee

`func (o *MyReferralView) GetReferee() string`

GetReferee returns the Referee field if non-nil, zero value otherwise.

### GetRefereeOk

`func (o *MyReferralView) GetRefereeOk() (*string, bool)`

GetRefereeOk returns a tuple with the Referee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferee

`func (o *MyReferralView) SetReferee(v string)`

SetReferee sets Referee field to given value.

### HasReferee

`func (o *MyReferralView) HasReferee() bool`

HasReferee returns a boolean if a field has been set.

### GetStatus

`func (o *MyReferralView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MyReferralView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MyReferralView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MyReferralView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


