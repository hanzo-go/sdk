# TrainVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** |  | [optional] 
**BaseSha256** | Pointer to **string** | BaseSHA256 is Kai&#39;s weights it was taught over. | [optional] 
**By** | Pointer to **string** |  | [optional] 
**Bytes** | Pointer to **int64** | Bytes is its artifact&#39;s size. | [optional] 
**Corrections** | Pointer to **[]string** | Corrections are the corrections it was taught from. | [optional] 
**Evaluation** | Pointer to **interface{}** |  | [optional] 
**Form** | Pointer to **string** | Form is calibration, head or pair; Rank its place, 1 the smallest. | [optional] 
**Passed** | Pointer to **bool** | Passed says its evaluation passed the gate: it is publishable. | [optional] 
**Questions** | Pointer to **[]string** | Questions are the question keys it answers. | [optional] 
**Rank** | Pointer to **int64** |  | [optional] 
**Representation** | Pointer to **string** | Representation is what it reads of Kai (the weights, tokenizer, layout and pooling, hashed): the only Kai it serves on, whatever the release. | [optional] 
**Sha256** | Pointer to **string** | SHA256 is its artifact&#39;s: the tar the decision service installs. | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewTrainVersion

`func NewTrainVersion() *TrainVersion`

NewTrainVersion instantiates a new TrainVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainVersionWithDefaults

`func NewTrainVersionWithDefaults() *TrainVersion`

NewTrainVersionWithDefaults instantiates a new TrainVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *TrainVersion) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *TrainVersion) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *TrainVersion) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *TrainVersion) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBaseSha256

`func (o *TrainVersion) GetBaseSha256() string`

GetBaseSha256 returns the BaseSha256 field if non-nil, zero value otherwise.

### GetBaseSha256Ok

`func (o *TrainVersion) GetBaseSha256Ok() (*string, bool)`

GetBaseSha256Ok returns a tuple with the BaseSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSha256

`func (o *TrainVersion) SetBaseSha256(v string)`

SetBaseSha256 sets BaseSha256 field to given value.

### HasBaseSha256

`func (o *TrainVersion) HasBaseSha256() bool`

HasBaseSha256 returns a boolean if a field has been set.

### GetBy

`func (o *TrainVersion) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *TrainVersion) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *TrainVersion) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *TrainVersion) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetBytes

`func (o *TrainVersion) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *TrainVersion) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *TrainVersion) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *TrainVersion) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCorrections

`func (o *TrainVersion) GetCorrections() []string`

GetCorrections returns the Corrections field if non-nil, zero value otherwise.

### GetCorrectionsOk

`func (o *TrainVersion) GetCorrectionsOk() (*[]string, bool)`

GetCorrectionsOk returns a tuple with the Corrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrections

`func (o *TrainVersion) SetCorrections(v []string)`

SetCorrections sets Corrections field to given value.

### HasCorrections

`func (o *TrainVersion) HasCorrections() bool`

HasCorrections returns a boolean if a field has been set.

### GetEvaluation

`func (o *TrainVersion) GetEvaluation() interface{}`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *TrainVersion) GetEvaluationOk() (*interface{}, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *TrainVersion) SetEvaluation(v interface{})`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *TrainVersion) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### SetEvaluationNil

`func (o *TrainVersion) SetEvaluationNil(b bool)`

 SetEvaluationNil sets the value for Evaluation to be an explicit nil

### UnsetEvaluation
`func (o *TrainVersion) UnsetEvaluation()`

UnsetEvaluation ensures that no value is present for Evaluation, not even an explicit nil
### GetForm

`func (o *TrainVersion) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *TrainVersion) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *TrainVersion) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *TrainVersion) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetPassed

`func (o *TrainVersion) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *TrainVersion) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *TrainVersion) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *TrainVersion) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### GetQuestions

`func (o *TrainVersion) GetQuestions() []string`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *TrainVersion) GetQuestionsOk() (*[]string, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *TrainVersion) SetQuestions(v []string)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *TrainVersion) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetRank

`func (o *TrainVersion) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TrainVersion) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TrainVersion) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *TrainVersion) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRepresentation

`func (o *TrainVersion) GetRepresentation() string`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *TrainVersion) GetRepresentationOk() (*string, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *TrainVersion) SetRepresentation(v string)`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *TrainVersion) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.

### GetSha256

`func (o *TrainVersion) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *TrainVersion) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *TrainVersion) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *TrainVersion) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetVersion

`func (o *TrainVersion) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrainVersion) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrainVersion) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrainVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


