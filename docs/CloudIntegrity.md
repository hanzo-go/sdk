# CloudIntegrity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokenAt** | Pointer to **int32** | BrokenAt is the seq of the FIRST record that failed verification, or -1 when OK. Reason describes the break (recomputed-hash mismatch, prev-hash discontinuity, or a seq gap). | [optional] 
**Count** | Pointer to **int32** | Count is the number of records walked. | [optional] 
**HeadHash** | Pointer to **string** | HeadHash is the hash of the last record (or the genesis anchor for an empty chain). Pin this externally over time to detect tail-truncation. | [optional] 
**Ok** | Pointer to **bool** | OK is true iff every record&#39;s stored hash equals the recomputed hash AND the chain links are continuous (each PrevHash &#x3D;&#x3D; the prior record&#39;s Hash, seqs gapless from 0). | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudIntegrity

`func NewCloudIntegrity() *CloudIntegrity`

NewCloudIntegrity instantiates a new CloudIntegrity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIntegrityWithDefaults

`func NewCloudIntegrityWithDefaults() *CloudIntegrity`

NewCloudIntegrityWithDefaults instantiates a new CloudIntegrity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokenAt

`func (o *CloudIntegrity) GetBrokenAt() int32`

GetBrokenAt returns the BrokenAt field if non-nil, zero value otherwise.

### GetBrokenAtOk

`func (o *CloudIntegrity) GetBrokenAtOk() (*int32, bool)`

GetBrokenAtOk returns a tuple with the BrokenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenAt

`func (o *CloudIntegrity) SetBrokenAt(v int32)`

SetBrokenAt sets BrokenAt field to given value.

### HasBrokenAt

`func (o *CloudIntegrity) HasBrokenAt() bool`

HasBrokenAt returns a boolean if a field has been set.

### GetCount

`func (o *CloudIntegrity) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudIntegrity) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudIntegrity) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudIntegrity) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHeadHash

`func (o *CloudIntegrity) GetHeadHash() string`

GetHeadHash returns the HeadHash field if non-nil, zero value otherwise.

### GetHeadHashOk

`func (o *CloudIntegrity) GetHeadHashOk() (*string, bool)`

GetHeadHashOk returns a tuple with the HeadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadHash

`func (o *CloudIntegrity) SetHeadHash(v string)`

SetHeadHash sets HeadHash field to given value.

### HasHeadHash

`func (o *CloudIntegrity) HasHeadHash() bool`

HasHeadHash returns a boolean if a field has been set.

### GetOk

`func (o *CloudIntegrity) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *CloudIntegrity) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *CloudIntegrity) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *CloudIntegrity) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetReason

`func (o *CloudIntegrity) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudIntegrity) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudIntegrity) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudIntegrity) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


