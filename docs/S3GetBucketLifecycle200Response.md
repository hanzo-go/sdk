# S3GetBucketLifecycle200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]S3LifecycleRule**](S3LifecycleRule.md) |  | [optional] 

## Methods

### NewS3GetBucketLifecycle200Response

`func NewS3GetBucketLifecycle200Response() *S3GetBucketLifecycle200Response`

NewS3GetBucketLifecycle200Response instantiates a new S3GetBucketLifecycle200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3GetBucketLifecycle200ResponseWithDefaults

`func NewS3GetBucketLifecycle200ResponseWithDefaults() *S3GetBucketLifecycle200Response`

NewS3GetBucketLifecycle200ResponseWithDefaults instantiates a new S3GetBucketLifecycle200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *S3GetBucketLifecycle200Response) GetRules() []S3LifecycleRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *S3GetBucketLifecycle200Response) GetRulesOk() (*[]S3LifecycleRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *S3GetBucketLifecycle200Response) SetRules(v []S3LifecycleRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *S3GetBucketLifecycle200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


