# O11yVolumeRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]string** |  | [optional] 
**PersistentVolumeClaimName** | Pointer to **string** |  | [optional] 
**VolumeAvailable** | Pointer to **float64** |  | [optional] 
**VolumeCapacity** | Pointer to **float64** |  | [optional] 
**VolumeInodes** | Pointer to **float64** |  | [optional] 
**VolumeInodesFree** | Pointer to **float64** |  | [optional] 
**VolumeInodesUsed** | Pointer to **float64** |  | [optional] 
**VolumeUsage** | Pointer to **float64** |  | [optional] 

## Methods

### NewO11yVolumeRecord

`func NewO11yVolumeRecord() *O11yVolumeRecord`

NewO11yVolumeRecord instantiates a new O11yVolumeRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yVolumeRecordWithDefaults

`func NewO11yVolumeRecordWithDefaults() *O11yVolumeRecord`

NewO11yVolumeRecordWithDefaults instantiates a new O11yVolumeRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *O11yVolumeRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yVolumeRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yVolumeRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yVolumeRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPersistentVolumeClaimName

`func (o *O11yVolumeRecord) GetPersistentVolumeClaimName() string`

GetPersistentVolumeClaimName returns the PersistentVolumeClaimName field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimNameOk

`func (o *O11yVolumeRecord) GetPersistentVolumeClaimNameOk() (*string, bool)`

GetPersistentVolumeClaimNameOk returns a tuple with the PersistentVolumeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaimName

`func (o *O11yVolumeRecord) SetPersistentVolumeClaimName(v string)`

SetPersistentVolumeClaimName sets PersistentVolumeClaimName field to given value.

### HasPersistentVolumeClaimName

`func (o *O11yVolumeRecord) HasPersistentVolumeClaimName() bool`

HasPersistentVolumeClaimName returns a boolean if a field has been set.

### GetVolumeAvailable

`func (o *O11yVolumeRecord) GetVolumeAvailable() float64`

GetVolumeAvailable returns the VolumeAvailable field if non-nil, zero value otherwise.

### GetVolumeAvailableOk

`func (o *O11yVolumeRecord) GetVolumeAvailableOk() (*float64, bool)`

GetVolumeAvailableOk returns a tuple with the VolumeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAvailable

`func (o *O11yVolumeRecord) SetVolumeAvailable(v float64)`

SetVolumeAvailable sets VolumeAvailable field to given value.

### HasVolumeAvailable

`func (o *O11yVolumeRecord) HasVolumeAvailable() bool`

HasVolumeAvailable returns a boolean if a field has been set.

### GetVolumeCapacity

`func (o *O11yVolumeRecord) GetVolumeCapacity() float64`

GetVolumeCapacity returns the VolumeCapacity field if non-nil, zero value otherwise.

### GetVolumeCapacityOk

`func (o *O11yVolumeRecord) GetVolumeCapacityOk() (*float64, bool)`

GetVolumeCapacityOk returns a tuple with the VolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCapacity

`func (o *O11yVolumeRecord) SetVolumeCapacity(v float64)`

SetVolumeCapacity sets VolumeCapacity field to given value.

### HasVolumeCapacity

`func (o *O11yVolumeRecord) HasVolumeCapacity() bool`

HasVolumeCapacity returns a boolean if a field has been set.

### GetVolumeInodes

`func (o *O11yVolumeRecord) GetVolumeInodes() float64`

GetVolumeInodes returns the VolumeInodes field if non-nil, zero value otherwise.

### GetVolumeInodesOk

`func (o *O11yVolumeRecord) GetVolumeInodesOk() (*float64, bool)`

GetVolumeInodesOk returns a tuple with the VolumeInodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeInodes

`func (o *O11yVolumeRecord) SetVolumeInodes(v float64)`

SetVolumeInodes sets VolumeInodes field to given value.

### HasVolumeInodes

`func (o *O11yVolumeRecord) HasVolumeInodes() bool`

HasVolumeInodes returns a boolean if a field has been set.

### GetVolumeInodesFree

`func (o *O11yVolumeRecord) GetVolumeInodesFree() float64`

GetVolumeInodesFree returns the VolumeInodesFree field if non-nil, zero value otherwise.

### GetVolumeInodesFreeOk

`func (o *O11yVolumeRecord) GetVolumeInodesFreeOk() (*float64, bool)`

GetVolumeInodesFreeOk returns a tuple with the VolumeInodesFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeInodesFree

`func (o *O11yVolumeRecord) SetVolumeInodesFree(v float64)`

SetVolumeInodesFree sets VolumeInodesFree field to given value.

### HasVolumeInodesFree

`func (o *O11yVolumeRecord) HasVolumeInodesFree() bool`

HasVolumeInodesFree returns a boolean if a field has been set.

### GetVolumeInodesUsed

`func (o *O11yVolumeRecord) GetVolumeInodesUsed() float64`

GetVolumeInodesUsed returns the VolumeInodesUsed field if non-nil, zero value otherwise.

### GetVolumeInodesUsedOk

`func (o *O11yVolumeRecord) GetVolumeInodesUsedOk() (*float64, bool)`

GetVolumeInodesUsedOk returns a tuple with the VolumeInodesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeInodesUsed

`func (o *O11yVolumeRecord) SetVolumeInodesUsed(v float64)`

SetVolumeInodesUsed sets VolumeInodesUsed field to given value.

### HasVolumeInodesUsed

`func (o *O11yVolumeRecord) HasVolumeInodesUsed() bool`

HasVolumeInodesUsed returns a boolean if a field has been set.

### GetVolumeUsage

`func (o *O11yVolumeRecord) GetVolumeUsage() float64`

GetVolumeUsage returns the VolumeUsage field if non-nil, zero value otherwise.

### GetVolumeUsageOk

`func (o *O11yVolumeRecord) GetVolumeUsageOk() (*float64, bool)`

GetVolumeUsageOk returns a tuple with the VolumeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsage

`func (o *O11yVolumeRecord) SetVolumeUsage(v float64)`

SetVolumeUsage sets VolumeUsage field to given value.

### HasVolumeUsage

`func (o *O11yVolumeRecord) HasVolumeUsage() bool`

HasVolumeUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


