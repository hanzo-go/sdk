# PlatformHealthTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **int64** | Degraded is how many CD reports Degraded. | [optional] 
**Healthy** | Pointer to **int64** | Healthy is how many CD reports Healthy. | [optional] 
**Missing** | Pointer to **int64** | Missing is how many have no live objects at all. | [optional] 
**Progressing** | Pointer to **int64** | Progressing is how many are still rolling out. | [optional] 
**Unknown** | Pointer to **int64** | Unknown is every other verdict, and apps CD could not be asked about. | [optional] 

## Methods

### NewPlatformHealthTally

`func NewPlatformHealthTally() *PlatformHealthTally`

NewPlatformHealthTally instantiates a new PlatformHealthTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformHealthTallyWithDefaults

`func NewPlatformHealthTallyWithDefaults() *PlatformHealthTally`

NewPlatformHealthTallyWithDefaults instantiates a new PlatformHealthTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *PlatformHealthTally) GetDegraded() int64`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *PlatformHealthTally) GetDegradedOk() (*int64, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *PlatformHealthTally) SetDegraded(v int64)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *PlatformHealthTally) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetHealthy

`func (o *PlatformHealthTally) GetHealthy() int64`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *PlatformHealthTally) GetHealthyOk() (*int64, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *PlatformHealthTally) SetHealthy(v int64)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *PlatformHealthTally) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetMissing

`func (o *PlatformHealthTally) GetMissing() int64`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *PlatformHealthTally) GetMissingOk() (*int64, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *PlatformHealthTally) SetMissing(v int64)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *PlatformHealthTally) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetProgressing

`func (o *PlatformHealthTally) GetProgressing() int64`

GetProgressing returns the Progressing field if non-nil, zero value otherwise.

### GetProgressingOk

`func (o *PlatformHealthTally) GetProgressingOk() (*int64, bool)`

GetProgressingOk returns a tuple with the Progressing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressing

`func (o *PlatformHealthTally) SetProgressing(v int64)`

SetProgressing sets Progressing field to given value.

### HasProgressing

`func (o *PlatformHealthTally) HasProgressing() bool`

HasProgressing returns a boolean if a field has been set.

### GetUnknown

`func (o *PlatformHealthTally) GetUnknown() int64`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *PlatformHealthTally) GetUnknownOk() (*int64, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *PlatformHealthTally) SetUnknown(v int64)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *PlatformHealthTally) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


