# CloudHelpTicketIntake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is the customer&#39;s message. Optional; it becomes the ticket&#39;s description AND the opening entry of its conversation thread. Clipped at 16 KiB. | [optional] 
**Email** | Pointer to **string** | Email is how the support team replies. Required; clipped at 320 characters (the RFC 5321 maximum). It is recorded as the ticket&#39;s customer, and it is not verified. | [optional] 
**Priority** | Pointer to **string** | Priority is Low, Medium, High or Urgent, case-insensitively. Anything else — including omitting it — is recorded as Medium rather than refused. | [optional] 
**Subject** | Pointer to **string** | Subject is the one-line summary of the problem. Required; longer than 300 characters is clipped rather than refused. | [optional] 

## Methods

### NewCloudHelpTicketIntake

`func NewCloudHelpTicketIntake() *CloudHelpTicketIntake`

NewCloudHelpTicketIntake instantiates a new CloudHelpTicketIntake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHelpTicketIntakeWithDefaults

`func NewCloudHelpTicketIntakeWithDefaults() *CloudHelpTicketIntake`

NewCloudHelpTicketIntakeWithDefaults instantiates a new CloudHelpTicketIntake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudHelpTicketIntake) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudHelpTicketIntake) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudHelpTicketIntake) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudHelpTicketIntake) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *CloudHelpTicketIntake) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudHelpTicketIntake) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudHelpTicketIntake) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudHelpTicketIntake) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPriority

`func (o *CloudHelpTicketIntake) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CloudHelpTicketIntake) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CloudHelpTicketIntake) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CloudHelpTicketIntake) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSubject

`func (o *CloudHelpTicketIntake) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudHelpTicketIntake) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudHelpTicketIntake) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudHelpTicketIntake) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


