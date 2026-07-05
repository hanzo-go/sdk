# KmsCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CaId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsCertificate

`func NewKmsCertificate() *KmsCertificate`

NewKmsCertificate instantiates a new KmsCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCertificateWithDefaults

`func NewKmsCertificateWithDefaults() *KmsCertificate`

NewKmsCertificateWithDefaults instantiates a new KmsCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCaId

`func (o *KmsCertificate) GetCaId() string`

GetCaId returns the CaId field if non-nil, zero value otherwise.

### GetCaIdOk

`func (o *KmsCertificate) GetCaIdOk() (*string, bool)`

GetCaIdOk returns a tuple with the CaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaId

`func (o *KmsCertificate) SetCaId(v string)`

SetCaId sets CaId field to given value.

### HasCaId

`func (o *KmsCertificate) HasCaId() bool`

HasCaId returns a boolean if a field has been set.

### GetStatus

`func (o *KmsCertificate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KmsCertificate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KmsCertificate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KmsCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFriendlyName

`func (o *KmsCertificate) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *KmsCertificate) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *KmsCertificate) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *KmsCertificate) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetCommonName

`func (o *KmsCertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *KmsCertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *KmsCertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *KmsCertificate) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *KmsCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *KmsCertificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *KmsCertificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *KmsCertificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetNotBefore

`func (o *KmsCertificate) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *KmsCertificate) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *KmsCertificate) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *KmsCertificate) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *KmsCertificate) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *KmsCertificate) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *KmsCertificate) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *KmsCertificate) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsCertificate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsCertificate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsCertificate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsCertificate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


