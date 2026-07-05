# GuardSanitizeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Sanitization result status | [optional] 
**Text** | Pointer to **string** | Sanitized text (present for clean and redacted) | [optional] 
**Reason** | Pointer to **string** | Block reason (present only for blocked) | [optional] 
**Redactions** | Pointer to [**[]GuardSanitizeResultRedactionsInner**](GuardSanitizeResultRedactionsInner.md) | PII redactions applied | [optional] 
**Injection** | Pointer to [**GuardSanitizeResultInjection**](GuardSanitizeResultInjection.md) |  | [optional] 
**ContentFilter** | Pointer to [**GuardSanitizeResultContentFilter**](GuardSanitizeResultContentFilter.md) |  | [optional] 
**ProcessingTimeUs** | Pointer to **int32** | Processing time in microseconds | [optional] 

## Methods

### NewGuardSanitizeResult

`func NewGuardSanitizeResult() *GuardSanitizeResult`

NewGuardSanitizeResult instantiates a new GuardSanitizeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardSanitizeResultWithDefaults

`func NewGuardSanitizeResultWithDefaults() *GuardSanitizeResult`

NewGuardSanitizeResultWithDefaults instantiates a new GuardSanitizeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GuardSanitizeResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GuardSanitizeResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GuardSanitizeResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GuardSanitizeResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetText

`func (o *GuardSanitizeResult) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GuardSanitizeResult) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GuardSanitizeResult) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GuardSanitizeResult) HasText() bool`

HasText returns a boolean if a field has been set.

### GetReason

`func (o *GuardSanitizeResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GuardSanitizeResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GuardSanitizeResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GuardSanitizeResult) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRedactions

`func (o *GuardSanitizeResult) GetRedactions() []GuardSanitizeResultRedactionsInner`

GetRedactions returns the Redactions field if non-nil, zero value otherwise.

### GetRedactionsOk

`func (o *GuardSanitizeResult) GetRedactionsOk() (*[]GuardSanitizeResultRedactionsInner, bool)`

GetRedactionsOk returns a tuple with the Redactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactions

`func (o *GuardSanitizeResult) SetRedactions(v []GuardSanitizeResultRedactionsInner)`

SetRedactions sets Redactions field to given value.

### HasRedactions

`func (o *GuardSanitizeResult) HasRedactions() bool`

HasRedactions returns a boolean if a field has been set.

### GetInjection

`func (o *GuardSanitizeResult) GetInjection() GuardSanitizeResultInjection`

GetInjection returns the Injection field if non-nil, zero value otherwise.

### GetInjectionOk

`func (o *GuardSanitizeResult) GetInjectionOk() (*GuardSanitizeResultInjection, bool)`

GetInjectionOk returns a tuple with the Injection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjection

`func (o *GuardSanitizeResult) SetInjection(v GuardSanitizeResultInjection)`

SetInjection sets Injection field to given value.

### HasInjection

`func (o *GuardSanitizeResult) HasInjection() bool`

HasInjection returns a boolean if a field has been set.

### GetContentFilter

`func (o *GuardSanitizeResult) GetContentFilter() GuardSanitizeResultContentFilter`

GetContentFilter returns the ContentFilter field if non-nil, zero value otherwise.

### GetContentFilterOk

`func (o *GuardSanitizeResult) GetContentFilterOk() (*GuardSanitizeResultContentFilter, bool)`

GetContentFilterOk returns a tuple with the ContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilter

`func (o *GuardSanitizeResult) SetContentFilter(v GuardSanitizeResultContentFilter)`

SetContentFilter sets ContentFilter field to given value.

### HasContentFilter

`func (o *GuardSanitizeResult) HasContentFilter() bool`

HasContentFilter returns a boolean if a field has been set.

### GetProcessingTimeUs

`func (o *GuardSanitizeResult) GetProcessingTimeUs() int32`

GetProcessingTimeUs returns the ProcessingTimeUs field if non-nil, zero value otherwise.

### GetProcessingTimeUsOk

`func (o *GuardSanitizeResult) GetProcessingTimeUsOk() (*int32, bool)`

GetProcessingTimeUsOk returns a tuple with the ProcessingTimeUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeUs

`func (o *GuardSanitizeResult) SetProcessingTimeUs(v int32)`

SetProcessingTimeUs sets ProcessingTimeUs field to given value.

### HasProcessingTimeUs

`func (o *GuardSanitizeResult) HasProcessingTimeUs() bool`

HasProcessingTimeUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


