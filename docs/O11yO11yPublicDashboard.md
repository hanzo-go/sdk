# O11yO11yPublicDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTimeRange** | Pointer to **string** | DefaultTimeRange is the fixed window when the range is not caller-chosen. | [optional] 
**PublicPath** | Pointer to **string** | PublicPath is the public URL path the share is reachable at. | [optional] 
**TimeRangeEnabled** | Pointer to **bool** | TimeRangeEnabled reports whether the public page may pick its own range. | [optional] 

## Methods

### NewO11yO11yPublicDashboard

`func NewO11yO11yPublicDashboard() *O11yO11yPublicDashboard`

NewO11yO11yPublicDashboard instantiates a new O11yO11yPublicDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yPublicDashboardWithDefaults

`func NewO11yO11yPublicDashboardWithDefaults() *O11yO11yPublicDashboard`

NewO11yO11yPublicDashboardWithDefaults instantiates a new O11yO11yPublicDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTimeRange

`func (o *O11yO11yPublicDashboard) GetDefaultTimeRange() string`

GetDefaultTimeRange returns the DefaultTimeRange field if non-nil, zero value otherwise.

### GetDefaultTimeRangeOk

`func (o *O11yO11yPublicDashboard) GetDefaultTimeRangeOk() (*string, bool)`

GetDefaultTimeRangeOk returns a tuple with the DefaultTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeRange

`func (o *O11yO11yPublicDashboard) SetDefaultTimeRange(v string)`

SetDefaultTimeRange sets DefaultTimeRange field to given value.

### HasDefaultTimeRange

`func (o *O11yO11yPublicDashboard) HasDefaultTimeRange() bool`

HasDefaultTimeRange returns a boolean if a field has been set.

### GetPublicPath

`func (o *O11yO11yPublicDashboard) GetPublicPath() string`

GetPublicPath returns the PublicPath field if non-nil, zero value otherwise.

### GetPublicPathOk

`func (o *O11yO11yPublicDashboard) GetPublicPathOk() (*string, bool)`

GetPublicPathOk returns a tuple with the PublicPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPath

`func (o *O11yO11yPublicDashboard) SetPublicPath(v string)`

SetPublicPath sets PublicPath field to given value.

### HasPublicPath

`func (o *O11yO11yPublicDashboard) HasPublicPath() bool`

HasPublicPath returns a boolean if a field has been set.

### GetTimeRangeEnabled

`func (o *O11yO11yPublicDashboard) GetTimeRangeEnabled() bool`

GetTimeRangeEnabled returns the TimeRangeEnabled field if non-nil, zero value otherwise.

### GetTimeRangeEnabledOk

`func (o *O11yO11yPublicDashboard) GetTimeRangeEnabledOk() (*bool, bool)`

GetTimeRangeEnabledOk returns a tuple with the TimeRangeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeEnabled

`func (o *O11yO11yPublicDashboard) SetTimeRangeEnabled(v bool)`

SetTimeRangeEnabled sets TimeRangeEnabled field to given value.

### HasTimeRangeEnabled

`func (o *O11yO11yPublicDashboard) HasTimeRangeEnabled() bool`

HasTimeRangeEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


