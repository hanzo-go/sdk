# O11yVolumeListRecord

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

### NewO11yVolumeListRecord

`func NewO11yVolumeListRecord() *O11yVolumeListRecord`

NewO11yVolumeListRecord instantiates a new O11yVolumeListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yVolumeListRecordWithDefaults

`func NewO11yVolumeListRecordWithDefaults() *O11yVolumeListRecord`

NewO11yVolumeListRecordWithDefaults instantiates a new O11yVolumeListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *O11yVolumeListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yVolumeListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yVolumeListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yVolumeListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPersistentVolumeClaimName

`func (o *O11yVolumeListRecord) GetPersistentVolumeClaimName() string`

GetPersistentVolumeClaimName returns the PersistentVolumeClaimName field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimNameOk

`func (o *O11yVolumeListRecord) GetPersistentVolumeClaimNameOk() (*string, bool)`

GetPersistentVolumeClaimNameOk returns a tuple with the PersistentVolumeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaimName

`func (o *O11yVolumeListRecord) SetPersistentVolumeClaimName(v string)`

SetPersistentVolumeClaimName sets PersistentVolumeClaimName field to given value.

### HasPersistentVolumeClaimName

`func (o *O11yVolumeListRecord) HasPersistentVolumeClaimName() bool`

HasPersistentVolumeClaimName returns a boolean if a field has been set.

### GetVolumeAvailable

`func (o *O11yVolumeListRecord) GetVolumeAvailable() float64`

GetVolumeAvailable returns the VolumeAvailable field if non-nil, zero value otherwise.

### GetVolumeAvailableOk

`func (o *O11yVolumeListRecord) GetVolumeAvailableOk() (*float64, bool)`

GetVolumeAvailableOk returns a tuple with the VolumeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAvailable

`func (o *O11yVolumeListRecord) SetVolumeAvailable(v float64)`

SetVolumeAvailable sets VolumeAvailable field to given value.

### HasVolumeAvailable

`func (o *O11yVolumeListRecord) HasVolumeAvailable() bool`

HasVolumeAvailable returns a boolean if a field has been set.

### GetVolumeCapacity

`func (o *O11yVolumeListRecord) GetVolumeCapacity() float64`

GetVolumeCapacity returns the VolumeCapacity field if non-nil, zero value otherwise.

### GetVolumeCapacityOk

`func (o *O11yVolumeListRecord) GetVolumeCapacityOk() (*float64, bool)`

GetVolumeCapacityOk returns a tuple with the VolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCapacity

`func (o *O11yVolumeListRecord) SetVolumeCapacity(v float64)`

SetVolumeCapacity sets VolumeCapacity field to given value.

### HasVolumeCapacity

`func (o *O11yVolumeListRecord) HasVolumeCapacity() bool`

HasVolumeCapacity returns a boolean if a field has been set.

### GetVolumeInodes

`func (o *O11yVolumeListRecord) GetVolumeInodes() float64`

GetVolumeInodes returns the VolumeInodes field if non-nil, zero value otherwise.

### GetVolumeInodesOk

`func (o *O11yVolumeListRecord) GetVolumeInodesOk() (*float64, bool)`

GetVolumeInodesOk returns a tuple with the VolumeInodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeInodes

`func (o *O11yVolumeListRecord) SetVolumeInodes(v float64)`

SetVolumeInodes sets VolumeInodes field to given value.

### HasVolumeInodes

`func (o *O11yVolumeListRecord) HasVolumeInodes() bool`

HasVolumeInodes returns a boolean if a field has been set.

### GetVolumeInodesFree

`func (o *O11yVolumeListRecord) GetVolumeInodesFree() float64`

GetVolumeInodesFree returns the VolumeInodesFree field if non-nil, zero value otherwise.

### GetVolumeInodesFreeOk

`func (o *O11yVolumeListRecord) GetVolumeInodesFreeOk() (*float64, bool)`

GetVolumeInodesFreeOk returns a tuple with the VolumeInodesFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeInodesFree

`func (o *O11yVolumeListRecord) SetVolumeInodesFree(v float64)`

SetVolumeInodesFree sets VolumeInodesFree field to given value.

### HasVolumeInodesFree

`func (o *O11yVolumeListRecord) HasVolumeInodesFree() bool`

HasVolumeInodesFree returns a boolean if a field has been set.

### GetVolumeInodesUsed

`func (o *O11yVolumeListRecord) GetVolumeInodesUsed() float64`

GetVolumeInodesUsed returns the VolumeInodesUsed field if non-nil, zero value otherwise.

### GetVolumeInodesUsedOk

`func (o *O11yVolumeListRecord) GetVolumeInodesUsedOk() (*float64, bool)`

GetVolumeInodesUsedOk returns a tuple with the VolumeInodesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeInodesUsed

`func (o *O11yVolumeListRecord) SetVolumeInodesUsed(v float64)`

SetVolumeInodesUsed sets VolumeInodesUsed field to given value.

### HasVolumeInodesUsed

`func (o *O11yVolumeListRecord) HasVolumeInodesUsed() bool`

HasVolumeInodesUsed returns a boolean if a field has been set.

### GetVolumeUsage

`func (o *O11yVolumeListRecord) GetVolumeUsage() float64`

GetVolumeUsage returns the VolumeUsage field if non-nil, zero value otherwise.

### GetVolumeUsageOk

`func (o *O11yVolumeListRecord) GetVolumeUsageOk() (*float64, bool)`

GetVolumeUsageOk returns a tuple with the VolumeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsage

`func (o *O11yVolumeListRecord) SetVolumeUsage(v float64)`

SetVolumeUsage sets VolumeUsage field to given value.

### HasVolumeUsage

`func (o *O11yVolumeListRecord) HasVolumeUsage() bool`

HasVolumeUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


