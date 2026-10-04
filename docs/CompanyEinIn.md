# CompanyEinIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expedited** | Pointer to **bool** | Expedited asks for prioritised handling. Only meaningful when the responsible party cannot file online. | [optional] 
**Naics** | Pointer to **string** | NAICS is the six-digit code for what the business does. | [optional] 
**Responsible** | Pointer to [**CompanyResponsible**](CompanyResponsible.md) | Responsible is the person the IRS holds answerable for the entity. | [optional] 

## Methods

### NewCompanyEinIn

`func NewCompanyEinIn() *CompanyEinIn`

NewCompanyEinIn instantiates a new CompanyEinIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyEinInWithDefaults

`func NewCompanyEinInWithDefaults() *CompanyEinIn`

NewCompanyEinInWithDefaults instantiates a new CompanyEinIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpedited

`func (o *CompanyEinIn) GetExpedited() bool`

GetExpedited returns the Expedited field if non-nil, zero value otherwise.

### GetExpeditedOk

`func (o *CompanyEinIn) GetExpeditedOk() (*bool, bool)`

GetExpeditedOk returns a tuple with the Expedited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedited

`func (o *CompanyEinIn) SetExpedited(v bool)`

SetExpedited sets Expedited field to given value.

### HasExpedited

`func (o *CompanyEinIn) HasExpedited() bool`

HasExpedited returns a boolean if a field has been set.

### GetNaics

`func (o *CompanyEinIn) GetNaics() string`

GetNaics returns the Naics field if non-nil, zero value otherwise.

### GetNaicsOk

`func (o *CompanyEinIn) GetNaicsOk() (*string, bool)`

GetNaicsOk returns a tuple with the Naics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaics

`func (o *CompanyEinIn) SetNaics(v string)`

SetNaics sets Naics field to given value.

### HasNaics

`func (o *CompanyEinIn) HasNaics() bool`

HasNaics returns a boolean if a field has been set.

### GetResponsible

`func (o *CompanyEinIn) GetResponsible() CompanyResponsible`

GetResponsible returns the Responsible field if non-nil, zero value otherwise.

### GetResponsibleOk

`func (o *CompanyEinIn) GetResponsibleOk() (*CompanyResponsible, bool)`

GetResponsibleOk returns a tuple with the Responsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsible

`func (o *CompanyEinIn) SetResponsible(v CompanyResponsible)`

SetResponsible sets Responsible field to given value.

### HasResponsible

`func (o *CompanyEinIn) HasResponsible() bool`

HasResponsible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


