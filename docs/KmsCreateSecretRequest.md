# KmsCreateSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretKey** | **string** |  | 
**SecretValue** | **string** |  | 
**SecretComment** | Pointer to **string** |  | [optional] 
**SecretPath** | **string** |  | [default to "/"]
**Environment** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "shared"]
**TagIds** | Pointer to **[]string** |  | [optional] 
**SecretMetadata** | Pointer to [**[]KmsCreateSecretRequestSecretMetadataInner**](KmsCreateSecretRequestSecretMetadataInner.md) |  | [optional] 
**SecretReminderNote** | Pointer to **string** |  | [optional] 
**SecretReminderRepeatDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewKmsCreateSecretRequest

`func NewKmsCreateSecretRequest(secretKey string, secretValue string, secretPath string, environment string, ) *KmsCreateSecretRequest`

NewKmsCreateSecretRequest instantiates a new KmsCreateSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateSecretRequestWithDefaults

`func NewKmsCreateSecretRequestWithDefaults() *KmsCreateSecretRequest`

NewKmsCreateSecretRequestWithDefaults instantiates a new KmsCreateSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretKey

`func (o *KmsCreateSecretRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *KmsCreateSecretRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *KmsCreateSecretRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetSecretValue

`func (o *KmsCreateSecretRequest) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *KmsCreateSecretRequest) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *KmsCreateSecretRequest) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.


### GetSecretComment

`func (o *KmsCreateSecretRequest) GetSecretComment() string`

GetSecretComment returns the SecretComment field if non-nil, zero value otherwise.

### GetSecretCommentOk

`func (o *KmsCreateSecretRequest) GetSecretCommentOk() (*string, bool)`

GetSecretCommentOk returns a tuple with the SecretComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretComment

`func (o *KmsCreateSecretRequest) SetSecretComment(v string)`

SetSecretComment sets SecretComment field to given value.

### HasSecretComment

`func (o *KmsCreateSecretRequest) HasSecretComment() bool`

HasSecretComment returns a boolean if a field has been set.

### GetSecretPath

`func (o *KmsCreateSecretRequest) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsCreateSecretRequest) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsCreateSecretRequest) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.


### GetEnvironment

`func (o *KmsCreateSecretRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsCreateSecretRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsCreateSecretRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *KmsCreateSecretRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsCreateSecretRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsCreateSecretRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KmsCreateSecretRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTagIds

`func (o *KmsCreateSecretRequest) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *KmsCreateSecretRequest) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *KmsCreateSecretRequest) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *KmsCreateSecretRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSecretMetadata

`func (o *KmsCreateSecretRequest) GetSecretMetadata() []KmsCreateSecretRequestSecretMetadataInner`

GetSecretMetadata returns the SecretMetadata field if non-nil, zero value otherwise.

### GetSecretMetadataOk

`func (o *KmsCreateSecretRequest) GetSecretMetadataOk() (*[]KmsCreateSecretRequestSecretMetadataInner, bool)`

GetSecretMetadataOk returns a tuple with the SecretMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretMetadata

`func (o *KmsCreateSecretRequest) SetSecretMetadata(v []KmsCreateSecretRequestSecretMetadataInner)`

SetSecretMetadata sets SecretMetadata field to given value.

### HasSecretMetadata

`func (o *KmsCreateSecretRequest) HasSecretMetadata() bool`

HasSecretMetadata returns a boolean if a field has been set.

### GetSecretReminderNote

`func (o *KmsCreateSecretRequest) GetSecretReminderNote() string`

GetSecretReminderNote returns the SecretReminderNote field if non-nil, zero value otherwise.

### GetSecretReminderNoteOk

`func (o *KmsCreateSecretRequest) GetSecretReminderNoteOk() (*string, bool)`

GetSecretReminderNoteOk returns a tuple with the SecretReminderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretReminderNote

`func (o *KmsCreateSecretRequest) SetSecretReminderNote(v string)`

SetSecretReminderNote sets SecretReminderNote field to given value.

### HasSecretReminderNote

`func (o *KmsCreateSecretRequest) HasSecretReminderNote() bool`

HasSecretReminderNote returns a boolean if a field has been set.

### GetSecretReminderRepeatDays

`func (o *KmsCreateSecretRequest) GetSecretReminderRepeatDays() int32`

GetSecretReminderRepeatDays returns the SecretReminderRepeatDays field if non-nil, zero value otherwise.

### GetSecretReminderRepeatDaysOk

`func (o *KmsCreateSecretRequest) GetSecretReminderRepeatDaysOk() (*int32, bool)`

GetSecretReminderRepeatDaysOk returns a tuple with the SecretReminderRepeatDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretReminderRepeatDays

`func (o *KmsCreateSecretRequest) SetSecretReminderRepeatDays(v int32)`

SetSecretReminderRepeatDays sets SecretReminderRepeatDays field to given value.

### HasSecretReminderRepeatDays

`func (o *KmsCreateSecretRequest) HasSecretReminderRepeatDays() bool`

HasSecretReminderRepeatDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


