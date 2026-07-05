# KmsCertificateAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Dn** | Pointer to **string** |  | [optional] 
**MaxPathLength** | Pointer to **int32** |  | [optional] 
**KeyAlgorithm** | Pointer to **string** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsCertificateAuthority

`func NewKmsCertificateAuthority() *KmsCertificateAuthority`

NewKmsCertificateAuthority instantiates a new KmsCertificateAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCertificateAuthorityWithDefaults

`func NewKmsCertificateAuthorityWithDefaults() *KmsCertificateAuthority`

NewKmsCertificateAuthorityWithDefaults instantiates a new KmsCertificateAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsCertificateAuthority) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsCertificateAuthority) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsCertificateAuthority) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsCertificateAuthority) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KmsCertificateAuthority) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsCertificateAuthority) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsCertificateAuthority) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KmsCertificateAuthority) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *KmsCertificateAuthority) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KmsCertificateAuthority) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KmsCertificateAuthority) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KmsCertificateAuthority) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFriendlyName

`func (o *KmsCertificateAuthority) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *KmsCertificateAuthority) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *KmsCertificateAuthority) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *KmsCertificateAuthority) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetOrganization

`func (o *KmsCertificateAuthority) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KmsCertificateAuthority) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KmsCertificateAuthority) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KmsCertificateAuthority) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCommonName

`func (o *KmsCertificateAuthority) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *KmsCertificateAuthority) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *KmsCertificateAuthority) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *KmsCertificateAuthority) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetDn

`func (o *KmsCertificateAuthority) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *KmsCertificateAuthority) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *KmsCertificateAuthority) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *KmsCertificateAuthority) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetMaxPathLength

`func (o *KmsCertificateAuthority) GetMaxPathLength() int32`

GetMaxPathLength returns the MaxPathLength field if non-nil, zero value otherwise.

### GetMaxPathLengthOk

`func (o *KmsCertificateAuthority) GetMaxPathLengthOk() (*int32, bool)`

GetMaxPathLengthOk returns a tuple with the MaxPathLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathLength

`func (o *KmsCertificateAuthority) SetMaxPathLength(v int32)`

SetMaxPathLength sets MaxPathLength field to given value.

### HasMaxPathLength

`func (o *KmsCertificateAuthority) HasMaxPathLength() bool`

HasMaxPathLength returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *KmsCertificateAuthority) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *KmsCertificateAuthority) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *KmsCertificateAuthority) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *KmsCertificateAuthority) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetNotBefore

`func (o *KmsCertificateAuthority) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *KmsCertificateAuthority) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *KmsCertificateAuthority) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *KmsCertificateAuthority) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *KmsCertificateAuthority) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *KmsCertificateAuthority) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *KmsCertificateAuthority) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *KmsCertificateAuthority) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsCertificateAuthority) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsCertificateAuthority) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsCertificateAuthority) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsCertificateAuthority) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


