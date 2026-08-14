# Registrant

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

### NewRegistrant

`func NewRegistrant() *Registrant`

NewRegistrant instantiates a new Registrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrantWithDefaults

`func NewRegistrantWithDefaults() *Registrant`

NewRegistrantWithDefaults instantiates a new Registrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *Registrant) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *Registrant) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *Registrant) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *Registrant) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *Registrant) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *Registrant) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *Registrant) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *Registrant) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *Registrant) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Registrant) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Registrant) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Registrant) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *Registrant) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Registrant) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Registrant) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Registrant) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *Registrant) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Registrant) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Registrant) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Registrant) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *Registrant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Registrant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Registrant) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Registrant) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFax

`func (o *Registrant) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Registrant) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Registrant) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Registrant) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetFirstName

`func (o *Registrant) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Registrant) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Registrant) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Registrant) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Registrant) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Registrant) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Registrant) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Registrant) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *Registrant) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Registrant) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Registrant) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Registrant) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *Registrant) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Registrant) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Registrant) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Registrant) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *Registrant) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *Registrant) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *Registrant) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *Registrant) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


