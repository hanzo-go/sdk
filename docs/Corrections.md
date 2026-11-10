# Corrections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Correction**](Correction.md) |  | [optional] 
**Questions** | Pointer to [**[]Captured**](Captured.md) | Questions are the decision&#39;s captured questions, by key; none when it was made with capture off. | [optional] 

## Methods

### NewCorrections

`func NewCorrections() *Corrections`

NewCorrections instantiates a new Corrections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrectionsWithDefaults

`func NewCorrectionsWithDefaults() *Corrections`

NewCorrectionsWithDefaults instantiates a new Corrections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Corrections) GetData() []Correction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Corrections) GetDataOk() (*[]Correction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Corrections) SetData(v []Correction)`

SetData sets Data field to given value.

### HasData

`func (o *Corrections) HasData() bool`

HasData returns a boolean if a field has been set.

### GetQuestions

`func (o *Corrections) GetQuestions() []Captured`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *Corrections) GetQuestionsOk() (*[]Captured, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *Corrections) SetQuestions(v []Captured)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *Corrections) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


