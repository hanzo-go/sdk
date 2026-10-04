# KnowledgeFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the org bucket the object is in, by the friendly name /v1/s3/buckets lists. | [optional] 
**Chars** | Pointer to **int64** | Chars is the length of the text read out of the file, in bytes. | [optional] 
**Clipped** | Pointer to **bool** | Clipped is true when only the file&#39;s beginning is indexed: its text ran past the org&#39;s bound or the room the index has. Note says how far. | [optional] 
**Created** | Pointer to **int64** | Created is when the file was first registered, in unix seconds. | [optional] 
**Done** | Pointer to **int64** | Done is how far the running stage has come, of Total: bytes of the file read (extract), sections summarized (toc), cut into passages (passages) and linked (graph), passages embedded (embed). Absent between stages and once the ingest is done. | [optional] 
**Embedded** | Pointer to **int64** | Embedded is how many of those passages carry a vector — Passages once the embed stage is done, unless Note says the file is embedded in part. | [optional] 
**Error** | Pointer to **string** | Error is why a stored or failed file was not indexed — or, on a ready file, why it is searched by its words alone — in words a person can act on. Absent otherwise. | [optional] 
**Id** | Pointer to **string** | ID names the file in its org. It is derived from the bucket and key, so registering the same object twice answers the same file. | [optional] 
**Key** | Pointer to **string** | Key is the object&#39;s key in that bucket. GET /v1/s3/buckets/{bucket}/objects/{key} answers a signed download URL for it. | [optional] 
**Name** | Pointer to **string** | Name is the object&#39;s file name, the last segment of its key. | [optional] 
**Note** | Pointer to **string** | Note says in words where the file is indexed less than whole and why: its text past the org&#39;s bound or the room the index has, its passages past the bound on vectors. Absent when the whole file is indexed every way. | [optional] 
**Parent** | Pointer to **string** | Parent is the id of the archive this file was unpacked from. Absent for a file uploaded on its own. | [optional] 
**Passages** | Pointer to **int64** | Passages is how many passages its text was cut into. | [optional] 
**Project** | Pointer to **string** | Project is the project scope it is indexed under. Absent for the org&#39;s own files. | [optional] 
**Sections** | Pointer to **int64** | Sections is how many nodes its table of contents has, the document&#39;s own root included. | [optional] 
**Size** | Pointer to **int64** | Size is the object&#39;s length in bytes, as the store reports it. | [optional] 
**Stage** | Pointer to **string** | Stage is the ingest stage the file is in: extract, toc, passages or graph while indexing, embed while a ready file&#39;s vectors are written. Absent once the ingest is done. | [optional] 
**Status** | Pointer to **string** | Status is queued, indexing, ready, stored (kept but not indexed — Error says why) or failed. A ready file&#39;s contents, passages and links are all readable and its full text is searched; its vectors may still be filling in (Stage embed), which adds search by meaning as it goes. | [optional] 
**Total** | Pointer to **int64** | Total is what the running stage has to do in all, in Done&#39;s units. | [optional] 
**Type** | Pointer to **string** | Type is the object&#39;s media type as the store holds it, or the one its name implies when the store holds only the generic default. | [optional] 
**Updated** | Pointer to **int64** | Updated is when its record last changed, in unix seconds. | [optional] 

## Methods

### NewKnowledgeFile

`func NewKnowledgeFile() *KnowledgeFile`

NewKnowledgeFile instantiates a new KnowledgeFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFileWithDefaults

`func NewKnowledgeFileWithDefaults() *KnowledgeFile`

NewKnowledgeFileWithDefaults instantiates a new KnowledgeFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *KnowledgeFile) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *KnowledgeFile) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *KnowledgeFile) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *KnowledgeFile) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetChars

`func (o *KnowledgeFile) GetChars() int64`

GetChars returns the Chars field if non-nil, zero value otherwise.

### GetCharsOk

`func (o *KnowledgeFile) GetCharsOk() (*int64, bool)`

GetCharsOk returns a tuple with the Chars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChars

`func (o *KnowledgeFile) SetChars(v int64)`

SetChars sets Chars field to given value.

### HasChars

`func (o *KnowledgeFile) HasChars() bool`

HasChars returns a boolean if a field has been set.

### GetClipped

`func (o *KnowledgeFile) GetClipped() bool`

GetClipped returns the Clipped field if non-nil, zero value otherwise.

### GetClippedOk

`func (o *KnowledgeFile) GetClippedOk() (*bool, bool)`

GetClippedOk returns a tuple with the Clipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipped

`func (o *KnowledgeFile) SetClipped(v bool)`

SetClipped sets Clipped field to given value.

### HasClipped

`func (o *KnowledgeFile) HasClipped() bool`

HasClipped returns a boolean if a field has been set.

### GetCreated

`func (o *KnowledgeFile) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *KnowledgeFile) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *KnowledgeFile) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *KnowledgeFile) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDone

`func (o *KnowledgeFile) GetDone() int64`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *KnowledgeFile) GetDoneOk() (*int64, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *KnowledgeFile) SetDone(v int64)`

SetDone sets Done field to given value.

### HasDone

`func (o *KnowledgeFile) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetEmbedded

`func (o *KnowledgeFile) GetEmbedded() int64`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *KnowledgeFile) GetEmbeddedOk() (*int64, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *KnowledgeFile) SetEmbedded(v int64)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *KnowledgeFile) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetError

`func (o *KnowledgeFile) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *KnowledgeFile) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *KnowledgeFile) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *KnowledgeFile) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *KnowledgeFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KnowledgeFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *KnowledgeFile) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KnowledgeFile) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KnowledgeFile) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KnowledgeFile) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *KnowledgeFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KnowledgeFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *KnowledgeFile) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *KnowledgeFile) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *KnowledgeFile) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *KnowledgeFile) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetParent

`func (o *KnowledgeFile) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *KnowledgeFile) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *KnowledgeFile) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *KnowledgeFile) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPassages

`func (o *KnowledgeFile) GetPassages() int64`

GetPassages returns the Passages field if non-nil, zero value otherwise.

### GetPassagesOk

`func (o *KnowledgeFile) GetPassagesOk() (*int64, bool)`

GetPassagesOk returns a tuple with the Passages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassages

`func (o *KnowledgeFile) SetPassages(v int64)`

SetPassages sets Passages field to given value.

### HasPassages

`func (o *KnowledgeFile) HasPassages() bool`

HasPassages returns a boolean if a field has been set.

### GetProject

`func (o *KnowledgeFile) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *KnowledgeFile) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *KnowledgeFile) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *KnowledgeFile) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSections

`func (o *KnowledgeFile) GetSections() int64`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *KnowledgeFile) GetSectionsOk() (*int64, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *KnowledgeFile) SetSections(v int64)`

SetSections sets Sections field to given value.

### HasSections

`func (o *KnowledgeFile) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSize

`func (o *KnowledgeFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *KnowledgeFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *KnowledgeFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *KnowledgeFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStage

`func (o *KnowledgeFile) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *KnowledgeFile) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *KnowledgeFile) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *KnowledgeFile) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStatus

`func (o *KnowledgeFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KnowledgeFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KnowledgeFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KnowledgeFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *KnowledgeFile) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *KnowledgeFile) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *KnowledgeFile) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *KnowledgeFile) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *KnowledgeFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KnowledgeFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KnowledgeFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KnowledgeFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *KnowledgeFile) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *KnowledgeFile) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *KnowledgeFile) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *KnowledgeFile) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


