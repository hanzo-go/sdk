# DataroomTrustDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int64** | Days is how long the grant stays open, from now. Optional; 14 by default and 365 at most — a longer release is describing a customer relationship rather than a document. | [optional] 
**Id** | Pointer to **string** | ID is the request to answer, taken from the path. | [optional] 
**Note** | Pointer to **string** | Note is why. Recorded on the request either way, and it is what the record shows a year later. | [optional] 

## Methods

### NewDataroomTrustDecision

`func NewDataroomTrustDecision() *DataroomTrustDecision`

NewDataroomTrustDecision instantiates a new DataroomTrustDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustDecisionWithDefaults

`func NewDataroomTrustDecisionWithDefaults() *DataroomTrustDecision`

NewDataroomTrustDecisionWithDefaults instantiates a new DataroomTrustDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *DataroomTrustDecision) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *DataroomTrustDecision) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *DataroomTrustDecision) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *DataroomTrustDecision) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetId

`func (o *DataroomTrustDecision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomTrustDecision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomTrustDecision) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomTrustDecision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNote

`func (o *DataroomTrustDecision) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *DataroomTrustDecision) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *DataroomTrustDecision) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *DataroomTrustDecision) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


