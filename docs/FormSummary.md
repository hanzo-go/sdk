# FormSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionRate** | Pointer to **float64** | CompletionRate is completions over starts, 0 to 1. | [optional] 
**Completions** | Pointer to **int64** | Completions is how many were completed. | [optional] 
**Fields** | Pointer to [**[]FormFieldSummary**](FormFieldSummary.md) | Fields are each asked field&#39;s figures, in form order. | [optional] 
**Form** | Pointer to **string** | Form is the form&#39;s id. | [optional] 
**Scanned** | Pointer to **int64** | Scanned is how many responses were read; Truncated says the form has more than one summary reads (50000), and the newest were read. | [optional] 
**Starts** | Pointer to **int64** | Starts is how many responses were begun: open and completed. | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**Views** | Pointer to **int64** | Views is null: a view is stated on the event stream as form_viewed and counted there, not kept with the responses. | [optional] 

## Methods

### NewFormSummary

`func NewFormSummary() *FormSummary`

NewFormSummary instantiates a new FormSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSummaryWithDefaults

`func NewFormSummaryWithDefaults() *FormSummary`

NewFormSummaryWithDefaults instantiates a new FormSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionRate

`func (o *FormSummary) GetCompletionRate() float64`

GetCompletionRate returns the CompletionRate field if non-nil, zero value otherwise.

### GetCompletionRateOk

`func (o *FormSummary) GetCompletionRateOk() (*float64, bool)`

GetCompletionRateOk returns a tuple with the CompletionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionRate

`func (o *FormSummary) SetCompletionRate(v float64)`

SetCompletionRate sets CompletionRate field to given value.

### HasCompletionRate

`func (o *FormSummary) HasCompletionRate() bool`

HasCompletionRate returns a boolean if a field has been set.

### GetCompletions

`func (o *FormSummary) GetCompletions() int64`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *FormSummary) GetCompletionsOk() (*int64, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *FormSummary) SetCompletions(v int64)`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *FormSummary) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### GetFields

`func (o *FormSummary) GetFields() []FormFieldSummary`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FormSummary) GetFieldsOk() (*[]FormFieldSummary, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FormSummary) SetFields(v []FormFieldSummary)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FormSummary) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetForm

`func (o *FormSummary) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *FormSummary) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *FormSummary) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *FormSummary) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetScanned

`func (o *FormSummary) GetScanned() int64`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *FormSummary) GetScannedOk() (*int64, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *FormSummary) SetScanned(v int64)`

SetScanned sets Scanned field to given value.

### HasScanned

`func (o *FormSummary) HasScanned() bool`

HasScanned returns a boolean if a field has been set.

### GetStarts

`func (o *FormSummary) GetStarts() int64`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *FormSummary) GetStartsOk() (*int64, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *FormSummary) SetStarts(v int64)`

SetStarts sets Starts field to given value.

### HasStarts

`func (o *FormSummary) HasStarts() bool`

HasStarts returns a boolean if a field has been set.

### GetTruncated

`func (o *FormSummary) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *FormSummary) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *FormSummary) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *FormSummary) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetViews

`func (o *FormSummary) GetViews() int64`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *FormSummary) GetViewsOk() (*int64, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *FormSummary) SetViews(v int64)`

SetViews sets Views field to given value.

### HasViews

`func (o *FormSummary) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


