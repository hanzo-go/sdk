# PatentPatentTimelineEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthoredBy** | Pointer to **string** | AuthoredBy is the validated subject who appended it. | [optional] 
**Canonical** | Pointer to **bool** | Canonical marks the version a listing would return: the latest non-retracted one. Exactly one entry per id is canonical, or none, when the latest retracts. | [optional] 
**ContentHash** | Pointer to **string** | ContentHash is this version&#39;s identity. | [optional] 
**RecordedAt** | Pointer to **string** | RecordedAt is the server clock at the append. | [optional] 
**Revision** | Pointer to **string** | Revision is original, corrected or retracted. | [optional] 
**Seq** | Pointer to **int64** | Seq is the append clock value this version was stamped with. | [optional] 
**Superseded** | Pointer to **bool** | Superseded marks a retained version that a later one replaced. Retained and superseded are both true of most of a long history, and that is the point. | [optional] 

## Methods

### NewPatentPatentTimelineEntry

`func NewPatentPatentTimelineEntry() *PatentPatentTimelineEntry`

NewPatentPatentTimelineEntry instantiates a new PatentPatentTimelineEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentTimelineEntryWithDefaults

`func NewPatentPatentTimelineEntryWithDefaults() *PatentPatentTimelineEntry`

NewPatentPatentTimelineEntryWithDefaults instantiates a new PatentPatentTimelineEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthoredBy

`func (o *PatentPatentTimelineEntry) GetAuthoredBy() string`

GetAuthoredBy returns the AuthoredBy field if non-nil, zero value otherwise.

### GetAuthoredByOk

`func (o *PatentPatentTimelineEntry) GetAuthoredByOk() (*string, bool)`

GetAuthoredByOk returns a tuple with the AuthoredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredBy

`func (o *PatentPatentTimelineEntry) SetAuthoredBy(v string)`

SetAuthoredBy sets AuthoredBy field to given value.

### HasAuthoredBy

`func (o *PatentPatentTimelineEntry) HasAuthoredBy() bool`

HasAuthoredBy returns a boolean if a field has been set.

### GetCanonical

`func (o *PatentPatentTimelineEntry) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *PatentPatentTimelineEntry) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *PatentPatentTimelineEntry) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *PatentPatentTimelineEntry) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetContentHash

`func (o *PatentPatentTimelineEntry) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *PatentPatentTimelineEntry) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *PatentPatentTimelineEntry) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *PatentPatentTimelineEntry) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetRecordedAt

`func (o *PatentPatentTimelineEntry) GetRecordedAt() string`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *PatentPatentTimelineEntry) GetRecordedAtOk() (*string, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *PatentPatentTimelineEntry) SetRecordedAt(v string)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *PatentPatentTimelineEntry) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetRevision

`func (o *PatentPatentTimelineEntry) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PatentPatentTimelineEntry) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PatentPatentTimelineEntry) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PatentPatentTimelineEntry) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSeq

`func (o *PatentPatentTimelineEntry) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PatentPatentTimelineEntry) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PatentPatentTimelineEntry) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PatentPatentTimelineEntry) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSuperseded

`func (o *PatentPatentTimelineEntry) GetSuperseded() bool`

GetSuperseded returns the Superseded field if non-nil, zero value otherwise.

### GetSupersededOk

`func (o *PatentPatentTimelineEntry) GetSupersededOk() (*bool, bool)`

GetSupersededOk returns a tuple with the Superseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperseded

`func (o *PatentPatentTimelineEntry) SetSuperseded(v bool)`

SetSuperseded sets Superseded field to given value.

### HasSuperseded

`func (o *PatentPatentTimelineEntry) HasSuperseded() bool`

HasSuperseded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


