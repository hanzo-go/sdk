# KmsUpdateSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretKey** | **string** |  | 
**SecretValue** | Pointer to **string** |  | [optional] 
**SecretComment** | Pointer to **string** |  | [optional] 
**SecretPath** | **string** |  | [default to "/"]
**Environment** | **string** |  | 
**NewSecretName** | Pointer to **string** |  | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 
**SecretMetadata** | Pointer to [**[]KmsCreateSecretRequestSecretMetadataInner**](KmsCreateSecretRequestSecretMetadataInner.md) |  | [optional] 
**SecretReminderNote** | Pointer to **string** |  | [optional] 
**SecretReminderRepeatDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewKmsUpdateSecretRequest

`func NewKmsUpdateSecretRequest(secretKey string, secretPath string, environment string, ) *KmsUpdateSecretRequest`

NewKmsUpdateSecretRequest instantiates a new KmsUpdateSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsUpdateSecretRequestWithDefaults

`func NewKmsUpdateSecretRequestWithDefaults() *KmsUpdateSecretRequest`

NewKmsUpdateSecretRequestWithDefaults instantiates a new KmsUpdateSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretKey

`func (o *KmsUpdateSecretRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *KmsUpdateSecretRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *KmsUpdateSecretRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetSecretValue

`func (o *KmsUpdateSecretRequest) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *KmsUpdateSecretRequest) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *KmsUpdateSecretRequest) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *KmsUpdateSecretRequest) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.

### GetSecretComment

`func (o *KmsUpdateSecretRequest) GetSecretComment() string`

GetSecretComment returns the SecretComment field if non-nil, zero value otherwise.

### GetSecretCommentOk

`func (o *KmsUpdateSecretRequest) GetSecretCommentOk() (*string, bool)`

GetSecretCommentOk returns a tuple with the SecretComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretComment

`func (o *KmsUpdateSecretRequest) SetSecretComment(v string)`

SetSecretComment sets SecretComment field to given value.

### HasSecretComment

`func (o *KmsUpdateSecretRequest) HasSecretComment() bool`

HasSecretComment returns a boolean if a field has been set.

### GetSecretPath

`func (o *KmsUpdateSecretRequest) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsUpdateSecretRequest) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsUpdateSecretRequest) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.


### GetEnvironment

`func (o *KmsUpdateSecretRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsUpdateSecretRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsUpdateSecretRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetNewSecretName

`func (o *KmsUpdateSecretRequest) GetNewSecretName() string`

GetNewSecretName returns the NewSecretName field if non-nil, zero value otherwise.

### GetNewSecretNameOk

`func (o *KmsUpdateSecretRequest) GetNewSecretNameOk() (*string, bool)`

GetNewSecretNameOk returns a tuple with the NewSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSecretName

`func (o *KmsUpdateSecretRequest) SetNewSecretName(v string)`

SetNewSecretName sets NewSecretName field to given value.

### HasNewSecretName

`func (o *KmsUpdateSecretRequest) HasNewSecretName() bool`

HasNewSecretName returns a boolean if a field has been set.

### GetTagIds

`func (o *KmsUpdateSecretRequest) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *KmsUpdateSecretRequest) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *KmsUpdateSecretRequest) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *KmsUpdateSecretRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSecretMetadata

`func (o *KmsUpdateSecretRequest) GetSecretMetadata() []KmsCreateSecretRequestSecretMetadataInner`

GetSecretMetadata returns the SecretMetadata field if non-nil, zero value otherwise.

### GetSecretMetadataOk

`func (o *KmsUpdateSecretRequest) GetSecretMetadataOk() (*[]KmsCreateSecretRequestSecretMetadataInner, bool)`

GetSecretMetadataOk returns a tuple with the SecretMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretMetadata

`func (o *KmsUpdateSecretRequest) SetSecretMetadata(v []KmsCreateSecretRequestSecretMetadataInner)`

SetSecretMetadata sets SecretMetadata field to given value.

### HasSecretMetadata

`func (o *KmsUpdateSecretRequest) HasSecretMetadata() bool`

HasSecretMetadata returns a boolean if a field has been set.

### GetSecretReminderNote

`func (o *KmsUpdateSecretRequest) GetSecretReminderNote() string`

GetSecretReminderNote returns the SecretReminderNote field if non-nil, zero value otherwise.

### GetSecretReminderNoteOk

`func (o *KmsUpdateSecretRequest) GetSecretReminderNoteOk() (*string, bool)`

GetSecretReminderNoteOk returns a tuple with the SecretReminderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretReminderNote

`func (o *KmsUpdateSecretRequest) SetSecretReminderNote(v string)`

SetSecretReminderNote sets SecretReminderNote field to given value.

### HasSecretReminderNote

`func (o *KmsUpdateSecretRequest) HasSecretReminderNote() bool`

HasSecretReminderNote returns a boolean if a field has been set.

### GetSecretReminderRepeatDays

`func (o *KmsUpdateSecretRequest) GetSecretReminderRepeatDays() int32`

GetSecretReminderRepeatDays returns the SecretReminderRepeatDays field if non-nil, zero value otherwise.

### GetSecretReminderRepeatDaysOk

`func (o *KmsUpdateSecretRequest) GetSecretReminderRepeatDaysOk() (*int32, bool)`

GetSecretReminderRepeatDaysOk returns a tuple with the SecretReminderRepeatDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretReminderRepeatDays

`func (o *KmsUpdateSecretRequest) SetSecretReminderRepeatDays(v int32)`

SetSecretReminderRepeatDays sets SecretReminderRepeatDays field to given value.

### HasSecretReminderRepeatDays

`func (o *KmsUpdateSecretRequest) HasSecretReminderRepeatDays() bool`

HasSecretReminderRepeatDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


