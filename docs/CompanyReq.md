# CompanyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arr** | Pointer to **int32** | ARR is annual recurring revenue in minor units (cents) of Currency. | [optional] 
**City** | Pointer to **string** | City is the head-office city. | [optional] 
**Country** | Pointer to **string** | Country is the head-office country. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code ARR is denominated in; empty defaults to USD. | [optional] 
**DomainName** | Pointer to **string** | DomainName is the company&#39;s primary domain, e.g. \&quot;acme.com\&quot;. | [optional] 
**Employees** | Pointer to **int32** | Employees is the headcount. | [optional] 
**Id** | Pointer to **string** | ID names the company to update and comes from the path. A create ignores it: the server mints the id. | [optional] 
**IdealCustomerProfile** | Pointer to **bool** | ICP marks the company as an ideal-customer-profile fit. | [optional] 
**LinkedinLink** | Pointer to **string** | Linkedin is the company&#39;s LinkedIn URL. | [optional] 
**Name** | Pointer to **string** | Name is the company name. Required. | [optional] 
**XLink** | Pointer to **string** | XLink is the company&#39;s X (Twitter) URL. | [optional] 

## Methods

### NewCompanyReq

`func NewCompanyReq() *CompanyReq`

NewCompanyReq instantiates a new CompanyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyReqWithDefaults

`func NewCompanyReqWithDefaults() *CompanyReq`

NewCompanyReqWithDefaults instantiates a new CompanyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArr

`func (o *CompanyReq) GetArr() int32`

GetArr returns the Arr field if non-nil, zero value otherwise.

### GetArrOk

`func (o *CompanyReq) GetArrOk() (*int32, bool)`

GetArrOk returns a tuple with the Arr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArr

`func (o *CompanyReq) SetArr(v int32)`

SetArr sets Arr field to given value.

### HasArr

`func (o *CompanyReq) HasArr() bool`

HasArr returns a boolean if a field has been set.

### GetCity

`func (o *CompanyReq) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanyReq) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanyReq) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanyReq) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CompanyReq) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyReq) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyReq) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyReq) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *CompanyReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompanyReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompanyReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CompanyReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDomainName

`func (o *CompanyReq) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CompanyReq) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CompanyReq) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CompanyReq) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEmployees

`func (o *CompanyReq) GetEmployees() int32`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *CompanyReq) GetEmployeesOk() (*int32, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *CompanyReq) SetEmployees(v int32)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *CompanyReq) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### GetId

`func (o *CompanyReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdealCustomerProfile

`func (o *CompanyReq) GetIdealCustomerProfile() bool`

GetIdealCustomerProfile returns the IdealCustomerProfile field if non-nil, zero value otherwise.

### GetIdealCustomerProfileOk

`func (o *CompanyReq) GetIdealCustomerProfileOk() (*bool, bool)`

GetIdealCustomerProfileOk returns a tuple with the IdealCustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdealCustomerProfile

`func (o *CompanyReq) SetIdealCustomerProfile(v bool)`

SetIdealCustomerProfile sets IdealCustomerProfile field to given value.

### HasIdealCustomerProfile

`func (o *CompanyReq) HasIdealCustomerProfile() bool`

HasIdealCustomerProfile returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *CompanyReq) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *CompanyReq) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *CompanyReq) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *CompanyReq) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetName

`func (o *CompanyReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXLink

`func (o *CompanyReq) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *CompanyReq) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *CompanyReq) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *CompanyReq) HasXLink() bool`

HasXLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


