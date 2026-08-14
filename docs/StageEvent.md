# StageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int32** | At is the unix second of the move. | [optional] 
**By** | Pointer to **string** | By is who moved it: \&quot;system\&quot; for intake and the AI auto-advance, else the validated staff user id. | [optional] 
**From** | Pointer to **string** | From is the stage moved out of; empty on the intake event that opens the log. | [optional] 
**Note** | Pointer to **string** | Note is the free-text comment recorded with the move. Absent when none. | [optional] 
**To** | Pointer to **string** | To is the stage moved into. | [optional] 

## Methods

### NewStageEvent

`func NewStageEvent() *StageEvent`

NewStageEvent instantiates a new StageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageEventWithDefaults

`func NewStageEventWithDefaults() *StageEvent`

NewStageEventWithDefaults instantiates a new StageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *StageEvent) GetAt() int32`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *StageEvent) GetAtOk() (*int32, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *StageEvent) SetAt(v int32)`

SetAt sets At field to given value.

### HasAt

`func (o *StageEvent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *StageEvent) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *StageEvent) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *StageEvent) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *StageEvent) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetFrom

`func (o *StageEvent) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *StageEvent) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *StageEvent) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *StageEvent) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetNote

`func (o *StageEvent) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *StageEvent) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *StageEvent) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *StageEvent) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTo

`func (o *StageEvent) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *StageEvent) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *StageEvent) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *StageEvent) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


