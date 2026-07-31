# CloudStreamUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discard** | Pointer to **string** |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**MaxBytes** | Pointer to **int32** |  | [optional] 
**MaxMsgs** | Pointer to **int32** |  | [optional] 
**Retention** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to **string** | Stream is the stream to update, from the path. | [optional] 
**Subjects** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudStreamUpdate

`func NewCloudStreamUpdate() *CloudStreamUpdate`

NewCloudStreamUpdate instantiates a new CloudStreamUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStreamUpdateWithDefaults

`func NewCloudStreamUpdateWithDefaults() *CloudStreamUpdate`

NewCloudStreamUpdateWithDefaults instantiates a new CloudStreamUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscard

`func (o *CloudStreamUpdate) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *CloudStreamUpdate) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *CloudStreamUpdate) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *CloudStreamUpdate) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetMaxAge

`func (o *CloudStreamUpdate) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CloudStreamUpdate) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CloudStreamUpdate) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CloudStreamUpdate) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxBytes

`func (o *CloudStreamUpdate) GetMaxBytes() int32`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *CloudStreamUpdate) GetMaxBytesOk() (*int32, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *CloudStreamUpdate) SetMaxBytes(v int32)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *CloudStreamUpdate) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *CloudStreamUpdate) GetMaxMsgs() int32`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *CloudStreamUpdate) GetMaxMsgsOk() (*int32, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *CloudStreamUpdate) SetMaxMsgs(v int32)`

SetMaxMsgs sets MaxMsgs field to given value.

### HasMaxMsgs

`func (o *CloudStreamUpdate) HasMaxMsgs() bool`

HasMaxMsgs returns a boolean if a field has been set.

### GetRetention

`func (o *CloudStreamUpdate) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *CloudStreamUpdate) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *CloudStreamUpdate) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *CloudStreamUpdate) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetStorage

`func (o *CloudStreamUpdate) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CloudStreamUpdate) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CloudStreamUpdate) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CloudStreamUpdate) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetStream

`func (o *CloudStreamUpdate) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CloudStreamUpdate) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CloudStreamUpdate) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CloudStreamUpdate) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSubjects

`func (o *CloudStreamUpdate) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CloudStreamUpdate) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CloudStreamUpdate) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *CloudStreamUpdate) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


