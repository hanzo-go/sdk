# DatasetRiskDataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when this version last changed state, RFC 3339 UTC. | [optional] 
**By** | Pointer to **string** | By is who moved it there: the validated user, or the org itself when the caller is a machine with no user behind it. | [optional] 
**Counts** | Pointer to [**DatasetRiskSplitCounts**](DatasetRiskSplitCounts.md) | Counts is how the rows fall across the splits. | [optional] 
**Digest** | Pointer to **string** | Digest fingerprints the SPEC and the ROWS together. Two materialisations of one spec agree on it or the plane says they do not. | [optional] 
**Name** | Pointer to **string** | Name identifies the dataset across all of its versions. | [optional] 
**Oversize** | Pointer to **int64** | Oversize is how many of the window&#39;s subjects this version could NOT carry because their subject identity exceeds the plane&#39;s per-subject byte bound.  It is on the wire, not only in a log, because it is the one degradation a caller cannot otherwise detect: the rows that are here look complete, and a dataset silently missing a population is a model silently blind to it. Non-zero does not make a version invalid — it makes it a version whose coverage is STATED. Zero is the normal case and omits. | [optional] 
**Refusal** | Pointer to **string** | Refusal names why there are no bytes, when there are none. | [optional] 
**Running** | Pointer to **bool** | Running is true while THIS process is materialising the version. A version that is &#x60;materializing&#x60; and not running was started by a process that is gone — two states the register cannot tell apart, because a register cannot know which processes are alive. | [optional] 
**Share** | Pointer to **int64** | Share is the fraction of the window&#39;s subjects admitted, in thousandths. 1000 means the whole window fitted under the cap; anything less means the version is a reproducible sample and says by how much. | [optional] 
**Spec** | Pointer to [**DatasetRiskDatasetSpec**](DatasetRiskDatasetSpec.md) | Spec is the bound query this version was built from, exactly as recorded. | [optional] 
**Status** | Pointer to **string** | Status is declared, materializing, ready or refused. Only &#x60;ready&#x60; has bytes, and &#x60;ready&#x60; is terminal: a published version is never rewritten. | [optional] 
**Truncated** | Pointer to **bool** | Truncated is true when the row cap bound before the window ran out. The trailing subject is dropped whole when that happens, because half a subject on one side of a split is exactly the leak the grouping prevents. | [optional] 
**Version** | Pointer to **int64** | Version is which version this is, from 1 and monotone within the dataset. A number is never reused — not even after a disposal, where the next declare continues the count — so \&quot;signups v3\&quot; means one thing forever, which is what makes a model&#39;s citation of it checkable. | [optional] 

## Methods

### NewDatasetRiskDataset

`func NewDatasetRiskDataset() *DatasetRiskDataset`

NewDatasetRiskDataset instantiates a new DatasetRiskDataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetRiskDatasetWithDefaults

`func NewDatasetRiskDatasetWithDefaults() *DatasetRiskDataset`

NewDatasetRiskDatasetWithDefaults instantiates a new DatasetRiskDataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *DatasetRiskDataset) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *DatasetRiskDataset) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *DatasetRiskDataset) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *DatasetRiskDataset) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *DatasetRiskDataset) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *DatasetRiskDataset) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *DatasetRiskDataset) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *DatasetRiskDataset) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCounts

`func (o *DatasetRiskDataset) GetCounts() DatasetRiskSplitCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *DatasetRiskDataset) GetCountsOk() (*DatasetRiskSplitCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *DatasetRiskDataset) SetCounts(v DatasetRiskSplitCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *DatasetRiskDataset) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetDigest

`func (o *DatasetRiskDataset) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DatasetRiskDataset) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DatasetRiskDataset) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *DatasetRiskDataset) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetName

`func (o *DatasetRiskDataset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetRiskDataset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetRiskDataset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasetRiskDataset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOversize

`func (o *DatasetRiskDataset) GetOversize() int64`

GetOversize returns the Oversize field if non-nil, zero value otherwise.

### GetOversizeOk

`func (o *DatasetRiskDataset) GetOversizeOk() (*int64, bool)`

GetOversizeOk returns a tuple with the Oversize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversize

`func (o *DatasetRiskDataset) SetOversize(v int64)`

SetOversize sets Oversize field to given value.

### HasOversize

`func (o *DatasetRiskDataset) HasOversize() bool`

HasOversize returns a boolean if a field has been set.

### GetRefusal

`func (o *DatasetRiskDataset) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *DatasetRiskDataset) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *DatasetRiskDataset) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *DatasetRiskDataset) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetRunning

`func (o *DatasetRiskDataset) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *DatasetRiskDataset) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *DatasetRiskDataset) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *DatasetRiskDataset) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetShare

`func (o *DatasetRiskDataset) GetShare() int64`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *DatasetRiskDataset) GetShareOk() (*int64, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *DatasetRiskDataset) SetShare(v int64)`

SetShare sets Share field to given value.

### HasShare

`func (o *DatasetRiskDataset) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSpec

`func (o *DatasetRiskDataset) GetSpec() DatasetRiskDatasetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DatasetRiskDataset) GetSpecOk() (*DatasetRiskDatasetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DatasetRiskDataset) SetSpec(v DatasetRiskDatasetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DatasetRiskDataset) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *DatasetRiskDataset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatasetRiskDataset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatasetRiskDataset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatasetRiskDataset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTruncated

`func (o *DatasetRiskDataset) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *DatasetRiskDataset) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *DatasetRiskDataset) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *DatasetRiskDataset) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetVersion

`func (o *DatasetRiskDataset) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatasetRiskDataset) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatasetRiskDataset) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatasetRiskDataset) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


