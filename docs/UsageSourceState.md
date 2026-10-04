# UsageSourceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is whether this side&#39;s ledger answered. False means its rows are missing, not that there were none. | [optional] 
**Note** | Pointer to **string** | Note is the human sentence that says what this side&#39;s numbers mean, so a board cannot present a plan percentage as a Hanzo charge. | [optional] 
**Scope** | Pointer to **string** | Scope is whose rows this side carries: \&quot;user\&quot; or \&quot;org\&quot;. | [optional] 
**Source** | Pointer to **string** | Source is the table of record the rows came from. | [optional] 

## Methods

### NewUsageSourceState

`func NewUsageSourceState() *UsageSourceState`

NewUsageSourceState instantiates a new UsageSourceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSourceStateWithDefaults

`func NewUsageSourceStateWithDefaults() *UsageSourceState`

NewUsageSourceStateWithDefaults instantiates a new UsageSourceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *UsageSourceState) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *UsageSourceState) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *UsageSourceState) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *UsageSourceState) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetNote

`func (o *UsageSourceState) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UsageSourceState) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UsageSourceState) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UsageSourceState) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetScope

`func (o *UsageSourceState) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UsageSourceState) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UsageSourceState) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UsageSourceState) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *UsageSourceState) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageSourceState) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageSourceState) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UsageSourceState) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


