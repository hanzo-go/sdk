# MqStreamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Stream name (alphanumeric, hyphens, underscores). | 
**Subjects** | **[]string** | Subjects bound to this stream. Supports wildcards.  | 
**Retention** | Pointer to **string** | Message retention policy. &#x60;limits&#x60; keeps messages until limits are hit. &#x60;interest&#x60; keeps messages while consumers exist. &#x60;workqueue&#x60; deletes messages after acknowledgment.  | [optional] [default to "limits"]
**MaxMsgs** | Pointer to **int64** | Maximum number of messages (-1 for unlimited). | [optional] [default to -1]
**MaxBytes** | Pointer to **int64** | Maximum total bytes (-1 for unlimited). | [optional] [default to -1]
**MaxAge** | Pointer to **string** | Maximum message age (e.g., \&quot;24h\&quot;, \&quot;7d\&quot;, \&quot;0\&quot; for unlimited).  | [optional] [default to "0"]
**MaxMsgSize** | Pointer to **int32** | Maximum single message size in bytes (-1 for default). | [optional] [default to -1]
**Storage** | Pointer to **string** | Storage backend for stream data. | [optional] [default to "file"]
**NumReplicas** | Pointer to **int32** | Number of replicas in the cluster. | [optional] [default to 1]
**Discard** | Pointer to **string** | Discard policy when limits are reached. &#x60;old&#x60; discards the oldest messages. &#x60;new&#x60; rejects new messages.  | [optional] [default to "old"]
**DuplicateWindow** | Pointer to **string** | Window for message deduplication based on Nats-Msg-Id header (e.g., \&quot;2m\&quot;). Defaults to \&quot;2m\&quot;.  | [optional] [default to "2m"]
**Description** | Pointer to **string** | Optional human-readable description. | [optional] 

## Methods

### NewMqStreamConfig

`func NewMqStreamConfig(name string, subjects []string, ) *MqStreamConfig`

NewMqStreamConfig instantiates a new MqStreamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqStreamConfigWithDefaults

`func NewMqStreamConfigWithDefaults() *MqStreamConfig`

NewMqStreamConfigWithDefaults instantiates a new MqStreamConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MqStreamConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqStreamConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqStreamConfig) SetName(v string)`

SetName sets Name field to given value.


### GetSubjects

`func (o *MqStreamConfig) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *MqStreamConfig) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *MqStreamConfig) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.


### GetRetention

`func (o *MqStreamConfig) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *MqStreamConfig) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *MqStreamConfig) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *MqStreamConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *MqStreamConfig) GetMaxMsgs() int64`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *MqStreamConfig) GetMaxMsgsOk() (*int64, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *MqStreamConfig) SetMaxMsgs(v int64)`

SetMaxMsgs sets MaxMsgs field to given value.

### HasMaxMsgs

`func (o *MqStreamConfig) HasMaxMsgs() bool`

HasMaxMsgs returns a boolean if a field has been set.

### GetMaxBytes

`func (o *MqStreamConfig) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *MqStreamConfig) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *MqStreamConfig) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *MqStreamConfig) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxAge

`func (o *MqStreamConfig) GetMaxAge() string`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *MqStreamConfig) GetMaxAgeOk() (*string, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *MqStreamConfig) SetMaxAge(v string)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *MqStreamConfig) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxMsgSize

`func (o *MqStreamConfig) GetMaxMsgSize() int32`

GetMaxMsgSize returns the MaxMsgSize field if non-nil, zero value otherwise.

### GetMaxMsgSizeOk

`func (o *MqStreamConfig) GetMaxMsgSizeOk() (*int32, bool)`

GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSize

`func (o *MqStreamConfig) SetMaxMsgSize(v int32)`

SetMaxMsgSize sets MaxMsgSize field to given value.

### HasMaxMsgSize

`func (o *MqStreamConfig) HasMaxMsgSize() bool`

HasMaxMsgSize returns a boolean if a field has been set.

### GetStorage

`func (o *MqStreamConfig) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *MqStreamConfig) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *MqStreamConfig) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *MqStreamConfig) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetNumReplicas

`func (o *MqStreamConfig) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *MqStreamConfig) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *MqStreamConfig) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.

### HasNumReplicas

`func (o *MqStreamConfig) HasNumReplicas() bool`

HasNumReplicas returns a boolean if a field has been set.

### GetDiscard

`func (o *MqStreamConfig) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *MqStreamConfig) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *MqStreamConfig) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *MqStreamConfig) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetDuplicateWindow

`func (o *MqStreamConfig) GetDuplicateWindow() string`

GetDuplicateWindow returns the DuplicateWindow field if non-nil, zero value otherwise.

### GetDuplicateWindowOk

`func (o *MqStreamConfig) GetDuplicateWindowOk() (*string, bool)`

GetDuplicateWindowOk returns a tuple with the DuplicateWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateWindow

`func (o *MqStreamConfig) SetDuplicateWindow(v string)`

SetDuplicateWindow sets DuplicateWindow field to given value.

### HasDuplicateWindow

`func (o *MqStreamConfig) HasDuplicateWindow() bool`

HasDuplicateWindow returns a boolean if a field has been set.

### GetDescription

`func (o *MqStreamConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MqStreamConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MqStreamConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MqStreamConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


