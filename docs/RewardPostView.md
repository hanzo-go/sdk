# RewardPostView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**[]RewardCheck**](RewardCheck.md) | Checks is every look at the post, oldest first. | [optional] 
**CountsAt** | Pointer to **time.Time** |  | [optional] 
**DaysLeft** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** | Note says why a post was removed. | [optional] 
**Since** | Pointer to **time.Time** | Since is when the post&#39;s week began; CountsAt when it counts, and DaysLeft the days until then. Both are absent unless the post is pending. | [optional] 
**Status** | Pointer to **string** | Status is pending (held for its week, counts for nothing), verified (it stayed up and was put to staff), removed (it earns nothing) or review (staff decide it by hand). | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Week** | Pointer to **string** | Week is the ISO week the post was published in; a post in a later week can earn again. | [optional] 

## Methods

### NewRewardPostView

`func NewRewardPostView() *RewardPostView`

NewRewardPostView instantiates a new RewardPostView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardPostViewWithDefaults

`func NewRewardPostViewWithDefaults() *RewardPostView`

NewRewardPostViewWithDefaults instantiates a new RewardPostView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *RewardPostView) GetChecks() []RewardCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *RewardPostView) GetChecksOk() (*[]RewardCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *RewardPostView) SetChecks(v []RewardCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *RewardPostView) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCountsAt

`func (o *RewardPostView) GetCountsAt() time.Time`

GetCountsAt returns the CountsAt field if non-nil, zero value otherwise.

### GetCountsAtOk

`func (o *RewardPostView) GetCountsAtOk() (*time.Time, bool)`

GetCountsAtOk returns a tuple with the CountsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsAt

`func (o *RewardPostView) SetCountsAt(v time.Time)`

SetCountsAt sets CountsAt field to given value.

### HasCountsAt

`func (o *RewardPostView) HasCountsAt() bool`

HasCountsAt returns a boolean if a field has been set.

### GetDaysLeft

`func (o *RewardPostView) GetDaysLeft() int64`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *RewardPostView) GetDaysLeftOk() (*int64, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *RewardPostView) SetDaysLeft(v int64)`

SetDaysLeft sets DaysLeft field to given value.

### HasDaysLeft

`func (o *RewardPostView) HasDaysLeft() bool`

HasDaysLeft returns a boolean if a field has been set.

### GetId

`func (o *RewardPostView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardPostView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardPostView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RewardPostView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetwork

`func (o *RewardPostView) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RewardPostView) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RewardPostView) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RewardPostView) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNote

`func (o *RewardPostView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RewardPostView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RewardPostView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RewardPostView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetSince

`func (o *RewardPostView) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *RewardPostView) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *RewardPostView) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *RewardPostView) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetStatus

`func (o *RewardPostView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RewardPostView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RewardPostView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RewardPostView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *RewardPostView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RewardPostView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RewardPostView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RewardPostView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWeek

`func (o *RewardPostView) GetWeek() string`

GetWeek returns the Week field if non-nil, zero value otherwise.

### GetWeekOk

`func (o *RewardPostView) GetWeekOk() (*string, bool)`

GetWeekOk returns a tuple with the Week field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeek

`func (o *RewardPostView) SetWeek(v string)`

SetWeek sets Week field to given value.

### HasWeek

`func (o *RewardPostView) HasWeek() bool`

HasWeek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


