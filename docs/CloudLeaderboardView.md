# CloudLeaderboardView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** | tokens|requests|cost | [optional] 
**Period** | Pointer to **string** | day|week|month|all|custom | [optional] 
**Rows** | Pointer to [**[]CloudLeaderboardRow**](CloudLeaderboardRow.md) |  | [optional] 
**Scope** | Pointer to **string** | personal|org|global | [optional] 
**Self** | Pointer to [**CloudSelfRank**](CloudSelfRank.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **string** | \&quot;\&quot; for all | [optional] 
**Subject** | Pointer to **string** | user|org | [optional] 
**Total** | Pointer to **int32** | ranked subjects in the window | [optional] 

## Methods

### NewCloudLeaderboardView

`func NewCloudLeaderboardView() *CloudLeaderboardView`

NewCloudLeaderboardView instantiates a new CloudLeaderboardView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLeaderboardViewWithDefaults

`func NewCloudLeaderboardViewWithDefaults() *CloudLeaderboardView`

NewCloudLeaderboardViewWithDefaults instantiates a new CloudLeaderboardView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CloudLeaderboardView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CloudLeaderboardView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CloudLeaderboardView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CloudLeaderboardView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetEnd

`func (o *CloudLeaderboardView) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudLeaderboardView) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudLeaderboardView) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudLeaderboardView) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetMetric

`func (o *CloudLeaderboardView) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CloudLeaderboardView) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CloudLeaderboardView) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *CloudLeaderboardView) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPeriod

`func (o *CloudLeaderboardView) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CloudLeaderboardView) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CloudLeaderboardView) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CloudLeaderboardView) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRows

`func (o *CloudLeaderboardView) GetRows() []CloudLeaderboardRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CloudLeaderboardView) GetRowsOk() (*[]CloudLeaderboardRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CloudLeaderboardView) SetRows(v []CloudLeaderboardRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *CloudLeaderboardView) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetScope

`func (o *CloudLeaderboardView) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudLeaderboardView) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudLeaderboardView) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudLeaderboardView) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *CloudLeaderboardView) GetSelf() CloudSelfRank`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CloudLeaderboardView) GetSelfOk() (*CloudSelfRank, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CloudLeaderboardView) SetSelf(v CloudSelfRank)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CloudLeaderboardView) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSource

`func (o *CloudLeaderboardView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudLeaderboardView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudLeaderboardView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudLeaderboardView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStart

`func (o *CloudLeaderboardView) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudLeaderboardView) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudLeaderboardView) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudLeaderboardView) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSubject

`func (o *CloudLeaderboardView) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudLeaderboardView) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudLeaderboardView) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudLeaderboardView) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTotal

`func (o *CloudLeaderboardView) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudLeaderboardView) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudLeaderboardView) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudLeaderboardView) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


