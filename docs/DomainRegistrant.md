# DomainRegistrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** | street address | [optional] 
**Address2** | Pointer to **string** | second address line | [optional] 
**City** | Pointer to **string** | city or locality | [optional] 
**CompanyName** | Pointer to **string** | the organisation the contact acts for | [optional] 
**Country** | Pointer to **string** | ISO-3166 alpha-2, e.g. \&quot;US\&quot; | [optional] 
**Email** | Pointer to **string** | where WHOIS correspondence is sent | [optional] 
**Fax** | Pointer to **string** | fax number, in the same form as phone | [optional] 
**FirstName** | Pointer to **string** | the contact&#39;s given name | [optional] 
**LastName** | Pointer to **string** | the contact&#39;s family name | [optional] 
**Phone** | Pointer to **string** | +NN.NNNNNNN | [optional] 
**State** | Pointer to **string** | state, province or region | [optional] 
**Zip** | Pointer to **string** | postal code | [optional] 

## Methods

### NewDomainRegistrant

`func NewDomainRegistrant() *DomainRegistrant`

NewDomainRegistrant instantiates a new DomainRegistrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRegistrantWithDefaults

`func NewDomainRegistrantWithDefaults() *DomainRegistrant`

NewDomainRegistrantWithDefaults instantiates a new DomainRegistrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *DomainRegistrant) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *DomainRegistrant) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *DomainRegistrant) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *DomainRegistrant) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *DomainRegistrant) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *DomainRegistrant) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *DomainRegistrant) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *DomainRegistrant) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *DomainRegistrant) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DomainRegistrant) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DomainRegistrant) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DomainRegistrant) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *DomainRegistrant) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DomainRegistrant) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DomainRegistrant) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DomainRegistrant) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *DomainRegistrant) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DomainRegistrant) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DomainRegistrant) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DomainRegistrant) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *DomainRegistrant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DomainRegistrant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DomainRegistrant) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DomainRegistrant) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFax

`func (o *DomainRegistrant) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *DomainRegistrant) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *DomainRegistrant) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *DomainRegistrant) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetFirstName

`func (o *DomainRegistrant) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DomainRegistrant) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DomainRegistrant) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DomainRegistrant) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *DomainRegistrant) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DomainRegistrant) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DomainRegistrant) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DomainRegistrant) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *DomainRegistrant) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DomainRegistrant) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DomainRegistrant) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *DomainRegistrant) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *DomainRegistrant) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainRegistrant) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainRegistrant) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DomainRegistrant) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *DomainRegistrant) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *DomainRegistrant) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *DomainRegistrant) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *DomainRegistrant) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


