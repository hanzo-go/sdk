# CompanyResponsible

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country is where they reside, ISO 3166-1 alpha-2. | [optional] 
**Email** | Pointer to **string** | Email reaches them for signature. | [optional] 
**Name** | Pointer to **string** | Name is their full legal name as the IRS will hold it. | [optional] 
**UsTaxId** | Pointer to **bool** | USTaxID reports that they hold an SSN or ITIN. It is a BOOLEAN on purpose: the number itself is never needed here and a field that could hold it is a field that will eventually be logged. | [optional] 

## Methods

### NewCompanyResponsible

`func NewCompanyResponsible() *CompanyResponsible`

NewCompanyResponsible instantiates a new CompanyResponsible object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponsibleWithDefaults

`func NewCompanyResponsibleWithDefaults() *CompanyResponsible`

NewCompanyResponsibleWithDefaults instantiates a new CompanyResponsible object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CompanyResponsible) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyResponsible) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyResponsible) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyResponsible) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *CompanyResponsible) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CompanyResponsible) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CompanyResponsible) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CompanyResponsible) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *CompanyResponsible) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyResponsible) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyResponsible) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyResponsible) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsTaxId

`func (o *CompanyResponsible) GetUsTaxId() bool`

GetUsTaxId returns the UsTaxId field if non-nil, zero value otherwise.

### GetUsTaxIdOk

`func (o *CompanyResponsible) GetUsTaxIdOk() (*bool, bool)`

GetUsTaxIdOk returns a tuple with the UsTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsTaxId

`func (o *CompanyResponsible) SetUsTaxId(v bool)`

SetUsTaxId sets UsTaxId field to given value.

### HasUsTaxId

`func (o *CompanyResponsible) HasUsTaxId() bool`

HasUsTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


