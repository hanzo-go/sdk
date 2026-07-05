# PubsubStreamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique stream name | 
**Subjects** | **[]string** | Subjects captured by this stream | 
**Storage** | Pointer to **string** | Storage backend | [optional] [default to "file"]
**Replicas** | Pointer to **int32** | Replication factor | [optional] [default to 1]
**Retention** | Pointer to **string** |  | [optional] [default to "limits"]
**MaxMsgs** | Pointer to **int32** | Maximum messages to retain | [optional] [default to -1]
**MaxBytes** | Pointer to **int32** | Maximum bytes to retain | [optional] [default to -1]
**MaxAge** | Pointer to **int32** | Maximum age in nanoseconds (0 &#x3D; unlimited) | [optional] [default to 0]
**Discard** | Pointer to **string** | Discard policy when limits are reached | [optional] [default to "old"]

## Methods

### NewPubsubStreamConfig

`func NewPubsubStreamConfig(name string, subjects []string, ) *PubsubStreamConfig`

NewPubsubStreamConfig instantiates a new PubsubStreamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubStreamConfigWithDefaults

`func NewPubsubStreamConfigWithDefaults() *PubsubStreamConfig`

NewPubsubStreamConfigWithDefaults instantiates a new PubsubStreamConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PubsubStreamConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PubsubStreamConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PubsubStreamConfig) SetName(v string)`

SetName sets Name field to given value.


### GetSubjects

`func (o *PubsubStreamConfig) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *PubsubStreamConfig) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *PubsubStreamConfig) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.


### GetStorage

`func (o *PubsubStreamConfig) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PubsubStreamConfig) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PubsubStreamConfig) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *PubsubStreamConfig) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetReplicas

`func (o *PubsubStreamConfig) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PubsubStreamConfig) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PubsubStreamConfig) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PubsubStreamConfig) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRetention

`func (o *PubsubStreamConfig) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *PubsubStreamConfig) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *PubsubStreamConfig) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *PubsubStreamConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *PubsubStreamConfig) GetMaxMsgs() int32`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *PubsubStreamConfig) GetMaxMsgsOk() (*int32, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *PubsubStreamConfig) SetMaxMsgs(v int32)`

SetMaxMsgs sets MaxMsgs field to given value.

### HasMaxMsgs

`func (o *PubsubStreamConfig) HasMaxMsgs() bool`

HasMaxMsgs returns a boolean if a field has been set.

### GetMaxBytes

`func (o *PubsubStreamConfig) GetMaxBytes() int32`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *PubsubStreamConfig) GetMaxBytesOk() (*int32, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *PubsubStreamConfig) SetMaxBytes(v int32)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *PubsubStreamConfig) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxAge

`func (o *PubsubStreamConfig) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *PubsubStreamConfig) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *PubsubStreamConfig) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *PubsubStreamConfig) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetDiscard

`func (o *PubsubStreamConfig) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *PubsubStreamConfig) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *PubsubStreamConfig) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *PubsubStreamConfig) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


