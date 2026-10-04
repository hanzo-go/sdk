# LinkSourceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available reports whether this ledger answered; false is honest \&quot;unavailable\&quot;, never a zero that would read as no usage. | [optional] 
**Note** | Pointer to **string** | Note says in prose what the ledger&#39;s numbers mean. | [optional] 
**Scope** | Pointer to **string** | Scope is whose usage the ledger measures: user or org. | [optional] 
**Source** | Pointer to **string** | Source is the table of record behind the ledger. | [optional] 

## Methods

### NewLinkSourceState

`func NewLinkSourceState() *LinkSourceState`

NewLinkSourceState instantiates a new LinkSourceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkSourceStateWithDefaults

`func NewLinkSourceStateWithDefaults() *LinkSourceState`

NewLinkSourceStateWithDefaults instantiates a new LinkSourceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LinkSourceState) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LinkSourceState) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LinkSourceState) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LinkSourceState) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetNote

`func (o *LinkSourceState) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LinkSourceState) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LinkSourceState) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *LinkSourceState) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetScope

`func (o *LinkSourceState) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LinkSourceState) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LinkSourceState) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *LinkSourceState) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *LinkSourceState) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LinkSourceState) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LinkSourceState) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LinkSourceState) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


