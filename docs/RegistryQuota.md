# RegistryQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Ref** | Pointer to [**RegistryQuotaRef**](RegistryQuotaRef.md) |  | [optional] 
**Hard** | Pointer to [**RegistryQuotaHard**](RegistryQuotaHard.md) |  | [optional] 
**Used** | Pointer to [**RegistryQuotaUsed**](RegistryQuotaUsed.md) |  | [optional] 

## Methods

### NewRegistryQuota

`func NewRegistryQuota() *RegistryQuota`

NewRegistryQuota instantiates a new RegistryQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryQuotaWithDefaults

`func NewRegistryQuotaWithDefaults() *RegistryQuota`

NewRegistryQuotaWithDefaults instantiates a new RegistryQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistryQuota) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryQuota) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryQuota) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryQuota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRef

`func (o *RegistryQuota) GetRef() RegistryQuotaRef`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RegistryQuota) GetRefOk() (*RegistryQuotaRef, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RegistryQuota) SetRef(v RegistryQuotaRef)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RegistryQuota) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetHard

`func (o *RegistryQuota) GetHard() RegistryQuotaHard`

GetHard returns the Hard field if non-nil, zero value otherwise.

### GetHardOk

`func (o *RegistryQuota) GetHardOk() (*RegistryQuotaHard, bool)`

GetHardOk returns a tuple with the Hard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHard

`func (o *RegistryQuota) SetHard(v RegistryQuotaHard)`

SetHard sets Hard field to given value.

### HasHard

`func (o *RegistryQuota) HasHard() bool`

HasHard returns a boolean if a field has been set.

### GetUsed

`func (o *RegistryQuota) GetUsed() RegistryQuotaUsed`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *RegistryQuota) GetUsedOk() (*RegistryQuotaUsed, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *RegistryQuota) SetUsed(v RegistryQuotaUsed)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *RegistryQuota) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


