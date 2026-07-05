# StreamTopicConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionMs** | Pointer to **int32** | Message retention in milliseconds | [optional] [default to -1]
**CleanupPolicy** | Pointer to **string** |  | [optional] [default to "delete"]
**CompressionType** | Pointer to **string** |  | [optional] [default to "none"]

## Methods

### NewStreamTopicConfig

`func NewStreamTopicConfig() *StreamTopicConfig`

NewStreamTopicConfig instantiates a new StreamTopicConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamTopicConfigWithDefaults

`func NewStreamTopicConfigWithDefaults() *StreamTopicConfig`

NewStreamTopicConfigWithDefaults instantiates a new StreamTopicConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionMs

`func (o *StreamTopicConfig) GetRetentionMs() int32`

GetRetentionMs returns the RetentionMs field if non-nil, zero value otherwise.

### GetRetentionMsOk

`func (o *StreamTopicConfig) GetRetentionMsOk() (*int32, bool)`

GetRetentionMsOk returns a tuple with the RetentionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionMs

`func (o *StreamTopicConfig) SetRetentionMs(v int32)`

SetRetentionMs sets RetentionMs field to given value.

### HasRetentionMs

`func (o *StreamTopicConfig) HasRetentionMs() bool`

HasRetentionMs returns a boolean if a field has been set.

### GetCleanupPolicy

`func (o *StreamTopicConfig) GetCleanupPolicy() string`

GetCleanupPolicy returns the CleanupPolicy field if non-nil, zero value otherwise.

### GetCleanupPolicyOk

`func (o *StreamTopicConfig) GetCleanupPolicyOk() (*string, bool)`

GetCleanupPolicyOk returns a tuple with the CleanupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupPolicy

`func (o *StreamTopicConfig) SetCleanupPolicy(v string)`

SetCleanupPolicy sets CleanupPolicy field to given value.

### HasCleanupPolicy

`func (o *StreamTopicConfig) HasCleanupPolicy() bool`

HasCleanupPolicy returns a boolean if a field has been set.

### GetCompressionType

`func (o *StreamTopicConfig) GetCompressionType() string`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *StreamTopicConfig) GetCompressionTypeOk() (*string, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *StreamTopicConfig) SetCompressionType(v string)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *StreamTopicConfig) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


