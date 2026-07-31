# CloudCompanyReq

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

### NewCloudCompanyReq

`func NewCloudCompanyReq() *CloudCompanyReq`

NewCloudCompanyReq instantiates a new CloudCompanyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCompanyReqWithDefaults

`func NewCloudCompanyReqWithDefaults() *CloudCompanyReq`

NewCloudCompanyReqWithDefaults instantiates a new CloudCompanyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArr

`func (o *CloudCompanyReq) GetArr() int32`

GetArr returns the Arr field if non-nil, zero value otherwise.

### GetArrOk

`func (o *CloudCompanyReq) GetArrOk() (*int32, bool)`

GetArrOk returns a tuple with the Arr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArr

`func (o *CloudCompanyReq) SetArr(v int32)`

SetArr sets Arr field to given value.

### HasArr

`func (o *CloudCompanyReq) HasArr() bool`

HasArr returns a boolean if a field has been set.

### GetCity

`func (o *CloudCompanyReq) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CloudCompanyReq) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CloudCompanyReq) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CloudCompanyReq) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CloudCompanyReq) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CloudCompanyReq) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CloudCompanyReq) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CloudCompanyReq) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudCompanyReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudCompanyReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudCompanyReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudCompanyReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDomainName

`func (o *CloudCompanyReq) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CloudCompanyReq) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CloudCompanyReq) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CloudCompanyReq) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEmployees

`func (o *CloudCompanyReq) GetEmployees() int32`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *CloudCompanyReq) GetEmployeesOk() (*int32, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *CloudCompanyReq) SetEmployees(v int32)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *CloudCompanyReq) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### GetId

`func (o *CloudCompanyReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCompanyReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCompanyReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCompanyReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdealCustomerProfile

`func (o *CloudCompanyReq) GetIdealCustomerProfile() bool`

GetIdealCustomerProfile returns the IdealCustomerProfile field if non-nil, zero value otherwise.

### GetIdealCustomerProfileOk

`func (o *CloudCompanyReq) GetIdealCustomerProfileOk() (*bool, bool)`

GetIdealCustomerProfileOk returns a tuple with the IdealCustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdealCustomerProfile

`func (o *CloudCompanyReq) SetIdealCustomerProfile(v bool)`

SetIdealCustomerProfile sets IdealCustomerProfile field to given value.

### HasIdealCustomerProfile

`func (o *CloudCompanyReq) HasIdealCustomerProfile() bool`

HasIdealCustomerProfile returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *CloudCompanyReq) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *CloudCompanyReq) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *CloudCompanyReq) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *CloudCompanyReq) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetName

`func (o *CloudCompanyReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCompanyReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCompanyReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCompanyReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXLink

`func (o *CloudCompanyReq) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *CloudCompanyReq) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *CloudCompanyReq) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *CloudCompanyReq) HasXLink() bool`

HasXLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


