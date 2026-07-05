# CrmCompanyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DomainName** | Pointer to **string** |  | [optional] 
**Employees** | Pointer to **int64** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Arr** | Pointer to **int64** | Annual recurring revenue in minor units (cents) of currency | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**IdealCustomerProfile** | Pointer to **bool** |  | [optional] 
**LinkedinLink** | Pointer to **string** |  | [optional] 
**XLink** | Pointer to **string** |  | [optional] 

## Methods

### NewCrmCompanyInput

`func NewCrmCompanyInput(name string, ) *CrmCompanyInput`

NewCrmCompanyInput instantiates a new CrmCompanyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmCompanyInputWithDefaults

`func NewCrmCompanyInputWithDefaults() *CrmCompanyInput`

NewCrmCompanyInputWithDefaults instantiates a new CrmCompanyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CrmCompanyInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrmCompanyInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrmCompanyInput) SetName(v string)`

SetName sets Name field to given value.


### GetDomainName

`func (o *CrmCompanyInput) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CrmCompanyInput) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CrmCompanyInput) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CrmCompanyInput) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEmployees

`func (o *CrmCompanyInput) GetEmployees() int64`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *CrmCompanyInput) GetEmployeesOk() (*int64, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *CrmCompanyInput) SetEmployees(v int64)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *CrmCompanyInput) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### GetCity

`func (o *CrmCompanyInput) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CrmCompanyInput) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CrmCompanyInput) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CrmCompanyInput) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CrmCompanyInput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CrmCompanyInput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CrmCompanyInput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CrmCompanyInput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetArr

`func (o *CrmCompanyInput) GetArr() int64`

GetArr returns the Arr field if non-nil, zero value otherwise.

### GetArrOk

`func (o *CrmCompanyInput) GetArrOk() (*int64, bool)`

GetArrOk returns a tuple with the Arr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArr

`func (o *CrmCompanyInput) SetArr(v int64)`

SetArr sets Arr field to given value.

### HasArr

`func (o *CrmCompanyInput) HasArr() bool`

HasArr returns a boolean if a field has been set.

### GetCurrency

`func (o *CrmCompanyInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CrmCompanyInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CrmCompanyInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CrmCompanyInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIdealCustomerProfile

`func (o *CrmCompanyInput) GetIdealCustomerProfile() bool`

GetIdealCustomerProfile returns the IdealCustomerProfile field if non-nil, zero value otherwise.

### GetIdealCustomerProfileOk

`func (o *CrmCompanyInput) GetIdealCustomerProfileOk() (*bool, bool)`

GetIdealCustomerProfileOk returns a tuple with the IdealCustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdealCustomerProfile

`func (o *CrmCompanyInput) SetIdealCustomerProfile(v bool)`

SetIdealCustomerProfile sets IdealCustomerProfile field to given value.

### HasIdealCustomerProfile

`func (o *CrmCompanyInput) HasIdealCustomerProfile() bool`

HasIdealCustomerProfile returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *CrmCompanyInput) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *CrmCompanyInput) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *CrmCompanyInput) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *CrmCompanyInput) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetXLink

`func (o *CrmCompanyInput) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *CrmCompanyInput) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *CrmCompanyInput) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *CrmCompanyInput) HasXLink() bool`

HasXLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


