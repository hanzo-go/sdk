# CaptableNotes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CaptableNote**](CaptableNote.md) | Data is every convertible note, newest first. | [optional] 

## Methods

### NewCaptableNotes

`func NewCaptableNotes() *CaptableNotes`

NewCaptableNotes instantiates a new CaptableNotes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableNotesWithDefaults

`func NewCaptableNotesWithDefaults() *CaptableNotes`

NewCaptableNotesWithDefaults instantiates a new CaptableNotes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CaptableNotes) GetData() []CaptableNote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CaptableNotes) GetDataOk() (*[]CaptableNote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CaptableNotes) SetData(v []CaptableNote)`

SetData sets Data field to given value.

### HasData

`func (o *CaptableNotes) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


