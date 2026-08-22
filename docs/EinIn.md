# EinIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expedited** | Pointer to **bool** | Expedited asks for prioritised handling. Only meaningful when the responsible party cannot file online. | [optional] 
**Naics** | Pointer to **string** | NAICS is the six-digit code for what the business does. | [optional] 
**Responsible** | Pointer to [**Responsible**](Responsible.md) | Responsible is the person the IRS holds answerable for the entity. | [optional] 

## Methods

### NewEinIn

`func NewEinIn() *EinIn`

NewEinIn instantiates a new EinIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEinInWithDefaults

`func NewEinInWithDefaults() *EinIn`

NewEinInWithDefaults instantiates a new EinIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpedited

`func (o *EinIn) GetExpedited() bool`

GetExpedited returns the Expedited field if non-nil, zero value otherwise.

### GetExpeditedOk

`func (o *EinIn) GetExpeditedOk() (*bool, bool)`

GetExpeditedOk returns a tuple with the Expedited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedited

`func (o *EinIn) SetExpedited(v bool)`

SetExpedited sets Expedited field to given value.

### HasExpedited

`func (o *EinIn) HasExpedited() bool`

HasExpedited returns a boolean if a field has been set.

### GetNaics

`func (o *EinIn) GetNaics() string`

GetNaics returns the Naics field if non-nil, zero value otherwise.

### GetNaicsOk

`func (o *EinIn) GetNaicsOk() (*string, bool)`

GetNaicsOk returns a tuple with the Naics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaics

`func (o *EinIn) SetNaics(v string)`

SetNaics sets Naics field to given value.

### HasNaics

`func (o *EinIn) HasNaics() bool`

HasNaics returns a boolean if a field has been set.

### GetResponsible

`func (o *EinIn) GetResponsible() Responsible`

GetResponsible returns the Responsible field if non-nil, zero value otherwise.

### GetResponsibleOk

`func (o *EinIn) GetResponsibleOk() (*Responsible, bool)`

GetResponsibleOk returns a tuple with the Responsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsible

`func (o *EinIn) SetResponsible(v Responsible)`

SetResponsible sets Responsible field to given value.

### HasResponsible

`func (o *EinIn) HasResponsible() bool`

HasResponsible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


