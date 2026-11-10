# EmailSMTPCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatchAll** | Pointer to **bool** | CatchAll reports whether the host accepts all recipient addresses. | [optional] 
**Deliverable** | Pointer to **bool** | Deliverable reports whether the mailbox accepted the test recipient address. | [optional] 
**Disabled** | Pointer to **bool** | Disabled reports whether the recipient mailbox is disabled or inactive. | [optional] 
**FullInbox** | Pointer to **bool** | FullInbox reports whether the recipient mailbox is over quota. | [optional] 
**HostExists** | Pointer to **bool** | HostExists reports whether the destination mail server host exists and responded. | [optional] 
**Reason** | Pointer to **string** | Reason describes the failure or skip reason when Status is not pass. | [optional] 
**Status** | Pointer to **string** | Status reports the check outcome: pass, fail, unknown, or skipped. | [optional] 

## Methods

### NewEmailSMTPCheck

`func NewEmailSMTPCheck() *EmailSMTPCheck`

NewEmailSMTPCheck instantiates a new EmailSMTPCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSMTPCheckWithDefaults

`func NewEmailSMTPCheckWithDefaults() *EmailSMTPCheck`

NewEmailSMTPCheckWithDefaults instantiates a new EmailSMTPCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatchAll

`func (o *EmailSMTPCheck) GetCatchAll() bool`

GetCatchAll returns the CatchAll field if non-nil, zero value otherwise.

### GetCatchAllOk

`func (o *EmailSMTPCheck) GetCatchAllOk() (*bool, bool)`

GetCatchAllOk returns a tuple with the CatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchAll

`func (o *EmailSMTPCheck) SetCatchAll(v bool)`

SetCatchAll sets CatchAll field to given value.

### HasCatchAll

`func (o *EmailSMTPCheck) HasCatchAll() bool`

HasCatchAll returns a boolean if a field has been set.

### GetDeliverable

`func (o *EmailSMTPCheck) GetDeliverable() bool`

GetDeliverable returns the Deliverable field if non-nil, zero value otherwise.

### GetDeliverableOk

`func (o *EmailSMTPCheck) GetDeliverableOk() (*bool, bool)`

GetDeliverableOk returns a tuple with the Deliverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverable

`func (o *EmailSMTPCheck) SetDeliverable(v bool)`

SetDeliverable sets Deliverable field to given value.

### HasDeliverable

`func (o *EmailSMTPCheck) HasDeliverable() bool`

HasDeliverable returns a boolean if a field has been set.

### GetDisabled

`func (o *EmailSMTPCheck) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *EmailSMTPCheck) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *EmailSMTPCheck) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *EmailSMTPCheck) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFullInbox

`func (o *EmailSMTPCheck) GetFullInbox() bool`

GetFullInbox returns the FullInbox field if non-nil, zero value otherwise.

### GetFullInboxOk

`func (o *EmailSMTPCheck) GetFullInboxOk() (*bool, bool)`

GetFullInboxOk returns a tuple with the FullInbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullInbox

`func (o *EmailSMTPCheck) SetFullInbox(v bool)`

SetFullInbox sets FullInbox field to given value.

### HasFullInbox

`func (o *EmailSMTPCheck) HasFullInbox() bool`

HasFullInbox returns a boolean if a field has been set.

### GetHostExists

`func (o *EmailSMTPCheck) GetHostExists() bool`

GetHostExists returns the HostExists field if non-nil, zero value otherwise.

### GetHostExistsOk

`func (o *EmailSMTPCheck) GetHostExistsOk() (*bool, bool)`

GetHostExistsOk returns a tuple with the HostExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostExists

`func (o *EmailSMTPCheck) SetHostExists(v bool)`

SetHostExists sets HostExists field to given value.

### HasHostExists

`func (o *EmailSMTPCheck) HasHostExists() bool`

HasHostExists returns a boolean if a field has been set.

### GetReason

`func (o *EmailSMTPCheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailSMTPCheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailSMTPCheck) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailSMTPCheck) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *EmailSMTPCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailSMTPCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailSMTPCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailSMTPCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


