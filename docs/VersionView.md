# VersionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when this revision was appended, RFC 3339 UTC. | [optional] 
**Type** | Pointer to **string** | Type is the kind this revision was written with, which may differ from the current one. | [optional] 
**Version** | Pointer to **int64** | Version is this revision&#39;s number, 1 for the first. Numbers are dense and never reused: deleting the prompt drops the whole history with it. | [optional] 

## Methods

### NewVersionView

`func NewVersionView() *VersionView`

NewVersionView instantiates a new VersionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionViewWithDefaults

`func NewVersionViewWithDefaults() *VersionView`

NewVersionViewWithDefaults instantiates a new VersionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *VersionView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VersionView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VersionView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VersionView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetType

`func (o *VersionView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VersionView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *VersionView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


