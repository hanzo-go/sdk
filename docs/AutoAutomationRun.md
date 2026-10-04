# AutoAutomationRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when it started and Finished when it ended (null while it runs), RFC 3339 UTC. | [optional] 
**Finished** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Status is succeeded, failed, running, queued, skipped for a start that found the previous run still going, or refused for a run whose person is no longer a member of the org. | [optional] 
**Summary** | Pointer to **string** | Summary is one line on how it went. | [optional] 
**Transcript** | Pointer to **string** | Transcript opens the run&#39;s Dev run; null when it started none. | [optional] 

## Methods

### NewAutoAutomationRun

`func NewAutoAutomationRun() *AutoAutomationRun`

NewAutoAutomationRun instantiates a new AutoAutomationRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoAutomationRunWithDefaults

`func NewAutoAutomationRunWithDefaults() *AutoAutomationRun`

NewAutoAutomationRunWithDefaults instantiates a new AutoAutomationRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *AutoAutomationRun) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AutoAutomationRun) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AutoAutomationRun) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *AutoAutomationRun) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetFinished

`func (o *AutoAutomationRun) GetFinished() string`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *AutoAutomationRun) GetFinishedOk() (*string, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *AutoAutomationRun) SetFinished(v string)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *AutoAutomationRun) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *AutoAutomationRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoAutomationRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoAutomationRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoAutomationRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AutoAutomationRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoAutomationRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoAutomationRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutoAutomationRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *AutoAutomationRun) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AutoAutomationRun) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AutoAutomationRun) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AutoAutomationRun) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTranscript

`func (o *AutoAutomationRun) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *AutoAutomationRun) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *AutoAutomationRun) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *AutoAutomationRun) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


