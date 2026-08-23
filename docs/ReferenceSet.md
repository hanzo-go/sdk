# ReferenceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to **string** | Age is how long ago that was. | [optional] 
**AsOf** | Pointer to **string** | AsOf is when the OLDEST contributing publisher was current, RFC 3339. The oldest and not the newest: a set is exactly as fresh as its weakest source. | [optional] 
**Keys** | Pointer to **int32** | Keys is how many members the baseline carries. | [optional] 
**Kind** | Pointer to **string** | Kind is how the baseline comes to exist: fetch (downloaded from a publisher), local (computed here), attest (held by the component that screens against it, freshness reported), or client (declared and NOT held, because the source needs a licence we do not have). | [optional] 
**Match** | Pointer to **string** | Match is how a key is tested: exact, domain, net, digits, pattern or range. | [optional] 
**MaxAge** | Pointer to **string** | MaxAge is how old this set may be before it is stale. | [optional] 
**Overrides** | Pointer to **int32** | Overrides is how many entries YOUR org has laid over this baseline. | [optional] 
**Refusal** | Pointer to **string** | Refusal names why the set cannot be relied on, when it cannot: never loaded, held elsewhere, or a licence we do not hold. Non-empty means a lookup against this set will not answer, rather than answering clean. | [optional] 
**Set** | Pointer to **string** | Set is the name this set is addressed by. | [optional] 
**Sources** | Pointer to [**[]ReferenceSource**](ReferenceSource.md) | Sources is each contributing publisher, its licence and its own freshness. | [optional] 
**Stale** | Pointer to **bool** | Stale is whether it is past that bound. A stale set still answers and says so, because yesterday&#39;s list beats none. | [optional] 
**Version** | Pointer to **string** | Version is the exact baseline consulted — every contributing publisher and its content digest. A decision records this and an auditor resolves it back. | [optional] 
**What** | Pointer to **string** | What the set holds, in one sentence. | [optional] 

## Methods

### NewReferenceSet

`func NewReferenceSet() *ReferenceSet`

NewReferenceSet instantiates a new ReferenceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceSetWithDefaults

`func NewReferenceSetWithDefaults() *ReferenceSet`

NewReferenceSetWithDefaults instantiates a new ReferenceSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *ReferenceSet) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ReferenceSet) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ReferenceSet) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *ReferenceSet) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAsOf

`func (o *ReferenceSet) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ReferenceSet) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ReferenceSet) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ReferenceSet) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetKeys

`func (o *ReferenceSet) GetKeys() int32`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ReferenceSet) GetKeysOk() (*int32, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ReferenceSet) SetKeys(v int32)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ReferenceSet) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetKind

`func (o *ReferenceSet) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReferenceSet) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReferenceSet) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ReferenceSet) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMatch

`func (o *ReferenceSet) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *ReferenceSet) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *ReferenceSet) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *ReferenceSet) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMaxAge

`func (o *ReferenceSet) GetMaxAge() string`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *ReferenceSet) GetMaxAgeOk() (*string, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *ReferenceSet) SetMaxAge(v string)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *ReferenceSet) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetOverrides

`func (o *ReferenceSet) GetOverrides() int32`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *ReferenceSet) GetOverridesOk() (*int32, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *ReferenceSet) SetOverrides(v int32)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *ReferenceSet) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetRefusal

`func (o *ReferenceSet) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *ReferenceSet) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *ReferenceSet) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *ReferenceSet) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetSet

`func (o *ReferenceSet) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *ReferenceSet) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *ReferenceSet) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *ReferenceSet) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetSources

`func (o *ReferenceSet) GetSources() []ReferenceSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ReferenceSet) GetSourcesOk() (*[]ReferenceSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ReferenceSet) SetSources(v []ReferenceSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *ReferenceSet) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetStale

`func (o *ReferenceSet) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ReferenceSet) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ReferenceSet) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ReferenceSet) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetVersion

`func (o *ReferenceSet) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReferenceSet) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReferenceSet) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReferenceSet) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWhat

`func (o *ReferenceSet) GetWhat() string`

GetWhat returns the What field if non-nil, zero value otherwise.

### GetWhatOk

`func (o *ReferenceSet) GetWhatOk() (*string, bool)`

GetWhatOk returns a tuple with the What field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhat

`func (o *ReferenceSet) SetWhat(v string)`

SetWhat sets What field to given value.

### HasWhat

`func (o *ReferenceSet) HasWhat() bool`

HasWhat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


