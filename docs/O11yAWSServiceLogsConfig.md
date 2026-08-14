# O11yAWSServiceLogsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**S3Buckets** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewO11yAWSServiceLogsConfig

`func NewO11yAWSServiceLogsConfig() *O11yAWSServiceLogsConfig`

NewO11yAWSServiceLogsConfig instantiates a new O11yAWSServiceLogsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAWSServiceLogsConfigWithDefaults

`func NewO11yAWSServiceLogsConfigWithDefaults() *O11yAWSServiceLogsConfig`

NewO11yAWSServiceLogsConfigWithDefaults instantiates a new O11yAWSServiceLogsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *O11yAWSServiceLogsConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yAWSServiceLogsConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yAWSServiceLogsConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yAWSServiceLogsConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetS3Buckets

`func (o *O11yAWSServiceLogsConfig) GetS3Buckets() map[string][]string`

GetS3Buckets returns the S3Buckets field if non-nil, zero value otherwise.

### GetS3BucketsOk

`func (o *O11yAWSServiceLogsConfig) GetS3BucketsOk() (*map[string][]string, bool)`

GetS3BucketsOk returns a tuple with the S3Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Buckets

`func (o *O11yAWSServiceLogsConfig) SetS3Buckets(v map[string][]string)`

SetS3Buckets sets S3Buckets field to given value.

### HasS3Buckets

`func (o *O11yAWSServiceLogsConfig) HasS3Buckets() bool`

HasS3Buckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


