# S3LifecycleRuleExpiration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 

## Methods

### NewS3LifecycleRuleExpiration

`func NewS3LifecycleRuleExpiration() *S3LifecycleRuleExpiration`

NewS3LifecycleRuleExpiration instantiates a new S3LifecycleRuleExpiration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LifecycleRuleExpirationWithDefaults

`func NewS3LifecycleRuleExpirationWithDefaults() *S3LifecycleRuleExpiration`

NewS3LifecycleRuleExpirationWithDefaults instantiates a new S3LifecycleRuleExpiration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *S3LifecycleRuleExpiration) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *S3LifecycleRuleExpiration) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *S3LifecycleRuleExpiration) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *S3LifecycleRuleExpiration) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDate

`func (o *S3LifecycleRuleExpiration) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *S3LifecycleRuleExpiration) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *S3LifecycleRuleExpiration) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *S3LifecycleRuleExpiration) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


