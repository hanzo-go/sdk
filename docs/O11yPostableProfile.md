# O11yPostableProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingObservabilityTool** | Pointer to **string** |  | [optional] 
**ExistingObservabilityToolPresent** | Pointer to **bool** |  | [optional] 
**LogsScalePerDayInGb** | Pointer to **int32** |  | [optional] 
**NumberOfHosts** | Pointer to **int32** |  | [optional] 
**NumberOfServices** | Pointer to **int32** |  | [optional] 
**ReasonsForInterestInO11y** | Pointer to **[]string** |  | [optional] 
**TimelineForMigratingToO11y** | Pointer to **string** |  | [optional] 
**UsesOtel** | Pointer to **bool** |  | [optional] 
**WhereDidYouDiscoverO11y** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yPostableProfile

`func NewO11yPostableProfile() *O11yPostableProfile`

NewO11yPostableProfile instantiates a new O11yPostableProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPostableProfileWithDefaults

`func NewO11yPostableProfileWithDefaults() *O11yPostableProfile`

NewO11yPostableProfileWithDefaults instantiates a new O11yPostableProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingObservabilityTool

`func (o *O11yPostableProfile) GetExistingObservabilityTool() string`

GetExistingObservabilityTool returns the ExistingObservabilityTool field if non-nil, zero value otherwise.

### GetExistingObservabilityToolOk

`func (o *O11yPostableProfile) GetExistingObservabilityToolOk() (*string, bool)`

GetExistingObservabilityToolOk returns a tuple with the ExistingObservabilityTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingObservabilityTool

`func (o *O11yPostableProfile) SetExistingObservabilityTool(v string)`

SetExistingObservabilityTool sets ExistingObservabilityTool field to given value.

### HasExistingObservabilityTool

`func (o *O11yPostableProfile) HasExistingObservabilityTool() bool`

HasExistingObservabilityTool returns a boolean if a field has been set.

### GetExistingObservabilityToolPresent

`func (o *O11yPostableProfile) GetExistingObservabilityToolPresent() bool`

GetExistingObservabilityToolPresent returns the ExistingObservabilityToolPresent field if non-nil, zero value otherwise.

### GetExistingObservabilityToolPresentOk

`func (o *O11yPostableProfile) GetExistingObservabilityToolPresentOk() (*bool, bool)`

GetExistingObservabilityToolPresentOk returns a tuple with the ExistingObservabilityToolPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingObservabilityToolPresent

`func (o *O11yPostableProfile) SetExistingObservabilityToolPresent(v bool)`

SetExistingObservabilityToolPresent sets ExistingObservabilityToolPresent field to given value.

### HasExistingObservabilityToolPresent

`func (o *O11yPostableProfile) HasExistingObservabilityToolPresent() bool`

HasExistingObservabilityToolPresent returns a boolean if a field has been set.

### GetLogsScalePerDayInGb

`func (o *O11yPostableProfile) GetLogsScalePerDayInGb() int32`

GetLogsScalePerDayInGb returns the LogsScalePerDayInGb field if non-nil, zero value otherwise.

### GetLogsScalePerDayInGbOk

`func (o *O11yPostableProfile) GetLogsScalePerDayInGbOk() (*int32, bool)`

GetLogsScalePerDayInGbOk returns a tuple with the LogsScalePerDayInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsScalePerDayInGb

`func (o *O11yPostableProfile) SetLogsScalePerDayInGb(v int32)`

SetLogsScalePerDayInGb sets LogsScalePerDayInGb field to given value.

### HasLogsScalePerDayInGb

`func (o *O11yPostableProfile) HasLogsScalePerDayInGb() bool`

HasLogsScalePerDayInGb returns a boolean if a field has been set.

### GetNumberOfHosts

`func (o *O11yPostableProfile) GetNumberOfHosts() int32`

GetNumberOfHosts returns the NumberOfHosts field if non-nil, zero value otherwise.

### GetNumberOfHostsOk

`func (o *O11yPostableProfile) GetNumberOfHostsOk() (*int32, bool)`

GetNumberOfHostsOk returns a tuple with the NumberOfHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfHosts

`func (o *O11yPostableProfile) SetNumberOfHosts(v int32)`

SetNumberOfHosts sets NumberOfHosts field to given value.

