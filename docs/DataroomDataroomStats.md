# DataroomDataroomStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataroomId** | Pointer to **string** | DataroomId is the room these counts are for. | [optional] 
**Links** | Pointer to [**[]DataroomDataroomLinkStats**](DataroomDataroomLinkStats.md) | Links is the same per-page breakdown for each link into the room. | [optional] 
**TotalPageViews** | Pointer to **int64** | TotalPageViews is the room&#39;s page views across every link. | [optional] 
**TotalViews** | Pointer to **int64** | TotalViews is the room&#39;s viewing sessions across every link. | [optional] 

## Methods

### NewDataroomDataroomStats

`func NewDataroomDataroomStats() *DataroomDataroomStats`

NewDataroomDataroomStats instantiates a new DataroomDataroomStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomStatsWithDefaults

`func NewDataroomDataroomStatsWithDefaults() *DataroomDataroomStats`

NewDataroomDataroomStatsWithDefaults instantiates a new DataroomDataroomStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataroomId

`func (o *DataroomDataroomStats) GetDataroomId() string`

GetDataroomId returns the DataroomId field if non-nil, zero value otherwise.

### GetDataroomIdOk

`func (o *DataroomDataroomStats) GetDataroomIdOk() (*string, bool)`

GetDataroomIdOk returns a tuple with the DataroomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataroomId

`func (o *DataroomDataroomStats) SetDataroomId(v string)`

SetDataroomId sets DataroomId field to given value.

### HasDataroomId

`func (o *DataroomDataroomStats) HasDataroomId() bool`

HasDataroomId returns a boolean if a field has been set.

### GetLinks

`func (o *DataroomDataroomStats) GetLinks() []DataroomDataroomLinkStats`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DataroomDataroomStats) GetLinksOk() (*[]DataroomDataroomLinkStats, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DataroomDataroomStats) SetLinks(v []DataroomDataroomLinkStats)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DataroomDataroomStats) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalPageViews

`func (o *DataroomDataroomStats) GetTotalPageViews() int64`

GetTotalPageViews returns the TotalPageViews field if non-nil, zero value otherwise.

### GetTotalPageViewsOk

`func (o *DataroomDataroomStats) GetTotalPageViewsOk() (*int64, bool)`

GetTotalPageViewsOk returns a tuple with the TotalPageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPageViews

`func (o *DataroomDataroomStats) SetTotalPageViews(v int64)`

SetTotalPageViews sets TotalPageViews field to given value.

### HasTotalPageViews

`func (o *DataroomDataroomStats) HasTotalPageViews() bool`

HasTotalPageViews returns a boolean if a field has been set.

### GetTotalViews

`func (o *DataroomDataroomStats) GetTotalViews() int64`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *DataroomDataroomStats) GetTotalViewsOk() (*int64, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *DataroomDataroomStats) SetTotalViews(v int64)`

SetTotalViews sets TotalViews field to given value.

### HasTotalViews

`func (o *DataroomDataroomStats) HasTotalViews() bool`

HasTotalViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


