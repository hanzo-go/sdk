# RiskDataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when this version last changed state, RFC 3339 UTC. | [optional] 
**By** | Pointer to **string** | By is who moved it there: the validated user, or the org itself when the caller is a machine with no user behind it. | [optional] 
**Counts** | Pointer to [**RiskSplitCounts**](RiskSplitCounts.md) | Counts is how the rows fall across the splits. | [optional] 
**Digest** | Pointer to **string** | Digest fingerprints the SPEC and the ROWS together. Two materialisations of one spec agree on it or the plane says they do not. | [optional] 
**Name** | Pointer to **string** | Name identifies the dataset across all of its versions. | [optional] 
**Oversize** | Pointer to **int32** | Oversize is how many of the window&#39;s subjects this version could NOT carry because their subject identity exceeds the plane&#39;s per-subject byte bound.  It is on the wire, not only in a log, because it is the one degradation a caller cannot otherwise detect: the rows that are here look complete, and a dataset silently missing a population is a model silently blind to it. Non-zero does not make a version invalid — it makes it a version whose coverage is STATED. Zero is the normal case and omits. | [optional] 
**Refusal** | Pointer to **string** | Refusal names why there are no bytes, when there are none. | [optional] 
**Running** | Pointer to **bool** | Running is true while THIS process is materialising the version. A version that is &#x60;materializing&#x60; and not running was started by a process that is gone — two states the register cannot tell apart, because a register cannot know which processes are alive. | [optional] 
**Share** | Pointer to **int32** | Share is the fraction of the window&#39;s subjects admitted, in thousandths. 1000 means the whole window fitted under the cap; anything less means the version is a reproducible sample and says by how much. | [optional] 
**Spec** | Pointer to [**RiskDatasetSpec**](RiskDatasetSpec.md) | Spec is the bound query this version was built from, exactly as recorded. | [optional] 
**Status** | Pointer to **string** | Status is declared, materializing, ready or refused. Only &#x60;ready&#x60; has bytes, and &#x60;ready&#x60; is terminal: a published version is never rewritten. | [optional] 
**Truncated** | Pointer to **bool** | Truncated is true when the row cap bound before the window ran out. The trailing subject is dropped whole when that happens, because half a subject on one side of a split is exactly the leak the grouping prevents. | [optional] 
**Version** | Pointer to **int32** | Version is which version this is, from 1 and monotone within the dataset. A number is never reused — not even after a disposal, where the next declare continues the count — so \&quot;signups v3\&quot; means one thing forever, which is what makes a model&#39;s citation of it checkable. | [optional] 

## Methods

### NewRiskDataset

`func NewRiskDataset() *RiskDataset`

NewRiskDataset instantiates a new RiskDataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDatasetWithDefaults

`func NewRiskDatasetWithDefaults() *RiskDataset`

NewRiskDatasetWithDefaults instantiates a new RiskDataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *RiskDataset) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RiskDataset) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RiskDataset) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *RiskDataset) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *RiskDataset) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *RiskDataset) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *RiskDataset) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *RiskDataset) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCounts

`func (o *RiskDataset) GetCounts() RiskSplitCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *RiskDataset) GetCountsOk() (*RiskSplitCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *RiskDataset) SetCounts(v RiskSplitCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *RiskDataset) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetDigest

`func (o *RiskDataset) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RiskDataset) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RiskDataset) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *RiskDataset) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetName

`func (o *RiskDataset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskDataset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskDataset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskDataset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOversize

`func (o *RiskDataset) GetOversize() int32`

GetOversize returns the Oversize field if non-nil, zero value otherwise.

### GetOversizeOk

`func (o *RiskDataset) GetOversizeOk() (*int32, bool)`

GetOversizeOk returns a tuple with the Oversize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversize

`func (o *RiskDataset) SetOversize(v int32)`

SetOversize sets Oversize field to given value.

### HasOversize

`func (o *RiskDataset) HasOversize() bool`

HasOversize returns a boolean if a field has been set.

### GetRefusal

`func (o *RiskDataset) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *RiskDataset) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *RiskDataset) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *RiskDataset) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetRunning

`func (o *RiskDataset) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *RiskDataset) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *RiskDataset) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *RiskDataset) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetShare

`func (o *RiskDataset) GetShare() int32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *RiskDataset) GetShareOk() (*int32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *RiskDataset) SetShare(v int32)`

SetShare sets Share field to given value.

### HasShare

`func (o *RiskDataset) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSpec

`func (o *RiskDataset) GetSpec() RiskDatasetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RiskDataset) GetSpecOk() (*RiskDatasetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RiskDataset) SetSpec(v RiskDatasetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RiskDataset) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *RiskDataset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RiskDataset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RiskDataset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RiskDataset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTruncated

`func (o *RiskDataset) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *RiskDataset) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *RiskDataset) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *RiskDataset) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetVersion

`func (o *RiskDataset) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RiskDataset) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RiskDataset) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RiskDataset) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


