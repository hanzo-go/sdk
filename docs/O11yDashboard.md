# O11yDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Panels** | Pointer to [**[]O11yPanel**](O11yPanel.md) |  | [optional] 
**Time** | Pointer to [**O11yDashboardTime**](O11yDashboardTime.md) |  | [optional] 
**Refresh** | Pointer to **string** | Auto-refresh interval (e.g. 30s, 1m, 5m) | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewO11yDashboard

`func NewO11yDashboard() *O11yDashboard`

NewO11yDashboard instantiates a new O11yDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDashboardWithDefaults

`func NewO11yDashboardWithDefaults() *O11yDashboard`

NewO11yDashboardWithDefaults instantiates a new O11yDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *O11yDashboard) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *O11yDashboard) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *O11yDashboard) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *O11yDashboard) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetTitle

`func (o *O11yDashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yDashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yDashboard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yDashboard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *O11yDashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yDashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yDashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yDashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *O11yDashboard) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yDashboard) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yDashboard) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yDashboard) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPanels

`func (o *O11yDashboard) GetPanels() []O11yPanel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *O11yDashboard) GetPanelsOk() (*[]O11yPanel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *O11yDashboard) SetPanels(v []O11yPanel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *O11yDashboard) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetTime

`func (o *O11yDashboard) GetTime() O11yDashboardTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *O11yDashboard) GetTimeOk() (*O11yDashboardTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *O11yDashboard) SetTime(v O11yDashboardTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *O11yDashboard) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRefresh

`func (o *O11yDashboard) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *O11yDashboard) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *O11yDashboard) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *O11yDashboard) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetVersion

`func (o *O11yDashboard) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *O11yDashboard) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *O11yDashboard) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *O11yDashboard) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yDashboard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yDashboard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yDashboard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yDashboard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yDashboard) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yDashboard) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yDashboard) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yDashboard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


