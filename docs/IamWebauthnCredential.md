# IamWebauthnCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aaguid** | Pointer to **string** |  | [optional] 
**Attachment** | Pointer to **string** |  | [optional] 
**AttestationFormat** | Pointer to **string** | AttestationFormat is the statement format the authenticator attested in (\&quot;packed\&quot;, \&quot;apple\&quot;, \&quot;none\&quot;, …), which is a DIFFERENT value from the attestation type above. The library reads it back when resolving the FIDO AppID extension, so a row that dropped it would round-trip a credential the verifier no longer recognises as the one it stored. | [optional] 
**AttestationType** | Pointer to **string** |  | [optional] 
**BackupEligible** | Pointer to **bool** |  | [optional] 
**BackupState** | Pointer to **bool** |  | [optional] 
**CloneWarning** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**CredentialId** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**SignCount** | Pointer to **int32** |  | [optional] 
**Transport** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**UserPresent** | Pointer to **bool** |  | [optional] 
**UserVerified** | Pointer to **bool** |  | [optional] 

## Methods

### NewIamWebauthnCredential

`func NewIamWebauthnCredential() *IamWebauthnCredential`

NewIamWebauthnCredential instantiates a new IamWebauthnCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamWebauthnCredentialWithDefaults

`func NewIamWebauthnCredentialWithDefaults() *IamWebauthnCredential`

NewIamWebauthnCredentialWithDefaults instantiates a new IamWebauthnCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAaguid

`func (o *IamWebauthnCredential) GetAaguid() string`

GetAaguid returns the Aaguid field if non-nil, zero value otherwise.

### GetAaguidOk

`func (o *IamWebauthnCredential) GetAaguidOk() (*string, bool)`

GetAaguidOk returns a tuple with the Aaguid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaguid

`func (o *IamWebauthnCredential) SetAaguid(v string)`

SetAaguid sets Aaguid field to given value.

### HasAaguid

`func (o *IamWebauthnCredential) HasAaguid() bool`

HasAaguid returns a boolean if a field has been set.

### GetAttachment

`func (o *IamWebauthnCredential) GetAttachment() string`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *IamWebauthnCredential) GetAttachmentOk() (*string, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *IamWebauthnCredential) SetAttachment(v string)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *IamWebauthnCredential) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetAttestationFormat

`func (o *IamWebauthnCredential) GetAttestationFormat() string`

GetAttestationFormat returns the AttestationFormat field if non-nil, zero value otherwise.

### GetAttestationFormatOk

`func (o *IamWebauthnCredential) GetAttestationFormatOk() (*string, bool)`

GetAttestationFormatOk returns a tuple with the AttestationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationFormat

`func (o *IamWebauthnCredential) SetAttestationFormat(v string)`

SetAttestationFormat sets AttestationFormat field to given value.

### HasAttestationFormat

`func (o *IamWebauthnCredential) HasAttestationFormat() bool`

HasAttestationFormat returns a boolean if a field has been set.

### GetAttestationType

`func (o *IamWebauthnCredential) GetAttestationType() string`

GetAttestationType returns the AttestationType field if non-nil, zero value otherwise.

### GetAttestationTypeOk

`func (o *IamWebauthnCredential) GetAttestationTypeOk() (*string, bool)`

GetAttestationTypeOk returns a tuple with the AttestationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationType

`func (o *IamWebauthnCredential) SetAttestationType(v string)`

SetAttestationType sets AttestationType field to given value.

### HasAttestationType

`func (o *IamWebauthnCredential) HasAttestationType() bool`

HasAttestationType returns a boolean if a field has been set.

### GetBackupEligible

`func (o *IamWebauthnCredential) GetBackupEligible() bool`

GetBackupEligible returns the BackupEligible field if non-nil, zero value otherwise.

### GetBackupEligibleOk

`func (o *IamWebauthnCredential) GetBackupEligibleOk() (*bool, bool)`

GetBackupEligibleOk returns a tuple with the BackupEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEligible

`func (o *IamWebauthnCredential) SetBackupEligible(v bool)`

SetBackupEligible sets BackupEligible field to given value.

### HasBackupEligible

`func (o *IamWebauthnCredential) HasBackupEligible() bool`

HasBackupEligible returns a boolean if a field has been set.

### GetBackupState

`func (o *IamWebauthnCredential) GetBackupState() bool`

