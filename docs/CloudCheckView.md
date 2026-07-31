# CloudCheckView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the verification was started. | [optional] 
**DecidedAt** | Pointer to **int32** | DecidedAt is the unix second a terminal status was recorded. | [optional] 
**DecidedBy** | Pointer to **string** | DecidedBy records who settled a terminal status: the provider name, or a reviewer&#39;s user id for a recorded manual decision. | [optional] 
**Id** | Pointer to **string** | ID is the verification&#39;s opaque id. | [optional] 
**Kind** | Pointer to **string** | Kind is the party type: \&quot;individual\&quot; (KYC) or \&quot;business\&quot; (KYB). | [optional] 
**Provider** | Pointer to **string** | Provider is the verification provider this check runs through. | [optional] 
**Status** | Pointer to **string** | Status is the check&#39;s state: pending, provider_verified, provider_rejected, manual_review, or expired (provider-reported), or reviewer_confirmed — the one value a privileged human reviewer records, never a provider. | [optional] 
**SubjectId** | Pointer to **string** | SubjectID is the opaque id of the subject under verification. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second the verification last changed. | [optional] 
**VerifyUrl** | Pointer to **string** | VerifyURL is the provider&#39;s hosted verification flow for the subject, when one exists. | [optional] 

## Methods

### NewCloudCheckView

`func NewCloudCheckView() *CloudCheckView`

NewCloudCheckView instantiates a new CloudCheckView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCheckViewWithDefaults

`func NewCloudCheckViewWithDefaults() *CloudCheckView`

NewCloudCheckViewWithDefaults instantiates a new CloudCheckView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudCheckView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCheckView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCheckView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudCheckView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDecidedAt

`func (o *CloudCheckView) GetDecidedAt() int32`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *CloudCheckView) GetDecidedAtOk() (*int32, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *CloudCheckView) SetDecidedAt(v int32)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *CloudCheckView) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDecidedBy

`func (o *CloudCheckView) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *CloudCheckView) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *CloudCheckView) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *CloudCheckView) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetId

`func (o *CloudCheckView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCheckView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCheckView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCheckView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudCheckView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudCheckView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudCheckView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudCheckView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProvider

`func (o *CloudCheckView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudCheckView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudCheckView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudCheckView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCheckView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCheckView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCheckView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCheckView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectId

`func (o *CloudCheckView) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CloudCheckView) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CloudCheckView) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *CloudCheckView) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudCheckView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudCheckView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudCheckView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudCheckView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerifyUrl

`func (o *CloudCheckView) GetVerifyUrl() string`

GetVerifyUrl returns the VerifyUrl field if non-nil, zero value otherwise.

### GetVerifyUrlOk

`func (o *CloudCheckView) GetVerifyUrlOk() (*string, bool)`

GetVerifyUrlOk returns a tuple with the VerifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyUrl

`func (o *CloudCheckView) SetVerifyUrl(v string)`

SetVerifyUrl sets VerifyUrl field to given value.

### HasVerifyUrl

`func (o *CloudCheckView) HasVerifyUrl() bool`

HasVerifyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


