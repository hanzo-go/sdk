# KmsSharedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**EncryptedValue** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAfterViews** | Pointer to **int32** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsSharedSecret

`func NewKmsSharedSecret() *KmsSharedSecret`

NewKmsSharedSecret instantiates a new KmsSharedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsSharedSecretWithDefaults

`func NewKmsSharedSecretWithDefaults() *KmsSharedSecret`

NewKmsSharedSecretWithDefaults instantiates a new KmsSharedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsSharedSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsSharedSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsSharedSecret) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsSharedSecret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEncryptedValue

`func (o *KmsSharedSecret) GetEncryptedValue() string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *KmsSharedSecret) GetEncryptedValueOk() (*string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *KmsSharedSecret) SetEncryptedValue(v string)`

SetEncryptedValue sets EncryptedValue field to given value.

### HasEncryptedValue

`func (o *KmsSharedSecret) HasEncryptedValue() bool`

HasEncryptedValue returns a boolean if a field has been set.

### GetExpiresAt

`func (o *KmsSharedSecret) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *KmsSharedSecret) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *KmsSharedSecret) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *KmsSharedSecret) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExpiresAfterViews

`func (o *KmsSharedSecret) GetExpiresAfterViews() int32`

GetExpiresAfterViews returns the ExpiresAfterViews field if non-nil, zero value otherwise.

### GetExpiresAfterViewsOk

`func (o *KmsSharedSecret) GetExpiresAfterViewsOk() (*int32, bool)`

GetExpiresAfterViewsOk returns a tuple with the ExpiresAfterViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAfterViews

`func (o *KmsSharedSecret) SetExpiresAfterViews(v int32)`

SetExpiresAfterViews sets ExpiresAfterViews field to given value.

### HasExpiresAfterViews

`func (o *KmsSharedSecret) HasExpiresAfterViews() bool`

HasExpiresAfterViews returns a boolean if a field has been set.

### GetAccessType

`func (o *KmsSharedSecret) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *KmsSharedSecret) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *KmsSharedSecret) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *KmsSharedSecret) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsSharedSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsSharedSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsSharedSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsSharedSecret) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