### HasNumberOfHosts

`func (o *O11yPostableProfile) HasNumberOfHosts() bool`

HasNumberOfHosts returns a boolean if a field has been set.

### GetNumberOfServices

`func (o *O11yPostableProfile) GetNumberOfServices() int32`

GetNumberOfServices returns the NumberOfServices field if non-nil, zero value otherwise.

### GetNumberOfServicesOk

`func (o *O11yPostableProfile) GetNumberOfServicesOk() (*int32, bool)`

GetNumberOfServicesOk returns a tuple with the NumberOfServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfServices

`func (o *O11yPostableProfile) SetNumberOfServices(v int32)`

SetNumberOfServices sets NumberOfServices field to given value.

### HasNumberOfServices

`func (o *O11yPostableProfile) HasNumberOfServices() bool`

HasNumberOfServices returns a boolean if a field has been set.

### GetReasonsForInterestInO11y

`func (o *O11yPostableProfile) GetReasonsForInterestInO11y() []string`

GetReasonsForInterestInO11y returns the ReasonsForInterestInO11y field if non-nil, zero value otherwise.

### GetReasonsForInterestInO11yOk

`func (o *O11yPostableProfile) GetReasonsForInterestInO11yOk() (*[]string, bool)`

GetReasonsForInterestInO11yOk returns a tuple with the ReasonsForInterestInO11y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonsForInterestInO11y

`func (o *O11yPostableProfile) SetReasonsForInterestInO11y(v []string)`

SetReasonsForInterestInO11y sets ReasonsForInterestInO11y field to given value.

### HasReasonsForInterestInO11y

`func (o *O11yPostableProfile) HasReasonsForInterestInO11y() bool`

HasReasonsForInterestInO11y returns a boolean if a field has been set.

### GetTimelineForMigratingToO11y

`func (o *O11yPostableProfile) GetTimelineForMigratingToO11y() string`

GetTimelineForMigratingToO11y returns the TimelineForMigratingToO11y field if non-nil, zero value otherwise.

### GetTimelineForMigratingToO11yOk

`func (o *O11yPostableProfile) GetTimelineForMigratingToO11yOk() (*string, bool)`

GetTimelineForMigratingToO11yOk returns a tuple with the TimelineForMigratingToO11y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineForMigratingToO11y

`func (o *O11yPostableProfile) SetTimelineForMigratingToO11y(v string)`

SetTimelineForMigratingToO11y sets TimelineForMigratingToO11y field to given value.

### HasTimelineForMigratingToO11y

`func (o *O11yPostableProfile) HasTimelineForMigratingToO11y() bool`

HasTimelineForMigratingToO11y returns a boolean if a field has been set.

### GetUsesOtel

`func (o *O11yPostableProfile) GetUsesOtel() bool`

GetUsesOtel returns the UsesOtel field if non-nil, zero value otherwise.

### GetUsesOtelOk

`func (o *O11yPostableProfile) GetUsesOtelOk() (*bool, bool)`

GetUsesOtelOk returns a tuple with the UsesOtel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesOtel

`func (o *O11yPostableProfile) SetUsesOtel(v bool)`

SetUsesOtel sets UsesOtel field to given value.

### HasUsesOtel

`func (o *O11yPostableProfile) HasUsesOtel() bool`

HasUsesOtel returns a boolean if a field has been set.

### GetWhereDidYouDiscoverO11y

`func (o *O11yPostableProfile) GetWhereDidYouDiscoverO11y() string`

GetWhereDidYouDiscoverO11y returns the WhereDidYouDiscoverO11y field if non-nil, zero value otherwise.

### GetWhereDidYouDiscoverO11yOk

`func (o *O11yPostableProfile) GetWhereDidYouDiscoverO11yOk() (*string, bool)`

GetWhereDidYouDiscoverO11yOk returns a tuple with the WhereDidYouDiscoverO11y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhereDidYouDiscoverO11y

`func (o *O11yPostableProfile) SetWhereDidYouDiscoverO11y(v string)`

SetWhereDidYouDiscoverO11y sets WhereDidYouDiscoverO11y field to given value.

### HasWhereDidYouDiscoverO11y

`func (o *O11yPostableProfile) HasWhereDidYouDiscoverO11y() bool`

HasWhereDidYouDiscoverO11y returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


