# IamCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessSecret** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**BitSize** | Pointer to **int32** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**CryptoAlgorithm** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DomainExpireTime** | Pointer to **string** |  | [optional] 
**ExpireInYears** | Pointer to **int32** |  | [optional] 
**ExpireTime** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIamCert

`func NewIamCert() *IamCert`

NewIamCert instantiates a new IamCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCertWithDefaults

`func NewIamCertWithDefaults() *IamCert`

NewIamCertWithDefaults instantiates a new IamCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *IamCert) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IamCert) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IamCert) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *IamCert) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessSecret

`func (o *IamCert) GetAccessSecret() string`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *IamCert) GetAccessSecretOk() (*string, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *IamCert) SetAccessSecret(v string)`

SetAccessSecret sets AccessSecret field to given value.

### HasAccessSecret

`func (o *IamCert) HasAccessSecret() bool`

HasAccessSecret returns a boolean if a field has been set.

### GetAccount

`func (o *IamCert) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamCert) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamCert) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamCert) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBitSize

`func (o *IamCert) GetBitSize() int32`

GetBitSize returns the BitSize field if non-nil, zero value otherwise.

### GetBitSizeOk

`func (o *IamCert) GetBitSizeOk() (*int32, bool)`

GetBitSizeOk returns a tuple with the BitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitSize

`func (o *IamCert) SetBitSize(v int32)`

SetBitSize sets BitSize field to given value.

### HasBitSize

`func (o *IamCert) HasBitSize() bool`

HasBitSize returns a boolean if a field has been set.

### GetCertificate

`func (o *IamCert) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IamCert) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IamCert) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IamCert) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamCert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamCert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamCert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamCert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamCert) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamCert) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamCert) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamCert) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCryptoAlgorithm

`func (o *IamCert) GetCryptoAlgorithm() string`

GetCryptoAlgorithm returns the CryptoAlgorithm field if non-nil, zero value otherwise.

### GetCryptoAlgorithmOk

`func (o *IamCert) GetCryptoAlgorithmOk() (*string, bool)`

GetCryptoAlgorithmOk returns a tuple with the CryptoAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAlgorithm

`func (o *IamCert) SetCryptoAlgorithm(v string)`

SetCryptoAlgorithm sets CryptoAlgorithm field to given value.

### HasCryptoAlgorithm

`func (o *IamCert) HasCryptoAlgorithm() bool`

HasCryptoAlgorithm returns a boolean if a field has been set.

### GetDeleted

`func (o *IamCert) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamCert) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamCert) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamCert) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamCert) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamCert) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamCert) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamCert) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomainExpireTime

`func (o *IamCert) GetDomainExpireTime() string`

GetDomainExpireTime returns the DomainExpireTime field if non-nil, zero value otherwise.

### GetDomainExpireTimeOk

`func (o *IamCert) GetDomainExpireTimeOk() (*string, bool)`

GetDomainExpireTimeOk returns a tuple with the DomainExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainExpireTime

`func (o *IamCert) SetDomainExpireTime(v string)`

SetDomainExpireTime sets DomainExpireTime field to given value.

### HasDomainExpireTime

`func (o *IamCert) HasDomainExpireTime() bool`

HasDomainExpireTime returns a boolean if a field has been set.

### GetExpireInYears

`func (o *IamCert) GetExpireInYears() int32`

GetExpireInYears returns the ExpireInYears field if non-nil, zero value otherwise.

### GetExpireInYearsOk

`func (o *IamCert) GetExpireInYearsOk() (*int32, bool)`

GetExpireInYearsOk returns a tuple with the ExpireInYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireInYears

`func (o *IamCert) SetExpireInYears(v int32)`

SetExpireInYears sets ExpireInYears field to given value.

### HasExpireInYears

`func (o *IamCert) HasExpireInYears() bool`

HasExpireInYears returns a boolean if a field has been set.

### GetExpireTime

`func (o *IamCert) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *IamCert) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *IamCert) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *IamCert) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetId

`func (o *IamCert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamCert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamCert) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamCert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamCert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamCert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamCert) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamCert) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamCert) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamCert) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamCert) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamCert) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProvider

`func (o *IamCert) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IamCert) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IamCert) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IamCert) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScope

`func (o *IamCert) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamCert) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamCert) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamCert) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetType

`func (o *IamCert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamCert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamCert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamCert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamCert) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamCert) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamCert) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamCert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


