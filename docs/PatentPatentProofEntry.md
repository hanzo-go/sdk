# PatentPatentProofEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthoredBy** | Pointer to **string** | AuthoredBy is the validated subject on the record. | [optional] 
**ContentHash** | Pointer to **string** | ContentHash is the server&#39;s digest of the recorded content. | [optional] 
**Digest** | Pointer to **string** | Digest is what the fold consumed: content hash and clock together, so two versions of one id with identical content still differ in the fold. | [optional] 
**Id** | Pointer to **string** | ID is the record&#39;s stable id. | [optional] 
**Kind** | Pointer to **string** | Kind is idea, implementation, disclosure or approval. | [optional] 
**RecordedAt** | Pointer to **string** | RecordedAt is the server clock at the append. | [optional] 
**Revision** | Pointer to **string** | Revision is original, corrected or retracted. | [optional] 
**Seq** | Pointer to **int64** | Seq is the append clock value. The ordering authority. | [optional] 

## Methods

### NewPatentPatentProofEntry

`func NewPatentPatentProofEntry() *PatentPatentProofEntry`

NewPatentPatentProofEntry instantiates a new PatentPatentProofEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentProofEntryWithDefaults

`func NewPatentPatentProofEntryWithDefaults() *PatentPatentProofEntry`

NewPatentPatentProofEntryWithDefaults instantiates a new PatentPatentProofEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthoredBy

`func (o *PatentPatentProofEntry) GetAuthoredBy() string`

GetAuthoredBy returns the AuthoredBy field if non-nil, zero value otherwise.

### GetAuthoredByOk

`func (o *PatentPatentProofEntry) GetAuthoredByOk() (*string, bool)`

GetAuthoredByOk returns a tuple with the AuthoredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredBy

`func (o *PatentPatentProofEntry) SetAuthoredBy(v string)`

SetAuthoredBy sets AuthoredBy field to given value.

### HasAuthoredBy

`func (o *PatentPatentProofEntry) HasAuthoredBy() bool`

HasAuthoredBy returns a boolean if a field has been set.

### GetContentHash

`func (o *PatentPatentProofEntry) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *PatentPatentProofEntry) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *PatentPatentProofEntry) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *PatentPatentProofEntry) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetDigest

`func (o *PatentPatentProofEntry) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PatentPatentProofEntry) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PatentPatentProofEntry) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *PatentPatentProofEntry) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetId

`func (o *PatentPatentProofEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatentPatentProofEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatentPatentProofEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatentPatentProofEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *PatentPatentProofEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatentPatentProofEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatentPatentProofEntry) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatentPatentProofEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRecordedAt

`func (o *PatentPatentProofEntry) GetRecordedAt() string`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *PatentPatentProofEntry) GetRecordedAtOk() (*string, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *PatentPatentProofEntry) SetRecordedAt(v string)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *PatentPatentProofEntry) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetRevision

`func (o *PatentPatentProofEntry) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PatentPatentProofEntry) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PatentPatentProofEntry) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PatentPatentProofEntry) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSeq

`func (o *PatentPatentProofEntry) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PatentPatentProofEntry) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PatentPatentProofEntry) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PatentPatentProofEntry) HasSeq() bool`

HasSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


