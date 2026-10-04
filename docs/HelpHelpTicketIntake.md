# HelpHelpTicketIntake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is the customer&#39;s message. Optional; it becomes the ticket&#39;s description AND the opening entry of its conversation thread. Clipped at 16 KiB. | [optional] 
**Email** | Pointer to **string** | Email is how the support team replies. Required; clipped at 320 characters (the RFC 5321 maximum). It is recorded as the ticket&#39;s customer, and it is not verified. | [optional] 
**Priority** | Pointer to **string** | Priority is Low, Medium, High or Urgent, case-insensitively. Anything else — including omitting it — is recorded as Medium rather than refused. | [optional] 
**Subject** | Pointer to **string** | Subject is the one-line summary of the problem. Required; longer than 300 characters is clipped rather than refused. | [optional] 

## Methods

### NewHelpHelpTicketIntake

`func NewHelpHelpTicketIntake() *HelpHelpTicketIntake`

NewHelpHelpTicketIntake instantiates a new HelpHelpTicketIntake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpHelpTicketIntakeWithDefaults

`func NewHelpHelpTicketIntakeWithDefaults() *HelpHelpTicketIntake`

NewHelpHelpTicketIntakeWithDefaults instantiates a new HelpHelpTicketIntake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HelpHelpTicketIntake) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelpHelpTicketIntake) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelpHelpTicketIntake) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelpHelpTicketIntake) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *HelpHelpTicketIntake) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *HelpHelpTicketIntake) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *HelpHelpTicketIntake) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *HelpHelpTicketIntake) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPriority

`func (o *HelpHelpTicketIntake) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *HelpHelpTicketIntake) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *HelpHelpTicketIntake) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *HelpHelpTicketIntake) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSubject

`func (o *HelpHelpTicketIntake) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *HelpHelpTicketIntake) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *HelpHelpTicketIntake) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *HelpHelpTicketIntake) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


