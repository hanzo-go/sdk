# FormSinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crm** | Pointer to [**FormCRM**](FormCRM.md) | CRM files the response on the timeline of the person it is from, or on the org&#39;s own timeline when it names nobody. Absent, nothing is filed. | [optional] 
**Events** | Pointer to **bool** | Events states form_viewed, form_started and form_submitted on the deployment&#39;s event stream: the form&#39;s ids and the refs of the fields answered, never an answer&#39;s text. | [optional] 
**Forward** | Pointer to [**FormForward**](FormForward.md) | Forward mails every completed response to one address. | [optional] 
**Welcome** | Pointer to [**FormWelcome**](FormWelcome.md) | Welcome mails the respondent once, the first time their address answers. | [optional] 

## Methods

### NewFormSinks

`func NewFormSinks() *FormSinks`

NewFormSinks instantiates a new FormSinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSinksWithDefaults

`func NewFormSinksWithDefaults() *FormSinks`

NewFormSinksWithDefaults instantiates a new FormSinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrm

`func (o *FormSinks) GetCrm() FormCRM`

GetCrm returns the Crm field if non-nil, zero value otherwise.

### GetCrmOk

`func (o *FormSinks) GetCrmOk() (*FormCRM, bool)`

GetCrmOk returns a tuple with the Crm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrm

`func (o *FormSinks) SetCrm(v FormCRM)`

SetCrm sets Crm field to given value.

### HasCrm

`func (o *FormSinks) HasCrm() bool`

HasCrm returns a boolean if a field has been set.

### GetEvents

`func (o *FormSinks) GetEvents() bool`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FormSinks) GetEventsOk() (*bool, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FormSinks) SetEvents(v bool)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *FormSinks) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetForward

`func (o *FormSinks) GetForward() FormForward`

GetForward returns the Forward field if non-nil, zero value otherwise.

### GetForwardOk

`func (o *FormSinks) GetForwardOk() (*FormForward, bool)`

GetForwardOk returns a tuple with the Forward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForward

`func (o *FormSinks) SetForward(v FormForward)`

SetForward sets Forward field to given value.

### HasForward

`func (o *FormSinks) HasForward() bool`

HasForward returns a boolean if a field has been set.

### GetWelcome

`func (o *FormSinks) GetWelcome() FormWelcome`

GetWelcome returns the Welcome field if non-nil, zero value otherwise.

### GetWelcomeOk

`func (o *FormSinks) GetWelcomeOk() (*FormWelcome, bool)`

GetWelcomeOk returns a tuple with the Welcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcome

`func (o *FormSinks) SetWelcome(v FormWelcome)`

SetWelcome sets Welcome field to given value.

### HasWelcome

`func (o *FormSinks) HasWelcome() bool`

HasWelcome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


