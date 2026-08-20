# Integrity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokenAt** | Pointer to **int32** | BrokenAt is the seq of the FIRST record that failed verification, and -1 whenever the walk found no break (including an unread chain, where no seq was reached). Reason describes the break (recomputed-hash mismatch, prev-hash discontinuity, or a seq gap) or why the chain could not be read. | [optional] 
**Count** | Pointer to **int32** | Count is the number of records walked. Zero on an unread chain, where it means \&quot;nothing was read\&quot;, not \&quot;the chain is empty\&quot;. | [optional] 
**Head** | Pointer to **string** | Head is the hash of the last record (or the genesis anchor for an empty chain). Pin this externally over time to detect tail-truncation. | [optional] 
**Name** | Pointer to **string** | Name is the chain this verdict is about, e.g. \&quot;audit\&quot; or \&quot;audit-iam\&quot;. It is carried because a verdict with no chain on it reads as the whole trail&#39;s, which is what a reader of a 128-chain deployment did. | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Verdict** | Pointer to **string** | Verdict is intact, broken or unread. | [optional] 

## Methods

### NewIntegrity

`func NewIntegrity() *Integrity`

NewIntegrity instantiates a new Integrity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrityWithDefaults

`func NewIntegrityWithDefaults() *Integrity`

NewIntegrityWithDefaults instantiates a new Integrity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokenAt

`func (o *Integrity) GetBrokenAt() int32`

GetBrokenAt returns the BrokenAt field if non-nil, zero value otherwise.

### GetBrokenAtOk

`func (o *Integrity) GetBrokenAtOk() (*int32, bool)`

GetBrokenAtOk returns a tuple with the BrokenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenAt

`func (o *Integrity) SetBrokenAt(v int32)`

SetBrokenAt sets BrokenAt field to given value.

### HasBrokenAt

`func (o *Integrity) HasBrokenAt() bool`

HasBrokenAt returns a boolean if a field has been set.

### GetCount

`func (o *Integrity) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Integrity) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Integrity) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Integrity) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHead

`func (o *Integrity) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *Integrity) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *Integrity) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *Integrity) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetName

`func (o *Integrity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integrity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integrity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Integrity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReason

`func (o *Integrity) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Integrity) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Integrity) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Integrity) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetVerdict

`func (o *Integrity) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *Integrity) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *Integrity) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *Integrity) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


