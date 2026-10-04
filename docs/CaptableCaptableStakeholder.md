# CaptableCaptableStakeholder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City is the stakeholder&#39;s city, if recorded. | [optional] 
**CompanyName** | Pointer to **string** | CompanyName is the name of the company whose cap table this is. | [optional] 
**Country** | Pointer to **string** | Country is the stakeholder&#39;s two-letter country code. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the stakeholder was added, in unix milliseconds. | [optional] 
**CurrentRelationship** | Pointer to **string** | CurrentRelationship is how the stakeholder relates to the company, e.g. FOUNDER, INVESTOR or EMPLOYEE. | [optional] 
**Email** | Pointer to **string** | Email is the stakeholder&#39;s email, unique within the company. | [optional] 
**Id** | Pointer to **string** | ID is the stakeholder id. | [optional] 
**InstitutionName** | Pointer to **string** | InstitutionName names the institution, when the stakeholder is one. | [optional] 
**Name** | Pointer to **string** | Name is the stakeholder&#39;s full name. | [optional] 
**StakeholderType** | Pointer to **string** | StakeholderType is INDIVIDUAL or INSTITUTION. | [optional] 
**State** | Pointer to **string** | State is the stakeholder&#39;s state or province, if recorded. | [optional] 
**StreetAddress** | Pointer to **string** | StreetAddress is the stakeholder&#39;s street address, if recorded. | [optional] 
**TaxId** | Pointer to **string** | TaxID is the stakeholder&#39;s tax identifier, if recorded. | [optional] 
**Zipcode** | Pointer to **string** | Zipcode is the stakeholder&#39;s postal code, if recorded. | [optional] 

## Methods

### NewCaptableCaptableStakeholder

`func NewCaptableCaptableStakeholder() *CaptableCaptableStakeholder`

NewCaptableCaptableStakeholder instantiates a new CaptableCaptableStakeholder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableStakeholderWithDefaults

`func NewCaptableCaptableStakeholderWithDefaults() *CaptableCaptableStakeholder`

NewCaptableCaptableStakeholderWithDefaults instantiates a new CaptableCaptableStakeholder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *CaptableCaptableStakeholder) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CaptableCaptableStakeholder) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CaptableCaptableStakeholder) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CaptableCaptableStakeholder) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *CaptableCaptableStakeholder) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CaptableCaptableStakeholder) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CaptableCaptableStakeholder) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CaptableCaptableStakeholder) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *CaptableCaptableStakeholder) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CaptableCaptableStakeholder) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CaptableCaptableStakeholder) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CaptableCaptableStakeholder) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CaptableCaptableStakeholder) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaptableCaptableStakeholder) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaptableCaptableStakeholder) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CaptableCaptableStakeholder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentRelationship

`func (o *CaptableCaptableStakeholder) GetCurrentRelationship() string`

GetCurrentRelationship returns the CurrentRelationship field if non-nil, zero value otherwise.

### GetCurrentRelationshipOk

`func (o *CaptableCaptableStakeholder) GetCurrentRelationshipOk() (*string, bool)`

GetCurrentRelationshipOk returns a tuple with the CurrentRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRelationship

`func (o *CaptableCaptableStakeholder) SetCurrentRelationship(v string)`

SetCurrentRelationship sets CurrentRelationship field to given value.

### HasCurrentRelationship

`func (o *CaptableCaptableStakeholder) HasCurrentRelationship() bool`

HasCurrentRelationship returns a boolean if a field has been set.

### GetEmail

`func (o *CaptableCaptableStakeholder) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CaptableCaptableStakeholder) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CaptableCaptableStakeholder) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CaptableCaptableStakeholder) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *CaptableCaptableStakeholder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableCaptableStakeholder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableCaptableStakeholder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableCaptableStakeholder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstitutionName

`func (o *CaptableCaptableStakeholder) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *CaptableCaptableStakeholder) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *CaptableCaptableStakeholder) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.

### HasInstitutionName

`func (o *CaptableCaptableStakeholder) HasInstitutionName() bool`

HasInstitutionName returns a boolean if a field has been set.

### GetName

`func (o *CaptableCaptableStakeholder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableCaptableStakeholder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableCaptableStakeholder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableCaptableStakeholder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStakeholderType

`func (o *CaptableCaptableStakeholder) GetStakeholderType() string`

GetStakeholderType returns the StakeholderType field if non-nil, zero value otherwise.

### GetStakeholderTypeOk

`func (o *CaptableCaptableStakeholder) GetStakeholderTypeOk() (*string, bool)`

GetStakeholderTypeOk returns a tuple with the StakeholderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderType

`func (o *CaptableCaptableStakeholder) SetStakeholderType(v string)`

SetStakeholderType sets StakeholderType field to given value.

### HasStakeholderType

`func (o *CaptableCaptableStakeholder) HasStakeholderType() bool`

HasStakeholderType returns a boolean if a field has been set.

### GetState

`func (o *CaptableCaptableStakeholder) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CaptableCaptableStakeholder) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CaptableCaptableStakeholder) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CaptableCaptableStakeholder) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreetAddress

`func (o *CaptableCaptableStakeholder) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *CaptableCaptableStakeholder) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *CaptableCaptableStakeholder) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *CaptableCaptableStakeholder) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetTaxId

`func (o *CaptableCaptableStakeholder) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CaptableCaptableStakeholder) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CaptableCaptableStakeholder) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CaptableCaptableStakeholder) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetZipcode

`func (o *CaptableCaptableStakeholder) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *CaptableCaptableStakeholder) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *CaptableCaptableStakeholder) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *CaptableCaptableStakeholder) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


