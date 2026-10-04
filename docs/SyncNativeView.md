# SyncNativeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedAt** | Pointer to **string** | AdvancedAt is when an import or a push last moved one of its refs, RFC3339 in UTC. Absent until one has. | [optional] 
**Branch** | Pointer to **string** | Branch is the forge&#39;s default branch, which an import sets to the upstream&#39;s once that branch has landed. | [optional] 
**Clone** | Pointer to **string** | Clone is the https address to clone it from. | [optional] 
**SizeBytes** | Pointer to **int64** | SizeBytes is the repository&#39;s size on the forge, in bytes. | [optional] 
**Ssh** | Pointer to **string** | SSH is the ssh address to clone it from, when the forge serves one. | [optional] 
**Status** | Pointer to **string** | Status is one of \&quot;synced\&quot; (the forge holds it and no ref is in conflict), \&quot;conflict\&quot; (an upstream ref diverged and the forge kept its own history), \&quot;pending\&quot; (the forge does not hold it yet) or \&quot;paused\&quot; (the link&#39;s direction is off). | [optional] 
**Url** | Pointer to **string** | URL is the repository&#39;s page on the forge. | [optional] 

## Methods

### NewSyncNativeView

`func NewSyncNativeView() *SyncNativeView`

NewSyncNativeView instantiates a new SyncNativeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncNativeViewWithDefaults

`func NewSyncNativeViewWithDefaults() *SyncNativeView`

NewSyncNativeViewWithDefaults instantiates a new SyncNativeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedAt

`func (o *SyncNativeView) GetAdvancedAt() string`

GetAdvancedAt returns the AdvancedAt field if non-nil, zero value otherwise.

### GetAdvancedAtOk

`func (o *SyncNativeView) GetAdvancedAtOk() (*string, bool)`

GetAdvancedAtOk returns a tuple with the AdvancedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedAt

`func (o *SyncNativeView) SetAdvancedAt(v string)`

SetAdvancedAt sets AdvancedAt field to given value.

### HasAdvancedAt

`func (o *SyncNativeView) HasAdvancedAt() bool`

HasAdvancedAt returns a boolean if a field has been set.

### GetBranch

`func (o *SyncNativeView) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *SyncNativeView) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *SyncNativeView) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *SyncNativeView) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetClone

`func (o *SyncNativeView) GetClone() string`

GetClone returns the Clone field if non-nil, zero value otherwise.

### GetCloneOk

`func (o *SyncNativeView) GetCloneOk() (*string, bool)`

GetCloneOk returns a tuple with the Clone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClone

`func (o *SyncNativeView) SetClone(v string)`

SetClone sets Clone field to given value.

### HasClone

`func (o *SyncNativeView) HasClone() bool`

HasClone returns a boolean if a field has been set.

### GetSizeBytes

`func (o *SyncNativeView) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *SyncNativeView) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *SyncNativeView) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *SyncNativeView) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetSsh

`func (o *SyncNativeView) GetSsh() string`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *SyncNativeView) GetSshOk() (*string, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *SyncNativeView) SetSsh(v string)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *SyncNativeView) HasSsh() bool`

HasSsh returns a boolean if a field has been set.

### GetStatus

`func (o *SyncNativeView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncNativeView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncNativeView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyncNativeView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *SyncNativeView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyncNativeView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyncNativeView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyncNativeView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


