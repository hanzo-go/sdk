# HelpTicketIntake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is the customer&#39;s message. Optional; it becomes the ticket&#39;s description AND the opening entry of its conversation thread. Clipped at 16 KiB. | [optional] 
**Email** | Pointer to **string** | Email is how the support team replies. Required; clipped at 320 characters (the RFC 5321 maximum). It is recorded as the ticket&#39;s customer, and it is not verified. | [optional] 
**Priority** | Pointer to **string** | Priority is Low, Medium, High or Urgent, case-insensitively. Anything else — including omitting it — is recorded as Medium rather than refused. | [optional] 
**Subject** | Pointer to **string** | Subject is the one-line summary of the problem. Required; longer than 300 characters is clipped rather than refused. | [optional] 

## Methods

### NewHelpTicketIntake

`func NewHelpTicketIntake() *HelpTicketIntake`

NewHelpTicketIntake instantiates a new HelpTicketIntake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpTicketIntakeWithDefaults

`func NewHelpTicketIntakeWithDefaults() *HelpTicketIntake`

NewHelpTicketIntakeWithDefaults instantiates a new HelpTicketIntake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HelpTicketIntake) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelpTicketIntake) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelpTicketIntake) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelpTicketIntake) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *HelpTicketIntake) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *HelpTicketIntake) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *HelpTicketIntake) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *HelpTicketIntake) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPriority

`func (o *HelpTicketIntake) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *HelpTicketIntake) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *HelpTicketIntake) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *HelpTicketIntake) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSubject

`func (o *HelpTicketIntake) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *HelpTicketIntake) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *HelpTicketIntake) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *HelpTicketIntake) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


