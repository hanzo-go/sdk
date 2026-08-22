# EIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expedited** | Pointer to **bool** | Expedited reports that prioritised handling was asked for. | [optional] 
**Forms** | Pointer to [**[]Form**](Form.md) | Forms are the forms this application owes, with what each is for. | [optional] 
**Naics** | Pointer to **string** | NAICS is the six-digit code for what the business does. The SS-4 asks it and the IRS will not process an application without one. | [optional] 
**Number** | Pointer to **string** | Number is the issued EIN, absent until the IRS issues it. | [optional] 
**Online** | Pointer to **bool** | Online reports that this application can be filed with the IRS online and issued in a sitting, rather than signed and posted. It is the single fact that decides how long a customer waits, so it is answered rather than implied by the absence of forms. | [optional] 
**Responsible** | Pointer to [**Responsible**](Responsible.md) | Responsible is the person the IRS holds answerable. | [optional] 
**Status** | Pointer to **string** | Status is how far it has got. | [optional] 

## Methods

### NewEIN

`func NewEIN() *EIN`

NewEIN instantiates a new EIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEINWithDefaults

`func NewEINWithDefaults() *EIN`

NewEINWithDefaults instantiates a new EIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpedited

`func (o *EIN) GetExpedited() bool`

GetExpedited returns the Expedited field if non-nil, zero value otherwise.

### GetExpeditedOk

`func (o *EIN) GetExpeditedOk() (*bool, bool)`

GetExpeditedOk returns a tuple with the Expedited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedited

`func (o *EIN) SetExpedited(v bool)`

SetExpedited sets Expedited field to given value.

### HasExpedited

`func (o *EIN) HasExpedited() bool`

HasExpedited returns a boolean if a field has been set.

### GetForms

`func (o *EIN) GetForms() []Form`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *EIN) GetFormsOk() (*[]Form, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *EIN) SetForms(v []Form)`

SetForms sets Forms field to given value.

### HasForms

`func (o *EIN) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetNaics

`func (o *EIN) GetNaics() string`

GetNaics returns the Naics field if non-nil, zero value otherwise.

### GetNaicsOk

`func (o *EIN) GetNaicsOk() (*string, bool)`

GetNaicsOk returns a tuple with the Naics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaics

`func (o *EIN) SetNaics(v string)`

SetNaics sets Naics field to given value.

### HasNaics

`func (o *EIN) HasNaics() bool`

HasNaics returns a boolean if a field has been set.

### GetNumber

`func (o *EIN) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *EIN) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *EIN) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *EIN) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOnline

`func (o *EIN) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *EIN) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *EIN) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *EIN) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetResponsible

`func (o *EIN) GetResponsible() Responsible`

GetResponsible returns the Responsible field if non-nil, zero value otherwise.

### GetResponsibleOk

`func (o *EIN) GetResponsibleOk() (*Responsible, bool)`

GetResponsibleOk returns a tuple with the Responsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsible

`func (o *EIN) SetResponsible(v Responsible)`

SetResponsible sets Responsible field to given value.

### HasResponsible

`func (o *EIN) HasResponsible() bool`

HasResponsible returns a boolean if a field has been set.

### GetStatus

`func (o *EIN) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EIN) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EIN) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EIN) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


