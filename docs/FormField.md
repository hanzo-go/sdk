# FormField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | Pointer to [**[]FormChoice**](FormChoice.md) | Choices are what a multiple_choice or dropdown field offers, in order. | [optional] 
**Description** | Pointer to **string** | Description is the line under the question. At most 1000 characters. | [optional] 
**Id** | Pointer to **string** | ID is the field&#39;s id: assigned when a write leaves it empty, and kept across edits by sending it back. | [optional] 
**Labels** | Pointer to [**FormLabels**](FormLabels.md) | Labels name the ends and middle of an opinion_scale. | [optional] 
**Multiple** | Pointer to **bool** | Multiple is whether a multiple_choice field takes several choices. | [optional] 
**Ref** | Pointer to **string** | Ref is the field&#39;s stable name, how an answer and a sink name it: lowercase letters, digits, dash and underscore, at most 64, unique in the form. Defaults to the id. | [optional] 
**Required** | Pointer to **bool** | Required is whether a completed response must answer it. A statement is never required. | [optional] 
**Start** | Pointer to **int64** | Start is where an opinion_scale begins: 0 (the default) or 1. | [optional] 
**Steps** | Pointer to **int64** | Steps is how many points a rating (3 to 10, default 5) or an opinion_scale (5 to 11, default 11) has. | [optional] 
**Title** | Pointer to **string** | Title is the question as asked. At most 500 characters. | [optional] 
**Type** | Pointer to **string** | Type is one of short_text, long_text, email, phone, url, number, rating, opinion_scale, yes_no, multiple_choice, dropdown, date and statement. | [optional] 
**Validations** | Pointer to [**FormValidations**](FormValidations.md) | Validations bound an answer beyond its type. | [optional] 

## Methods

### NewFormField

`func NewFormField() *FormField`

NewFormField instantiates a new FormField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldWithDefaults

`func NewFormFieldWithDefaults() *FormField`

NewFormFieldWithDefaults instantiates a new FormField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *FormField) GetChoices() []FormChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *FormField) GetChoicesOk() (*[]FormChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *FormField) SetChoices(v []FormChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *FormField) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetDescription

`func (o *FormField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FormField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *FormField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *FormField) GetLabels() FormLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FormField) GetLabelsOk() (*FormLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FormField) SetLabels(v FormLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FormField) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMultiple

`func (o *FormField) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *FormField) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *FormField) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *FormField) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetRef

`func (o *FormField) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *FormField) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *FormField) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *FormField) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRequired

`func (o *FormField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetStart

`func (o *FormField) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FormField) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FormField) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *FormField) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSteps

`func (o *FormField) GetSteps() int64`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *FormField) GetStepsOk() (*int64, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *FormField) SetSteps(v int64)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *FormField) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTitle

`func (o *FormField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormField) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormField) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *FormField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidations

`func (o *FormField) GetValidations() FormValidations`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *FormField) GetValidationsOk() (*FormValidations, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *FormField) SetValidations(v FormValidations)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *FormField) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


