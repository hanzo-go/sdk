# DataroomDataroomLinkStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | Pointer to **string** | LinkId is the link these counts are for. | [optional] 
**Pages** | Pointer to [**[]DataroomDataroomPageStat**](DataroomDataroomPageStat.md) | Pages is the per-page breakdown, in page order. | [optional] 
**TotalPageViews** | Pointer to **int64** | TotalPageViews is how many page views the link received. | [optional] 
**TotalViews** | Pointer to **int64** | TotalViews is how many viewing sessions the link opened. | [optional] 

## Methods

### NewDataroomDataroomLinkStats

`func NewDataroomDataroomLinkStats() *DataroomDataroomLinkStats`

NewDataroomDataroomLinkStats instantiates a new DataroomDataroomLinkStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomLinkStatsWithDefaults

`func NewDataroomDataroomLinkStatsWithDefaults() *DataroomDataroomLinkStats`

NewDataroomDataroomLinkStatsWithDefaults instantiates a new DataroomDataroomLinkStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *DataroomDataroomLinkStats) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *DataroomDataroomLinkStats) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *DataroomDataroomLinkStats) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *DataroomDataroomLinkStats) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetPages

`func (o *DataroomDataroomLinkStats) GetPages() []DataroomDataroomPageStat`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *DataroomDataroomLinkStats) GetPagesOk() (*[]DataroomDataroomPageStat, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *DataroomDataroomLinkStats) SetPages(v []DataroomDataroomPageStat)`

SetPages sets Pages field to given value.

### HasPages

`func (o *DataroomDataroomLinkStats) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotalPageViews

`func (o *DataroomDataroomLinkStats) GetTotalPageViews() int64`

GetTotalPageViews returns the TotalPageViews field if non-nil, zero value otherwise.

### GetTotalPageViewsOk

`func (o *DataroomDataroomLinkStats) GetTotalPageViewsOk() (*int64, bool)`

GetTotalPageViewsOk returns a tuple with the TotalPageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPageViews

`func (o *DataroomDataroomLinkStats) SetTotalPageViews(v int64)`

SetTotalPageViews sets TotalPageViews field to given value.

### HasTotalPageViews

`func (o *DataroomDataroomLinkStats) HasTotalPageViews() bool`

HasTotalPageViews returns a boolean if a field has been set.

### GetTotalViews

`func (o *DataroomDataroomLinkStats) GetTotalViews() int64`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *DataroomDataroomLinkStats) GetTotalViewsOk() (*int64, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *DataroomDataroomLinkStats) SetTotalViews(v int64)`

SetTotalViews sets TotalViews field to given value.

### HasTotalViews

`func (o *DataroomDataroomLinkStats) HasTotalViews() bool`

HasTotalViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


