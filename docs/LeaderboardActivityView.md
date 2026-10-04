# LeaderboardActivityView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when nothing could be read: the warehouse is not connected, the rollup is not ready, or the subject is one the ledger cannot attribute (see Note). Days is then empty because there is no answer, not because there was no activity. | [optional] 
**Days** | Pointer to [**[]LeaderboardActivityPoint**](LeaderboardActivityPoint.md) | Days is the gap-filled series, one point per calendar day from From up to (not including) To, in ascending order, zero-valued days included. Always a list, never null. | [optional] 
**From** | Pointer to **string** | From is the first day in Days, \&quot;2006-01-02\&quot; inclusive. | [optional] 
**Id** | Pointer to **string** | ID is the subject the server actually read, after resolving \&quot;me\&quot;/empty to the caller and bounding it to what they may see — a ledger \&quot;owner/name\&quot; for a user, an org id for an org. Echoed so a client can confirm whose series it holds. | [optional] 
**Note** | Pointer to **string** | Note explains an empty-but-not-broken answer in plain words — today only subject&#x3D;project, which the usage ledger records no column for. Present only when there is something to say; show it instead of an empty chart. | [optional] 
**Source** | Pointer to **string** | Source names the table the series was aggregated from (the derived daily rollup, hanzo.usage_rollup_daily). | [optional] 
**Subject** | Pointer to **string** | Subject echoes what the series is about: user|org|project. | [optional] 
**To** | Pointer to **string** | To is the EXCLUSIVE upper bound, \&quot;2006-01-02\&quot; — the day AFTER the last point in Days. A request for to&#x3D;2026-03-31 answers to&#x3D;2026-04-01 with 2026-03-31 last. | [optional] 
**Totals** | Pointer to [**LeaderboardActivityTotals**](LeaderboardActivityTotals.md) | Totals are the window&#39;s sums plus the busiest-day ceilings a heatmap scales against. Derived from Days — nothing here is read separately. | [optional] 

## Methods

### NewLeaderboardActivityView

`func NewLeaderboardActivityView() *LeaderboardActivityView`

NewLeaderboardActivityView instantiates a new LeaderboardActivityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardActivityViewWithDefaults

`func NewLeaderboardActivityViewWithDefaults() *LeaderboardActivityView`

NewLeaderboardActivityViewWithDefaults instantiates a new LeaderboardActivityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LeaderboardActivityView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LeaderboardActivityView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LeaderboardActivityView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LeaderboardActivityView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetDays

`func (o *LeaderboardActivityView) GetDays() []LeaderboardActivityPoint`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *LeaderboardActivityView) GetDaysOk() (*[]LeaderboardActivityPoint, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *LeaderboardActivityView) SetDays(v []LeaderboardActivityPoint)`

SetDays sets Days field to given value.

### HasDays

`func (o *LeaderboardActivityView) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetFrom

`func (o *LeaderboardActivityView) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LeaderboardActivityView) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LeaderboardActivityView) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *LeaderboardActivityView) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *LeaderboardActivityView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeaderboardActivityView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeaderboardActivityView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LeaderboardActivityView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNote

`func (o *LeaderboardActivityView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LeaderboardActivityView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LeaderboardActivityView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *LeaderboardActivityView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetSource

`func (o *LeaderboardActivityView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LeaderboardActivityView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LeaderboardActivityView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LeaderboardActivityView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubject

`func (o *LeaderboardActivityView) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LeaderboardActivityView) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LeaderboardActivityView) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *LeaderboardActivityView) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTo

`func (o *LeaderboardActivityView) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LeaderboardActivityView) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LeaderboardActivityView) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *LeaderboardActivityView) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotals

`func (o *LeaderboardActivityView) GetTotals() LeaderboardActivityTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *LeaderboardActivityView) GetTotalsOk() (*LeaderboardActivityTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *LeaderboardActivityView) SetTotals(v LeaderboardActivityTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *LeaderboardActivityView) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


