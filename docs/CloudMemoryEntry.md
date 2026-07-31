# CloudMemoryEntry

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

### NewCloudMemoryEntry

`func NewCloudMemoryEntry() *CloudMemoryEntry`

NewCloudMemoryEntry instantiates a new CloudMemoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMemoryEntryWithDefaults

`func NewCloudMemoryEntryWithDefaults() *CloudMemoryEntry`

NewCloudMemoryEntryWithDefaults instantiates a new CloudMemoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *CloudMemoryEntry) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudMemoryEntry) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudMemoryEntry) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudMemoryEntry) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetGlossaryVersion

`func (o *CloudMemoryEntry) GetGlossaryVersion() string`

GetGlossaryVersion returns the GlossaryVersion field if non-nil, zero value otherwise.

### GetGlossaryVersionOk

`func (o *CloudMemoryEntry) GetGlossaryVersionOk() (*string, bool)`

GetGlossaryVersionOk returns a tuple with the GlossaryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossaryVersion

`func (o *CloudMemoryEntry) SetGlossaryVersion(v string)`

SetGlossaryVersion sets GlossaryVersion field to given value.

### HasGlossaryVersion

`func (o *CloudMemoryEntry) HasGlossaryVersion() bool`

HasGlossaryVersion returns a boolean if a field has been set.

### GetSource

`func (o *CloudMemoryEntry) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudMemoryEntry) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudMemoryEntry) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudMemoryEntry) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *CloudMemoryEntry) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudMemoryEntry) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudMemoryEntry) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudMemoryEntry) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTarget

`func (o *CloudMemoryEntry) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudMemoryEntry) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudMemoryEntry) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudMemoryEntry) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetText

`func (o *CloudMemoryEntry) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CloudMemoryEntry) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CloudMemoryEntry) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CloudMemoryEntry) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTier

`func (o *CloudMemoryEntry) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CloudMemoryEntry) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CloudMemoryEntry) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CloudMemoryEntry) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudMemoryEntry) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudMemoryEntry) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudMemoryEntry) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudMemoryEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


