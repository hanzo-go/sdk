# CloudGcOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Maintained** | Pointer to **bool** | Maintained is always true; the call fails rather than reporting false. | [optional] 
**Repo** | Pointer to **string** | Repo is the repo that was repacked. | [optional] 
**SizeBytes** | Pointer to **int32** | SizeBytes is the size measured AFTER the repack — usually smaller, since repacking drops the packs it supersedes. | [optional] 

## Methods

### NewCloudGcOut

`func NewCloudGcOut() *CloudGcOut`

NewCloudGcOut instantiates a new CloudGcOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGcOutWithDefaults

`func NewCloudGcOutWithDefaults() *CloudGcOut`

NewCloudGcOutWithDefaults instantiates a new CloudGcOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintained

`func (o *CloudGcOut) GetMaintained() bool`

GetMaintained returns the Maintained field if non-nil, zero value otherwise.

### GetMaintainedOk

`func (o *CloudGcOut) GetMaintainedOk() (*bool, bool)`

GetMaintainedOk returns a tuple with the Maintained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintained

`func (o *CloudGcOut) SetMaintained(v bool)`

SetMaintained sets Maintained field to given value.

### HasMaintained

`func (o *CloudGcOut) HasMaintained() bool`

HasMaintained returns a boolean if a field has been set.

### GetRepo

`func (o *CloudGcOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudGcOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudGcOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudGcOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSizeBytes

`func (o *CloudGcOut) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *CloudGcOut) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *CloudGcOut) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *CloudGcOut) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


