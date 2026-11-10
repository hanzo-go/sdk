# Correction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is the right answer; Base what Kai answered. | [optional] 
**At** | Pointer to **int64** |  | [optional] 
**Base** | Pointer to **string** |  | [optional] 
**By** | Pointer to **string** |  | [optional] 
**Cohort** | Pointer to **string** |  | [optional] 
**Decision** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | State is captured when the decision&#39;s state is held and the correction can teach, missing when the decision was made with capture off. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCorrection

`func NewCorrection() *Correction`

NewCorrection instantiates a new Correction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrectionWithDefaults

`func NewCorrectionWithDefaults() *Correction`

NewCorrectionWithDefaults instantiates a new Correction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *Correction) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *Correction) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *Correction) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *Correction) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetAt

`func (o *Correction) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *Correction) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *Correction) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *Correction) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBase

`func (o *Correction) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *Correction) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *Correction) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *Correction) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBy

`func (o *Correction) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *Correction) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *Correction) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *Correction) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCohort

`func (o *Correction) GetCohort() string`

GetCohort returns the Cohort field if non-nil, zero value otherwise.

### GetCohortOk

`func (o *Correction) GetCohortOk() (*string, bool)`

GetCohortOk returns a tuple with the Cohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohort

`func (o *Correction) SetCohort(v string)`

SetCohort sets Cohort field to given value.

### HasCohort

`func (o *Correction) HasCohort() bool`

HasCohort returns a boolean if a field has been set.

### GetDecision

`func (o *Correction) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *Correction) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *Correction) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *Correction) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetId

`func (o *Correction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Correction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Correction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Correction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *Correction) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Correction) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Correction) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Correction) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNote

`func (o *Correction) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Correction) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Correction) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Correction) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetQuestion

`func (o *Correction) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *Correction) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *Correction) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *Correction) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetState

`func (o *Correction) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Correction) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Correction) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Correction) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *Correction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Correction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Correction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Correction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


