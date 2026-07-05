# KmsSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**SecretValue** | Pointer to **string** |  | [optional] 
**SecretComment** | Pointer to **string** |  | [optional] 
**SecretPath** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]KmsSecretTag**](KmsSecretTag.md) |  | [optional] 
**SecretMetadata** | Pointer to [**[]KmsCreateSecretRequestSecretMetadataInner**](KmsCreateSecretRequestSecretMetadataInner.md) |  | [optional] 
**SecretReminderNote** | Pointer to **string** |  | [optional] 
**SecretReminderRepeatDays** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsSecret

`func NewKmsSecret() *KmsSecret`

NewKmsSecret instantiates a new KmsSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsSecretWithDefaults

`func NewKmsSecretWithDefaults() *KmsSecret`

NewKmsSecretWithDefaults instantiates a new KmsSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsSecret) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsSecret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *KmsSecret) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KmsSecret) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KmsSecret) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KmsSecret) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *KmsSecret) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsSecret) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsSecret) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KmsSecret) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSecretKey

`func (o *KmsSecret) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *KmsSecret) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *KmsSecret) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *KmsSecret) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSecretValue

`func (o *KmsSecret) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *KmsSecret) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *KmsSecret) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *KmsSecret) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.

### GetSecretComment

`func (o *KmsSecret) GetSecretComment() string`

GetSecretComment returns the SecretComment field if non-nil, zero value otherwise.

### GetSecretCommentOk

`func (o *KmsSecret) GetSecretCommentOk() (*string, bool)`

GetSecretCommentOk returns a tuple with the SecretComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretComment

`func (o *KmsSecret) SetSecretComment(v string)`

SetSecretComment sets SecretComment field to given value.

### HasSecretComment

`func (o *KmsSecret) HasSecretComment() bool`

HasSecretComment returns a boolean if a field has been set.

### GetSecretPath

`func (o *KmsSecret) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsSecret) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsSecret) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *KmsSecret) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.

### GetEnvironment

`func (o *KmsSecret) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsSecret) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsSecret) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KmsSecret) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetTags

`func (o *KmsSecret) GetTags() []KmsSecretTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KmsSecret) GetTagsOk() (*[]KmsSecretTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KmsSecret) SetTags(v []KmsSecretTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KmsSecret) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSecretMetadata

`func (o *KmsSecret) GetSecretMetadata() []KmsCreateSecretRequestSecretMetadataInner`

GetSecretMetadata returns the SecretMetadata field if non-nil, zero value otherwise.

### GetSecretMetadataOk

`func (o *KmsSecret) GetSecretMetadataOk() (*[]KmsCreateSecretRequestSecretMetadataInner, bool)`

GetSecretMetadataOk returns a tuple with the SecretMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretMetadata

`func (o *KmsSecret) SetSecretMetadata(v []KmsCreateSecretRequestSecretMetadataInner)`

SetSecretMetadata sets SecretMetadata field to given value.

### HasSecretMetadata

`func (o *KmsSecret) HasSecretMetadata() bool`

HasSecretMetadata returns a boolean if a field has been set.

### GetSecretReminderNote

`func (o *KmsSecret) GetSecretReminderNote() string`

GetSecretReminderNote returns the SecretReminderNote field if non-nil, zero value otherwise.

### GetSecretReminderNoteOk

`func (o *KmsSecret) GetSecretReminderNoteOk() (*string, bool)`

GetSecretReminderNoteOk returns a tuple with the SecretReminderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretReminderNote

`func (o *KmsSecret) SetSecretReminderNote(v string)`

SetSecretReminderNote sets SecretReminderNote field to given value.

### HasSecretReminderNote

`func (o *KmsSecret) HasSecretReminderNote() bool`

HasSecretReminderNote returns a boolean if a field has been set.

### GetSecretReminderRepeatDays

`func (o *KmsSecret) GetSecretReminderRepeatDays() int32`

GetSecretReminderRepeatDays returns the SecretReminderRepeatDays field if non-nil, zero value otherwise.

### GetSecretReminderRepeatDaysOk

`func (o *KmsSecret) GetSecretReminderRepeatDaysOk() (*int32, bool)`

GetSecretReminderRepeatDaysOk returns a tuple with the SecretReminderRepeatDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretReminderRepeatDays

`func (o *KmsSecret) SetSecretReminderRepeatDays(v int32)`

SetSecretReminderRepeatDays sets SecretReminderRepeatDays field to given value.

### HasSecretReminderRepeatDays

`func (o *KmsSecret) HasSecretReminderRepeatDays() bool`

HasSecretReminderRepeatDays returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsSecret) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KmsSecret) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KmsSecret) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KmsSecret) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KmsSecret) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


