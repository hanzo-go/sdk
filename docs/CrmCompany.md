# CrmCompany

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
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** | Unix seconds | [optional] 
**UpdatedAt** | Pointer to **int64** | Unix seconds | [optional] 

## Methods

### NewCrmCompany

`func NewCrmCompany(name string, ) *CrmCompany`

NewCrmCompany instantiates a new CrmCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmCompanyWithDefaults

`func NewCrmCompanyWithDefaults() *CrmCompany`

NewCrmCompanyWithDefaults instantiates a new CrmCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CrmCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrmCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrmCompany) SetName(v string)`

SetName sets Name field to given value.


### GetDomainName

`func (o *CrmCompany) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CrmCompany) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CrmCompany) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CrmCompany) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEmployees

`func (o *CrmCompany) GetEmployees() int64`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *CrmCompany) GetEmployeesOk() (*int64, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *CrmCompany) SetEmployees(v int64)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *CrmCompany) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### GetCity

`func (o *CrmCompany) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CrmCompany) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CrmCompany) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CrmCompany) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CrmCompany) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CrmCompany) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CrmCompany) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CrmCompany) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetArr

`func (o *CrmCompany) GetArr() int64`

GetArr returns the Arr field if non-nil, zero value otherwise.

### GetArrOk

`func (o *CrmCompany) GetArrOk() (*int64, bool)`

GetArrOk returns a tuple with the Arr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArr

`func (o *CrmCompany) SetArr(v int64)`

SetArr sets Arr field to given value.

### HasArr

`func (o *CrmCompany) HasArr() bool`

HasArr returns a boolean if a field has been set.

### GetCurrency

`func (o *CrmCompany) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CrmCompany) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CrmCompany) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CrmCompany) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIdealCustomerProfile

`func (o *CrmCompany) GetIdealCustomerProfile() bool`

GetIdealCustomerProfile returns the IdealCustomerProfile field if non-nil, zero value otherwise.

### GetIdealCustomerProfileOk

`func (o *CrmCompany) GetIdealCustomerProfileOk() (*bool, bool)`

GetIdealCustomerProfileOk returns a tuple with the IdealCustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdealCustomerProfile

`func (o *CrmCompany) SetIdealCustomerProfile(v bool)`

SetIdealCustomerProfile sets IdealCustomerProfile field to given value.

### HasIdealCustomerProfile

`func (o *CrmCompany) HasIdealCustomerProfile() bool`

HasIdealCustomerProfile returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *CrmCompany) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *CrmCompany) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *CrmCompany) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *CrmCompany) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetXLink

`func (o *CrmCompany) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *CrmCompany) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *CrmCompany) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *CrmCompany) HasXLink() bool`

HasXLink returns a boolean if a field has been set.

### GetId

`func (o *CrmCompany) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrmCompany) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrmCompany) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CrmCompany) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CrmCompany) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrmCompany) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrmCompany) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CrmCompany) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CrmCompany) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CrmCompany) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CrmCompany) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CrmCompany) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


