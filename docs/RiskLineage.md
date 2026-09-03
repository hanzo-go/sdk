# RiskLineage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** | Dataset is the dataset traced. | [optional] 
**Digest** | Pointer to **string** | Digest is the version&#39;s fingerprint, repeated here so a lineage answer is self-contained. | [optional] 
**From** | Pointer to **string** | From is where the window actually read opens, RFC 3339. Same as the spec&#39;s. | [optional] 
**Holds** | Pointer to **int64** | Holds is what the source holds for the same window NOW. The difference between it and Rows is the whole of the reproducibility claim. | [optional] 
**Oversize** | Pointer to **int64** | Oversize is how many subjects the window held that were too large to represent when this version was built. It is part of the fingerprint, so it is part of what \&quot;reproducible\&quot; is measured over. | [optional] 
**Refusal** | Pointer to **string** | Refusal says which way it failed — the window expired, or the source now holds a different count. Absent when Reproducible is true. | [optional] 
**Reproducible** | Pointer to **bool** | Reproducible is true when the source still holds what this version was built from — measured by asking it again, not recalled. False is ordinary: the source is fed by a rollup that runs behind the events, so \&quot;it holds more now\&quot; is the common case and it means re-running the spec would not produce this version. | [optional] 
**Retention** | Pointer to **string** | Retention is the source&#39;s own expiry rule as the store reports it, read at materialisation time rather than assumed. A source whose retention is shorter than this window cannot re-derive it. | [optional] 
**Rows** | Pointer to **int64** | Rows is how many rows the source held for that window at materialisation time. Holds is the same question asked now, and the difference between them is the whole of the reproducibility claim. | [optional] 
**Share** | Pointer to **int64** | Share is the fraction of subjects admitted, in thousandths. | [optional] 
**Source** | Pointer to **string** | Source is the plane the rows were derived from. | [optional] 
**Subjects** | Pointer to **int64** | Subjects is how many distinct subjects those rows belonged to. It is the real sample size — the row count flatters it whenever a subject is active. | [optional] 
**To** | Pointer to **string** | To is where it ends: the spec&#39;s own end pulled BACK by the maturity horizon, so it is usually earlier than the spec says. This is the window a reproduction has to ask for — asking the spec&#39;s would not return these rows. | [optional] 
**Version** | Pointer to **int64** | Version is the version traced — the one asked for, or the newest published one when the request named none. | [optional] 

## Methods

### NewRiskLineage

`func NewRiskLineage() *RiskLineage`

NewRiskLineage instantiates a new RiskLineage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLineageWithDefaults

`func NewRiskLineageWithDefaults() *RiskLineage`

NewRiskLineageWithDefaults instantiates a new RiskLineage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *RiskLineage) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *RiskLineage) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *RiskLineage) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *RiskLineage) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetDigest

`func (o *RiskLineage) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RiskLineage) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RiskLineage) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *RiskLineage) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetFrom

`func (o *RiskLineage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RiskLineage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RiskLineage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RiskLineage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHolds

`func (o *RiskLineage) GetHolds() int64`

GetHolds returns the Holds field if non-nil, zero value otherwise.

### GetHoldsOk

`func (o *RiskLineage) GetHoldsOk() (*int64, bool)`

GetHoldsOk returns a tuple with the Holds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolds

`func (o *RiskLineage) SetHolds(v int64)`

SetHolds sets Holds field to given value.

### HasHolds

`func (o *RiskLineage) HasHolds() bool`

HasHolds returns a boolean if a field has been set.

### GetOversize

`func (o *RiskLineage) GetOversize() int64`

GetOversize returns the Oversize field if non-nil, zero value otherwise.

### GetOversizeOk

`func (o *RiskLineage) GetOversizeOk() (*int64, bool)`

GetOversizeOk returns a tuple with the Oversize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversize

`func (o *RiskLineage) SetOversize(v int64)`

SetOversize sets Oversize field to given value.

### HasOversize

`func (o *RiskLineage) HasOversize() bool`

HasOversize returns a boolean if a field has been set.

### GetRefusal

`func (o *RiskLineage) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *RiskLineage) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *RiskLineage) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *RiskLineage) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetReproducible

`func (o *RiskLineage) GetReproducible() bool`

GetReproducible returns the Reproducible field if non-nil, zero value otherwise.

### GetReproducibleOk

`func (o *RiskLineage) GetReproducibleOk() (*bool, bool)`

GetReproducibleOk returns a tuple with the Reproducible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReproducible

`func (o *RiskLineage) SetReproducible(v bool)`

SetReproducible sets Reproducible field to given value.

### HasReproducible

`func (o *RiskLineage) HasReproducible() bool`

HasReproducible returns a boolean if a field has been set.

### GetRetention

`func (o *RiskLineage) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *RiskLineage) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *RiskLineage) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *RiskLineage) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetRows

`func (o *RiskLineage) GetRows() int64`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *RiskLineage) GetRowsOk() (*int64, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *RiskLineage) SetRows(v int64)`

SetRows sets Rows field to given value.

### HasRows

`func (o *RiskLineage) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetShare

`func (o *RiskLineage) GetShare() int64`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *RiskLineage) GetShareOk() (*int64, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *RiskLineage) SetShare(v int64)`

SetShare sets Share field to given value.

### HasShare

`func (o *RiskLineage) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSource

`func (o *RiskLineage) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskLineage) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskLineage) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskLineage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubjects

`func (o *RiskLineage) GetSubjects() int64`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RiskLineage) GetSubjectsOk() (*int64, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RiskLineage) SetSubjects(v int64)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *RiskLineage) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetTo

`func (o *RiskLineage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RiskLineage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RiskLineage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RiskLineage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVersion

`func (o *RiskLineage) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RiskLineage) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RiskLineage) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RiskLineage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


