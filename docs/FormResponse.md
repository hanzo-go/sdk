# FormResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]FormAnswer**](FormAnswer.md) | Answers are the answers, in the form&#39;s field order. | [optional] 
**Completed** | Pointer to **bool** | Completed is whether the response was submitted whole; false is one saved part way. | [optional] 
**Created** | Pointer to **string** | Created is when the record was first written, RFC 3339. | [optional] 
**Duration** | Pointer to **int64** | Duration is how many seconds lay between landing and submitting. | [optional] 
**External** | Pointer to **string** | External is an imported response&#39;s id where it came from. | [optional] 
**Filed** | Pointer to **bool** | Filed is whether a completed response is on the CRM&#39;s timeline. A filing that failed is tried again on a schedule, and POST .../file tries it now. | [optional] 
**Filing** | Pointer to **string** | Filing is why the last filing failed; empty once it is filed. | [optional] 
**Form** | Pointer to **string** | Form is the form&#39;s id. | [optional] 
**Hidden** | Pointer to **map[string]string** | Hidden are the values the page passed in, by name. | [optional] 
**Id** | Pointer to **string** | ID is the response&#39;s id: derived from the form and the submission&#39;s token (from the form and the address on a subscribe form), so a retried submission is this same response. | [optional] 
**LandedAt** | Pointer to **string** | LandedAt is when the respondent opened the form, RFC 3339, as their page said. | [optional] 
**Path** | Pointer to **string** | Path is the page the form was answered on. | [optional] 
**Respondent** | Pointer to [**FormRespondent**](FormRespondent.md) | Respondent is who it is from. | [optional] 
**Source** | Pointer to **string** | Source is web, or import for a response brought in from elsewhere. | [optional] 
**SubmittedAt** | Pointer to **string** | SubmittedAt is when the response was completed, RFC 3339. | [optional] 
**Token** | Pointer to **string** | Token is the submission token the response was made under. | [optional] 
**Updated** | Pointer to **string** | Updated is when it last changed, RFC 3339. | [optional] 

## Methods

### NewFormResponse

`func NewFormResponse() *FormResponse`

NewFormResponse instantiates a new FormResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormResponseWithDefaults

`func NewFormResponseWithDefaults() *FormResponse`

NewFormResponseWithDefaults instantiates a new FormResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *FormResponse) GetAnswers() []FormAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *FormResponse) GetAnswersOk() (*[]FormAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *FormResponse) SetAnswers(v []FormAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *FormResponse) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetCompleted

`func (o *FormResponse) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *FormResponse) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *FormResponse) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *FormResponse) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCreated

`func (o *FormResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FormResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FormResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FormResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDuration

`func (o *FormResponse) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *FormResponse) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *FormResponse) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *FormResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetExternal

`func (o *FormResponse) GetExternal() string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *FormResponse) GetExternalOk() (*string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *FormResponse) SetExternal(v string)`

SetExternal sets External field to given value.

### HasExternal

`func (o *FormResponse) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetFiled

`func (o *FormResponse) GetFiled() bool`

GetFiled returns the Filed field if non-nil, zero value otherwise.

### GetFiledOk

`func (o *FormResponse) GetFiledOk() (*bool, bool)`

GetFiledOk returns a tuple with the Filed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiled

`func (o *FormResponse) SetFiled(v bool)`

SetFiled sets Filed field to given value.

### HasFiled

`func (o *FormResponse) HasFiled() bool`

HasFiled returns a boolean if a field has been set.

### GetFiling

`func (o *FormResponse) GetFiling() string`

GetFiling returns the Filing field if non-nil, zero value otherwise.

### GetFilingOk

`func (o *FormResponse) GetFilingOk() (*string, bool)`

GetFilingOk returns a tuple with the Filing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiling

`func (o *FormResponse) SetFiling(v string)`

SetFiling sets Filing field to given value.

### HasFiling

`func (o *FormResponse) HasFiling() bool`

HasFiling returns a boolean if a field has been set.

### GetForm

`func (o *FormResponse) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *FormResponse) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *FormResponse) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *FormResponse) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetHidden

`func (o *FormResponse) GetHidden() map[string]string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FormResponse) GetHiddenOk() (*map[string]string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FormResponse) SetHidden(v map[string]string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FormResponse) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *FormResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLandedAt

`func (o *FormResponse) GetLandedAt() string`

GetLandedAt returns the LandedAt field if non-nil, zero value otherwise.

### GetLandedAtOk

`func (o *FormResponse) GetLandedAtOk() (*string, bool)`

GetLandedAtOk returns a tuple with the LandedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandedAt

`func (o *FormResponse) SetLandedAt(v string)`

SetLandedAt sets LandedAt field to given value.

### HasLandedAt

`func (o *FormResponse) HasLandedAt() bool`

HasLandedAt returns a boolean if a field has been set.

### GetPath

`func (o *FormResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FormResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FormResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FormResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRespondent

`func (o *FormResponse) GetRespondent() FormRespondent`

GetRespondent returns the Respondent field if non-nil, zero value otherwise.

### GetRespondentOk

`func (o *FormResponse) GetRespondentOk() (*FormRespondent, bool)`

GetRespondentOk returns a tuple with the Respondent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondent

`func (o *FormResponse) SetRespondent(v FormRespondent)`

SetRespondent sets Respondent field to given value.

### HasRespondent

`func (o *FormResponse) HasRespondent() bool`

HasRespondent returns a boolean if a field has been set.

### GetSource

`func (o *FormResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FormResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FormResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FormResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *FormResponse) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *FormResponse) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *FormResponse) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *FormResponse) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetToken

`func (o *FormResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FormResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FormResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FormResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdated

`func (o *FormResponse) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FormResponse) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FormResponse) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FormResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


