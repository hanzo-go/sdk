# O11yDashboardCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Panels** | Pointer to [**[]O11yPanel**](O11yPanel.md) |  | [optional] 
**Time** | Pointer to [**O11yDashboardCreateTime**](O11yDashboardCreateTime.md) |  | [optional] 
**Refresh** | Pointer to **string** |  | [optional] [default to "30s"]

## Methods

### NewO11yDashboardCreate

`func NewO11yDashboardCreate(title string, ) *O11yDashboardCreate`

NewO11yDashboardCreate instantiates a new O11yDashboardCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDashboardCreateWithDefaults

`func NewO11yDashboardCreateWithDefaults() *O11yDashboardCreate`

NewO11yDashboardCreateWithDefaults instantiates a new O11yDashboardCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *O11yDashboardCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yDashboardCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yDashboardCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *O11yDashboardCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yDashboardCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yDashboardCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yDashboardCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *O11yDashboardCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yDashboardCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yDashboardCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yDashboardCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPanels

`func (o *O11yDashboardCreate) GetPanels() []O11yPanel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *O11yDashboardCreate) GetPanelsOk() (*[]O11yPanel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *O11yDashboardCreate) SetPanels(v []O11yPanel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *O11yDashboardCreate) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetTime

`func (o *O11yDashboardCreate) GetTime() O11yDashboardCreateTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *O11yDashboardCreate) GetTimeOk() (*O11yDashboardCreateTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *O11yDashboardCreate) SetTime(v O11yDashboardCreateTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *O11yDashboardCreate) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRefresh

`func (o *O11yDashboardCreate) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *O11yDashboardCreate) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *O11yDashboardCreate) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *O11yDashboardCreate) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


