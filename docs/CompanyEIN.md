# CompanyEIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expedited** | Pointer to **bool** | Expedited reports that prioritised handling was asked for. | [optional] 
**Forms** | Pointer to [**[]CompanyForm**](CompanyForm.md) | Forms are the forms this application owes, with what each is for. | [optional] 
**Naics** | Pointer to **string** | NAICS is the six-digit code for what the business does. The SS-4 asks it and the IRS will not process an application without one. | [optional] 
**Number** | Pointer to **string** | Number is the issued EIN, absent until the IRS issues it. | [optional] 
**Online** | Pointer to **bool** | Online reports that this application can be filed with the IRS online and issued in a sitting, rather than signed and posted. It is the single fact that decides how long a customer waits, so it is answered rather than implied by the absence of forms. | [optional] 
**Responsible** | Pointer to [**CompanyResponsible**](CompanyResponsible.md) | Responsible is the person the IRS holds answerable. | [optional] 
**Status** | Pointer to **string** | Status is how far it has got. | [optional] 

## Methods

### NewCompanyEIN

`func NewCompanyEIN() *CompanyEIN`

NewCompanyEIN instantiates a new CompanyEIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyEINWithDefaults

`func NewCompanyEINWithDefaults() *CompanyEIN`

NewCompanyEINWithDefaults instantiates a new CompanyEIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpedited

`func (o *CompanyEIN) GetExpedited() bool`

GetExpedited returns the Expedited field if non-nil, zero value otherwise.

### GetExpeditedOk

`func (o *CompanyEIN) GetExpeditedOk() (*bool, bool)`

GetExpeditedOk returns a tuple with the Expedited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedited

`func (o *CompanyEIN) SetExpedited(v bool)`

SetExpedited sets Expedited field to given value.

### HasExpedited

`func (o *CompanyEIN) HasExpedited() bool`

HasExpedited returns a boolean if a field has been set.

### GetForms

`func (o *CompanyEIN) GetForms() []CompanyForm`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *CompanyEIN) GetFormsOk() (*[]CompanyForm, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *CompanyEIN) SetForms(v []CompanyForm)`

SetForms sets Forms field to given value.

### HasForms

`func (o *CompanyEIN) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetNaics

`func (o *CompanyEIN) GetNaics() string`

GetNaics returns the Naics field if non-nil, zero value otherwise.

### GetNaicsOk

`func (o *CompanyEIN) GetNaicsOk() (*string, bool)`

GetNaicsOk returns a tuple with the Naics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaics

`func (o *CompanyEIN) SetNaics(v string)`

SetNaics sets Naics field to given value.

### HasNaics

`func (o *CompanyEIN) HasNaics() bool`

HasNaics returns a boolean if a field has been set.

### GetNumber

`func (o *CompanyEIN) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CompanyEIN) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CompanyEIN) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CompanyEIN) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOnline

`func (o *CompanyEIN) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *CompanyEIN) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *CompanyEIN) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *CompanyEIN) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetResponsible

`func (o *CompanyEIN) GetResponsible() CompanyResponsible`

GetResponsible returns the Responsible field if non-nil, zero value otherwise.

### GetResponsibleOk

`func (o *CompanyEIN) GetResponsibleOk() (*CompanyResponsible, bool)`

GetResponsibleOk returns a tuple with the Responsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsible

`func (o *CompanyEIN) SetResponsible(v CompanyResponsible)`

SetResponsible sets Responsible field to given value.

### HasResponsible

`func (o *CompanyEIN) HasResponsible() bool`

HasResponsible returns a boolean if a field has been set.

### GetStatus

`func (o *CompanyEIN) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyEIN) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyEIN) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyEIN) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