GetBackupState returns the BackupState field if non-nil, zero value otherwise.

### GetBackupStateOk

`func (o *IamWebauthnCredential) GetBackupStateOk() (*bool, bool)`

GetBackupStateOk returns a tuple with the BackupState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupState

`func (o *IamWebauthnCredential) SetBackupState(v bool)`

SetBackupState sets BackupState field to given value.

### HasBackupState

`func (o *IamWebauthnCredential) HasBackupState() bool`

HasBackupState returns a boolean if a field has been set.

### GetCloneWarning

`func (o *IamWebauthnCredential) GetCloneWarning() bool`

GetCloneWarning returns the CloneWarning field if non-nil, zero value otherwise.

### GetCloneWarningOk

`func (o *IamWebauthnCredential) GetCloneWarningOk() (*bool, bool)`

GetCloneWarningOk returns a tuple with the CloneWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneWarning

`func (o *IamWebauthnCredential) SetCloneWarning(v bool)`

SetCloneWarning sets CloneWarning field to given value.

### HasCloneWarning

`func (o *IamWebauthnCredential) HasCloneWarning() bool`

HasCloneWarning returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamWebauthnCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamWebauthnCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamWebauthnCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamWebauthnCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamWebauthnCredential) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamWebauthnCredential) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamWebauthnCredential) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamWebauthnCredential) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCredentialId

`func (o *IamWebauthnCredential) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *IamWebauthnCredential) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *IamWebauthnCredential) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *IamWebauthnCredential) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetDeleted

`func (o *IamWebauthnCredential) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamWebauthnCredential) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamWebauthnCredential) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamWebauthnCredential) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *IamWebauthnCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamWebauthnCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamWebauthnCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamWebauthnCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamWebauthnCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamWebauthnCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamWebauthnCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamWebauthnCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamWebauthnCredential) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamWebauthnCredential) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamWebauthnCredential) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamWebauthnCredential) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPublicKey

`func (o *IamWebauthnCredential) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *IamWebauthnCredential) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *IamWebauthnCredential) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *IamWebauthnCredential) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSignCount

`func (o *IamWebauthnCredential) GetSignCount() int32`

GetSignCount returns the SignCount field if non-nil, zero value otherwise.

### GetSignCountOk

`func (o *IamWebauthnCredential) GetSignCountOk() (*int32, bool)`

GetSignCountOk returns a tuple with the SignCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignCount

`func (o *IamWebauthnCredential) SetSignCount(v int32)`

SetSignCount sets SignCount field to given value.

### HasSignCount

`func (o *IamWebauthnCredential) HasSignCount() bool`

HasSignCount returns a boolean if a field has been set.

### GetTransport

`func (o *IamWebauthnCredential) GetTransport() []string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *IamWebauthnCredential) GetTransportOk() (*[]string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *IamWebauthnCredential) SetTransport(v []string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *IamWebauthnCredential) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamWebauthnCredential) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamWebauthnCredential) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamWebauthnCredential) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamWebauthnCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *IamWebauthnCredential) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamWebauthnCredential) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamWebauthnCredential) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamWebauthnCredential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserPresent

`func (o *IamWebauthnCredential) GetUserPresent() bool`

GetUserPresent returns the UserPresent field if non-nil, zero value otherwise.

### GetUserPresentOk

`func (o *IamWebauthnCredential) GetUserPresentOk() (*bool, bool)`

GetUserPresentOk returns a tuple with the UserPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPresent

`func (o *IamWebauthnCredential) SetUserPresent(v bool)`

SetUserPresent sets UserPresent field to given value.

### HasUserPresent

`func (o *IamWebauthnCredential) HasUserPresent() bool`

HasUserPresent returns a boolean if a field has been set.

### GetUserVerified

`func (o *IamWebauthnCredential) GetUserVerified() bool`

GetUserVerified returns the UserVerified field if non-nil, zero value otherwise.

### GetUserVerifiedOk

`func (o *IamWebauthnCredential) GetUserVerifiedOk() (*bool, bool)`

GetUserVerifiedOk returns a tuple with the UserVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerified

`func (o *IamWebauthnCredential) SetUserVerified(v bool)`

SetUserVerified sets UserVerified field to given value.

### HasUserVerified

`func (o *IamWebauthnCredential) HasUserVerified() bool`

HasUserVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


