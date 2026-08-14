# MemoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the validated user id that last wrote this entry by hand. Empty on an entry an engine produced, and on one written before attribution existed. | [optional] 
**GlossaryVersion** | Pointer to **string** | Glossary is the glossary VERSION the entry was translated under — the digest version() derives from the terms, so changing a term changes the key and the stale rendering can never be served. | [optional] 
**Source** | Pointer to **string** | Source is the ORIGINAL string this entry translates. Part of the identity. | [optional] 
**State** | Pointer to **string** | State is the entry&#39;s position on the review ladder: machine, suggested, approved or published. | [optional] 
**Target** | Pointer to **string** | Target is the target language tag (BCP-47, e.g. \&quot;es\&quot; or \&quot;pt-BR\&quot;). Part of the identity. | [optional] 
**Text** | Pointer to **string** | Text is the stored translation. A memory hit returns it verbatim, which is the idempotence contract. | [optional] 
**Tier** | Pointer to **string** | Tier is the engine tier the entry belongs to, quality or bulk. Part of the identity: the two tiers keep separate renderings of the same source. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second the entry last changed. | [optional] 

## Methods

### NewMemoryEntry

`func NewMemoryEntry() *MemoryEntry`

NewMemoryEntry instantiates a new MemoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryEntryWithDefaults

`func NewMemoryEntryWithDefaults() *MemoryEntry`

NewMemoryEntryWithDefaults instantiates a new MemoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *MemoryEntry) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *MemoryEntry) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *MemoryEntry) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *MemoryEntry) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetGlossaryVersion

`func (o *MemoryEntry) GetGlossaryVersion() string`

GetGlossaryVersion returns the GlossaryVersion field if non-nil, zero value otherwise.

### GetGlossaryVersionOk

`func (o *MemoryEntry) GetGlossaryVersionOk() (*string, bool)`

GetGlossaryVersionOk returns a tuple with the GlossaryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossaryVersion

`func (o *MemoryEntry) SetGlossaryVersion(v string)`

SetGlossaryVersion sets GlossaryVersion field to given value.

### HasGlossaryVersion

`func (o *MemoryEntry) HasGlossaryVersion() bool`

HasGlossaryVersion returns a boolean if a field has been set.

### GetSource

`func (o *MemoryEntry) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MemoryEntry) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MemoryEntry) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MemoryEntry) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *MemoryEntry) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MemoryEntry) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MemoryEntry) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MemoryEntry) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTarget

`func (o *MemoryEntry) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MemoryEntry) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MemoryEntry) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MemoryEntry) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetText

`func (o *MemoryEntry) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MemoryEntry) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MemoryEntry) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MemoryEntry) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTier

`func (o *MemoryEntry) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *MemoryEntry) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *MemoryEntry) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *MemoryEntry) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MemoryEntry) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MemoryEntry) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MemoryEntry) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MemoryEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


