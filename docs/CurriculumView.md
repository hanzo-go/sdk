# CurriculumView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Curriculum** | Pointer to [**Curriculum**](Curriculum.md) | Curriculum is the enabled journey: its version, title and ordered steps. | [optional] 
**Custom** | Pointer to **bool** | Custom is true when the org&#39;s OWN curriculum override is active; false when the journey comes from the brand blueprint or the embedded fixture. | [optional] 

## Methods

### NewCurriculumView

`func NewCurriculumView() *CurriculumView`

NewCurriculumView instantiates a new CurriculumView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurriculumViewWithDefaults

`func NewCurriculumViewWithDefaults() *CurriculumView`

NewCurriculumViewWithDefaults instantiates a new CurriculumView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurriculum

`func (o *CurriculumView) GetCurriculum() Curriculum`

GetCurriculum returns the Curriculum field if non-nil, zero value otherwise.

### GetCurriculumOk

`func (o *CurriculumView) GetCurriculumOk() (*Curriculum, bool)`

GetCurriculumOk returns a tuple with the Curriculum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculum

`func (o *CurriculumView) SetCurriculum(v Curriculum)`

SetCurriculum sets Curriculum field to given value.

### HasCurriculum

`func (o *CurriculumView) HasCurriculum() bool`

HasCurriculum returns a boolean if a field has been set.

### GetCustom

`func (o *CurriculumView) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CurriculumView) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CurriculumView) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CurriculumView) HasCustom() bool`

HasCustom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


