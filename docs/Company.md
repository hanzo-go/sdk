# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arr** | Pointer to **int32** | ARR is annual recurring revenue in minor units (cents) of Currency. | [optional] 
**City** | Pointer to **string** | City is the head-office city. | [optional] 
**Country** | Pointer to **string** | Country is the head-office country. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the company was created. Server-owned. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code ARR is denominated in; a write that names none stores USD. | [optional] 
**DomainName** | Pointer to **string** | DomainName is the company&#39;s primary domain, e.g. \&quot;acme.com\&quot;. | [optional] 
**Employees** | Pointer to **int32** | Employees is the headcount. | [optional] 
**Id** | Pointer to **string** | ID is the server-minted company id (\&quot;comp_\&quot; + 128 random bits). | [optional] 
**IdealCustomerProfile** | Pointer to **bool** | ICP marks the company as an ideal-customer-profile fit. | [optional] 
**LinkedinLink** | Pointer to **string** | Linkedin is the company&#39;s LinkedIn URL. | [optional] 
**Name** | Pointer to **string** | Name is the company name. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second of the last write. Server-owned. | [optional] 
**XLink** | Pointer to **string** | XLink is the company&#39;s X (Twitter) URL. | [optional] 

## Methods

### NewCompany

`func NewCompany() *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArr

`func (o *Company) GetArr() int32`

GetArr returns the Arr field if non-nil, zero value otherwise.

### GetArrOk

`func (o *Company) GetArrOk() (*int32, bool)`

GetArrOk returns a tuple with the Arr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArr

`func (o *Company) SetArr(v int32)`

SetArr sets Arr field to given value.

### HasArr

`func (o *Company) HasArr() bool`

HasArr returns a boolean if a field has been set.

### GetCity

`func (o *Company) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Company) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Company) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Company) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *Company) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Company) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Company) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Company) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Company) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Company) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Company) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Company) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Company) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Company) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Company) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Company) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDomainName

`func (o *Company) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Company) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Company) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Company) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEmployees

`func (o *Company) GetEmployees() int32`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *Company) GetEmployeesOk() (*int32, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *Company) SetEmployees(v int32)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *Company) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### GetId

`func (o *Company) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdealCustomerProfile

`func (o *Company) GetIdealCustomerProfile() bool`

GetIdealCustomerProfile returns the IdealCustomerProfile field if non-nil, zero value otherwise.

### GetIdealCustomerProfileOk

`func (o *Company) GetIdealCustomerProfileOk() (*bool, bool)`

GetIdealCustomerProfileOk returns a tuple with the IdealCustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdealCustomerProfile

`func (o *Company) SetIdealCustomerProfile(v bool)`

SetIdealCustomerProfile sets IdealCustomerProfile field to given value.

### HasIdealCustomerProfile

`func (o *Company) HasIdealCustomerProfile() bool`

HasIdealCustomerProfile returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *Company) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *Company) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *Company) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *Company) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Company) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Company) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Company) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Company) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Company) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetXLink

`func (o *Company) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *Company) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *Company) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *Company) HasXLink() bool`

HasXLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


