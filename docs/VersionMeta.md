# VersionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Brand is the white-label key this revision was authored under; empty is the shared base playbook. Revisions of two brands never share a number line. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when this revision was written, as Unix seconds — the \&quot;who changed the playbook, and when\&quot; half of the audit trail. | [optional] 
**Version** | Pointer to **int64** | Version is the store&#39;s own revision counter for that brand, starting at 1 for the seeded playbook and incrementing on every edit. Nothing is overwritten, so the highest number is the live one and every lower number is still readable. It is not the playbook&#39;s authored &#x60;version&#x60; string. | [optional] 

## Methods

### NewVersionMeta

`func NewVersionMeta() *VersionMeta`

NewVersionMeta instantiates a new VersionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionMetaWithDefaults

`func NewVersionMetaWithDefaults() *VersionMeta`

NewVersionMetaWithDefaults instantiates a new VersionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *VersionMeta) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *VersionMeta) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *VersionMeta) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *VersionMeta) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VersionMeta) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VersionMeta) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VersionMeta) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VersionMeta) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *VersionMeta) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionMeta) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionMeta) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionMeta) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


