# Responsible

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country is where they reside, ISO 3166-1 alpha-2. | [optional] 
**Email** | Pointer to **string** | Email reaches them for signature. | [optional] 
**Name** | Pointer to **string** | Name is their full legal name as the IRS will hold it. | [optional] 
**UsTaxId** | Pointer to **bool** | USTaxID reports that they hold an SSN or ITIN. It is a BOOLEAN on purpose: the number itself is never needed here and a field that could hold it is a field that will eventually be logged. | [optional] 

## Methods

### NewResponsible

`func NewResponsible() *Responsible`

NewResponsible instantiates a new Responsible object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsibleWithDefaults

`func NewResponsibleWithDefaults() *Responsible`

NewResponsibleWithDefaults instantiates a new Responsible object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *Responsible) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Responsible) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Responsible) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Responsible) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *Responsible) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Responsible) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Responsible) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Responsible) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Responsible) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Responsible) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Responsible) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Responsible) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsTaxId

`func (o *Responsible) GetUsTaxId() bool`

GetUsTaxId returns the UsTaxId field if non-nil, zero value otherwise.

### GetUsTaxIdOk

`func (o *Responsible) GetUsTaxIdOk() (*bool, bool)`

GetUsTaxIdOk returns a tuple with the UsTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsTaxId

`func (o *Responsible) SetUsTaxId(v bool)`

SetUsTaxId sets UsTaxId field to given value.

### HasUsTaxId

`func (o *Responsible) HasUsTaxId() bool`

HasUsTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


