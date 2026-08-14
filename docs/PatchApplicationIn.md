# PatchApplicationIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the application to move, from the path. | [optional] 
**Note** | Pointer to **string** | Note is a free-text comment recorded on the timeline, with or without a stage change. | [optional] 
**Reason** | Pointer to **string** | Reason records WHY, and is required to reject. | [optional] 
**Stage** | Pointer to **string** | Stage is the stage to move to: applied, screened, qualified, credits-offered, onboarded or rejected. Omit to leave the stage alone. | [optional] 

## Methods

### NewPatchApplicationIn

`func NewPatchApplicationIn() *PatchApplicationIn`

NewPatchApplicationIn instantiates a new PatchApplicationIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApplicationInWithDefaults

`func NewPatchApplicationInWithDefaults() *PatchApplicationIn`

NewPatchApplicationInWithDefaults instantiates a new PatchApplicationIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchApplicationIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchApplicationIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchApplicationIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchApplicationIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNote

`func (o *PatchApplicationIn) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PatchApplicationIn) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PatchApplicationIn) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PatchApplicationIn) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReason

`func (o *PatchApplicationIn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PatchApplicationIn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PatchApplicationIn) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PatchApplicationIn) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStage

`func (o *PatchApplicationIn) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PatchApplicationIn) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PatchApplicationIn) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PatchApplicationIn) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


