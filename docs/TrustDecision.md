# TrustDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int32** | Days is how long the grant stays open, from now. Optional; 14 by default and 365 at most — a longer release is describing a customer relationship rather than a document. | [optional] 
**Id** | Pointer to **string** | ID is the request to answer, taken from the path. | [optional] 
**Note** | Pointer to **string** | Note is why. Recorded on the request either way, and it is what the record shows a year later. | [optional] 

## Methods

### NewTrustDecision

`func NewTrustDecision() *TrustDecision`

NewTrustDecision instantiates a new TrustDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustDecisionWithDefaults

`func NewTrustDecisionWithDefaults() *TrustDecision`

NewTrustDecisionWithDefaults instantiates a new TrustDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *TrustDecision) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *TrustDecision) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *TrustDecision) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *TrustDecision) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetId

`func (o *TrustDecision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustDecision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustDecision) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrustDecision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNote

`func (o *TrustDecision) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TrustDecision) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TrustDecision) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *TrustDecision) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


