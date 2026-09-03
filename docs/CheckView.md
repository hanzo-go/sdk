# CheckView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is the unix second the verification was started. | [optional] 
**DecidedAt** | Pointer to **int64** | DecidedAt is the unix second a terminal status was recorded. | [optional] 
**DecidedBy** | Pointer to **string** | DecidedBy records who settled a terminal status: the provider name, or a reviewer&#39;s user id for a recorded manual decision. | [optional] 
**Id** | Pointer to **string** | ID is the verification&#39;s opaque id. | [optional] 
**Kind** | Pointer to **string** | Kind is the party type: \&quot;individual\&quot; (KYC) or \&quot;business\&quot; (KYB). | [optional] 
**Provider** | Pointer to **string** | Provider is the verification provider this check runs through. | [optional] 
**Status** | Pointer to **string** | Status is the check&#39;s state: pending, provider_verified, provider_rejected, manual_review, or expired (provider-reported), or reviewer_confirmed — the one value a privileged human reviewer records, never a provider. | [optional] 
**SubjectId** | Pointer to **string** | SubjectID is the opaque id of the subject under verification. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the unix second the verification last changed. | [optional] 
**VerifyUrl** | Pointer to **string** | VerifyURL is the provider&#39;s hosted verification flow for the subject, when one exists. | [optional] 

## Methods

### NewCheckView

`func NewCheckView() *CheckView`

NewCheckView instantiates a new CheckView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckViewWithDefaults

`func NewCheckViewWithDefaults() *CheckView`

NewCheckViewWithDefaults instantiates a new CheckView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CheckView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CheckView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDecidedAt

`func (o *CheckView) GetDecidedAt() int64`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *CheckView) GetDecidedAtOk() (*int64, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *CheckView) SetDecidedAt(v int64)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *CheckView) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDecidedBy

`func (o *CheckView) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *CheckView) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *CheckView) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *CheckView) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetId

`func (o *CheckView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CheckView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CheckView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CheckView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CheckView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProvider

`func (o *CheckView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CheckView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CheckView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CheckView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *CheckView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectId

`func (o *CheckView) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CheckView) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CheckView) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *CheckView) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CheckView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CheckView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CheckView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CheckView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerifyUrl

`func (o *CheckView) GetVerifyUrl() string`

GetVerifyUrl returns the VerifyUrl field if non-nil, zero value otherwise.

### GetVerifyUrlOk

`func (o *CheckView) GetVerifyUrlOk() (*string, bool)`

GetVerifyUrlOk returns a tuple with the VerifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyUrl

`func (o *CheckView) SetVerifyUrl(v string)`

SetVerifyUrl sets VerifyUrl field to given value.

### HasVerifyUrl

`func (o *CheckView) HasVerifyUrl() bool`

HasVerifyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


