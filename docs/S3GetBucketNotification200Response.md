# S3GetBucketNotification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configurations** | Pointer to [**[]S3EventConfig**](S3EventConfig.md) |  | [optional] 

## Methods

### NewS3GetBucketNotification200Response

`func NewS3GetBucketNotification200Response() *S3GetBucketNotification200Response`

NewS3GetBucketNotification200Response instantiates a new S3GetBucketNotification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3GetBucketNotification200ResponseWithDefaults

`func NewS3GetBucketNotification200ResponseWithDefaults() *S3GetBucketNotification200Response`

NewS3GetBucketNotification200ResponseWithDefaults instantiates a new S3GetBucketNotification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurations

`func (o *S3GetBucketNotification200Response) GetConfigurations() []S3EventConfig`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *S3GetBucketNotification200Response) GetConfigurationsOk() (*[]S3EventConfig, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *S3GetBucketNotification200Response) SetConfigurations(v []S3EventConfig)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *S3GetBucketNotification200Response) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


