# RiskLabelOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplicate** | Pointer to **int64** | Duplicate is how many members this tenant already held, byte for byte. The idempotency key is the assertion&#39;s CONTENT digest — kind, subject, at, seen, disposition, source, evidence, the asserting identity and confidence, folded in length-prefixed — so a webhook redelivering one chargeback is a duplicate and costs nothing, while an assertion differing in ANY of those fields is a DIFFERENT assertion and is recorded beside the first. Nothing was written and nothing was overwritten; it is an outcome, never an error. The asserting identity is in the digest, so the same claim filed by a second credential is two assertions and not a redelivery. | [optional] 
**Mirror** | Pointer to **string** | Mirror names why the columnar copy did not take this batch, when it did not. The record is already durable in the tenant&#39;s own store by then — the warehouse copy exists to make a training join cheap, and its absence is a gap in that join, never a lost label. | [optional] 
**Pending** | Pointer to **int64** | Pending is how many assertions the derived copy is still to take. Every write attempt carries the backlog forward as well as its own batch, so a warehouse that was unreachable closes its gap on the next write rather than leaving a hole in a training join nothing would report. It is counted under a cap and saturates there: zero means caught up, and a large number means a backlog to work through rather than an inventory to reconcile. | [optional] 
**Recorded** | Pointer to **int64** | Recorded is how many members became a NEW row in the tenant&#39;s record. Recorded + Duplicate + Refused is exactly the number of labels sent, so a caller reconciling a webhook delivery can do it on the counts alone. | [optional] 
**Refused** | Pointer to **int64** | Refused is how many members failed admission and were NOT recorded. Refusal is per member and never discards the rest of the batch: an empty or over-512-byte subject or evidence, a kind, disposition or source outside the closed vocabulary, an &#x60;at&#x60; or &#x60;seen&#x60; that is not RFC 3339, a &#x60;seen&#x60; before the &#x60;at&#x60; it judges, either instant more than five minutes past the server clock, or a confidence outside [0,1]. Results names which member and why, so the refused ones are exactly the ones to fix and resend. | [optional] 
**Results** | Pointer to [**[]RiskLabelResult**](RiskLabelResult.md) | Results is per fact, in the order sent, so a caller can retry exactly the members that were refused and can log the content digest of the ones that landed. | [optional] 

## Methods

### NewRiskLabelOut

`func NewRiskLabelOut() *RiskLabelOut`

NewRiskLabelOut instantiates a new RiskLabelOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLabelOutWithDefaults

`func NewRiskLabelOutWithDefaults() *RiskLabelOut`

NewRiskLabelOutWithDefaults instantiates a new RiskLabelOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplicate

`func (o *RiskLabelOut) GetDuplicate() int64`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *RiskLabelOut) GetDuplicateOk() (*int64, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *RiskLabelOut) SetDuplicate(v int64)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *RiskLabelOut) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetMirror

`func (o *RiskLabelOut) GetMirror() string`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *RiskLabelOut) GetMirrorOk() (*string, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *RiskLabelOut) SetMirror(v string)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *RiskLabelOut) HasMirror() bool`

HasMirror returns a boolean if a field has been set.

### GetPending

`func (o *RiskLabelOut) GetPending() int64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *RiskLabelOut) GetPendingOk() (*int64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *RiskLabelOut) SetPending(v int64)`

SetPending sets Pending field to given value.

### HasPending

`func (o *RiskLabelOut) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetRecorded

`func (o *RiskLabelOut) GetRecorded() int64`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *RiskLabelOut) GetRecordedOk() (*int64, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *RiskLabelOut) SetRecorded(v int64)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *RiskLabelOut) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetRefused

`func (o *RiskLabelOut) GetRefused() int64`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *RiskLabelOut) GetRefusedOk() (*int64, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *RiskLabelOut) SetRefused(v int64)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *RiskLabelOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetResults

`func (o *RiskLabelOut) GetResults() []RiskLabelResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RiskLabelOut) GetResultsOk() (*[]RiskLabelResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RiskLabelOut) SetResults(v []RiskLabelResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *RiskLabelOut) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


