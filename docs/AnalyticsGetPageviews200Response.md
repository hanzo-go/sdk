# AnalyticsGetPageviews200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageviews** | Pointer to [**[]AnalyticsPageviewSeries**](AnalyticsPageviewSeries.md) |  | [optional] 
**Sessions** | Pointer to [**[]AnalyticsPageviewSeries**](AnalyticsPageviewSeries.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Compare** | Pointer to [**AnalyticsGetPageviews200ResponseCompare**](AnalyticsGetPageviews200ResponseCompare.md) |  | [optional] 

## Methods

### NewAnalyticsGetPageviews200Response

`func NewAnalyticsGetPageviews200Response() *AnalyticsGetPageviews200Response`

NewAnalyticsGetPageviews200Response instantiates a new AnalyticsGetPageviews200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsGetPageviews200ResponseWithDefaults

`func NewAnalyticsGetPageviews200ResponseWithDefaults() *AnalyticsGetPageviews200Response`

NewAnalyticsGetPageviews200ResponseWithDefaults instantiates a new AnalyticsGetPageviews200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageviews

`func (o *AnalyticsGetPageviews200Response) GetPageviews() []AnalyticsPageviewSeries`

GetPageviews returns the Pageviews field if non-nil, zero value otherwise.

### GetPageviewsOk

`func (o *AnalyticsGetPageviews200Response) GetPageviewsOk() (*[]AnalyticsPageviewSeries, bool)`

GetPageviewsOk returns a tuple with the Pageviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageviews

`func (o *AnalyticsGetPageviews200Response) SetPageviews(v []AnalyticsPageviewSeries)`

SetPageviews sets Pageviews field to given value.

### HasPageviews

`func (o *AnalyticsGetPageviews200Response) HasPageviews() bool`

HasPageviews returns a boolean if a field has been set.

### GetSessions

`func (o *AnalyticsGetPageviews200Response) GetSessions() []AnalyticsPageviewSeries`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *AnalyticsGetPageviews200Response) GetSessionsOk() (*[]AnalyticsPageviewSeries, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *AnalyticsGetPageviews200Response) SetSessions(v []AnalyticsPageviewSeries)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *AnalyticsGetPageviews200Response) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetStartDate

`func (o *AnalyticsGetPageviews200Response) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AnalyticsGetPageviews200Response) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AnalyticsGetPageviews200Response) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AnalyticsGetPageviews200Response) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AnalyticsGetPageviews200Response) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AnalyticsGetPageviews200Response) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AnalyticsGetPageviews200Response) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AnalyticsGetPageviews200Response) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCompare

`func (o *AnalyticsGetPageviews200Response) GetCompare() AnalyticsGetPageviews200ResponseCompare`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *AnalyticsGetPageviews200Response) GetCompareOk() (*AnalyticsGetPageviews200ResponseCompare, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *AnalyticsGetPageviews200Response) SetCompare(v AnalyticsGetPageviews200ResponseCompare)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *AnalyticsGetPageviews200Response) HasCompare() bool`

HasCompare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


