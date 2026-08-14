# DataroomPageStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgDuration** | Pointer to **int32** | AvgDuration is totalDuration divided by views, rounded; 0 when unviewed. | [optional] 
**PageNumber** | Pointer to **int32** | PageNumber is the page these counts are for. | [optional] 
**TotalDuration** | Pointer to **int32** | TotalDuration is the summed dwell measure reported for the page. | [optional] 
**Views** | Pointer to **int32** | Views is how many times the page was viewed. | [optional] 

## Methods

### NewDataroomPageStat

`func NewDataroomPageStat() *DataroomPageStat`

NewDataroomPageStat instantiates a new DataroomPageStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomPageStatWithDefaults

`func NewDataroomPageStatWithDefaults() *DataroomPageStat`

NewDataroomPageStatWithDefaults instantiates a new DataroomPageStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgDuration

`func (o *DataroomPageStat) GetAvgDuration() int32`

GetAvgDuration returns the AvgDuration field if non-nil, zero value otherwise.

### GetAvgDurationOk

`func (o *DataroomPageStat) GetAvgDurationOk() (*int32, bool)`

GetAvgDurationOk returns a tuple with the AvgDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDuration

`func (o *DataroomPageStat) SetAvgDuration(v int32)`

SetAvgDuration sets AvgDuration field to given value.

### HasAvgDuration

`func (o *DataroomPageStat) HasAvgDuration() bool`

HasAvgDuration returns a boolean if a field has been set.

### GetPageNumber

`func (o *DataroomPageStat) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DataroomPageStat) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DataroomPageStat) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DataroomPageStat) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetTotalDuration

`func (o *DataroomPageStat) GetTotalDuration() int32`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *DataroomPageStat) GetTotalDurationOk() (*int32, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *DataroomPageStat) SetTotalDuration(v int32)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *DataroomPageStat) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.

### GetViews

`func (o *DataroomPageStat) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *DataroomPageStat) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *DataroomPageStat) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *DataroomPageStat) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


