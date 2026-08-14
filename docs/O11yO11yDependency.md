# O11yO11yDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallCount** | Pointer to **int32** | CallCount is how many calls crossed the edge in the window. | [optional] 
**CallRate** | Pointer to **float32** | CallRate is calls per second. | [optional] 
**Child** | Pointer to **string** | Child is the called service. | [optional] 
**ErrorRate** | Pointer to **float32** | ErrorRate is the percentage of calls that erred. | [optional] 
**P50** | Pointer to **float32** | P50 is the median call duration, in nanoseconds. | [optional] 
**P75** | Pointer to **float32** | P75 is the 75th-percentile call duration, in nanoseconds. | [optional] 
**P90** | Pointer to **float32** | P90 is the 90th-percentile call duration, in nanoseconds. | [optional] 
**P95** | Pointer to **float32** | P95 is the 95th-percentile call duration, in nanoseconds. | [optional] 
**P99** | Pointer to **float32** | P99 is the 99th-percentile call duration, in nanoseconds. | [optional] 
**Parent** | Pointer to **string** | Parent is the calling service. | [optional] 

## Methods

### NewO11yO11yDependency

`func NewO11yO11yDependency() *O11yO11yDependency`

NewO11yO11yDependency instantiates a new O11yO11yDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDependencyWithDefaults

`func NewO11yO11yDependencyWithDefaults() *O11yO11yDependency`

NewO11yO11yDependencyWithDefaults instantiates a new O11yO11yDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallCount

`func (o *O11yO11yDependency) GetCallCount() int32`

GetCallCount returns the CallCount field if non-nil, zero value otherwise.

### GetCallCountOk

`func (o *O11yO11yDependency) GetCallCountOk() (*int32, bool)`

GetCallCountOk returns a tuple with the CallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallCount

`func (o *O11yO11yDependency) SetCallCount(v int32)`

SetCallCount sets CallCount field to given value.

### HasCallCount

`func (o *O11yO11yDependency) HasCallCount() bool`

HasCallCount returns a boolean if a field has been set.

### GetCallRate

`func (o *O11yO11yDependency) GetCallRate() float32`

GetCallRate returns the CallRate field if non-nil, zero value otherwise.

### GetCallRateOk

`func (o *O11yO11yDependency) GetCallRateOk() (*float32, bool)`

GetCallRateOk returns a tuple with the CallRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRate

`func (o *O11yO11yDependency) SetCallRate(v float32)`

SetCallRate sets CallRate field to given value.

### HasCallRate

`func (o *O11yO11yDependency) HasCallRate() bool`

HasCallRate returns a boolean if a field has been set.

### GetChild

`func (o *O11yO11yDependency) GetChild() string`

GetChild returns the Child field if non-nil, zero value otherwise.

### GetChildOk

`func (o *O11yO11yDependency) GetChildOk() (*string, bool)`

GetChildOk returns a tuple with the Child field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChild

`func (o *O11yO11yDependency) SetChild(v string)`

SetChild sets Child field to given value.

### HasChild

`func (o *O11yO11yDependency) HasChild() bool`

HasChild returns a boolean if a field has been set.

### GetErrorRate

`func (o *O11yO11yDependency) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *O11yO11yDependency) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *O11yO11yDependency) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *O11yO11yDependency) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetP50

`func (o *O11yO11yDependency) GetP50() float32`

GetP50 returns the P50 field if non-nil, zero value otherwise.

### GetP50Ok

`func (o *O11yO11yDependency) GetP50Ok() (*float32, bool)`

GetP50Ok returns a tuple with the P50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50

`func (o *O11yO11yDependency) SetP50(v float32)`

SetP50 sets P50 field to given value.

### HasP50

`func (o *O11yO11yDependency) HasP50() bool`

HasP50 returns a boolean if a field has been set.

### GetP75

`func (o *O11yO11yDependency) GetP75() float32`

GetP75 returns the P75 field if non-nil, zero value otherwise.

### GetP75Ok

`func (o *O11yO11yDependency) GetP75Ok() (*float32, bool)`

GetP75Ok returns a tuple with the P75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP75

`func (o *O11yO11yDependency) SetP75(v float32)`

SetP75 sets P75 field to given value.

### HasP75

`func (o *O11yO11yDependency) HasP75() bool`

HasP75 returns a boolean if a field has been set.

### GetP90

`func (o *O11yO11yDependency) GetP90() float32`

GetP90 returns the P90 field if non-nil, zero value otherwise.

### GetP90Ok

`func (o *O11yO11yDependency) GetP90Ok() (*float32, bool)`

GetP90Ok returns a tuple with the P90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP90

`func (o *O11yO11yDependency) SetP90(v float32)`

SetP90 sets P90 field to given value.

### HasP90

`func (o *O11yO11yDependency) HasP90() bool`

HasP90 returns a boolean if a field has been set.

### GetP95

`func (o *O11yO11yDependency) GetP95() float32`

GetP95 returns the P95 field if non-nil, zero value otherwise.

### GetP95Ok

`func (o *O11yO11yDependency) GetP95Ok() (*float32, bool)`

GetP95Ok returns a tuple with the P95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95

`func (o *O11yO11yDependency) SetP95(v float32)`

SetP95 sets P95 field to given value.

### HasP95

`func (o *O11yO11yDependency) HasP95() bool`

HasP95 returns a boolean if a field has been set.

### GetP99

`func (o *O11yO11yDependency) GetP99() float32`

GetP99 returns the P99 field if non-nil, zero value otherwise.

### GetP99Ok

`func (o *O11yO11yDependency) GetP99Ok() (*float32, bool)`

GetP99Ok returns a tuple with the P99 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99

`func (o *O11yO11yDependency) SetP99(v float32)`

SetP99 sets P99 field to given value.

### HasP99

`func (o *O11yO11yDependency) HasP99() bool`

HasP99 returns a boolean if a field has been set.

### GetParent

`func (o *O11yO11yDependency) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *O11yO11yDependency) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *O11yO11yDependency) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *O11yO11yDependency) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


