# ReferralBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualByLevel** | Pointer to [**LevelSplit**](LevelSplit.md) | AccrualByLevel splits the lifetime accrual across the three upline levels — how much of the liability comes from direct referrals versus the chain above. | [optional] 
**Conversion** | Pointer to [**Funnel**](Funnel.md) | Conversion is the funnel: referred orgs against those that actually earned. | [optional] 
**Summary** | Pointer to [**Tally**](Tally.md) | Summary is the fleet tally — population by status, and lifetime accrued, paid and still-owed commission. | [optional] 
**TopReferrers** | Pointer to [**[]ReferrerRow**](ReferrerRow.md) | TopReferrers is the 25 affiliates with the most lifetime accrued commission, descending, orgs named. | [optional] 

## Methods

### NewReferralBoard

`func NewReferralBoard() *ReferralBoard`

NewReferralBoard instantiates a new ReferralBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralBoardWithDefaults

`func NewReferralBoardWithDefaults() *ReferralBoard`

NewReferralBoardWithDefaults instantiates a new ReferralBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrualByLevel

`func (o *ReferralBoard) GetAccrualByLevel() LevelSplit`

GetAccrualByLevel returns the AccrualByLevel field if non-nil, zero value otherwise.

### GetAccrualByLevelOk

`func (o *ReferralBoard) GetAccrualByLevelOk() (*LevelSplit, bool)`

GetAccrualByLevelOk returns a tuple with the AccrualByLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualByLevel

`func (o *ReferralBoard) SetAccrualByLevel(v LevelSplit)`

SetAccrualByLevel sets AccrualByLevel field to given value.

### HasAccrualByLevel

`func (o *ReferralBoard) HasAccrualByLevel() bool`

HasAccrualByLevel returns a boolean if a field has been set.

### GetConversion

`func (o *ReferralBoard) GetConversion() Funnel`

GetConversion returns the Conversion field if non-nil, zero value otherwise.

### GetConversionOk

`func (o *ReferralBoard) GetConversionOk() (*Funnel, bool)`

GetConversionOk returns a tuple with the Conversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversion

`func (o *ReferralBoard) SetConversion(v Funnel)`

SetConversion sets Conversion field to given value.

### HasConversion

`func (o *ReferralBoard) HasConversion() bool`

HasConversion returns a boolean if a field has been set.

### GetSummary

`func (o *ReferralBoard) GetSummary() Tally`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReferralBoard) GetSummaryOk() (*Tally, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReferralBoard) SetSummary(v Tally)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ReferralBoard) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTopReferrers

`func (o *ReferralBoard) GetTopReferrers() []ReferrerRow`

GetTopReferrers returns the TopReferrers field if non-nil, zero value otherwise.

### GetTopReferrersOk

`func (o *ReferralBoard) GetTopReferrersOk() (*[]ReferrerRow, bool)`

GetTopReferrersOk returns a tuple with the TopReferrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopReferrers

`func (o *ReferralBoard) SetTopReferrers(v []ReferrerRow)`

SetTopReferrers sets TopReferrers field to given value.

### HasTopReferrers

`func (o *ReferralBoard) HasTopReferrers() bool`

HasTopReferrers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


