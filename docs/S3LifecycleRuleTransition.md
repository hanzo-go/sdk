# S3LifecycleRuleTransition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int32** |  | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 

## Methods

### NewS3LifecycleRuleTransition

`func NewS3LifecycleRuleTransition() *S3LifecycleRuleTransition`

NewS3LifecycleRuleTransition instantiates a new S3LifecycleRuleTransition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LifecycleRuleTransitionWithDefaults

`func NewS3LifecycleRuleTransitionWithDefaults() *S3LifecycleRuleTransition`

NewS3LifecycleRuleTransitionWithDefaults instantiates a new S3LifecycleRuleTransition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *S3LifecycleRuleTransition) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *S3LifecycleRuleTransition) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *S3LifecycleRuleTransition) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *S3LifecycleRuleTransition) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetStorageClass

`func (o *S3LifecycleRuleTransition) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *S3LifecycleRuleTransition) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *S3LifecycleRuleTransition) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *S3LifecycleRuleTransition) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


