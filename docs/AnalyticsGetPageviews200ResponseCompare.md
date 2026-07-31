# AnalyticsGetPageviews200ResponseCompare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageviews** | Pointer to [**[]AnalyticsPageviewSeries**](AnalyticsPageviewSeries.md) |  | [optional] 
**Sessions** | Pointer to [**[]AnalyticsPageviewSeries**](AnalyticsPageviewSeries.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAnalyticsGetPageviews200ResponseCompare

`func NewAnalyticsGetPageviews200ResponseCompare() *AnalyticsGetPageviews200ResponseCompare`

NewAnalyticsGetPageviews200ResponseCompare instantiates a new AnalyticsGetPageviews200ResponseCompare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsGetPageviews200ResponseCompareWithDefaults

`func NewAnalyticsGetPageviews200ResponseCompareWithDefaults() *AnalyticsGetPageviews200ResponseCompare`

NewAnalyticsGetPageviews200ResponseCompareWithDefaults instantiates a new AnalyticsGetPageviews200ResponseCompare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageviews

`func (o *AnalyticsGetPageviews200ResponseCompare) GetPageviews() []AnalyticsPageviewSeries`

GetPageviews returns the Pageviews field if non-nil, zero value otherwise.

### GetPageviewsOk

`func (o *AnalyticsGetPageviews200ResponseCompare) GetPageviewsOk() (*[]AnalyticsPageviewSeries, bool)`

GetPageviewsOk returns a tuple with the Pageviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageviews

`func (o *AnalyticsGetPageviews200ResponseCompare) SetPageviews(v []AnalyticsPageviewSeries)`

SetPageviews sets Pageviews field to given value.

### HasPageviews

`func (o *AnalyticsGetPageviews200ResponseCompare) HasPageviews() bool`

HasPageviews returns a boolean if a field has been set.

### GetSessions

`func (o *AnalyticsGetPageviews200ResponseCompare) GetSessions() []AnalyticsPageviewSeries`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *AnalyticsGetPageviews200ResponseCompare) GetSessionsOk() (*[]AnalyticsPageviewSeries, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *AnalyticsGetPageviews200ResponseCompare) SetSessions(v []AnalyticsPageviewSeries)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *AnalyticsGetPageviews200ResponseCompare) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetStartDate

`func (o *AnalyticsGetPageviews200ResponseCompare) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AnalyticsGetPageviews200ResponseCompare) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AnalyticsGetPageviews200ResponseCompare) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AnalyticsGetPageviews200ResponseCompare) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AnalyticsGetPageviews200ResponseCompare) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AnalyticsGetPageviews200ResponseCompare) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AnalyticsGetPageviews200ResponseCompare) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AnalyticsGetPageviews200ResponseCompare) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


