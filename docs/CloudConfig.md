# CloudConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAge** | Pointer to **string** | MaxAge caps message age, e.g. \&quot;24h\&quot; or \&quot;7d\&quot;; \&quot;0\&quot; (default) is unlimited. | [optional] 
**MaxBytes** | Pointer to **int32** | MaxBytes caps the stream&#39;s total stored bytes; -1 (default) is unlimited. | [optional] 
**MaxMsgSize** | Pointer to **int32** | MaxMsgSize caps one message&#39;s size in bytes; -1 (default) is the broker&#39;s limit. | [optional] 
**MaxMsgs** | Pointer to **int32** | MaxMsgs caps the number of stored messages; -1 (default) is unlimited. | [optional] 
**Name** | Pointer to **string** | Name is the stream name, unique within the org (alphanumeric, hyphens, underscores). | [optional] 
**NumReplicas** | Pointer to **int32** | Replicas is the number of stream replicas (1–5); this plane runs 1. | [optional] 
**Retention** | Pointer to **string** | Retention is the retention policy: limits (default), interest, or workqueue. | [optional] 
**Storage** | Pointer to **string** | Storage is the storage backend: file (default) or memory. | [optional] 
**Subjects** | Pointer to **[]string** | Subjects are the org-relative subjects bound to this stream (wildcards supported). Default: the stream name. | [optional] 

## Methods

### NewCloudConfig

`func NewCloudConfig() *CloudConfig`

NewCloudConfig instantiates a new CloudConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConfigWithDefaults

`func NewCloudConfigWithDefaults() *CloudConfig`

NewCloudConfigWithDefaults instantiates a new CloudConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAge

`func (o *CloudConfig) GetMaxAge() string`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CloudConfig) GetMaxAgeOk() (*string, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CloudConfig) SetMaxAge(v string)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CloudConfig) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxBytes

`func (o *CloudConfig) GetMaxBytes() int32`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *CloudConfig) GetMaxBytesOk() (*int32, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *CloudConfig) SetMaxBytes(v int32)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *CloudConfig) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxMsgSize

`func (o *CloudConfig) GetMaxMsgSize() int32`

GetMaxMsgSize returns the MaxMsgSize field if non-nil, zero value otherwise.

### GetMaxMsgSizeOk

`func (o *CloudConfig) GetMaxMsgSizeOk() (*int32, bool)`

GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSize

`func (o *CloudConfig) SetMaxMsgSize(v int32)`

SetMaxMsgSize sets MaxMsgSize field to given value.

### HasMaxMsgSize

`func (o *CloudConfig) HasMaxMsgSize() bool`

HasMaxMsgSize returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *CloudConfig) GetMaxMsgs() int32`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *CloudConfig) GetMaxMsgsOk() (*int32, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *CloudConfig) SetMaxMsgs(v int32)`

SetMaxMsgs sets MaxMsgs field to given value.

### HasMaxMsgs

`func (o *CloudConfig) HasMaxMsgs() bool`

HasMaxMsgs returns a boolean if a field has been set.

### GetName

`func (o *CloudConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumReplicas

`func (o *CloudConfig) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *CloudConfig) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *CloudConfig) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.

### HasNumReplicas

`func (o *CloudConfig) HasNumReplicas() bool`

HasNumReplicas returns a boolean if a field has been set.

### GetRetention

`func (o *CloudConfig) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *CloudConfig) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *CloudConfig) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *CloudConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetStorage

`func (o *CloudConfig) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CloudConfig) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CloudConfig) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CloudConfig) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSubjects

`func (o *CloudConfig) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CloudConfig) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CloudConfig) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *CloudConfig) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


